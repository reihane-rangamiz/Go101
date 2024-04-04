# Arrays 101

## Solution

1. Create an empty array of integers of size 10

```go
package main

func main(){

    empArr := [10]int{}
}
```
---

2. Create an arrays of strings and fill it with 4 different strings

```go
package main

func main(){

    Strings := [4]string{"first","second","third","forth"}
}

```
---
3. Can you mix different data types in one array?

- No, can't have mixed data types in the same array.

---

4. You've created an empty array with the size of 10. Once declared can you add items to it? Can you add 50 items?


 - You can add items to the array up to its size once declared. But can't add 50 items to an array with the size of 10.

 ---

 5. There is an array called 'games'. Assign the string "Mario" to the first place in the array

 ```go
package main

func main(){
   
	games := [10]string{}
	fmt.Println(games)

	games[0] = "Mario"
	fmt.Println(games)
}

 ```
---

6. Create an array called "movies". Add some data to the array. Then, Print the following:
   1. The array itself
   2. The second item in the array
   3. The length of the array

```go
package main

import "fmt"

func main() {

    movies := [10]string{"DeadPool", "THOR", "spider-Man", "Bat-Man", "Iron-Man"}

    fmt.Println("Array:", movies)
    fmt.Println("Second item in the array:", movies[1])
    fmt.Println("length of the array is: ", len(movies))

}


```