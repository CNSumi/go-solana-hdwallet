# go-solana-hdwallet
> Solana HD Wallet derivations from [mnemonic] seed in Go (golang).

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/cnsumi/go-solana-hdwallet/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/cnsumi/go-solana-hdwallet?)](https://goreportcard.com/report/github.com/cnsumi/go-solana-hdwallet)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](#contributing)


## Install
```bash
go get -u github.com/cnsumi/go-solana-hdwallet
```

## Getting started

```go
package main

import (
	"fmt"
	hdwallet "github.com/cnsumi/go-solana-hdwallet"
)

func main() {
	// derivation from default path
	mnemonic := "beach liar addict wrap pause runway evolve front grab print jungle mimic"
	node := hdwallet.NewNode(mnemonic)

	fmt.Println(node.SecretKey())
	fmt.Println(node.Address())
	
	// derivation according solana account index
	node = hdwallet.NewNode(mnemonic, hdwallet.WithIndex(1))
	fmt.Println(node.SecretKey())
	fmt.Println(node.Address())
	
	// ...
}
```

## CLI

## Test

```bash
make test
```

## Contributing

Pull requests are welcome!

For contributions please create a new branch and submit a pull request for review.

## License

Released under the [MIT](./LICENSE) license.

© [CNSumi](https://github.com/CNSumi)