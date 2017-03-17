package prompt

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/howeyc/gopass"
)

// Writer is the default writer for the package.
var Writer io.Writer = os.Stdout

// String prompt.
func String(prompt string, args ...interface{}) string {
	fmt.Fprintf(Writer, prompt+": ", args...)
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	return string(bytes)
}

// MustString reads a string from the input until a non empty one is
// provided.
func MustString(prompt string, args ...interface{}) (s string) {
	for strings.Trim(s, " ") == "" {
		s = String(prompt, args...)
	}
	return s
}

// Float64 reads a string from the input and parses it as a float64.
func Float64(prompt string, args ...interface{}) (float64, error) {
	s := String(prompt, args...)

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}

	return f, nil
}

// MustFloat64 reads a string from the input and parses it as a float64 until it
// is valid.
func MustFloat64(prompt string, args ...interface{}) float64 {
	for {
		f, err := Float64(prompt, args...)
		if err != nil {
			fmt.Fprintln(Writer, err.Error())
			continue
		}

		return f
	}
}

// Confirm continues prompting until the input is boolean-ish.
func Confirm(prompt string, args ...interface{}) bool {
	for {
		switch String(prompt, args...) {
		case "Yes", "yes", "y", "Y":
			return true
		case "No", "no", "n", "N":
			return false
		}
	}
}

// Choose prompts for a single selection from `list`, returning in the index.
func Choose(prompt string, list []string) int {
	fmt.Fprintln(Writer)
	for i, val := range list {
		fmt.Fprintf(Writer, "  %d) %s\n", i+1, val)
	}

	fmt.Fprintln(Writer)
	i := -1

	for {
		s := String(prompt)

		// index
		n, err := strconv.Atoi(s)
		if err == nil {
			if n > 0 && n <= len(list) {
				i = n - 1
				break
			} else {
				continue
			}
		}

		// value
		i = indexOf(s, list)
		if i != -1 {
			break
		}
	}

	return i
}

// Password prompt.
func Password(prompt string, args ...interface{}) string {
	fmt.Fprintf(Writer, prompt+": ", args...)
	password, _ := gopass.GetPasswdPrompt("", false, os.Stdin, Writer)
	s := string(password[0:])
	return s
}

// PasswordMasked is a password prompt with mask.
func PasswordMasked(prompt string, args ...interface{}) string {
	fmt.Fprintf(Writer, prompt+": ", args...)
	password, _ := gopass.GetPasswdPrompt("", true, os.Stdin, Writer)
	s := string(password[0:])
	return s
}

// index of `s` in `list`.
func indexOf(s string, list []string) int {
	for i, val := range list {
		if val == s {
			return i
		}
	}
	return -1
}
