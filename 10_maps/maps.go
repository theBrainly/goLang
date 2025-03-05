package main
import "fmt"

func main() {

	// creating maps
	m := make(map[string]string)
	
	map["name"]="akash"
	map["age"]="20"
	  fmt.Println(map)

	// // getting element from map
	// fmt.Println(map["name"])
    // fmt.Println(map["age"])
    // // deleting element from map
    // delete(map, "age")
    // fmt.Println(map)
    // // checking if key exists in map
    // if val, ok := map["age"]; ok {
    //     fmt.Println("key exists", val)
    // } else {
    //     fmt.Println("key does not exist")
    // }
    // // iterating over map
    // for key, value := range map {
    //     fmt.Println("key:", key, "value:", value)
    // }
    // // map with string as key and slice of strings as value
    // map2:=make(map[string][]string)
    // map2["names"] = []string{"akash", "rahul", "deepak"}
    // fmt.Println(map2["names"])
}