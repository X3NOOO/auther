package values

import (
	"os"
)

const (
	URL     string = "github.com/X3NOOO/auther"
	NAME    string = "auther"
	VERSION        = "0.9.0"

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

// this should be shield, if anyone can do better ascii art pls edit it
const ascii_art string = `
 .==_==.
:|&&&&&|: FIGLET
|&&' '&&| NAME
'\&&.&&/' HERE ` + VERSION +`
  ",&,"`

const HELLO_STRING string = ascii_art + "\nhttps://github.com/X3NOOO/auther\n" + "Copyright (C) 2022 X3NO <X3NO@disroot.org> [https://github.com/X3NOOO]" + "\nauther is program to manage your 2fa (totp) tokens released under GNU GPL v3 license."