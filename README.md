# random-string
Generate random string in Golang

## Install
```bash
go get TheYahya/random-string
```

## Usage
```go
package main

import (
	"fmt"
	"github.com/theyahya/random-string"
)

func main() {
	randStr1 := randomstring.New().Generate(10)
	fmt.Println(randStr1) //hGhfVRtFH9

	// Or with spicific characters
	// Default is: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	randStr2 := randomstring.New().Letters("abcdf").Generate(10)
	fmt.Println(randStr2) //abdcbfbdfb
}
```

## License
Licensed under the [MIT License](https://github.com/TheYahya/random-string/blob/main/README.md).