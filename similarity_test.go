package similigo

import (
	"github.com/Ojelaidi/similigo/utils"
	"reflect"
	"testing"
)

func TestCalculateHybridSimilarity(t *testing.T) {
	type args struct {
		text1 string
		text2 string
		opts  []Option
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Identical strings with default options",
			args: args{
				text1: "Bonjour le monde",
				text2: "Bonjour le monde",
				opts:  nil, // no options, use defaults
			},
			want: 1.0,
		},
		{
			name: "Identical strings with 2-gram size",
			args: args{
				text1: "Bonjour le monde",
				text2: "Bonjour le monde",
				opts:  []Option{WithNgramSize(3)},
			},
			want: 0.9999999999999998,
		},
		{
			name: "Similar strings with custom stop words - basic test",
			args: args{
				text1: "Bonjour le monde",
				text2: "Salut le monde",
				opts:  []Option{WithCustomStopWords([]string{"stage", "custom", "test"})},
			},
			want: 0.4886750490563072,
		},
		{
			name: "Similar strings with custom stop words - sim score should be 1.0",
			args: args{
				text1: "Bonjour le monde",
				text2: "Salut le monde",
				opts:  []Option{WithCustomStopWords([]string{"Bonjour", "Salut"})},
			},
			want: 1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateHybridSimilarity(tt.args.text1, tt.args.text2, tt.args.opts...); got != tt.want {
				t.Errorf("CalculateHybridSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBestNMatchesInList(t *testing.T) {
	type args struct {
		targetText string
		texts      []string
		n          int
		opts       []Option
	}
	tests := []struct {
		name string
		args args
		want []utils.Match
	}{
		{
			name: "top 3 matches from list",
			args: args{
				targetText: "hello world",
				texts:      []string{"hello world", "hello", "world", "hola mundo", "hallo welt"},
				n:          3,
				opts:       nil, // No special options
			},
			want: []utils.Match{
				{Text: "hello world", Score: 0.9999999999999998},
				{Text: "world", Score: 0.7432900502033766},
				{Text: "hello", Score: 0.7432900502033766},
			},
		},
		{
			name: "top 2 matches with non-default options",
			args: args{
				targetText: "hello world",
				texts:      []string{"hello world", "hello", "world", "hola mundo", "hallo welt"},
				n:          2,
				opts:       []Option{WithNgramSize(3)},
			},
			want: []utils.Match{
				{Text: "hello world", Score: 0.9999999999999998},
				{Text: "world", Score: 0.7267584713501614},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBestNMatchesInList(tt.args.targetText, tt.args.texts, tt.args.n, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindBestNMatchesInList() = %v, want %v", got, tt.want)
			}
		})
	}
}
