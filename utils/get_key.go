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