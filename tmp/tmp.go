package tmp

import "fmt"

func main() {
	a := "12345"
	for i, val := range a {
		fmt.Println(i, val)
	}
}
