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
	// "errors"
	"fmt"
	"os"

	"github.com/X3NOOO/auther/values"
	"github.com/X3NOOO/logger"
	"github.com/spf13/cobra"
)

var (
	Verbose int
	DB_path string
	Testing bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "auther",
	Short: "Manage your otp tokens",
	Long:  values.HELLO_STRING,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.NewLogger("root.go")
		l.SetVerbosity(Verbose)
		l.Debugln("Verbosity:", Verbose)
		
		// hello message
		hello()
	},
}

func hello(){
	fmt.Println(values.HELLO_STRING)
}

//func fileExists(name string) bool {
//	_, err := os.Stat(name)
//	if err == nil {
//		return true
//	}
//	if errors.Is(err, os.ErrNotExist) {
//		return false
//	}
//	return false
//}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.auther.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().IntVarP(&Verbose, "verbose", "v", 3, "verbosity of output (0-5)")
	rootCmd.PersistentFlags().StringVarP(&DB_path, "database", "d", values.DB_path, "path to database")
	rootCmd.PersistentFlags().BoolVar(&Testing, "testing", false, "disable writing to database")

}

