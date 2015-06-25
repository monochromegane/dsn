package main

import (
	"fmt"

	"github.com/monochromegane/dsn"
)

func main() {
	s, _ := dsn.FromRailsConfig("/Users/miyakey/Documents/rails/hothub/config/database.yml", "production")
	fmt.Printf("%s\n", s)
}
