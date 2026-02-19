package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	os.Args = []string{
		"http://localhost:10055/marks/login",
		"http://localhost:10055/marks/login",
		"http://localhost:10055/marks/login",
	}
	main()
}
