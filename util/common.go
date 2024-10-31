package util

import (
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

func RelativePath(name string) string {
	if filepath.IsAbs(name) {
		return name
	}
	e, _ := os.Executable()
	return filepath.Join(filepath.Dir(e), name)
}

func Exists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return !fi.IsDir()
}

/*
*
替换content的某些字段

	result := util.Replace(map[string]string{
	"xxx": "111",
	"yyy": "222",
	}, "xxx + yyy")

output: 111 + 222
*/
func Replace(table map[string]string, content string) string {
	for k, v := range table {
		content = strings.Replace(content, k, v, -1)
	}
	return content
}

/*
*
util.RandStringRunes(5)
output:a6hro
*/
func RandStringRunes(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !Exists(basePath) {
		err := os.MkdirAll(basePath, 0700) //权限，默认0777
		if err != nil {
			Log().Error("Failed to create nested file directory: %s", err)
			return nil, err
		}
	}
	return os.Create(path)
}
