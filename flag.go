// Package flagemoji provides a minimal utility for converting ISO 3166-1 alpha-2
// country codes into their corresponding Unicode emoji flag representations.
package flagemoji

// Get converts an ISO 3166-1 alpha-2 country code (e.g., "US", "FR", "JP") into
// its corresponding emoji flag. It returns an empty string if the input is not
// exactly 2 characters.
//
// The conversion works by mapping each letter to a Unicode Regional Indicator Symbol.
// Regional Indicator Symbols range from U+1F1E6 (🇦) to U+1F1FF (🇿), where U+1F1E6
// corresponds to 'A'. When two Regional Indicator Symbols are combined, platforms
// that support emoji flags render them as a single flag glyph.
//
// Example:
//
//	flagemoji.Get("FR") // Returns "🇫🇷" (French flag)
//	flagemoji.Get("US") // Returns "🇺🇸" (US flag)
//	flagemoji.Get("X")  // Returns "" (invalid: not 2 characters)
func Get(countryCode string) string {
	if len(countryCode) != 2 {
		return ""
	}

	// Convert each letter to its Regional Indicator Symbol:
	// 'A' (0x41) maps to U+1F1E6, 'B' to U+1F1E7, etc.
	// The offset (countryCode[n] - 'A') gives us 0 for 'A', 1 for 'B', etc.
	flag := []rune{0x1F1E6 + rune(countryCode[0]-'A'), 0x1F1E6 + rune(countryCode[1]-'A')}
	return string(flag)
}
