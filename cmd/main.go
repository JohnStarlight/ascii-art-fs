package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/internal"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: invalid usage")
		fmt.Println("Usage: go run ./cmd \"your-text-here\" \"banner-style-here\"")
		os.Exit(1)
	}

	text := os.Args[1]
	// ToLower so the user can pass any casing ("Shadow", "SHADOW", etc.).
	filename := strings.ToLower(os.Args[2])
	if filename != "shadow" && filename != "standard" && filename != "thinkertoy" {
		fmt.Println("Error: invalid choice (must be \"Shadow\", \"Standard\" or \"Thinkertoy\")")
		os.Exit(1)
	}

	// Banner files only define characters 32–126; reject anything outside that range.
	for _, r := range text {
		if r < 32 || r > 126 {
			fmt.Printf("Error: invalid character %q (only printable ASCII is supported)\n", r)
			os.Exit(1)
		}
	}

	// The shell does not interpret \n inside double quotes, so callers pass the
	// literal two-character sequence \n to represent a line break (e.g. "Hi\nThere").
	lines := strings.Split(text, "\\n")

	if err := internal.PrintAscii(lines, "banners/"+filename+".txt"); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
