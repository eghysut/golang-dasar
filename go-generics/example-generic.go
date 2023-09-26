package main

import (
	"fmt"
)

type TypeData interface{
    ~int | ~int32 | ~float32 | ~float64 | ~string
}

type GenericMap[K TypeData, V TypeData] map[K]V

func (g GenericMap[K, V]) sum() V {
    var angka V
    for _, v := range g {
        angka += v
    }
    return angka
}

func (g GenericMap[_, V]) values() {
    for _, val := range g {
        fmt.Println(val)
    }
}


func (g GenericMap[K, _]) keys() {
    for key, _ := range g {
        fmt.Println(key)
    }
}


func (g GenericMap[K, V]) get(s K) V {
    //fmt.Println(s)
    var val V

    for i, v := range g {
        if i == s {
            return v
        }
    }
    return val
}

func main() {

    x := GenericMap[string, int]{"foo": 10, "bar": 20}
    fmt.Println(x.sum())
    // Output: 30

    x.values()
    // Output: 
    // 10
    // 20

    x.keys()
    // Output:
    // foo
    // bar

    fmt.Println(x.get("hello"))
    // Output: 0

    fmt.Println(x.get("foo"))
    // Output: 10

}
