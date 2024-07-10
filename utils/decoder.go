package utils

import (
	"hash/crc32"
	"strconv"
)

func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func CheckHashCode(crc32Code string) string {
	for i := 10000000; i < 1000000000; i++ {
		if strconv.FormatUint(uint64(Crc32(strconv.Itoa(i))), 16) == crc32Code {
			return strconv.Itoa(i)
		}
	}
	return "Unknown"
}
