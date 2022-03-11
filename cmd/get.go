/*
Copyright © 2022 X3NO <X3NO@disroot.org> [https://github.com/X3NOOO]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/X3NOOO/auther/utils"
	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Generate otp codes",
	Long:  `Generate otp codes from arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		Get(args)
	},
}

// get totp codes based on args
func Get(args []string) {
	// configure logger
	l := logger.NewLogger("get.go")
	l.SetVerbosity(Verbose)
	l.Debugln("get called")

	l.Debugln("args:", args)

	// read database
	db, err := utils.ReadDB(DB_path)
	if err != nil {
		l.Fatalln(1, err)
	}
	l.Debugln("json database:", db)

	// get through all args
	for i := 0; i <= len(args)-1; i++ {
		// get through all database names
		for j := 0; j <= len(db)-1; j++ {
			if db[j].Name == args[i] {
				// l.Debugln("match: " + db[j].Name) // TODO change match to generated otp
				if strings.ToLower(db[j].Type) == "totp" {
					code, err := utils.GenTOTP([]byte(db[i].Secret.Secret))
					if err != nil {
						l.Fatalln(1, err)
					}
					fmt.Println(code)
					// TODO add copying code to clipboard
				} else if strings.ToLower(db[j].Type) == "hotp" {
					code, err := utils.GenHOTP([]byte(db[i].Secret.Secret), db[i].Secret.Counter)
					if err != nil {
						l.Fatalln(1, err)
					}
					fmt.Println(code)
					// TODO add copying code to clipboard
				}
			}
		}

		// if argument have uri-format get totp from it
		if len(args[i]) >= 15 {
			if args[i][:15] == "otpauth://totp/" || args[i][:15] == "otpauth://hotp/" {
				l.Infoln("argument have uri format - trying to generate code from it")
				code, err := utils.GenFromURI(args[i])
				if err != nil {
					l.Fatalln(1, err)
				}
				fmt.Println(code)
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
