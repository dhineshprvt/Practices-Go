package main
import "fmt"
func main() {
	var x complex64 = 3.4 + 2.9i
	var y = 5 + 7i
	fmt.Println(x, y)
	var a1 = 4.5
	var a2 = 7.1
	var c = complex(a1, a2)
	fmt.Println(c)
	var a = 3 + 5i
	var b = 2 + 4i
	var res1 = a + b
	var res2 = a - b
	var res3 = a * b
	var res4 = a / b
	fmt.Println(res1, res2, res3, res4)
}
