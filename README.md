# go-sh

## example

```go
package main

import(
  "fmt"

  sh "github.com/ohkinozomu/go-sh"
)

func main() {
	if err := sh.Run("echo 'hello world' > tmp.txt"); err != nil {
		panic(err)
	}

	result, err := sh.RunR("cat tmp.txt")
	if err != nil {
		panic(err)
	}
	
  fmt.Println(result) 
}
```