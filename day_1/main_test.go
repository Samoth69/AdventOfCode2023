package day_1

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func BenchmarkDecodeNumber(b *testing.B) {
	file, err := os.OpenFile("../data/01.txt", os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	arr := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for index, _ := range arr {
			ret, err := decodeNumber(arr[index])
			if err != nil {
				panic(err)
			}
			if ret == 0 {
				panic(fmt.Errorf("invalid return value"))
			}
		}
	}
}

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

func TestIsDigitInText(t *testing.T) {
	type args struct {
		text       string
		startIndex int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"01", args{"one", 0}, "1"},
		{"02", args{"one", 1}, ""},
		{"03", args{"onetwo", 0}, "1"},
		{"04", args{"onetwo", 1}, ""},
		{"05", args{"onetwo", 2}, ""},
		{"06", args{"onetwo", 3}, "2"},
		{"07", args{"seventhree1eightztszfourfivesix", 0}, "7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigitInText(tt.args.text, tt.args.startIndex); got != tt.want {
				t.Errorf("isDigitInText() = %v, want %v", got, tt.want)
			}
		})
	}
}
