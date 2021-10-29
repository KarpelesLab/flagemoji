package flagemoji_test

import (
	"testing"

	"github.com/KarpelesLab/flagemoji"
)

func TestFlag(t *testing.T) {
	if flagemoji.Get("FR") != "🇫🇷" {
		t.Errorf("Tried to get flag for France, got %s", flagemoji.Get("FR"))
	}
}
