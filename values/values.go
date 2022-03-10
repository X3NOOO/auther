package values

import "os"

const (
	URL string = "github.com/X3NOOO/auther"
	NAME string = "auther"
	VERSION = 0.1

	AUTHOR string = "X3NO"
	MAIL string = "X3NO@disroot.org"
	WEBSITE string = "https://github.com/X3NOOO"

)

var (
	HOME, _ = os.UserHomeDir()
	DB_PATH string = HOME + "/.auther_db"
)


type Db_secret_struct struct {
	Secret string		`json:"secret"`
	Algorithm string	`json:"algorithm"`
	Digits int			`json:"digits"`
	Period int			`json:"period"`
}

type Db_struct struct {
	Type	string				`json:"type"`
	Name	string				`json:"name"`
	Issuer	string				`json:"issuer"`
	Secret	Db_secret_struct 	`json:"secret"`
}
