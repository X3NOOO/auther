package utils

import (
	"errors"
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