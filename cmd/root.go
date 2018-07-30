// Copyright © 2018 Hemied <hazem.hemied@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// DBDir directort
var DBDir = filepath.Join(os.Getenv("HOME"), ".mypass")

// DBFile secret file
var DBFile = filepath.Join(DBDir, ".db")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myPass",
	Short: "myPass mange your secrets",
	Long: `This program is to help you create a strong password and store them in a secret DB,
no more need to save passwords. For example:

MyPass is a CLI Program for GNU/Linux, OSX and MS Windows.
This application is a tool to generate  a strong and unique password,
make it related to a username, email and website.
the needed files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	DBExist()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myPass.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// DBExist to check if the DB file "~/.mypass/.db" is exist or not.
// is not exist, go to another FUNC to create the DB structure.
func DBExist() {
	if _, err := os.Stat(DBDir); os.IsNotExist(err) {
		color.Magenta("This program made for you with ❤️ ...")
		color.Green("Starting for the first time")
		color.Green("Creating [ %v ] ...", DBDir)
		os.Mkdir(DBDir, 0700)

	}
	if _, err := os.Stat(DBFile); os.IsNotExist(err) {
		color.Green("Creating [ %v ] ...", DBFile)
		os.Create(DBFile)
		os.Chmod(DBFile, 0600)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".myPass" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".myPass")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
