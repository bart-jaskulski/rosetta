package isbn

import (
	"testing"
)

func TestValidIsbn13(t *testing.T) {
	tests := []struct {
		isbn  string
		valid bool
	}{
		{
			isbn:  "978-0596528126",
			valid: true,
		},
		{
			isbn:  "978-0596528120",
			valid: false,
		},
		{
			isbn:  "978-1788399081",
			valid: true,
		},
		{
			isbn:  "978-1788399083",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.isbn, func(t *testing.T) {
			got := ValidISBN(tt.isbn)
			if got != tt.valid {
				t.Errorf("ValidISBN(%s) = %v", tt.isbn, got)
			}
		})
	}
}
