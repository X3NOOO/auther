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
	"io/ioutil"
	"os"

	// "io/ioutil"
	// "os"

	"github.com/X3NOOO/auther/utils"
	"github.com/X3NOOO/auther/values"
	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

// remCmd represents the rem command
var remCmd = &cobra.Command{
	Use:   "rem",
	Short: "Remove token",
	Long:  `Remove otp token from database`,
	Run: func(cmd *cobra.Command, args []string) {
		Rem(args)
	},
}

func RemoveIndex(s []values.Db_struct, index int) []values.Db_struct {
	return append(s[:index], s[index+1:]...)
}

func Rem(args []string) {
	// configure logger
	l := logger.NewLogger("rem.go")
	l.SetVerbosity(Verbose)
	l.Debugln("rem called")

	// read database
	db_encrypted, err := utils.ReadDB(DB_path)
	if err != nil {
		l.Fatalln(1, err)
	}

	// decrypt database
	// TODO add encryption
	db := db_encrypted

	l.Debugln("json database:", db)

	var db_new []values.Db_struct = db
	// get through all args
	for i := 0; i <= len(args)-1; i++ {
		// get through all database names
		for j := 0; j <= len(db_new)-1; j++ {
			if db[j].Name == args[i] {
				db_new = RemoveIndex(db_new, j)
				l.Debugln("db_new:", db_new)
			}
		}
	}

	// if we didnt found args in db end program
	// if(db_new == nil){
		// l.Fatalln(1, "Not found")
	// }

	// convert db_new to json
	db_new_json, err := json.Marshal(db_new)
	if err != nil {
		l.Fatalln(1, err)
	}
	l.Debugln("db_new_json:", string(db_new_json))

	// encrypt db_new_json
	// TODO add encryption
	db_new_encrypted := db_new_json

	// write db_new_json to DB_path
	stat, err := os.Stat(DB_path)
	if err != nil {
		l.Fatalln(1, err)
	}
	mode := stat.Mode().Perm()
	ioutil.WriteFile(DB_path, []byte(db_new_encrypted), mode)
}

func init() {
	rootCmd.AddCommand(remCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// remCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// remCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
