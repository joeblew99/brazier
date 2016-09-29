package bind

import (
	"log"

	"github.com/asdine/brazier/cli"
)

// Bind ...
func Bind() {
	cmd := cli.New()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
