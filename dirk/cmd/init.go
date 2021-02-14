/*
Copyright © 2021 Malte Groth <malte.groth@gmx.net>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/grothesk/go-dirk/dirk/internal/logging"
	direnv "github.com/grothesk/go-dirk/dirk/pkg/direnv"
	"github.com/grothesk/go-dirk/dirk/pkg/file/config"
	"github.com/grothesk/go-dirk/dirk/pkg/file/config/kubeconfig"
	"github.com/grothesk/go-dirk/dirk/pkg/file/rc"
	envrc "github.com/grothesk/go-dirk/dirk/pkg/file/rc/envrc"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes dirk in the desired directory",
	Long: `init sets up an .envrc file in the directory passed as an argument 
und refers a kubeconfig file.`,
	Args: initArgs,
	Run:  initRun,
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	initCmd.Flags().StringP("configfile", "c", "", "Config file to copy to kubeconfig")
	if err := viper.BindPFlag("configfile", initCmd.Flags().Lookup("configfile")); err != nil {
		logging.Logger.Fatal("Unable to bind flag: " + err.Error())
	}

	initCmd.Flags().StringP("mode", "m", "skip", "")
	if err := viper.BindPFlag("mode", initCmd.Flags().Lookup("mode")); err != nil {
		logging.Logger.Fatal("Unable to bind flag: " + err.Error())
	}
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// There has to be exact one argument and it has to be a directory
func initArgs(cmd *cobra.Command, args []string) error {
	err := cobra.ExactArgs(1)(cmd, args)
	if err != nil {
		return err
	}

	dir, err := os.Stat(args[0])
	if os.IsNotExist(err) {
		return fmt.Errorf("%s: no such directory", args[0])
	}
	if dir.Mode().IsRegular() {
		return fmt.Errorf("%s: exists as a file", args[0])
	}

	return nil
}

func initRun(cmd *cobra.Command, args []string) {
	dir, err := filepath.Abs(args[0])
	if err != nil {
		logging.Logger.Fatal(err.Error())
	}
	logging.Logger.Info(fmt.Sprintf("init dirk in %s.", dir))

	logging.Logger.Info("check if direnv is present on PATH.")
	if direnv.Exists() {
		logging.Logger.Info("direnv is on PATH.")

		ef := envrc.NewFile(dir)
		err = rc.SetupFile(&ef)
		if err != nil {
			logging.Logger.Fatal(err.Error())
		}

		kf := kubeconfig.NewFile(dir)
		configfile := viper.GetString("configfile")
		mode := viper.GetString("mode")
		err = config.SetupFile(&kf, configfile, mode)
		if err != nil {
			logging.Logger.Fatal(err.Error())
		}
	} else {
		logging.Logger.Info("cannot find direnv on PATH")
		logging.Logger.Info("please make sure that direnv has been installed.")
	}
}
