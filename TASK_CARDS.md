# Task: Empty Input Behavior

**Goal:** Keep one clear and tested behavior for empty input (`""`) so output is predictable.

**Depends on:** none

## Acceptance criteria
- [x] `PrintAscii([]string{""}, ...)` produces no output (empty string is skipped at index 0).
- [x] Related unit test matches the chosen behavior.
- [x] `go test ./...` passes.

## Tests written
- [x] `TestEmptyLine`: confirms empty input produces no output.
- [x] `TestNewlineSeparator`: confirms `\n` separators still work correctly.

## Notes
- Empty input and `"\\n"` are different cases and stay distinct.

---

# Task: CLI Input Validation

**Goal:** Make CLI failures clear and consistent for wrong arguments and style choices.

**Depends on:** none

## Acceptance criteria
- [x] Missing or extra arguments show usage guidance and exit non-zero.
- [x] Invalid style input (not `standard`, `shadow`, or `thinkertoy`) exits with clear error.
- [x] Error paths return non-zero exit code.

## Notes
- Style is passed as the second CLI argument (not an interactive prompt).
- Style matching is case-insensitive via `strings.ToLower`.

---

# Task: Character Validation

**Goal:** Reject unsupported characters safely and consistently.

**Depends on:** none

## Acceptance criteria
- [x] Only ASCII `32–126` is accepted.
- [x] Accented letters and emoji are rejected with a clear error message.
- [x] Program exits with non-zero status on invalid characters.

## Notes
- Validation happens in `main.go` before `PrintAscii` is called, so the renderer
  can safely skip bounds checks on the banner index.

---

# Task: CI Checks

**Goal:** Automatically verify formatting and tests on each push/PR.

**Depends on:** CLI Input Validation, Character Validation

## Acceptance criteria
- [x] CI runs `go test ./...`.
- [x] CI fails when tests fail.
- [x] CI checks formatting (`gofmt`).

## Notes
- Keep CI simple and fast.

---

# Task: Documentation Refresh

**Goal:** Make README, PRD, and TASK_CARDS clear and aligned with real behavior.

**Depends on:** Empty Input Behavior, CLI Input Validation

## Acceptance criteria
- [x] README shows correct 2-argument command syntax.
- [x] README explains style argument and `\n` usage.
- [x] README input rules match actual code behavior.
- [x] PRD CLI contract reflects 2-arg interface (no interactive prompt).
- [x] PRD acceptance criteria and milestones marked complete where applicable.

## Notes
- Interactive stdin style selection was removed; style is now a CLI argument.
