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
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"github.com/tijori/tijori"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetches the chosen password and writes to clipboard.",
	Long: `Use this to input the choice to fetch the password. The selected password is written to clipboard and can be Pasted 
	into destination of your choice.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		savedPasswords := tijori.LoadSavedPasswords()
		if len(savedPasswords) == 0 {
			fmt.Println("No saved passwords")
			return
		} else {
			for k, v := range savedPasswords {
				fmt.Printf("%d. UserName : %s \n", k+1, v.UserName)
				if len(v.AdditionalInfo) != 0 {
					fmt.Printf("Additional Info: %s \n", v.AdditionalInfo)

				}
			}
		}
		strNum, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("caught error ", err.Error())
			return
		}
		numInput, _ := strconv.Atoi(strNum)

		password := tijori.FetchSavedPasswordFor(numInput + 1)
		err = writeToClipBoard(password.Password)
		if err != nil {
			fmt.Println("caught error ", err.Error())
			return
		} else {
			fmt.Println("Password is copied to clipboard")
		}
	},
}

func writeToClipBoard(password string) error {

	return clipboard.WriteAll(password)

}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
