package similigo

type Option func(*SimilarityOptions)

// SimilarityOptions represents optional settings for hybrid similarity calculation.
// - NgramSize: The n-gram size used for n-gram similarity calculation.
// - WordSimWeight: Weight for word similarity in the final score.
// - NgramSimWeight: Weight for n-gram similarity in the final score.
// - ContainmentSimWeight: Weight for containment similarity in the final score.
type SimilarityOptions struct {
	NgramSize            int
	WordSimWeight        float64
	NgramSimWeight       float64
	ContainmentSimWeight float64
}

const (
	DefaultNgramSize            = 3
	DefaultWordSimWeight        = 0.5
	DefaultNgramSimWeight       = 0.3
	DefaultContainmentSimWeight = 0.2
)

func DefaultSimilarityOptions() *SimilarityOptions {
	return &SimilarityOptions{
		NgramSize:            DefaultNgramSize,
		WordSimWeight:        DefaultWordSimWeight,
		NgramSimWeight:       DefaultNgramSimWeight,
		ContainmentSimWeight: DefaultContainmentSimWeight,
	}
}

// WithNgramSize sets the ngramSize in similarityOptions.
func WithNgramSize(n int) Option {
	return func(opts *SimilarityOptions) {
		opts.NgramSize = n
	}
}

// WithWordSimWeight sets the wordSimWeight in similarityOptions.
func WithWordSimWeight(w float64) Option {
	return func(opts *SimilarityOptions) {
		opts.WordSimWeight = w
	}
}

// WithNgramSimWeight sets the ngramSimWeight in similarityOptions.
func WithNgramSimWeight(w float64) Option {
	return func(opts *SimilarityOptions) {
		opts.NgramSimWeight = w
	}
}

// WithContainmentSimWeight sets the containmentSimWeight in similarityOptions.
func WithContainmentSimWeight(w float64) Option {
	return func(opts *SimilarityOptions) {
		opts.ContainmentSimWeight = w
	}
}
