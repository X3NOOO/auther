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
	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		get(args)
	},
}

func get(names []string){
	// configure logger
	l := logger.NewLogger("get.go")
	l.SetVerbosity(Verbose)
	l.Debugln("get called")

	l.Debugln("args:", names)

	// read database
	db_json, err := ReadDb(Db_path)
	if(err != nil){
		l.Fatalln(1, err)
	}
	l.Debugln("json database: ", db_json)

	// get through all names
	for i := 0; i <= len(names)-1; i++{
		// get through all database names
		for j := 0; j <= len(db_json)-1;j++{
			if(db_json[j].Name==names[i]){
				l.Debugln("match: " + db_json[j].Name) //TODO: change match to generated otp
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
