package utils

import (
	"crypto/sha256"
	"syscall"

	"golang.org/x/term"
)

// return password from stdin hashed with sha256
func GetKey()([]byte,error){
	pass, err := term.ReadPassword(int(syscall.Stdin))
	if(err != nil){
	    return nil, err
	}
	var pass256 = sha256.Sum256(pass);	
	pass256_slice := pass256[:]
	return pass256_slice, nil
}