package main

import (
	"hexlet-go/greeting"

	"github.com/fatih/color"
)

func main() {
    msg := color.New(color.Bold, color.FgGreen).PrintlnFunc()
    msg(greeting.Hello())
}
