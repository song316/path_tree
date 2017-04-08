// ┏ ┳ ┓
// ┣ ╋ ┫
// ┗ ┻ ┛
// ┃ ━
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	tab            = " "
	namePrefix     = " "
	tabRepeatTimes = 2
	leftDown       = "┗"
	midCenter      = "┣"
	horizontalLine = "━━"
	verticalLine   = "┃"
	showHide       = true
)

func main() {
	//获取目录参数
	if len(os.Args) == 1 {
		fmt.Println("pleace input the dir")
		os.Exit(1)
	}
	rootPath := os.Args[1]
	if !isDir(rootPath) {
		fmt.Println("path error!")
	}
	fmt.Println(filepath.Base(rootPath))
	listFolder(rootPath, "")
}

//递归遍历目录
func listFolder(path string, begin string) {
	files, _ := ioutil.ReadDir(path)
	for index, file := range files {
		if !showHide && strings.Index(file.Name(), ".") == 0 {
			continue
		}
		var isLast bool
		if index == len(files)-1 {
			isLast = true
		}
		var split string
		if isLast {
			split = leftDown
		} else {
			split = midCenter
		}
		if file.IsDir() {
			fmt.Println(begin + split + horizontalLine + namePrefix + file.Name())
			var newBegin string
			if isLast {
				//如果是最后一个文件夹,多加一个namePrefix为了补齐
				newBegin = begin + strings.Repeat(tab, tabRepeatTimes) + strings.Repeat(namePrefix, 2)
			} else {
				newBegin = begin + verticalLine + strings.Repeat(tab, tabRepeatTimes) + namePrefix
			}
			listFolder(path+"/"+file.Name(), newBegin)
		} else {
			fmt.Println(begin + split + horizontalLine + namePrefix + file.Name())
		}
	}
}

//判断path参数是否是目录.
func isDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return fi.Mode().IsDir()
}
