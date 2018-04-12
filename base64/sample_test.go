package doc_base64

import (
	"encoding/base64"
	"io"
	"os"
	"testing"
)

type Buf struct {
	data []byte
	size int
}

func Test_ed(t *testing.T) {
	//StdEncoding Encode
	var data string = "I am learning Go!"
	str := base64.StdEncoding.EncodeToString([]byte(data))
	//str="SSBhbSBsZWFybmluZyBHbyE="

	//StdEncoding Decode
	byteArray, _ := base64.StdEncoding.DecodeString(str)
	str = string(byteArray)
	//str = "I am learning Go!"

	//URLEncoding Encode
	var data string = "I am learning Go!"
	str := base64.URLEncoding.EncodeToString([]byte(data))
	//str="SSBhbSBsZWFybmluZyBHbyE="
	fmt.Println(str)

	//URLEncoding Decode
	byteArray, _ := base64.URLEncoding.DecodeString(str)
	str = string(byteArray)
	//str = "I am learning Go!"
	fmt.Println(str)

	//RawStdEncoding Encode
	var data string = "I am learning Go!"
	str := base64.RawStdEncoding.EncodeToString([]byte(data))
	//str="SSBhbSBsZWFybmluZyBHbyE"
	fmt.Println(str)

	//RawStdEncoding Decode
	byteArray, _ := base64.RawStdEncoding.DecodeString(str)
	str = string(byteArray)
	//str = "I am learning Go!"
	fmt.Println(str)
}

//流式操作
func Test_ioed(t *testing.T) {
	f, err := os.Open("E:\\testfile.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	fEnc, err := os.Create("E:\\testfile_enc.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer fEnc.Close()

	w := base64.NewEncoder(base64.StdEncoding, fEnc)
	io.Copy(w, f)
	w.Close()
}

//自定义编码集
func Test_selfSet(t *testing.T) {
	//与标准的集合不同的在于大小写位置反了
	const CodingSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
	encoder := base64.NewEncoding(CodingSet)

	var data string = "I am learning Go!"
	str := encoder.EncodeToString([]byte(data))
	fmt.Println(str)
	//str=ssbHBsbSzwfYBMLUzYbhBYe=
}
