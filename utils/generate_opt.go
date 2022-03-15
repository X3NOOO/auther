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
func GenTOTP(secret []byte, digits int64, algorithm crypto.Hash)(string, error){
	// totp := gotp.NewDefaultTOTP(secret)
	// timest := time.Now()
	totp := gotp.NewTOTPHash(secret, int(digits), 30, 0, algorithm)
	code := totp.Now()

	err := totp.Verify(code, time.Now().Unix())
	if(!err){
		return "", errors.New("cannot verify totp")
	}

	return code, nil
}

// generate hotp code based on secret
func GenHOTP(secret []byte, counter int64)(string, error){
	hotp := gotp.NewDefaultHOTP(secret, counter)
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
		return GenTOTP(otp.OTP.GetSecret(), int64(otp.OTP.GetDigits()), otp.OTP.GetHash())
	} else if(strings.ToLower(uri[10:14]) == "hotp"){
		return GenHOTP(otp.OTP.GetSecret(), int64(otp.OTP.GetHash()))
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