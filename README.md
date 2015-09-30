# gemist
Package gemist provides functionality for working with Uitzending Gemist content.

## Example
```go
package main

import "github.com/dwlnetnl/gemist"

func main() {
  url := "http://www.npo.nl/radio-bergeijk/05-06-2004/POMS_VPRO_397233"

  if b, err := GetBroadcast(url); err == nil {
  	fmt.Println(b.Title)
  	fmt.Println(b.Date)
  	fmt.Println(b.Type)
  }

  // Output:
  // Het radiostation voor Bergeijk - Radio Bergeijk
  // 2004-06-05 13:32:00 +0200 CEST
  // Audio
}

```
