# flagemoji

A minimal, zero-dependency JavaScript library for converting ISO 3166-1 alpha-2 country codes into emoji flags.

## Installation

```bash
npm install flagemoji
```

## Usage

### CommonJS

```javascript
const { flagEmoji } = require('flagemoji');

console.log(flagEmoji('FR')); // 🇫🇷
console.log(flagEmoji('US')); // 🇺🇸
console.log(flagEmoji('JP')); // 🇯🇵
```

### ES Modules

```javascript
import { flagEmoji } from 'flagemoji';

console.log(flagEmoji('FR')); // 🇫🇷
```

## API

### `flagEmoji(countryCode)`

Converts a 2-letter ISO 3166-1 alpha-2 country code to its corresponding emoji flag.

**Parameters:**
- `countryCode` (string): A 2-letter uppercase country code (e.g., "US", "FR", "JP")

**Returns:**
- (string): The emoji flag representation

## How It Works

Emoji flags are composed of two Unicode Regional Indicator Symbols. Each letter A-Z maps to a Regional Indicator Symbol (U+1F1E6 to U+1F1FF). When two symbols are combined and correspond to a valid country code, platforms render them as a flag.

## Notes

- Country codes should be uppercase
- Flag rendering depends on platform/browser support
- Invalid codes will produce Regional Indicator Symbol pairs that may not display as flags

## License

This is free and unencumbered software released into the public domain.
