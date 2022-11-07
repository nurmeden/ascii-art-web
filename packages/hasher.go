package packages

import (
	"crypto/sha256"
	"io/ioutil"
)

// func HashCheck(a string) bool {
// 	hasher := sha256.New()
// 	s, err := ioutil.ReadFile(a)
// 	hasher.Write(s)
// 	CheckErr(err)
// 	l := hasher.Sum(nil)
// 	hashstandard := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
// 	hashshadow := []byte{38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129, 116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115}
// 	hashthinkertoy := []byte{100, 40, 94, 73, 96, 209, 153, 244, 129, 147, 35, 196, 220, 99, 25, 186, 52, 241, 240, 221, 157, 161, 77, 7, 17, 19, 69, 245, 215, 108, 63, 163}
// 	if a == "./bannerFiles/standard.txt" {
// 		return string(l) == string(hashstandard)
// 	} else if a == "./bannerFiles/shadow.txt" {
// 		return string(l) == string(hashshadow)
// 	} else if a == "./bannerFiles/thinkertoy.txt" {
// 		return string(l) == string(hashthinkertoy)
// 	}
// 	return false
// }
func HashCheck(a string) (bool, error) {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(a)
	hasher.Write(s)
	// CheckErr(err)
	l := hasher.Sum(nil)
	hashstandard := []byte{225, 148, 241, 3, 52, 66, 97, 122, 184, 167, 142, 28, 166, 58, 32, 97, 245, 204, 7, 163, 240, 90, 194, 38, 237, 50, 235, 157, 253, 34, 166, 191}
	hashshadow := []byte{38, 185, 77, 11, 19, 75, 119, 233, 253, 35, 224, 54, 11, 253, 129, 116, 15, 128, 251, 127, 101, 65, 209, 216, 197, 216, 94, 115, 238, 85, 15, 115}
	hashthinkertoy := []byte{71, 44, 250, 176, 86, 116, 6, 209, 3, 19, 171, 34, 32, 170, 94, 100, 183, 251, 241, 121, 109, 150, 193, 28, 119, 82, 70, 115, 45, 206, 126, 122}
	if a == "./bannerFiles/standard.txt" {
		return string(l) == string(hashstandard), nil
	} else if a == "./bannerFiles/shadow.txt" {
		return string(l) == string(hashshadow), nil
	} else if a == "./bannerFiles/thinkertoy.txt" {
		return string(l) == string(hashthinkertoy), nil
	}
	return false, err
}
