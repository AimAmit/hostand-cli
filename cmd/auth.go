package cmd

import (
	"github.com/aimamit/hostand-cli/api"
	"github.com/aimamit/hostand-cli/ui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	authCmd = &cobra.Command{
		Use:   "auth",
		Short: "user authentication",
	}

	signupCmd = &cobra.Command{
		Use:   "signup",
		Short: "Signup",
		Long:  "Signup with your credentials.",
		Run: func(cmd *cobra.Command, args []string) {
			email, err := ui.Validate("Email", "email")
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			password, err := ui.Validate("Password", "name")
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			stringToken, err := api.Signup(email, password)
			if err != nil {
				ui.Danger.Println(err.Error())
				return
			}
			viper.Set("email", email)
			viper.Set("token", stringToken)
			err = viper.WriteConfig()
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			ui.Success.Println("Successfully signed up!")
		},
	}

	signinCmd = &cobra.Command{
		Use:   "signin",
		Short: "Signin",
		Long:  "Signin with your credentials.",
		Run: func(cmd *cobra.Command, args []string) {
			email, err := ui.Validate("Email", "email")
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			password, err := ui.Validate("Password", "name")
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			stringToken, err := api.Signin(email, password)
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			viper.Set("email", email)
			viper.Set("token", stringToken)
			err = viper.WriteConfig()
			if err != nil {
				ui.Danger.Println(err)
				return
			}
			ui.Success.Println("Successfully signed in!")
		},
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	ui.Success.Println("pre run")
		// },
	}
)
