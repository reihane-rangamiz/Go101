# Solution

### 1.

 This code will lead to an error since no return value has been defined while the code is trying to return an integer value.

### 2. 

You just have to define a return value in the function:

```Go
package main
        
import "fmt"
        
func add(x, y int) (z int) {
    z = x + y                     
    return z
}
       
func main() {
    fmt.Print(add(100, 320))
}
```

### 3. 

**Yes indeed!**!! Surprised? 😲</br> HAHA 😄</br>

In Go, "naked returns" refer to a style of writing return statements in a function where the named return values are simply listed without specifying them in the return statement. It makes the code more concise but may reduce the readability ! 