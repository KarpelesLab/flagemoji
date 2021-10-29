'use strict'

function flagEmoji(countryCode) {
	return String.fromCodePoint(0x1F1E6+countryCode.charCodeAt(0)-0x41, 0x1F1E6+countryCode.charCodeAt(1)-0x41);
}

module.exports.flagEmoji = flagEmoji;
