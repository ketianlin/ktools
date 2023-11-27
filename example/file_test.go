package example

import (
	"fmt"
	"github.com/ketianlin/ktools"
	"testing"
)

func TestFile(t *testing.T) {
	selfPath := ktools.File.SelfPath()
	fmt.Println("selfPath:", selfPath)

	selfDir := ktools.File.SelfDir()
	fmt.Println("selfDir:", selfDir)

	fmt.Println("-------------------------")

	src := "/home/ke666/my_codes/sj_go_codes/ktools/example/file_test.go"

	basename := ktools.File.GetBasename(src)
	fmt.Println("Basename:", basename)

	dir := ktools.File.GetDir(src)
	fmt.Println("Dir:", dir)

	ext := ktools.File.GetExt(src)
	fmt.Println("Ext:", ext)

	fmt.Println("-------------------------")

	mTime, err := ktools.File.GetMTime(src)
	if err != nil {
		fmt.Println("GetMTime err:", err)
	}
	fmt.Println("MTime:", mTime)

	size, err := ktools.File.GetSize(src)
	if err != nil {
		fmt.Println("GetSize err:", err)
	}
	fmt.Println("Size:", size)

	fmt.Println("-------------------------")

	isExist := ktools.File.IsExist(src)
	fmt.Println("IsExist:", isExist)

	isFile := ktools.File.IsFile(src)
	fmt.Println("IsFile:", isFile)

	isDir := ktools.File.IsDir(src)
	fmt.Println("IsDir:", isDir)

	fmt.Println("-------------------------")

	contentBytes, err := ktools.File.ReadBytes(src)
	if err != nil {
		fmt.Println("ReadBytes err:", err)
	}
	fmt.Println("Content:", string(contentBytes))

	contentString, err := ktools.File.ReadString(src)
	if err != nil {
		fmt.Println("ReadString err:", err)
	}
	fmt.Println("Content:", contentString)

	fmt.Println("-------------------------")

	contentLines, err := ktools.File.ReadLines(src)
	if err != nil {
		fmt.Println("ReadLines err:", err)
	}
	fmt.Println("Content:", contentLines)

	fmt.Println("-------------------------")

	_, err = ktools.File.WriteString(dir+"/hello1.txt", "hello world")
	if err != nil {
		fmt.Println("WriteString err:", err)
	}

	fmt.Println("-------------------------")

	_, err = ktools.File.WriteBytes(dir+"/hello2.txt", []byte("hello world"))
	if err != nil {
		fmt.Println("WriteBytes err:", err)
	}

	err = ktools.File.Remove(dir + "/hello1.txt")
	if err != nil {
		fmt.Println("Remove err:", err)
	}

	err = ktools.File.Remove(dir + "/hello2.txt")
	if err != nil {
		fmt.Println("Remove err:", err)
	}

	fmt.Println("-------------------------")
}
