package bootstrap

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
)

func InitApplication() {
	fmt.Print(`

 / ___| | ___  _   _  __| |/ ___| ___/ ___|(_) __ _| |__ | |_
| |   | |/ _ \| | | |/ _  | |  _ / _ \___ \| |/ _  | '_ \| __|
| |___| | (_) | |_| | (_| | |_| | (_) |__) | | (_| | | | | |_
 \____|_|\___/ \__,_|\__,_|\____|\___/____/|_|\__, |_| |_|\__|
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
