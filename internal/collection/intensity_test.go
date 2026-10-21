package collection

import "testing"

func TestParseIntensity(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Intensity
		wantErr bool
	}{
		{name: "moderate", input: "moderate", want: Moderate},
		{name: "light", input: "light", want: Light},
		{name: "strong", input: "strong", want: Strong},
		{name: "beast", input: "beast", want: Beast},
		{name: "typo", input: "moderat", wantErr: true},
		{name: "empty", input: "", wantErr: true},
		{name: "capitalized", input: "Moderate", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIntensity(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseIntensity(%q): expected error, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseIntensity(%q): unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseIntensity(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
