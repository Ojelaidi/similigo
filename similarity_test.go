package similigo

import "testing"

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
