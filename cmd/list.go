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
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print stored passwords",
	Long: `Print out all passwords and secrets with its related data, > can be used with pipline | . For example:

myPass-cmd list
mypass-cmd list > [filename].`,
	Run: func(cmd *cobra.Command, args []string) {
		listSecrets()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listSecrets() {
	f, err := ioutil.ReadFile(DBFile)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(f))
	}
}
