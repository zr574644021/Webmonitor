package main

import (
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"fmt"
)

func main () {
	password := "123456"
	salt1 := "fi22.ij5.,2432!i"
	salt2 := "fo2.43o5h2f(juaz"
	md5_obj := md5.New()
	md5_obj.Write([]byte(salt1 + password + salt2))
	md5_encode := md5_obj.Sum(nil)
	md5_string := hex.EncodeToString(md5_encode)

	salt3 := "easfcvadwa"
	salt4 := "ofkafjdisa"
	sha1_obj := sha1.New()
	sha1_obj.Write([]byte(salt3 + md5_string + salt4))
	sha1_encode := sha1_obj.Sum(nil)
	fmt.Println(hex.EncodeToString(sha1_encode))
}