package main
2. 
3. import "fmt"
4. 
5. var a string = "foo"
6. 
7. func main() {
8.    var b string = "bar"
9.    fmt.Println(a)
10.   fmt.Println(b)
11.   fmt.Println(c)
12.  {
13.       var c string = "loreum"
14.       fmt.Println(a)
15.       fmt.Println(b)
16.       fmt.Println(c)
17.  }
18. }