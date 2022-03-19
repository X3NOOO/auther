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

	"github.com/X3NOOO/auther/utils"
	"github.com/X3NOOO/auther/values"
	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

var (
	flag_type      string
	flag_name      string
	flag_issuer    string
	flag_secret    string
	flag_algorithm string
	flag_period    int
	flag_counter   int
	flag_digits    int
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add token",
	Long:  `Add otp token to database`,
	Run: func(cmd *cobra.Command, args []string) {
		Add()
	},
}

/*
 * 1. Read db
 * 2. Generate object db struct from flags
 * 3. Append object to db_json array
 * 4. Encrypt db_json
 * 5. Write db_json to DB_path
 */
func Add() {
	// configure logger
	l := logger.NewLogger("add.go")
	l.SetVerbosity(Verbose)
	l.Debugln("add called")

	// read database
	db_encrypted, err := utils.ReadDB(DB_path)
	if err != nil {
		l.Fatalln(1, err)
	}

	// decrypt db
	// TODO add encryption
	db := db_encrypted

	l.Debugln("json database:", db)
	l.Debugln("database length:", len(db))

	// create object based on flags
	new_entry := values.Db_struct{
		Type:   flag_type,
		Name:   flag_name,
		Issuer: flag_issuer,
		Secret: values.Db_secret_struct{
			Secret:    flag_secret,
			Algorithm: flag_algorithm,
			Digits:    flag_digits,
			Period:    flag_period,
			Counter:   int64(flag_counter),
		},
	}

	// db_new = db + new_entry
	db_new := append(db, new_entry)
	l.Debugln("new_entry:", db_new)

	// convert db_new to json
	db_new_json, err := json.Marshal(db_new)
	if err != nil {
		l.Fatalln(1, err)
	}
	l.Debugln("new db:", string(db_new_json))

	// encrypt db_new_json
	// TODO add encryption
	db_new_encrypted := db_new_json

	// write db_new_encrypted to DB_path
	if(!Testing){
		stat, err := os.Stat(DB_path)
		if err != nil {
			l.Fatalln(1, err)
		}
		mode := stat.Mode().Perm()
		ioutil.WriteFile(DB_path, []byte(db_new_encrypted), mode)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().StringVarP(&flag_type, "type", "t", "totp", "totp/hotp")
	addCmd.Flags().StringVarP(&flag_name, "name", "n", "", "name of your 2fa token")
	addCmd.Flags().StringVarP(&flag_issuer, "issuer", "i", "", "issuer of your 2fa token")
	addCmd.Flags().StringVarP(&flag_secret, "secret", "s", "", "secret of your 2fa token")
	addCmd.Flags().StringVarP(&flag_algorithm, "algorithm", "a", "SHA1", "algorithm of your 2fa token")
	addCmd.Flags().IntVarP(&flag_period, "period", "p", 30, "period [totp only]")
	addCmd.Flags().IntVarP(&flag_counter, "counter", "c", 0, "counter [hotp only]")
	addCmd.Flags().IntVarP(&flag_digits, "digits", "g", 0, "digits [totp only]")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("secret")
}
