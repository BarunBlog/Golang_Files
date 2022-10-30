package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	//var languages = map[string]string{} // same
	languages := make(map[string]string) // map [keyType]valueType

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages:", languages)

	// iterate map
	for key, value := range languages {
		fmt.Printf("For key %s, value is %s\n", key, value)
	}
}
