// gfw project gfw.go
package gfw

import (
	_ "net/http"
	_ "os"
	_ "strings"
	"log"
	"bytes"
	"fmt"
)

const VERSION = "0.0.1"

//func Run(params ...string) {
//	if len(params) > 0 && params[0] != "" {
//		strs := strings.Split(params[0], ":")
//		if len(strs) > 0 && strs[0] != "" {
//			http.H
//		}
//	}
//}

func init() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file! from gfw.go")
	fmt.Print(&buf)
}