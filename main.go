// asciiart.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: asciiart <text>")
		os.Exit(1)
	}

	art, err := AsciiArt(os.Args[1])
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	fmt.Print(art)
}

// AsciiArt يرسم النص المدخَل بخط standard.txt
func AsciiArt(input string) (string, error) {
	// 1) حوِّل السلسلة "\n" النصّية إلى محرف سطر جديد فعلي
	input = strings.ReplaceAll(input, `\n`, "\n")

	// 2) تحميل الخط
	fontBytes, err := os.ReadFile("standard.txt")
	if err != nil {
		return "", fmt.Errorf("read font: %w", err)
	}
	fontLines := strings.Split(string(fontBytes), "\n")

	var out strings.Builder
	words := strings.Split(input, "\n") // الأسطر الفعلية بعد التحويل

	for _, word := range words {
		if word == "" { // سطر فارغ → سطر واحد فارغ في الخرج
			out.WriteByte('\n')
			continue
		}

		for row := 0; row < 8; row++ { // طباعة الصفوف 0..7
			for _, ch := range word { // طباعة أحرف السطر الحالي
				if ch < ' ' || ch > '~' {
					return "", fmt.Errorf("unsupported character: %q", ch)
				}
				idx := row + (int(ch-' ')*9) + 1
				out.WriteString(fontLines[idx])
			}
			out.WriteByte('\n')
		}
	}
	return out.String(), nil
}
