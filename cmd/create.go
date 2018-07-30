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
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new strong password",
	Long: `create a new password which will be stored in your secrets. For example:

myPass create -> for interaction
myPass create [username] [email] [website] -> no interaction`,
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		writeToStore()

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

}

// writeToStore is a function to inject all data into the secret file
func writeToStore() {
	f, err := os.OpenFile(DBFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(strings.Join(passGenerator(), "") + "," + createSecrets() + "\n"); err != nil {
		log.Fatal(err)
	}
}

func createSecrets() string {
	var username string
	var email string
	var website string
	var data []string
	fmt.Printf("Please Enter username: ")
	fmt.Scanf("%v", &username)
	fmt.Printf("Please Enter email: ")
	fmt.Scanf("%v", &email)
	fmt.Printf("Please Enter website: ")
	fmt.Scanf("%v", &website)
	data = append(data, username, email, website)
	return strings.Join(data, ",")
}

// PassGenerator to generate a strong Password and return it as string.
//    Use ASCI
func passGenerator() []string {
	var ASCI = "QWERTYUIOPLKJHGFDSAZXCVBNMmnbvcxzasdfghjklopiuytrewq7869543210"
	var Sign = "#!@&()_-][><~;{}"
	var Pass = []string{}
	for i := 0; i < 8; i++ {
		Pass = append(Pass, string(ASCI[rand.Intn(len(ASCI))]))
	}
	for i := 0; i < rand.Intn(3); i++ {
		Pass[rand.Intn(10)] = string(Sign[rand.Intn(len(Sign))])
	}
	return Pass
}
