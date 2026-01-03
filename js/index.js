'use strict'

/**
 * Converts an ISO 3166-1 alpha-2 country code to its corresponding emoji flag.
 *
 * The conversion maps each letter to a Unicode Regional Indicator Symbol.
 * Regional Indicator Symbols range from U+1F1E6 (🇦) to U+1F1FF (🇿).
 * When two Regional Indicator Symbols are combined, platforms that support
 * emoji flags render them as a single flag glyph.
 *
 * @param {string} countryCode - A 2-letter ISO 3166-1 alpha-2 country code (e.g., "US", "FR", "JP")
 * @returns {string} The emoji flag representation of the country
 * @example
 * flagEmoji('FR') // Returns "🇫🇷" (French flag)
 * flagEmoji('US') // Returns "🇺🇸" (US flag)
 */
function flagEmoji(countryCode) {
    // 0x1F1E6 is the Unicode code point for Regional Indicator Symbol Letter A
    // 0x41 is the ASCII code for 'A'
    // For each letter, we calculate: base + (letter - 'A') to get the Regional Indicator Symbol
    return String.fromCodePoint(0x1F1E6+countryCode.charCodeAt(0)-0x41, 0x1F1E6+countryCode.charCodeAt(1)-0x41);
}

module.exports.flagEmoji = flagEmoji;
