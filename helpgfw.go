// helpgfw
package gfw

import (
	"log"
	"bytes"
	"fmt"
)

func init() {
	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file! from helpgfw.go")
	fmt.Print(&buf)
}
