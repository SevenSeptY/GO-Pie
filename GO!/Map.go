package main

import "fmt"

func test4() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var k1 = "name"
	var k2 = "age 1"
	v, ok := m1[k1]
	fmt.Printf("v:%v\n", v)
	fmt.Printf("ok:%v\n", ok)
	fmt.Println("--------------------------")
	v, ok = m1[k2]
	fmt.Printf("v:%v\n", v)
	fmt.Printf("ok:%v\n", ok)

}

func test3() {
	var m1 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	var key = "name"

	var value = m1[key]
	fmt.Printf("Value: %v\n", value)
}

func test2() {
	var m2 = map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
	fmt.Printf("m2: %v\n", m2)

	m3 := make(map[string]string)
	m3["name"] = "tom"
	m3["age"] = "20"
	m3["email"] = "tom@hotmail.com"
	fmt.Printf("m3: %v\n", m3)
}

func test1() {
	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["age"] = "20"
	m1["email"] = "tom@gmail.com"

	fmt.Printf("m1: %v\n", m1)
	//var m1 map[string]string
	//m1 = make(map[string]string)
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}


//----------------------------------------------
package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(len(m))
	m = make(map[string]int) // ==m = make(map[string]int, 0)
	fmt.Println(len(m))
	m = make(map[string]int, 10) //cap = 10 , not length
	fmt.Println(len(m))
	m = map[string]int{"A": 3, "B": 2}
	fmt.Println(len(m))
	m["D"] = 18
	fmt.Println(len(m))

	delete(m, "B")
	fmt.Println(len(m))

	for key, value := range m {
		fmt.Printf("key %s, value %d\n", key, value)
	}
}

/*


 */
