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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	l.Debugln("database:", db)

	// decrypt database
	// TODO add encryption
	// db = db

	// get through all args
	var code string = ""
	for i := 0; i <= len(args)-1; i++ {
		// get through all database names
		for j := 0; j <= len(db)-1; j++ {
			if db[j].Name == args[i] {
				if strings.ToLower(db[j].Type) == "totp" {
					algo := utils.GetHash(db[j].Secret.Algorithm)
					l.Debugln("Algo:", algo)
					code, err = utils.GenTOTP([]byte(db[j].Secret.Secret), db[j].Secret.Digits, algo, db[j].Secret.Period)
					if err != nil {
						l.Fatalln(1, err)
					}
				} else if strings.ToLower(db[j].Type) == "hotp" {
					algo := utils.GetHash(db[j].Secret.Algorithm)
					l.Debugln("Algo:", algo)
					code, err = utils.GenHOTP([]byte(db[j].Secret.Secret), db[j].Secret.Digits, algo, db[j].Secret.Counter)
					l.Debugln("counter:",db[j].Secret.Counter)
					if err != nil {
						l.Fatalln(1, err)
					}
					// increase counter by 1 every time this block is called
					db[j].Secret.Counter++

					// marshal db
					db_new_json, err := json.Marshal(db)
					if err != nil {
						l.Fatalln(1, err)
					}
					l.Debugln("new db:", string(db_new_json))

					// encrypt db_new_json
					// TODO add encryption
					db_new_encrypted := db_new_json

					// write db_new_encrypted to DB_path
					stat, err := os.Stat(DB_path)
					if err != nil {
						l.Fatalln(1, err)
					}
					mode := stat.Mode().Perm()
					ioutil.WriteFile(DB_path, []byte(db_new_encrypted), mode)
				}
			}
		}

		// if code still is nil and argument have uri-format get code from it
		if code == "" && len(args[i]) >= 15{
			if args[i][:15] == "otpauth://totp/" || args[i][:15] == "otpauth://hotp/" {
				l.Infoln("argument have uri format - trying to generate code from it")
				code, err = utils.GenFromURI(args[i])
				if(args[i][:15] == "otpauth://hotp/"){
					l.Warningln("getting counter from HOTP uri isn't supported yet - using value of 0")
				}
				if err != nil {
					l.Fatalln(1, err)
				}
				
			}
		}

		// if code still is nil and args[i] is a digit get code from db[i]y
		if code == "" {
			num, err := strconv.Atoi(args[i])

			if(!(err != nil || num <= 0 || num > len(db))){
				// if args[num-1] is valid entry
				if strings.ToLower(db[num-1].Type) == "totp" {
					algo := utils.GetHash(db[num-1].Secret.Algorithm)
					l.Debugln("Algo:", algo)
					code, err = utils.GenTOTP([]byte(db[num-1].Secret.Secret), db[num-1].Secret.Digits, algo, db[num-1].Secret.Period)
					if err != nil {
						l.Fatalln(1, err)
					}
				} else if strings.ToLower(db[num-1].Type) == "hotp" {
					algo := utils.GetHash(db[num-1].Secret.Algorithm)
					l.Debugln("Algo:", algo)
					code, err = utils.GenHOTP([]byte(db[num-1].Secret.Secret), db[num-1].Secret.Digits, algo, db[num-1].Secret.Counter)
					l.Debugln("counter:",db[num-1].Secret.Counter)
					if err != nil {
						l.Fatalln(1, err)
					}
					// increase counter by 1 every time this block is called
					db[num-1].Secret.Counter++

					// marshal db
					db_new_json, err := json.Marshal(db)
					if err != nil {
						l.Fatalln(1, err)
					}
					l.Debugln("new db:", string(db_new_json))

					// encrypt db_new_json
					// TODO add encryption
					db_new_encrypted := db_new_json

					// write db_new_encrypted to DB_path
					stat, err := os.Stat(DB_path)
					if err != nil {
						l.Fatalln(1, err)
					}
					mode := stat.Mode().Perm()
					ioutil.WriteFile(DB_path, []byte(db_new_encrypted), mode)
				}
			} else {
				// if args[i] is not valid entry
				l.Warning(args[i],"has not been recognized as a database entry")
			}
		}
	}

	// print and copy code
	fmt.Println(code)
	err = utils.Copy(code)
	if(err != nil){
		l.Warningln(1, err)
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
