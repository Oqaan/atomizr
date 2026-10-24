package collection

import "testing"

func TestParseSeason(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Season
		wantErr bool
	}{
		{name: "spring", input: "spring", want: Spring},
		{name: "summer", input: "summer", want: Summer},
		{name: "autumn", input: "autumn", want: Autumn},
		{name: "winter", input: "winter", want: Winter},
		{name: "typo", input: "sumer", wantErr: true},
		{name: "empty", input: "", wantErr: true},
		{name: "capitalized", input: "Winter", wantErr: true},
		{name: "american", input: "fall", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSeason(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseSeason(%q): expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseSeason(%q): unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseSeason(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
