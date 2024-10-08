package armor64

import (
	"encoding/base64"
	"errors"
	"regexp"
)

var (
	alphabet          = "-0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz"
	valid             = regexp.MustCompile("^[0-9A-Za-z_-]*$")
	encoding          = base64.NewEncoding(alphabet).Strict().WithPadding(base64.NoPadding)
	invalidCharsError = errors.New("input has invalid characters")
)

// EncodedLen returns the length in bytes of the armor64 encoding
// of an input buffer of length n.
func EncodedLen(n int) int {
	return n/3*4 + (n%3*8+5)/6
}

// DecodedLen returns the maximum length in bytes of the decoded data
// corresponding to n bytes of armor64-encoded data.
func DecodedLen(n int) int {
	return n/4*3 + n%4*6/8
}

// Encode encodes src, writing [EncodedLen](len(src)) bytes to dst.
func Encode(dst []byte, src []byte) {
	encoding.Encode(dst, src)
}

// Decode decodes src. It writes at most [DecodedLen](len(src)) bytes to dst
// and returns the number of bytes written.
// If src contains invalid armor64 data, it will return the
// number of bytes successfully written and [base64.CorruptInputError].
func Decode(dst []byte, src []byte) (int, error) {
	if !valid.Match(src) {
		return 0, base64.CorruptInputError(0)
	}
	return encoding.Decode(dst, src)
}

// EncodeToString returns the armor64 encoding of src.
func EncodeToString(src []byte) string {
	return encoding.EncodeToString(src)
}

// DecodeString returns the bytes represented by the armor64 string s.
func DecodeString(s string) ([]byte, error) {
	if !valid.MatchString(s) {
		return nil, invalidCharsError
	}
	return encoding.DecodeString(s)
}
