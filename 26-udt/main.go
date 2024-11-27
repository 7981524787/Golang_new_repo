package main

import (
	"errors"
	"fmt"
)

func main() {

	var map1 mymap

	map1 = make(mymap)

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["560034"] = "Bangalore-3"
	map1["560082"] = "Bangalore-4"

	keys, err := map1.GetKeys()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("keys of map1:", keys)
	}

	values, err := map1.GetValues()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("values of map1:", values)
	}

	var map2 map[string]any

	map2 = make(map[string]any)

	map2["560086"] = "Bangalore-1"
	map2["560096"] = "Bangalore-2"
	map2["560034"] = "Bangalore-3"
	map2["560082"] = "Bangalore-4"

	keys, err = mymap(map2).GetKeys()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("keys of map2:", keys)
	}

	values, err = mymap(map2).GetValues()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("values of map2:", values)
	}
}

type mymap map[string]any

func (m mymap) GetKeys() ([]string, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	keys := make([]string, len(m))
	c := 0
	for k := range m {
		keys[c] = k
		c++
	}
	return keys, nil
}

func (m mymap) GetValues() ([]any, error) {
	if m == nil {
		return nil, errors.New("nil map")
	}
	values := make([]any, len(m))
	c := 0
	for _, v := range m {
		values[c] = v
		c++
	}
	return values, nil
}

// print map
// Key:-> key Value-> value

// package management
// interfaces
