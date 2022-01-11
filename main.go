package main

import (
	"fmt"

	"github.com/duysqubix/godict/dict"
)

func main() {
	dict1 := dict.Dict()
	dict2 := dict.Dict()

	dict1.Update("A", 1)
	dict2.Update("B", 2)

	dict1.Update(dict2)

	dict1.Update("C", 3)

	fmt.Println(dict1.String())

	d1 := dict.RawMap{"Hello": "World"}

	dict3 := dict.Dict(d1)
	dict3.Update("other", dict1)
	dict3.Update("other2", dict2)
	fmt.Println(dict3)

	dict2.Update("B", 5)
	fmt.Println(dict3)

}
