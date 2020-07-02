package util

import (
	"os"

	"github.com/jenkins-x/jx-logging/pkg/log"
)

func init() {
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)
}

// SetLevel sets the logging level
func SetLevel(s string) error {
	err := log.SetLevel(s)
	if err != nil {
		return err
	}
	return nil
}
