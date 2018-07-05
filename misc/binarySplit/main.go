package main

import (
	"log"
)

func binarySplit(src []byte) []int {

	sourceLen := len(src)

	log.Printf("len is %d", sourceLen)

	if sourceLen%4 != 0 {
		log.Println("return nil")
		return nil
	}

	resultLen := sourceLen / 4
	result := make([]int, resultLen)

	for i := 0; i < resultLen; i++ {
		offset := i * 4

		// log.Printf("offset is %d", offset)

		b := src[offset : offset+4]

		for _, res := range b {
			log.Printf("b is %#v", int(res))
		}

		result[i] = int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])
	}

	return result
}

func main() {
	log.Printf("%#v", binarySplit([]byte("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")))

	// const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	// for i := 0; i < len(sample); i++ {
	// 	log.Printf("%+q \n", sample[i])
	// }
	// fmt.Println(sample)
}
