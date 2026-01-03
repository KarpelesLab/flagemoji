package flagemoji_test

import (
	"testing"

	"github.com/KarpelesLab/flagemoji"
)

// TestFlag verifies that the Get function correctly converts
// an ISO 3166-1 alpha-2 country code to its emoji flag representation.
func TestFlag(t *testing.T) {
	// "FR" (France) should produce the French flag emoji (🇫🇷),
	// which consists of Regional Indicator Symbol Letters F and R.
	if flagemoji.Get("FR") != "🇫🇷" {
		t.Errorf("Tried to get flag for France, got %s", flagemoji.Get("FR"))
	}
}
