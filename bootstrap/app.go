package bootstrap

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func InitApplication() {
	fmt.Print(`
      _                 _
  ___| | ___  _   _  __| | __ _  ___
 / __| |/ _ \| | | |/ _  |/ _  |/ _ \
| (__| | (_) | |_| | (_| | (_| | (_) |
 \___|_|\___/ \__,_|\__,_|\__, |\___/
                          |___/
`)
	fmt.Println("init Goroutine ID:", GetGid())
	go checkUpdate()
}

func checkUpdate() {
	fmt.Println("checkUpdate Goroutine ID:", GetGid())
}

func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}
