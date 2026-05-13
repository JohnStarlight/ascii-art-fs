# ASCII Art CLI

A simple Go command-line tool that turns text into ASCII art.

## What This Project Does

- Takes a text string and a banner style as CLI arguments
- Converts each character to ASCII art (8 lines tall)
- Supports 3 styles: `standard`, `shadow`, `thinkertoy`
- Supports multi-line output via the literal `\n` sequence

## Quick Start

From the project root:

```bash
go run ./cmd "Hello" "standard"
```

Always pass both arguments inside quotes.

## Usage

```bash
go run ./cmd "<text>" "<style>"
```

- `<text>` — the string to render (printable ASCII only, use `\n` for line breaks)
- `<style>` — one of `standard`, `shadow`, or `thinkertoy` (case-insensitive)

### Examples

Single line:

```bash
go run ./cmd "Hello World" "shadow"
```

Multi-line (literal `\n` in the argument):

```bash
go run ./cmd "Hello\nThere" "thinkertoy"
```

## Input Rules

- Exactly 2 arguments are required (text and style)
- Only printable ASCII characters are accepted (`32` to `126`)
- Non-ASCII characters (for example `é` or emoji) are rejected with a clear error
- The literal sequence `\n` (backslash + n) creates a line break in the output
- Style name is case-insensitive (`Shadow`, `SHADOW`, and `shadow` all work)

## Run Tests

```bash
go test ./...
```

With verbose output:

```bash
go test ./... -v
```

Force a fresh run (skip cache):

```bash
go test ./... -count=1
```

## Project Structure

```
cmd/main.go                    — CLI entry point: arg parsing and validation
internal/printascii.go         — ASCII rendering logic
banners/standard.txt           — standard banner font
banners/shadow.txt             — shadow banner font
banners/thinkertoy.txt         — thinkertoy banner font
test/printascii_test.go        — core unit tests
test/audit_examples_test.go    — audit/subject example tests
```

## Known Limitations

- Only printable ASCII (`32`–`126`) is supported; Unicode is rejected
- No color, alignment, or animation support

## License

MIT. See `LICENSE`.
