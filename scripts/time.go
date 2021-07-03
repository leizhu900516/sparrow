package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)
import "fmt"
import "bytes"
//判断文件是否存在
func CheckFileState(filepath string)  bool{
	_,err := os.Stat(filepath)
	if err != nil{
		return true
	}
	return false
}
func main(){
	nowDate :=  time.Now().Format("2006-01-02")
	fmt.Println(nowDate)
	var DataBuffer bytes.Buffer
	DataBuffer.WriteString("helo")
	DataBuffer.WriteString("work")
	fmt.Println(DataBuffer.String())

	filename := "/Users/chenhuachao/go-dir/sparrow/scripts/enc.go"
	f,err := os.Stat(filename)
	fmt.Println(f,err)
	arr := strings.Split("a.gpg",".")
	fmt.Println(arr)
	fmt.Println(CheckFileState(filename))


	s := `[1, 2, 3, 4]`
	var a []int
	// 将字符串反解析为数组
	json.Unmarshal([]byte(s), &a)
	fmt.Println(a)  // [1 2 3 4]
	aa,err:= json.Marshal(a)
	fmt.Println(string(aa))
}
