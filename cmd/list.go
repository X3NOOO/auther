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
	"strconv"
	"strings"

	"github.com/X3NOOO/auther/utils"
	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List information",
	Long:  `List information about all your tokens`,
	Run: func(cmd *cobra.Command, args []string) {
		List()
	},
}

/*
* 1. read database from DB_path
* 2. decrypt database
* 3. read decrypted database
* 4. unmarshal json (or yaml?)
* 5. print name and issuer
 */
func List() {
	// run all things from here, not from Run: func

	// configure logger
	l := logger.NewLogger("list.go")
	l.SetVerbosity(Verbose)
	l.Debugln("list called")

	// read database
	db, err := utils.ReadDB(DB_path)
	if err != nil {
		l.Fatalln(1, err)
	}

	// decrypt database
	// db = db

	l.Debugln("json database: ", db)
	l.Debugln("entries in database: ", len(db))

	// get non-secret info from database and put it into variables so we can print it
	for i := 0; i <= len(db)-1; i++ {
		fmt.Println(strings.ToUpper(strconv.Itoa(i+1)+". "+db[i].Type) + ": " + db[i].Name + " @ " + db[i].Issuer)
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
