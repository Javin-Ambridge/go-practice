package arraysEx;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
);

func ArrayEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Arrays");

	myArr := [5]int8 {1,2,3,4,5};
	for i := 0; i < len(myArr); i++ {
		fmt.Println("at myArr[", i, "]=", myArr[i]);
	}

	//A slice is an array without a predifined size
	mySlice := []int {1,2,3};
	//Make mySlice2=mySlice[1 - 3 exclusive]
	mySlice2 := mySlice[1:3];
	for i := 0; i < len(mySlice2); i++ {
		fmt.Println(i, mySlice2[i]);
	}
	//Print out from indexs 0 - 1 exclusive
	fmt.Println(mySlice2[:1]);
	//Print out from index 1 onwards
	fmt.Println(mySlice[1:]);

	//Empty slice initially make(type, initial length, max length [default none])
	mySlice3 := make([]int, 0);
	fmt.Println(mySlice3);
	mySlice3 = append(mySlice3, 10);
	fmt.Println(mySlice3);
	mySlice3 = append(mySlice3, 11,12);
	fmt.Println(mySlice3);
	//Appending a slice onto another slice (need the '...')
	mySlice3 = append(mySlice3, mySlice...);
	fmt.Println(mySlice3);

	//Copy mySlice into mySlice3
	mySlice4 := make([]int, len(mySlice));
	copy(mySlice4, mySlice);
	fmt.Println(mySlice4);
}
