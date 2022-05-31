/*func function_name([parameter list])[return_types]
{
	body
}
*/

package main

import "fmt"

func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func main() {
	r := sum(3, 5)
	fmt.Printf("r: %v\n", r)
}
