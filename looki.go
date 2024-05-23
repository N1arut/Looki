package main

import (
	"fmt"
	"net"
	"os"

	"github.com/fatih/color"
	"github.com/projectdiscovery/cdncheck"
)

func main() {
	count := len(os.Args)
	if count < 2 {
		fmt.Printf("Usage: %s [HostName]\n", os.Args[0])
		os.Exit(0)
	}
	name, err := net.LookupIP(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	for index, ip := range name {
		if ipv4 := ip.To4(); ipv4 != nil {
			client := cdncheck.New()
			resp, kindof, ty, err := client.Check(ipv4)
			if err != nil {
				fmt.Printf("Error: %s", err)
				continue
			}
			if resp {
				output := fmt.Sprintf("[%d] %s %s %s",
					index+1,
					color.CyanString("%v", ipv4),
					color.RedString("%v", kindof),
					color.YellowString("%v", ty),
				)
				fmt.Println(output)
			}

		}
	}
}
