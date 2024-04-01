### 1.
```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNumber := rand.Intn(100)
	fmt.Printf("The random number: %d\n", randomNumber)

	if randomNumber > 50 {
		fmt.Println("It's closer to 100")
	} else {
		fmt.Println("It's closer to 0")
	}
}
```
---
### 2.
```go 
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	randomNumber := rand.Intn(100)
	fmt.Printf("The random number: %d\n", randomNumber)


	if randomNumber > 50 {
		fmt.Println("It's closer to 100")
	} else if randomNumber < 50 {

		fmt.Println("It's closer to 0")
	}else {
		fmt.Println("It's 50")
	}

}
```
---
### 3.

```go 
package main

import (
	"fmt"
	"math/rand"
)

func main() {
    
	randomNumber := rand.Intn(100)

	if randomNumber > 50 && randomNumber%2 == 0 {
		fmt.Println("It's closer to 100")
	} else if randomNumber < 50 {

		fmt.Println("It's closer to 0")
	} else if randomNumber == 50 {
		fmt.Println("It's 50")
	}
	fmt.Printf("The random number: %d\n", randomNumber)


}
```