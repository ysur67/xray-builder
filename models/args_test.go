package models_test

import (
	"os"
	"strings"
	"testing"
	"xraybuilder/models"

	"github.com/alexflint/go-arg"
)

func split(s string) []string {
	return strings.Split(s, " ")
}

func TestBaseArgumentsParsing(t *testing.T) {
	var args models.InstallArgs
	// os.Args = split("install -i -u 12")
	os.Args = split("-d gfd --help")
	arg.MustParse(&args)
}
