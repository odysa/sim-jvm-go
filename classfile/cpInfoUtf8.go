package classfile

import (
	"fmt"
	"unicode/utf16"
)

type constantUTF8Info struct {
	val string
}

func (c *constantUTF8Info) readInfo(reader *classReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	c.val = decodeMUTF8(bytes)
}

func decodeMUTF8(bytearr []byte) string {
	utfLen := len(bytearr)
	charArr := make([]uint16, utfLen)

	var c, char2, char3 uint16
	count := 0
	charArrCount := 0

	for count < utfLen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		charArr[charArrCount] = c
		charArrCount++
	}

	for count < utfLen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			charArr[charArrCount] = c
			charArrCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			charArr[charArrCount] = c&0x1F<<6 | char2&0x3F
			charArrCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utfLen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			charArr[charArrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			charArrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	charArr = charArr[0:charArrCount]
	runes := utf16.Decode(charArr)
	return string(runes)
}
