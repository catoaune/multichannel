package multichannel_test

import (
	"testing"

	"github.com/catoaune/multichannel"
)

func TestConfig(t *testing.T) {
	expected := 0
	result := multichannel.Config()
	if len(result) != expected {
		t.Errorf("Expected %d but got %d.", expected, len(result))
	}
}
