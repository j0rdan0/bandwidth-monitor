package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/j0rdan0/bandwidth"
)

func main() {
	color.Set(color.FgRed)
	fmt.Print("[*]")
	color.Unset()

	fmt.Println(" Started collecting data")
	for {
		s := bandwidth.ComposeData()
		color.Set(color.FgRed)
		fmt.Print("[*]")
		color.Unset()
		fmt.Println(" Collected data")
		bandwidth.PrintData(s)
		bandwidth.Plot()

		color.Set(color.FgRed)
		fmt.Print("[*]")
		color.Unset()
		fmt.Println(" Created plots")
		bandwidth.CreateFile(s)
	}
}
