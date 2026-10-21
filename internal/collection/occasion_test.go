package collection

import "testing"

func TestParseOccasion(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Occasion
		wantErr bool
	}{
		{name: "office", input: "office", want: Office},
		{name: "casual", input: "casual", want: Casual},
		{name: "evening", input: "evening", want: Evening},
		{name: "typo", input: "ofice", wantErr: true},
		{name: "empty", input: "", wantErr: true},
		{name: "capitalized", input: "Office", wantErr: true},
		{name: "fragella", input: "night out", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseOccasion(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseOccasion(%q): expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseOccasion(%q): unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseOccasion(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
