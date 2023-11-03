# Similigo

A Go client library to calculate similarity between two strings

<img src="https://img.shields.io/github/go-mod/go-version/Ojelaidi/similigo">

## Usage

```
go get -u github.com/Ojelaidi/similigo
```

#### Example

Default Usage
```go
package main

import "github.com/Ojelaidi/similigo"

func main() {

	similarityScore := similigo.CalculateHybridSimilarity("text1", "text2")
	fmt.Printf("Similarity Score: %.2f\n", similarityScore)
}

```

Usage with options
```go
package main

import "github.com/Ojelaidi/similigo"

func main() {
	
	similarityScore := similigo.CalculateHybridSimilarity(
		"text1",
		"text2",
		similigo.WithNgramSize(4),
		similigo.WithWordSimWeight(0.4),
		similigo.WithNgramSimWeight(0.4),
		similigo.WithContainmentSimWeight(0.2),
	)
	fmt.Printf("Similarity Score: %.2f\n", similarityScore)
}

```

## How to Contribute

If you want to contribute you can read [Contributing](CONTRIBUTING.md)


## License

This project is under [BSD 3-Clause License](LICENSE)

