package main

import (
	"testing"
)

func TestUnix2iso(t *testing.T) {
	tests := []struct {
		unix  int64
		human bool
		want  string
	}{
		{unix: 0, human: false, want: "1970-01-01T00:00:00Z"},
		{unix: 0, human: true, want: "1970-01-01 00:00:00 UTC"},
		{unix: 1609459200, human: false, want: "2021-01-01T00:00:00Z"},
		{unix: 1609459200, human: true, want: "2021-01-01 00:00:00 UTC"},
		{unix: 1672531199, human: false, want: "2022-12-31T23:59:59Z"},
		{unix: 1672531199, human: true, want: "2022-12-31 23:59:59 UTC"},
	}
	for _, tt := range tests {
		got := Unix2iso(tt.unix, tt.human)
		if got != tt.want {
			t.Errorf("Unix2iso(%d, %v) = %q; want %q", tt.unix, tt.human, got, tt.want)
		}
	}
}
