## Solution

### 1.
This code will produce an error indicating `invalid type`.
- We need to tell the compiler the variables's data type when declaring one, since Go is statically typed. Here is the fixed version of this code:

```go
package main

import "fmt"

func main(){
    // here you will get warning for merging variable declaration with assignment. But in some cases this kind of declaration will be useful
    var userName string 
    userName = "user" 
    // var userName = "user"   --> this will also be right(Go /performing type inference)
    fmt.Println(userName)
}
```
- You can also use `short declaration` instead of regular declaration:
    - The type of the variable is inferred from the value on the right-hand side.(Go performing type inference)
```go

package main

import "fmt"

func main(){
    userName := "user"
    fmt.Println(userName)
}
```
---
### 2. 
The `%T` format specifier in the `fmt.Printf` function is used to display the type of the value being printed:

```Go
package main

import "fmt"

func main() {
    var animal = "Bear"
    var myNumber = 369369
    var isAvailable = true

    fmt.Printf("food is %T\nslices is %T\npineappleOnPizza is %T", animal, myNumber, isAvailable)
}
```