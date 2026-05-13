package internal

import (
	"fmt"
	"os"
	"strings"
)

// PrintAscii renders each line of text as ASCII art using the given banner file.
// Each character maps to an 8-row block in the banner file, where blocks are
// separated by a blank line (9 lines per character total).
// An empty string in lines prints a blank line; the first element is skipped if
// empty to avoid a leading blank line when the original input starts with \n.
func PrintAscii(lines []string, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not open banner file: %w", err)
	}

	// 95 printable ASCII characters × 9 lines each = 855 newlines.
	// The file must also end with a trailing newline, giving 856 split elements.
	if strings.Count(string(data), "\n") != 855 {
		return fmt.Errorf("banner file is corrupt or invalid (expected 856 lines, got %d)", strings.Count(string(data), "\n")+1)
	}

	// Normalise \r\n (Windows) to \n so both line-ending styles work.
	bannerLines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	for i, line := range lines {
		// An empty string means the input had a \n at this position.
		// Skip index 0 to avoid a leading blank line when input starts with \n.
		if line == "" {
			if i > 0 {
				fmt.Println()
			}
			continue
		}

		// Row 0 of each character block is the blank separator; art rows are 1–8.
		// Index formula: character c occupies rows (c-32)*9+1 through (c-32)*9+8.
		// No bounds check needed: main.go guarantees c is 32–126, and the file
		// length check above guarantees bannerLines has at least 856 elements,
		// so the max possible index (126-32)*9+8 = 854 is always in range.
		for row := 1; row <= 8; row++ {
			var sb strings.Builder
			for _, r := range line {
				sb.WriteString(bannerLines[(int(r)-32)*9+row])
			}
			fmt.Println(sb.String())
		}
	}

	return nil
}
