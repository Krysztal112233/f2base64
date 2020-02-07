package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	a := os.Args
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	newSlice := make([]string, 0)
	if len(a) == 1 {
		return
	}
	for _, v := range a[1:] {
		if string(v[0]) == "/" {
			fmt.Println("is /")
			newSlice = append(newSlice, v)
		} else if string(v[1]) == ":" {
			newSlice = append(newSlice, v)
		} else {
			newSlice = append(newSlice, pwd+"/"+v)
		}
	}
	filePointerSlice := make([]*os.File, 0)
	for _, v := range newSlice {
		fp, err := os.Open(v)
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			filePointerSlice = append(filePointerSlice, fp)
		}
	}
	for _, v := range filePointerSlice {
		fmt.Println("读取文件:", v.Name(), "\n")
		b, err := ioutil.ReadAll(v)
		if err != nil {
			fmt.Println("读取失败...\n")
		} else {
			fmt.Println(base64.StdEncoding.EncodeToString(b) + "\n")
		}
	}
}
