package main

import (
	"fmt"

	"github.com/leofigy/valk/windows"
)

func main() {
	dll := windows.NewUser32()
	fmt.Println(dll)
}
