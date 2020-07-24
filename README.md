# timefmt-go [![CI Status](https://github.com/itchyny/timefmt-go/workflows/CI/badge.svg)](https://github.com/itchyny/timefmt-go/actions) [![Go Report Card](https://goreportcard.com/badge/github.com/itchyny/timefmt-go)](https://goreportcard.com/report/github.com/itchyny/timefmt-go) [![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/itchyny/timefmt-go/blob/master/LICENSE) [![PkgGoDev](https://pkg.go.dev/badge/github.com/itchyny/timefmt-go)](https://pkg.go.dev/github.com/itchyny/timefmt-go)
### ― Simple time formatting library (strftime, strptime) for Golang ―

## Usage
```go
package main

import (
	"fmt"
	"log"

	"github.com/itchyny/timefmt-go"
)

func main() {
	t, err := timefmt.Parse("2020/07/24 09:07:29", "%Y/%m/%d %H:%M:%S")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t) // 2020-07-24 09:07:29 +0000 UTC

	str := timefmt.Format(t, "%Y/%m/%d %H:%M:%S")
	fmt.Println(str) // 2020/07/24 09:07:29

	str = timefmt.Format(t, "%a, %d %b %Y %T %z")
	fmt.Println(str) // Fri, 24 Jul 2020 09:07:29 +0000
}
```

Please refer to [`man 3 strftime`](https://linux.die.net/man/3/strftime) and
[`man 3 strptime`](https://linux.die.net/man/3/strptime) for formatters.
Note that `E` and `O` modifier characters are not supported.

## Bug Tracker
Report bug at [Issues・itchyny/timefmt-go - GitHub](https://github.com/itchyny/timefmt-go/issues).

## Author
itchyny (https://github.com/itchyny)

## License
This software is released under the MIT License, see LICENSE.
