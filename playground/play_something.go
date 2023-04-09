package main

import (
	"fmt"

	"github.com/yunerou/oauth2-client/singleton"
)

func init() {
	// Read config file
	wViper := singleton.GetViper()
	wViper.LoadConfigFile([]string{"..", "."}, "config")
}

// main ... example.
func main() {
	fmt.Println("--- playing ---")

}
