# ASCII Art

A command-line tool written in Go that turns any text into large ASCII-art banners using a bitmap-style font file.

![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white)
![CLI](https://img.shields.io/badge/type-CLI-informational)

```text
$ go run . "Hi there"
 _    _    _           _      _
| |  | |  (_)         | |    | |
| |__| |   _          | |_   | |__      ___    _ __     ___
|  __  |  | |         | __|  |  _ \    / _ \  | '__|   / _ \
| |  | |  | |         \ |_   | | | |  |  __/  | |     |  __/
|_|  |_|  |_|          \__|  |_| |_|   \___|  |_|      \___|
```

## Overview

Each printable ASCII character (codes 32–126) is drawn as an 8-line glyph stored in a banner file. The program reads the input string, looks up every character's glyph and prints them side by side, line by line, to build the final banner.

## Features

- Renders any printable ASCII text as an 8-line banner
- Supports multi-line output using the `\n` escape sequence (`"Hello\nWorld"`)
- Ships with three banner fonts: `standard`, `shadow` and `thinkertoy`
- Handles both Unix (`LF`) and Windows (`CRLF`) font files
- Clear error messages for invalid arguments or missing font files

## Getting Started

**Requirements:** Go 1.20 or newer

```bash
git clone https://github.com/ahmed-jawher/ascii-art.git
cd ascii-art
go run . "Hello"
```

### Examples

```bash
go run . "Hello"            # single line
go run . "Hello\nThere"     # two lines
go run . ""                 # prints nothing
```

## Project Structure

```text
ascii-art/
├── main.go          # argument handling and rendering logic
├── standard.txt     # default banner font
├── shadow.txt       # alternative font
└── thinkertoy.txt   # alternative font
```

## How It Works

1. Replace literal `\n` sequences in the input with real line breaks.
2. Load the banner file and split it into lines (normalizing `\r\n`).
3. For each character, compute its glyph offset: `(char - 32) * 9 + 1`.
4. Print the 8 rows of every glyph next to each other, one input line at a time.

## Related Projects

This project is the foundation for [ascii-art-web](https://github.com/ahmed-jawher/ascii-art-web), which puts the same renderer behind a web interface.

## Author

**Ahmed Jawher** — [@ahmed-jawher](https://github.com/ahmed-jawher)
