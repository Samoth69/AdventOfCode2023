package day_2

import "testing"

func Test_decodeParty(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"01", args{"  8 green, 3 blue  "}, 0, false},
		{"02", args{"8 green, 3 blue"}, 0, false},
		{"03", args{"12 red, 13 green, 14 blue"}, 0, false},
		{"04", args{"12 red, 14 green, 7 blue"}, 1, false},
		{"05", args{"12 red, 25 purple, 14 green"}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeParty(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeParty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeParty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"game 01", args{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}, 0, false},
		{"game 02", args{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"}, 0, false},
		{"game 03", args{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"}, 1, false},
		{"game 04", args{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"}, 1, false},
		{"game 05", args{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("decodeLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
