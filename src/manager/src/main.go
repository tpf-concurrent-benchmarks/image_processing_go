package main

import (
	"fmt"
	"shared/config"
)

func main() {
	managerConfig := config.GetConfig()
	fmt.Println(managerConfig)
}
