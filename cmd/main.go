package main

import (
	"fmt"

	"github.com/monochromegane/dsn"
)

func main() {
	n, s, _ := dsn.FromRailsConfig("/Users/miyakey/Documents/rails/hothub/config/database.yml", "production")
	fmt.Printf("%s -> %s\n", n, s)
}
