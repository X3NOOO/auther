package utils

import (
	"github.com/atotto/clipboard"
)

func Copy(to_copy string)(error){
	return clipboard.WriteAll(to_copy)
}