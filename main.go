// asciiart.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Error: Wrong number of arguments")
		os.Exit(1)
	}

	art, err := AsciiArt(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print(art)
}

func AsciiArt(input string) (string, error) {
	// Replace \n with actual newlines
	input = strings.ReplaceAll(input, `\n`, "\n")

	fontBytes_txt, err := os.ReadFile("standard.txt")
	if err != nil {
		return "", fmt.Errorf("read font: (\" %w \") ", err)
	}

	// solve windows problem with \r\n
	fontContent := strings.ReplaceAll(string(fontBytes_txt), "\r\n", "\n")
	// split font.txt content into lines
	fontLines := strings.Split(fontContent, "\n")

	// string.Builder better for storing strings
	var out strings.Builder
	// split input into words by newlines
	lines := strings.Split(input, "\n")

	for i, Line := range lines {
		// if the line is empty, just write a newline
		// but only if it's not the last line
		if Line == "" {
			if i < len(lines)-1 {
				out.WriteByte('\n')
			}
			continue
		}

		// this loop for drawing each word (8 lines for each word)
		for row := 0; row < 8; row++ {
			for i, char := range Line {
				if char < ' ' || char > '~' {
					return "", fmt.Errorf("unsupported character: %q", char)
				}
				idx := row + (int(char-' ') * 9) + 1
				out.WriteString(fontLines[idx])
				if i != len([]rune(Line))-1 {
					out.WriteByte(' ')
				}
			}
			out.WriteByte('\n')
		}
	}
	return out.String(), nil
}
