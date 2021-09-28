package cmd

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/aimamit/hostand-cli/api"
	"github.com/aimamit/hostand-cli/ui"
	"github.com/spf13/cobra"
)

type HaConfig struct {
	IgnoreFile []string `json:"ignoreFile"`
}

var isDone bool
var haConfig HaConfig

var (
	buildCmd = &cobra.Command{
		Use:   "build",
		Short: "Deploy your app",
		Long:  "Build and deploy your app using our platform.",
		RunE: func(cmd *cobra.Command, args []string) error {
			isDone = false
			go progress()
			var buf bytes.Buffer
			byteValue, _ := ioutil.ReadFile("haconfig.json")
			json.Unmarshal(byteValue, &haConfig)

			_ = compress(".", &buf)
			err := api.BuildImage("cli-domain", "v2", buf)
			if err != nil {
				return err
			}
			isDone = true
			time.Sleep(300 * time.Millisecond)
			return nil
		},
	}
)

func compress(src string, buf io.Writer) error {
	// tar > gzip > buf
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {

		for _, pat := range haConfig.IgnoreFile {
			var matched bool
			m, er := filepath.Match(pat, filepath.ToSlash(file))
			if er != nil {
				return er
			}
			matched = m || filepath.Dir(file) == pat
			if matched {
				// fmt.Println(filepath.ToSlash(file))
				return nil
			}
		}
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}
		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)
		// fmt.Println(header.Name)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})
	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}

func progress() {
	progressString := `|/-\`
	i := 0
	for true {
		if isDone {
			break
		}
		if i == len(progressString) {
			i = 0
		}
		ui.Cyan.Printf("\033[2K\r%s", string(progressString[i]))
		time.Sleep(300 * time.Millisecond)
		i++
	}
	ui.Success.Printf("\033[H\033[2J%s\n", "File Uploaded")

}
