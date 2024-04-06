package main

import (
	"github.com/sqkam/codejournal/freecar"
)

func main() {
	freecar.GetIP()
	_, _ = freecar.GetOutboundIP()
}
