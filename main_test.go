package main

import (
	"fmt"
	"testing"
)

// Test_countBits makes sure this function works properly
func Test_countBits(t *testing.T) {
	tests := []struct {
		num     uint64
		leading int
		zeros   int
		ones    int
	}{
		{0, 64, 0, 0},
		{1, 63, 0, 1},
		{2, 62, 1, 1},
		{3, 62, 0, 2},
		{0xFFFF_FFFF_FFFF_FFFF, 0, 0, 64},
		{0x8000_0000_0000_0000, 0, 63, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt), func(t *testing.T) {
			gotLeading, zeros, ones := countBits(tt.num)
			if gotLeading != tt.leading {
				t.Errorf("countBits() leading = %v, want %v", gotLeading, tt.leading)
			}
			if zeros != tt.zeros {
				t.Errorf("countBits() zeros = %v, want %v", zeros, tt.zeros)
			}
			if ones != tt.ones {
				t.Errorf("countBits() ones = %v, want %v", ones, tt.ones)
			}
		})
	}
}
