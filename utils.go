package dataCrud

import (
	"errors"
	"io"
	"os"
	"strings"
)

func DirOrFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

//创建目录
func CreateDir(path string) bool {
	if DirOrFileExist(path) == false {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false
		}
	}
	return true
}

func GenerateDir(path string) (string, error) {
	if len(path) == 0 {
		return "", errors.New("目录为空")
	}
	last := path[len(path)-1:]
	if !strings.EqualFold(last, string(os.PathSeparator)) {
		path = path + string(os.PathSeparator)
	}
	if !IsDir(path) {
		if CreateDir(path) {
			return path, nil
		}
		return "", errors.New(path + "目录创建失败")
	}
	return path, nil
}

func WriteFile(filename string, data string) (count int, err error) {
	var f *os.File
	if DirOrFileExist(filename) == false {
		f, err = os.Create(filename)
		if err != nil {
			return
		}
	} else {
		f, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	}
	defer f.Close()
	count, err = io.WriteString(f, data)
	if err != nil {
		return
	}
	return
}


func UnderlineToHump(str string) string {
	temps := strings.Split(str, "_")
	var tempStr string
	for _, temp := range temps {
		tempStr += strings.ToUpper(string(temp[0]))
		tempStr += temp[1:]
	}
	return tempStr
}