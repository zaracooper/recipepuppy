[![release](https://img.shields.io/badge/release-v3.0.5-black)](https://github.com/zaracooper/recipepuppy/releases/tag/v3.0.5) [![pkg.go.dev](https://img.shields.io/badge/pgk.go.dev-reference-black)](https://pkg.go.dev/github.com/zaracooper/recipepuppy/v3@v3.0.5?tab=overview) [![GoDoc](https://godoc.org/github.com/zaracooper/recipepuppy?status.svg)](https://godoc.org/github.com/zaracooper/recipepuppy) [![Go mod version](https://img.shields.io/github/go-mod/go-version/zaracooper/recipepuppy?filename=v3%2Fgo.mod)](https://golang.org/doc/go1.13) [![Go Report Card](https://goreportcard.com/badge/github.com/zaracooper/recipepuppy)](https://goreportcard.com/report/github.com/zaracooper/recipepuppy) [![MIT License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/zaracooper/recipepuppy/blob/master/v3/LICENSE)
<p align="center">
    <img width="263" height="108" src="https://github.com/zaracooper/recipepuppy/blob/master/recipepuppy.png?raw=true" alt="centered image" />
</p>

<h2 align="center">Recipe Puppy Go Client</h2>

This is a Go client for the [recipepuppy.com](http://www.recipepuppy.com/about/api/) REST API. I made this to illustrate downgrading and upgrading modules for [this article](https://zaracooper.github.io/blog/posts/go-upgrades-downgrades/). 

## Usage
Add this as a dependency:
```
go get github.com/zaracooper/recipepuppy
```

To find recipes that match a title:
```go
recipes, err := recipepuppy.FindRecipes('steak', 1)
```

To find recipes that use certain ingredients:
```go
recipes, err := recipepuppy.FindRecipesByIngredients([]string{'egg', 'bacon'}, 1)
```

To find recipes that match a title and use the specified ingredients:
```go
recipes,  err := recipepuppy.FindRecipesWithIngredients('steak', []string{'eggs'}, 1)
```
