package main 
import "fmt"

func main() {
	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(x)
	*a = *a**a
	fmt.Println(x)
	fmt.Println(*a)

	
}