package authgateway

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TrimWhitespace removes leading and trailing whitespace from a string
func TrimWhitespace(input string) string {
	return strings.TrimSpace(input)
}

// ExtractToken removes the 'Bearer ' prefix from a token and returns it
func ExtractToken(authHeader string) (string, error) {
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid auth header")
	}
	return parts[1], nil
}

// DecodeInt64 attempts to decode a string into an int64
func DecodeInt64(input string) (int64, error) {
	id, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		log.Printf("Failed to parse int64: %v", err)
		return 0, err
	}
	return id, nil
}

// HTTPError returns an HTTP error response
func HTTPError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

// Empty returns true if a string is empty
func Empty(input string) bool {
	return input == ""
}