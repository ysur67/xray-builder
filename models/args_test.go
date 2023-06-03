package models_test

import (
	"os"
	"strings"
	"testing"
)

func split(s string) []string {
	return strings.Split(s, " ")
}

func TestBaseArgumentsParsing(t *testing.T) {
	// var args models.InstallArgs
	os.Args = split("-d gfd --help")
	// arg.MustParse(&args)
}
