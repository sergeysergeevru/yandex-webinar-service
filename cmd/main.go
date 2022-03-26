package main

import (
	"fmt"
	privatePkg "github.com/sergeysergeevru/yandex-webinar-private-pkg"
)

var version string = "v0.0.0"

func main() {
	fmt.Printf("the current version is %s\n", version)
	fmt.Println(privatePkg.SayHello())
}

