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
	"bytes"
	"crypto/aes"
)

// padding and unpadding functions by github.com/hothero
func pad(data []byte, block_size int)([]byte) {
	padding_len := block_size - len(data)%block_size
	padding := bytes.Repeat([]byte{byte(padding_len)}, padding_len)
	
	return append(data, padding...)
}

func unpad(data []byte)([]byte){
	padding := data[len(data)-1]

	return data[:len(data)-int(padding)]
}

// return []byte encrypted with key
func Encrypt(data []byte, key []byte)([]byte, error){
	var encrypted []byte
	
	aesBlock, err := aes.NewCipher(key)
	if(err != nil){
		return encrypted, err
	}

	// pad data
	data = pad(data, aesBlock.BlockSize())

	var tmp []byte = make([]byte, aesBlock.BlockSize())

	// encrypt data by blocks and append this blocks to encrypted
	for i := 0; i < len(data); i += aesBlock.BlockSize() {
		aesBlock.Encrypt(tmp, data[i:i+aesBlock.BlockSize()])
		encrypted = append(encrypted, tmp...)
	}

	return encrypted, err
}

// return []byte decrypted with key
func Decrypt(data []byte, key []byte)([]byte, error){
	var decrypted []byte

	aesBlock, err := aes.NewCipher(key)
	if(err != nil){
		return decrypted, err
	}

	var tmp []byte = make([]byte, aesBlock.BlockSize())

	// decrypt data by blocks and append this blocks to encrypted
	for i := 0; i < len(data); i += aesBlock.BlockSize() {
		aesBlock.Decrypt(tmp, data[i:i+aesBlock.BlockSize()])
		decrypted = append(decrypted, tmp...)
	}

	// unpad data before return
	decrypted = unpad(decrypted)

	return decrypted, err
}