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
package values

import (
	"os"
)

const (
	URL     string = "github.com/X3NOOO/auther"
	NAME    string = "auther"
	VERSION        = "1.0.0"

	AUTHOR  string = "X3NO"
	MAIL    string = "X3NO@disroot.org"
	WEBSITE string = "https://github.com/X3NOOO"
)

var (
	HOME, _        = os.UserHomeDir()
	DB_path string = HOME + "/.auther_db"
)

type Db_secret_struct struct {
	Secret    string `json:"secret"`
	Algorithm string `json:"algorithm"`
	Digits    int    `json:"digits"`
	Period    int    `json:"period"`
	Counter   int64  `json:"counter"`
}

type Db_struct struct {
	Type   string           `json:"type"`
	Name   string           `json:"name"`
	Issuer string           `json:"issuer"`
	Secret Db_secret_struct `json:"secret"`
}

// ╔═║║ ║═╔╝║ ║╔═╝╔═║
// ╔═║║ ║ ║ ╔═║╔═╝╔╔╝
// ╝ ╝══╝ ╝ ╝ ╝══╝╝ ╝

// this should be shield, if anyone can do better ascii art pls edit it
const ascii_art string = `
 .==_==.   dBBBBBb    dBP dB dBBBBBB dBP dB dBBB dBBBBBb 
:|&&&&&|:  BB                                   dBP      
|&&' '&&|  dBP BB  dBP dBP   dBP  dBBBBB dBBP   dBBBBK   
'\&&.&&/'  dBP  BB dBP_dBP   dBP  dBP dB dBP    dBP  BB  
  ",&,"    dBBBBBB dBBBBBP   dBP  dBP dB dBBBBP dBP  dB'  ` + VERSION + `
`

const HELLO_STRING string = ascii_art + "\nhttps://github.com/X3NOOO/auther\n" + "Copyright (C) 2022 X3NO <X3NO@disroot.org> [https://github.com/X3NOOO]" + "\nauther is program to manage your 2fa (totp) tokens released under GNU GPL v3 license."