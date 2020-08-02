package testhash

import (
	"fmt"
	"hash/crc32"
)

func Crc32_1(s string) {
	fmt.Println(crc32.ChecksumIEEE([]byte(s)))
}

func Crc32_2(s string) {
	tab := crc32.MakeTable(crc32.IEEE)
	fmt.Println(crc32.Checksum([]byte(s), tab))
}
