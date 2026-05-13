package test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"ascii-art/internal"
)

// captureOutput redirects stdout and returns whatever the function printed.
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Empty input (index 0, empty string) must produce no output.
func TestEmptyLine(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{""}, "../banners/standard.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if output != "" {
		t.Errorf("expected empty output, got %q", output)
	}
}

// A single word must produce exactly 8 lines of ASCII art.
func TestSingleWord(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{"Hi"}, "../banners/standard.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines of ASCII art, got %d", len(lines))
	}
}

// A leading \n (empty first element) is skipped to avoid a leading blank line.
func TestLeadingNewline(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{"", "Hi"}, "../banners/standard.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines (no leading blank), got %d", len(lines))
	}
}

// \n in the input separates two 8-line blocks with one blank line between them.
func TestNewlineSeparator(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{"Hi", "", "There"}, "../banners/standard.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	// "Hi" = 8 lines, "" = 1 blank, "There" = 8 lines → 17 total
	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 17 {
		t.Errorf("expected 17 lines (8 + 1 + 8), got %d", len(lines))
	}
}

// The 5th row of 'A' in the standard font contains an underscore.
func TestSpecificCharacter(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{"A"}, "../banners/standard.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(output, "\n")
	if !strings.Contains(lines[4], "_") {
		t.Errorf("expected 5th line of 'A' to contain '_', got %q", lines[4])
	}
}

// Shadow style must render without error and produce 8 lines.
func TestShadowStyle(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{"Hi"}, "../banners/shadow.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

// Thinkertoy style must render without error and produce 8 lines.
func TestThinkertoyStyle(t *testing.T) {
	output := captureOutput(func() {
		if err := internal.PrintAscii([]string{"Hi"}, "../banners/thinkertoy.txt"); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	lines := strings.Split(strings.TrimRight(output, "\n"), "\n")
	if len(lines) != 8 {
		t.Errorf("expected 8 lines, got %d", len(lines))
	}
}

// A missing banner file must return an error.
func TestBannerFileNotFound(t *testing.T) {
	err := internal.PrintAscii([]string{"Hi"}, "../banners/nonexistent.txt")
	if err == nil {
		t.Fatal("expected error for missing banner file, got nil")
	}
}

// A banner file with the wrong number of lines must return an error.
func TestCorruptBannerFile(t *testing.T) {
	f, err := os.CreateTemp(t.TempDir(), "banner-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("not a valid banner file\n")
	f.Close()

	err = internal.PrintAscii([]string{"Hi"}, f.Name())
	if err == nil {
		t.Fatal("expected error for corrupt banner file, got nil")
	}
}
