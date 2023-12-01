package day_1

import "testing"

func TestGetFirstDigit(t *testing.T) {
	type args struct {
		text    string
		reverse bool
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"01", args{"test1test2", false}, "1", false},
		{"02", args{"test1test2", true}, "2", false},
		{"03", args{"testtest", false}, "", true},
		{"04", args{"testtest", true}, "", true},
		{"05", args{"123test456test789", true}, "9", false},
		{"06", args{"123test456test789", false}, "1", false},
		{"07", args{"test1", false}, "1", false},
		{"08", args{"test1", true}, "1", false},
		{"09", args{"1test", false}, "1", false},
		{"10", args{"1test", true}, "1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFirstDigit(tt.args.text, tt.args.reverse)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFirstDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFirstDigit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
