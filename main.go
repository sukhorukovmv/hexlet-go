package main

import (
	"hexlet-go/v2/greeting"

	"github.com/fatih/color"
)

func main() {
    msg := color.New(color.Bold, color.FgGreen).PrintlnFunc()
    msg(greeting.Hello())
}
