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
	"strings"

	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List information",
	Long:  `List information about all your tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

type db_secret_struct struct {
	Secret string		`json:"secret"`
	Algorithm string	`json:"algorithm"`
	Digits int			`json:"digits"`
	Period int			`json:"period"`
}

type db_struct struct {
	Type	string				`json:"type"`
	Name	string				`json:"name"`
	Issuer	string				`json:"issuer"`
	Secret	db_secret_struct 	`json:"secret"`
}

/*
* 1. read database from db_path
* 2. decrypt database 			//TODO: or maybe encrypt only secret so you can use list without entering password?
* 3. read decrypted database
* 4. unmarshal json (or yaml?)
* 5. print name and issuer
 */
func list() {
	// run all things from here, not from Run: func

	// configure logger
	l := logger.NewLogger("list.go")
	l.SetVerbosity(Verbose)
	l.Debugln("list called")

	// read database
	db_encrypted, err := ioutil.ReadFile(db_path)
	if err != nil {
		l.Fatalln(1, err)
	}
	l.Debugln("encrypted database: ", db_encrypted)

	// decryption
	db_decrypted := db_encrypted
	l.Debugln("decrypted database: ", string(db_decrypted))

	// parsing json
	var db_json []db_struct
	err = json.Unmarshal(db_decrypted, &db_json)
	if err != nil {
		l.Fatalln(1, err)
	}
	l.Debugln("json database: ", db_json)
	l.Debugln("entries in database: ", len(db_json))

	// get non-secret info from database and put it into variables so we can print ir
	// info_name := make([]string, len(db_json))
	// info_type := make([]string, len(db_json))
	// info_issuer := make([]string, len(db_json))
	for i := 0; i<=len(db_json)-1; i++{
		fmt.Println(strings.ToUpper(db_json[i].Type) + ": " + db_json[i].Name + "@" + db_json[i].Issuer)
	}

}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
