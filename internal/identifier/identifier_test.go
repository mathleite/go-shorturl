package identifier_test

import (
	"mathleite/short-url/internal/identifier"
	"testing"
)

type identifierExpectation struct {
	input, expected string
}

func provideValidIndentifierExpectations() []identifierExpectation {
	return []identifierExpectation{
		{"https://www.google.com", "772e676f"},
		{"https://www.google.com/search?q=golang", "2e636f6d"},
		{"HTTP://www.google.com", "72e676f6"},
	}
}

func provideNotValidIndentifierExpectations() []identifierExpectation {
	return []identifierExpectation{
		{"https://www.google.com", ""},
		{"https://www.google.com", "_"},
		{"https://www.google.com/search?q=golang", "d34"},
		{"HTTP://www.google.com", "4343564564"},
	}
}

func TestCreateIdentifier_ShouldCreateHexIdentifiers(t *testing.T) {
	for _, expectation := range provideValidIndentifierExpectations() {
		identifier := identifier.CreateIdentifier(expectation.input)
		if identifier != expectation.expected {
			t.Errorf("Expected %s, got %s", expectation.expected, identifier)
		}
	}
}

func TestValidateIdentifier_GivenInputShouldValidateIdentifierAsEquals(t *testing.T) {
	for _, expectation := range provideValidIndentifierExpectations() {
		isValid := identifier.ValidateIdentifier(expectation.input, expectation.expected)
		if isValid != true {
			t.Errorf("Fail to validate `%s` as `%s`", expectation.input, expectation.expected)
		}
	}
}

func TestValidateIdentifier_GivenInputShouldReturnFalseWhenNoHasExactlyEightLength(t *testing.T) {
	for _, expectation := range provideNotValidIndentifierExpectations() {
		isValid := identifier.ValidateIdentifier(expectation.input, expectation.expected)
		if isValid == true {
			t.Errorf("Fail to validate `%s` as `%s`", expectation.input, expectation.expected)
		}
	}
}
