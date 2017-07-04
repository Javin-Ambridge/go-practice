package maps;

import "fmt";
import div "divider";

func MapEx() {
	defer div.PrintDivider();
	//Make a map with a string key, and int value.
	myMap := make(map[string] int);

	//Add to the map
	myMap["first key"] = 1;
	myMap["second key"] = 2;
	fmt.Println(len(myMap));
	fmt.Println(myMap);
	fmt.Println(myMap["first key"]);
	myMap["third key"] = 3;
	fmt.Println(myMap);

	//update a key
	myMap["second key"] = 12;
	fmt.Println(myMap);
	fmt.Println(myMap["second key"]);

	//deleting a key delete(map, key)
	delete(myMap, "second key");
	fmt.Println(len(myMap));
	fmt.Println(myMap);

	//map returns two values, both the value and a boolean if it exists
	//This if statements first gets the values, then runs if ok==true
	if val, ok := myMap["second key"]; ok {
		fmt.Println("Value in map: ", val);
	} else {
		fmt.Println("Value not in map: ", val);
	}
}