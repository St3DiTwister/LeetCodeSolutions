package _012_Integer_to_Roman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{
			name: "example 1",
			num:  3,
			want: "III",
		},
		{
			name: "example 2",
			num:  4,
			want: "IV",
		},
		{
			name: "example 3",
			num:  9,
			want: "IX",
		},
		{
			name: "example 4",
			num:  58,
			want: "LVIII",
			// Explanation:
			// L = 50, V = 5, III = 3
		},
		{
			name: "example 5",
			num:  1994,
			want: "MCMXCIV",
			// Explanation:
			// M = 1000
			// CM = 900
			// XC = 90
			// IV = 4
		},
		{
			name: "1000 -> M",
			num:  1000,
			want: "M",
		},
		{
			name: "3999 -> MMMCMXCIX",
			num:  3999,
			want: "MMMCMXCIX",
		},
		{
			name: "40 -> XL",
			num:  40,
			want: "XL",
		},
		{
			name: "90 -> XC",
			num:  90,
			want: "XC",
		},
		{
			name: "400 -> CD",
			num:  400,
			want: "CD",
		},
		{
			name: "900 -> CM",
			num:  900,
			want: "CM",
		},
		{
			name: "extended 2021",
			num:  2021,
			want: "MMXXI",
		},
		{
			name: "edge lowest",
			num:  1,
			want: "I",
		},
		{
			name: "edge upper-middle",
			num:  944,
			want: "CMXLIV",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if got != tt.want {
				t.Fatalf("intToRoman(%d) = %s, want %s", tt.num, got, tt.want)
			}
		})
	}
}
