package main

import "fmt"

func main() {
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)
	fmt.Println(sammy["name"])

	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}

	keys := []string{}

	for key := range sammy {
		fmt.Println(key)
		keys = append(keys, key)
	}
	fmt.Printf("%q\n", keys)

	items := make([]string, len(sammy))
	var i int

	for _, value := range sammy {
		items[i] = value
		i++
	}
	fmt.Printf("%q\n", items)

	fmt.Println(len(sammy))

	fmt.Println(sammy["sammy"])
	// text, ok := sammy["name"]
	// if ok {
	// 	fmt.Printf("Sammy has a text of %q\n", text)
	// } else {
	// 	fmt.Println("Sammy was not found")
	// }

	if text, ok := sammy["name"]; ok {
		fmt.Printf("Sammy has a text of %q\n", text)
	} else {
		fmt.Println("Sammy was not found")
	}

	sammy["contains"] = "not 0"
	fmt.Println(sammy)
	
	sammy["color"] = "red"
	fmt.Println(sammy)

	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16:"modify"}
	delete(permissions, 16)
	fmt.Println(permissions)
	permissions = map[int]string{}
	fmt.Println(permissions)

	

}
