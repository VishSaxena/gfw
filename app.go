// app.go
package gfw

import (
	"log"
	"bytes"
	"fmt"
)

func init() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file! from app.go")
	fmt.Print(&buf)
}
