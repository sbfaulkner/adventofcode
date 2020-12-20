package passport

import (
	"strings"
	"testing"
)

func readPassports(t *testing.T) []*Passport {
	in := `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in
`

	p, err := ReadPassports(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	return p
}

func TestReadPassports(t *testing.T) {
	p := readPassports(t)

	got := len(p)
	want := 2

	if got != want {
		t.Errorf("ReadPassports: got %#v, want %#v", got, want)
	}
}

func TestValid(t *testing.T) {
	testCases := []struct {
		data string
		want bool
	}{
		{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm", true},
		{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929", false},
		{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm", true},
		{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in", false},
	}

	for _, tc := range testCases {
		p := NewPassport(tc.data)
		got := p.Valid()
		if got != tc.want {
			t.Errorf("p.Valid: got %#v, want %#v (%#v)", got, tc.want, tc.data)
		}
	}
}
