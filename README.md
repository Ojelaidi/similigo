# Similigo

A Go client library to calculate similarity between two strings

<img src="https://img.shields.io/github/go-mod/go-version/Ojelaidi/similigo">

## Usage

```
go get -u github.com/Ojelaidi/similigo
```


#### Example
```go
package main

import "github.com/Ojelaidi/similigo"

func main() {
	options := similigo.SimilarityOptions{
		NgramSize:         3,
		WordSimWeight:     0.5,
		NgramSimWeight:    0.3,
		ContainmentWeight: 0.2,
	}

	result := similigo.CalculateHybridSimilarity("text1", "text2", &options)
	fmt.Printf("Similarity Score: %.2f\n", result)
}

```

## How to Contribute

If you want to contribute you can read [Contributing](CONTRIBUTING.md)


## License

This project is under [BSD 3-Clause License](LICENSE)

