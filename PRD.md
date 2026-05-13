# PRD - ascii-art

Keep this short. The project subject + audit cases are the source of truth.

---

## 1. Problem Statement

This CLI receives text and a banner style as arguments and prints the text as ASCII art using the chosen banner template.

---

## 2. CLI Contract

- Command: `go run ./cmd "<text>" "<style>"`
- Args:
  - exactly 2 arguments required: text string and banner style
  - style is one of `standard`, `shadow`, `thinkertoy` (case-insensitive)
  - text supports letters, numbers, spaces, special chars, and literal `\n`
- Wrong number of args:
  - print usage guidance
  - exit with non-zero status
- Invalid style:
  - print clear error message
  - exit with non-zero status
- Invalid chars (outside ASCII 32–126):
  - print clear error message
  - exit with non-zero status
- Banner file read failure:
  - print `could not open banner file: ...`
  - exit with non-zero status

---

## 3. Rendering Logic

### 3.1 Style Selection

- Style is passed as the second CLI argument (case-insensitive).
- Accepted values: `standard`, `shadow`, `thinkertoy`.
- Any other value is rejected before rendering begins.

Example:
- `go run ./cmd "Hello" "standard"` → prints `Hello` in the standard font.

### 3.2 Newline Handling

- Input is split on the literal two-character sequence `\n` (backslash + n).
- Each segment is rendered as a separate 8-line ASCII art block.
- An empty segment (from consecutive `\n`) produces one blank line separator,
  except a leading empty segment (input starting with `\n`) is skipped.

Examples:
- `Hello\nThere` → two 8-line blocks separated by one blank line.
- `Hello\n\nThere` → two 8-line blocks with two blank lines between them.

### 3.3 Character-to-Glyph Mapping

- Each printable ASCII character (32–126) maps to a 8-line glyph in the banner file.
- The banner file has 95 character definitions × 9 lines each = 855 newlines total.
- For each output row (1–8), the program concatenates the matching row from each glyph.
- Index formula: character `c` occupies rows `(c-32)*9+1` through `(c-32)*9+8`.

---

## 4. Non-Goals

- Unicode support beyond printable ASCII
- GUI/Web interface
- Custom user-uploaded fonts
- Rich text features (color, alignment, animation)
- Interactive stdin prompts

---

## 5. Acceptance Criteria

### Audit Cases

- [x] `hello` prints expected ASCII-art output in `standard` style.
- [x] Mixed case and spaces (for example `HeLlo HuMaN`) render correctly.
- [x] Special characters (for example `{|}~` and punctuation sets) render correctly.
- [x] Literal `\n` and `\n\n` create correct multi-line block separation.

### Extra Golden Tests

- [x] Empty input (`""`) produces no output.
- [x] Invalid character input (for example emoji) returns clear error and non-zero exit.
- [x] Invalid style choice returns clear error and non-zero exit.
- [x] Wrong number of arguments returns usage guidance and non-zero exit.

---

## 6. Architecture

- Pattern: sequential pipeline
- Because: input validation, style parsing, splitting, and rendering happen in a simple linear flow.
- Tradeoff: less flexible than a parser/FSM, but straightforward to read and maintain.

Pipeline:

`CLI args → validate arg count → validate style → validate chars → load banner → split by \n → render each line → stdout`

---

## 7. Milestones

1. [x] Stabilize expected behavior for empty/newline edge cases.
2. [x] Ensure functional output matches subject examples for key inputs.
3. [x] Add/clean tests for validation, multiline behavior, and special chars.
4. [x] Final docs pass (README usage + examples aligned with real behavior).

---

## 8. Known Limitations

- Rendering is limited to printable ASCII (`32`–`126`).
- Input with Unicode symbols (for example Greek characters or emoji) is not supported.
