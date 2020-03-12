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
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tijori/config"
	"github.com/tijori/tijori"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves the UserName and Password combination to a file. The File is encoded and not human readable.",

	Run: func(cmd *cobra.Command, args []string) {
		//tijorifmt.Println("save called")
		var passwordInfo config.SavedPassword
		//passwordInfo.UserName
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("username (Required)")

		username, err := reader.ReadString('\n')
		username = strings.TrimSuffix(username, "\n")
		if err != nil {
			fmt.Println("caught error ", err.Error())
		}
		if isEmpty(username) {
			fmt.Println("username cannot be empty")
			os.Exit(1)
		}
		fmt.Printf("\npassword (Required) ")
		password, _ := reader.ReadString('\n')
		password = strings.TrimSuffix(password, "\n")
		if isEmpty(password) {
			fmt.Println("password cannot be empty")
			os.Exit(1)
		}

		fmt.Printf("\nadditional info (Optional)")

		additionalInfo, _ := reader.ReadString('\n')

		passwordInfo.AdditionalInfo = additionalInfo
		passwordInfo.Password = password
		passwordInfo.UserName = username

		tijori.AddtoSavedPasswords(passwordInfo)
		fmt.Println("Saved!")

	},
}

func isEmpty(s string) bool {
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
