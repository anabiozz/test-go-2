package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

func stringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

func read_int32(data []byte) (ret int32) {
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &ret)
	return
}

func bytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func intToStrings(values int32) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&values))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func intToString(values int) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%d", values)
	return buf.String()
}

func ip2str(ip int32) (result string) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, ip)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	byteArray := buf.Bytes()
	for i, b := range byteArray {
		result += fmt.Sprintf("%d", b)
		if i < len(byteArray)-1 {
			result += "."
		}
	}

	return result
}

func main() {

	fmt.Println(ip2str(0x7f000001))

	fmt.Printf("%d\n", intToStrings(56446))

	/****************************************/

	// fmt.Println(read_int32([]byte{0x7F, 0x00, 0x00, 0x01}))

	// /****************************************/

	// fmt.Println(stringToBin("≠"))

	// /***************************************/
	// byteArray1 := []byte{'J', 'O', 'H', 'N'}
	// str1 := bytesToString(byteArray1)
	// fmt.Println("String:", str1)

	// /****************************************/

	// str2 := string(byteArray1[:])
	// fmt.Println("String:", str2)

	// /****************************************/

	// str3 := bytes.NewBuffer(byteArray1).String()
	// fmt.Println("String:", str3)

	// /****************************************/

	// fmt.Println(intToString(56484648))
}
