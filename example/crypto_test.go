package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestCrypto(t *testing.T) {
	crypto := ktools.Crypto

	src := "hello"
	encrypt := crypto.MD5.EncodeString(src)
	fmt.Println("MD5 Encrypt:", encrypt)

	encrypt = crypto.MD5.EncodeBytes([]byte(src))
	fmt.Println("MD5 Encrypt:", encrypt)

	encrypt = crypto.Base64.EncodeString(src)
	fmt.Println("Base64 Encrypt:", encrypt)

	encrypt = string(crypto.Base64.EncodeByte([]byte(src)))
	fmt.Println("Base64 Encrypt:", encrypt)

	decrypt := crypto.Base64.DecodeString(encrypt)
	fmt.Println("Base64 Decrypt:", decrypt)

	decrypt = string(crypto.Base64.DecodeByte([]byte(encrypt)))
	fmt.Println("Base64 Decrypt:", decrypt)
	fmt.Println("------------------------------------------")

	es := crypto.AES.EncryptString(src, "abcdefghijklmnopqrstuvwxyzABCDEF")
	fmt.Println("AES Encrypt:", es)

	ds := crypto.AES.DecryptString(es, "abcdefghijklmnopqrstuvwxyzABCDEF")
	fmt.Println("AES Decrypt:", ds)
	fmt.Println("------------------------------------------")
}
