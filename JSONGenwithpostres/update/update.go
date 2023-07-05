package update

import (
	"hz_gen"
	"log"
	"os"
)

func main() {
	err := os.Chdir("../")
	if err != nil {
		log.Fatalf("move to directory folder failed with %s\n", err)
	}

	hz_gen.ClearGateway()

}
