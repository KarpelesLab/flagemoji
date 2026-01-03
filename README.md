# flagemoji

A minimal, zero-dependency library for converting ISO 3166-1 alpha-2 country codes into emoji flags. Available for both Go and JavaScript/Node.js.

## How It Works

Emoji flags are composed of two Unicode Regional Indicator Symbols. Each letter A-Z has a corresponding Regional Indicator Symbol (U+1F1E6 to U+1F1FF). When two of these symbols are placed together and the combination corresponds to a valid ISO 3166-1 alpha-2 country code, platforms that support emoji flags render them as a single flag glyph.

For example, "FR" becomes:
- 'F' -> U+1F1EB (Regional Indicator Symbol Letter F)
- 'R' -> U+1F1F7 (Regional Indicator Symbol Letter R)
- Combined: U+1F1EB U+1F1F7 -> Displayed as a French flag

## Installation

### Go

```bash
go get github.com/KarpelesLab/flagemoji
```

### JavaScript/Node.js

```bash
npm install flagemoji
```

## Usage

### Go

```go
package main

import (
    "fmt"
    "github.com/KarpelesLab/flagemoji"
)

func main() {
    fmt.Println(flagemoji.Get("FR")) // 🇫🇷
    fmt.Println(flagemoji.Get("US")) // 🇺🇸
    fmt.Println(flagemoji.Get("JP")) // 🇯🇵
}
```

### JavaScript

```javascript
const { flagEmoji } = require('flagemoji');

console.log(flagEmoji('FR')); // 🇫🇷
console.log(flagEmoji('US')); // 🇺🇸
console.log(flagEmoji('JP')); // 🇯🇵
```

Or with ES modules:

```javascript
import { flagEmoji } from 'flagemoji';

console.log(flagEmoji('FR')); // 🇫🇷
```

## API

### Go: `Get(countryCode string) string`

Converts a 2-letter country code to its emoji flag. Returns an empty string if the input is not exactly 2 characters.

### JavaScript: `flagEmoji(countryCode)`

Converts a 2-letter country code to its emoji flag.

## Notes

- Country codes should be uppercase (e.g., "US" not "us")
- Only valid ISO 3166-1 alpha-2 codes will display as recognizable flags
- Flag rendering depends on platform support; some platforms may show the individual letter symbols instead

## License

This is free and unencumbered software released into the public domain. See [LICENSE](LICENSE) for details.
