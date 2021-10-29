package flagemoji

func Get(countryCode string) string {
	// convert given country code to a flag
	if len(countryCode) != 2 {
		return ""
	}

	flag := []rune{0x1F1E6 + rune(countryCode[0]-'A'), 0x1F1E6 + rune(countryCode[1]-'A')}
	return string(flag)
}
