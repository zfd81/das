package util

import (
	"testing"
)

var key = []byte("daszfd198143dasp")

func Test_encryptAES(t *testing.T) {

	str := "1234567890qazwsxedc你好"
	str1, err := EncryptAES([]byte(str), key)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(str1))
	str2, err := DecryptAES(str1, key)
	t.Log(string(str2))
}
