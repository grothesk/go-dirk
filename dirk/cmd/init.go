/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		configfile := viper.GetString("configfile")
		mode := viper.GetString("mode")

		fmt.Printf("configfile: %s\n", configfile)
		fmt.Printf("mode: %s\n", mode)

		for i, v := range args {
			fmt.Printf("argument %d: %s\n", i, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().StringP("configfile", "c", "kubeconfig", "Config file to copy to ")
	if err := viper.BindPFlag("configfile", initCmd.Flags().Lookup("configfile")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}

	initCmd.Flags().StringP("mode", "m", "skip", "")
	if err := viper.BindPFlag("mode", initCmd.Flags().Lookup("mode")); err != nil {
		log.Fatal("Unable to bind flag:", err)
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
