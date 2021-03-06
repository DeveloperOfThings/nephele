// Package utils contains all the utility functions used by nephele
package utils

import (
	"fmt"

	"gopkg.in/cheggaaa/pb.v1"

	"github.com/bharath-srinivas/nephele/internal/colors"
)

// constants for string color wrapping
const (
	escape = "\x1b"
	reset  = 0
)

// ColorString wraps the provided string with appropriate color according to its state and returns the colored string.
// It should be used only for highlighting instance states and not recommended for any other purpose.
func ColorString(str string) string {
	var prefix, suffix string
	if str == "running" || str == "available" {
		prefix = fmt.Sprintf("%s[%dm", escape, colors.Green)
		suffix = fmt.Sprintf("%s[%dm", escape, reset)
		return prefix + str + suffix
	} else if str == "stopped" {
		prefix = fmt.Sprintf("%s[%dm", escape, colors.Red)
		suffix = fmt.Sprintf("%s[%dm", escape, reset)
		return prefix + str + suffix
	}
	prefix = fmt.Sprintf("%s[%dm", escape, colors.Yellow)
	suffix = fmt.Sprintf("%s[%dm", escape, reset)
	return prefix + str + suffix
}

// GetProgressBar returns an instance of ProgressBar with predefined config.
func GetProgressBar(totalSize int) *pb.ProgressBar {
	progressBar := pb.New(totalSize).SetUnits(pb.U_BYTES)
	progressBar.ShowPercent = true
	progressBar.ShowBar = true
	progressBar.ShowTimeLeft = true
	progressBar.ShowSpeed = true
	return progressBar
}

// WordWrap wraps the given string according to the provided parts with the separator sep and returns the wrapped
// string if and only if the given string has the character `.` or `-`. Currently WordWrap is very naive and it'll
// break the string if the separator position is greater than the half length of the provided string. It has been
// written solely for the purpose of wrapping text for rendering in table writer and not recommended for normal use.
func WordWrap(s string, sep byte, parts int) string {
	var wrapped []byte

	if parts <= 0 || !hasSeparator(s) {
		return s
	}

	halfLength := len(s) / parts
	if halfLength <= 10 {
		return s
	}

	broken := false
	for i, char := range s {
		wrapped = append(wrapped, byte(char))
		if char == rune(sep) && !broken {
			if i >= halfLength {
				wrapped = append(wrapped, byte('\n'))
				broken = true
			}
		}
	}

	return string(wrapped)
}

// hasSeparator is a helper function for WordWrap which will return true if the given string has any one of the
// `.` or `-` separator.
func hasSeparator(s string) bool {
	for _, c := range s {
		if c == '.' || c == '-' {
			return true
		}
	}
	return false
}
