package testhash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// Md5_1 MD5的第一种实现
func Md5_1(s string) {
	m := md5.New()
	m.Write([]byte(s))
	fmt.Println(hex.EncodeToString(m.Sum(nil)))
}

func Md5_2(s string) {
	m := md5.Sum([]byte(s))
	fmt.Println(hex.EncodeToString(m[:]))
}

func Md5_3(s string) {
	m := md5.Sum([]byte(s))
	fmt.Printf("%x\n", m)
}

func Md5_4(s string) {
	m := md5.New()
	io.WriteString(m, s)
	fmt.Println(hex.EncodeToString(m.Sum(nil)))
}
