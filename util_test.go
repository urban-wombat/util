package util

import (
	"testing"
)

func TestFormatSource(t *testing.T) {
	source := `
	if err != nil { return nil }
	`

	expected := `
	if err != nil {
		return nil
	}
	`

	formatted, err := FormatSource(source)
	if err != nil {
		t.Error(err)
	}

	if formatted != expected {
		t.Errorf("expecting %s but got %s", expected, formatted)
	}
}
