package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/uaraven/gotp"
)

// generate totp code based on secret
func GenTOTP(secret []byte)(string, error){
	timestamp := time.Now()
	// timestamp := time.Date(t, time.UTC)
	totp := gotp.NewDefaultTOTP(secret)
	code := totp.At(timestamp)

	err := totp.Verify(code, timestamp.Unix())
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
		return GenTOTP(otp.OTP.GetSecret())
	} else if(strings.ToLower(uri[10:14]) == "hotp"){
		return GenHOTP(otp.OTP.GetSecret(), int64(otp.OTP.GetHash()))
	}
	return "", errors.New("uri isn't valid")
}
