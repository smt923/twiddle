package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	extraFlag := flag.Bool("e", false, "extra: Should we add some extra string fiddling?")
	lightFlag := flag.Bool("l", false, "light: Print just a few common substitutions (overrides -e)")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			twiddle(text, *extraFlag, *lightFlag)
		}
	}
}

func twiddle(t string, extra bool, light bool) {
	if !light {
		fmt.Println(t + t)
		fmt.Println(t + "0")
		fmt.Println(t + "1")
		if extra {
			fmt.Println(strings.ToUpper(t))
			fmt.Println(strings.ToUpper(string(t[0])) + t[1:])
			fmt.Println(strings.ToUpper(string(t[0])) + t)
			fmt.Println(string(t[len(t)-1]) + t[:len(t)-1])
			fmt.Println(string(t[len(t)-1]) + string(t[len(t)-1]) + t)
			fmt.Println(string(t[len(t)-1]) + string(t[len(t)-1]) + t[:len(t)-1])
			fmt.Println(string(t[len(t)-1]) + t + "0")
			fmt.Println(string(t[len(t)-1]) + t + "1")
			if len(t) > 1 {
				fmt.Println(t[:2] + t)
			}
		}
	}
	fmt.Println(t)
	fmt.Println(string(t[0]) + t)        // aliceh -> aaliceh
	fmt.Println(string(t[len(t)-1]) + t) // aliceh -> haliceh
}
