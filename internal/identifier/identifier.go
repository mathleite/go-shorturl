package identifier

import (
	"encoding/hex"
	"strings"
)

const IDENTIFIER_LENGTH = 8
const QUANTITY_RIGHT_OUTPUT_CROP_CHARACTER = 6
const QUANTITY_LEFT_OUTPUT_CROP_CHARACTER = 2

func CreateIdentifier(url string) string {
	output := hex.EncodeToString([]byte(normalizeInput(url)))
	return cropOutputValue(output)
}

func ValidateIdentifier(input string, identifier string) bool {
	if len(identifier) != IDENTIFIER_LENGTH {
		return false
	}
	return CreateIdentifier(input) == identifier
}

func normalizeInput(input string) string {
	return strings.ToLower(input)
}

func cropOutputValue(output string) string {
	outputHalf := len(output) / 2
	return output[outputHalf-QUANTITY_LEFT_OUTPUT_CROP_CHARACTER : outputHalf+QUANTITY_RIGHT_OUTPUT_CROP_CHARACTER]
}
