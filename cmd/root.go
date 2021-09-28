package cmd

import (
	"fmt"
	"net"

	"github.com/aimamit/hostand-cli/ui"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "hello",
		Version: "1.0.0",
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	// name, err := prompt.Validate("Your name")
		// 	// if err != nil {
		// 	// 	return err
		// 	// }
		// 	// fmt.Printf("hello %s\n", name)
		// 	return nil
		// },
	}

	cmdTimes = &cobra.Command{
		Use:   "ip",
		Short: "Get IP address",
		Long:  `Get IP address`,
		// Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			_, url, err := ui.Select("URLs", []string{"google.com", "facebook.com", "amazon.in"})
			if err != nil {
				return err
			}
			ui.Success.Printf("%s | ", url)
			ips, err := net.LookupIP(url)
			if err != nil {
				return err
			}
			// ipsString := strings.Join(fmt.Sprintf("%s", ips), " ")
			ui.Cyan.Printf("%s\n", fmt.Sprintf("%s", ips))
			return nil
		},
		Example: "example",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetVersionTemplate("1.0.0")
	cmdTimes.Flags().StringP("name", "n", "world", "Set you name")

	authCmd.AddCommand(signupCmd, signinCmd)
	rootCmd.AddCommand(cmdTimes, buildCmd, authCmd)

}
