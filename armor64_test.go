package armor64

import (
	"testing"
)

func TestArmor64(t *testing.T) {
	work := []struct{ src, dst string }{
		{"", ""},
		{"JP", "H_-"},
		{"Hello, World!", "H5KgQ5wg74SjRalZ7F"},
		{"armor64 is safe, strict, and stable. It is specified and easy to test. Do not settle for lesser encodings.", "NM8hQr7qC10dRm0nNLO_A10nS68dNrFg754iO10nS54XQ5Ji73_o75_n76CkOLCdOa__O10WQaFVOL4nTH0oQm0oOMCoAX03Qm0iQrFVRqKoS5l_75OjRX0gOMCnOM7VOLtYQqGdQaSnAV"},
	}
	for _, test := range work {
		encoded := EncodeToString([]byte(test.src))
		if encoded != test.dst {
			t.Fatalf("encoding failed: %v instead of %v", encoded, test.dst)
		}
		decoded, err := DecodeString(test.dst)
		if err != nil {
			t.Fatalf("decoding %v failed: %v", test.dst, err)
		}
		if string(decoded) != test.src {
			t.Fatalf("decoding failed: %v", string(decoded))
		}
	}
	doNotDecode := []string{
		" ",
		"\r",
		"\n",
		"__==",
	}
	for _, test := range doNotDecode {
		if _, err := DecodeString(test); err == nil {
			t.Fatalf("decoding %v should have failed", test)
		}
	}
}
