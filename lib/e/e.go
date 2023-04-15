// Package e keep error wrapper
package e

import "fmt"

// Wrap is wrapper error with message of error
func Wrap(msg string, err error) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// WrapIfErr called Wrap when error != nil
func WrapIfErr(msg string, err error) error {
	if err == nil {
		return nil
	}
	return Wrap(msg, err)
}
