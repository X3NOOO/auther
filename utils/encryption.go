package utils

import (
	"crypto/aes"
	"fmt"
)

// return []byte encrypted by key
func Encrypt(data []byte, key []byte)([]byte, error){
	var encrypted []byte
	
	aesBlock, err := aes.NewCipher(key)
	if(err!=nil){
		return nil, err
	}

	var tmp []byte = make([]byte, aesBlock.BlockSize())

	// encrypt data by blocks and append this blocks to encrypted
	fmt.Println(len(data))
	for i := 0; i < len(data); i += aesBlock.BlockSize() {
		aesBlock.Encrypt(tmp, data[i:i+aesBlock.BlockSize()])
		if err != nil {
			return encrypted, err
		}
		encrypted = append(encrypted, tmp...)
	}

	return encrypted, err
}
