package main

import (
	"fmt"

	"github.com/hanks177/podman/v4/version"
)

func main() {
	fmt.Print(version.Version.String())
}
