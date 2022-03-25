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
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"errors"
	"strings"
	"time"

	// "time"

	"github.com/uaraven/gotp"
)

// generate totp code based on secret
func GenTOTP(secret []byte, digits int, algorithm crypto.Hash, duration int)(string, error){
	// totp := gotp.NewDefaultTOTP(secret)
	// timest := time.Now()
	totp := gotp.NewTOTPHash(secret, int(digits), duration, 0, algorithm)
	code := totp.Now()

	err := totp.Verify(code, time.Now().Unix())
	if(!err){
		return "", errors.New("cannot verify totp")
	}

	return code, nil
}

// generate hotp code based on secret
func GenHOTP(secret []byte, digits int, algorithm crypto.Hash, counter int64)(string, error){
	// hotp := gotp.NewDefaultHOTP(secret, counter)
	hotp := gotp.NewHOTPHash(secret, counter, digits, -1, algorithm)
    code := hotp.CurrentOTP()

	err := hotp.Verify(code, counter)
	if(!err){
		return "", errors.New("cannot verify hotp")
	}

	return code, nil
}

// generate otp code form uri
func GenFromURI(uri string)(string, error){
	otp, err := gotp.OTPFromUri(uri)
	if(err != nil){
		return "", err
	}
	// check ifuri is totp or hotp
	if(strings.ToLower(uri[10:14]) == "totp"){
		return GenTOTP(otp.OTP.GetSecret(), otp.OTP.GetDigits(), otp.OTP.GetHash(), 30)
	} else if(strings.ToLower(uri[10:14]) == "hotp"){
		return GenHOTP(otp.OTP.GetSecret(), otp.OTP.GetDigits(), otp.OTP.GetHash(), 0) // TODO replace 0 with counter
	}
	return "", errors.New("uri isn't valid")
}

func GetHash(algo string)(crypto.Hash){
	switch strings.ToUpper(algo){
	case "SHA1", "SHA-1":
		return crypto.SHA1
	case "SHA256", "SHA-256":
		return crypto.SHA256
	case "MD5", "MD-5":
		return crypto.MD5
	default:
		return crypto.SHA1
	}
}