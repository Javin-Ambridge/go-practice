package functions;

import "fmt";
import div "divider";

func FunctionEx() {
	defer div.PrintDivider();
	//multiple return values.
	num1, num2 := returnTwoNums();
	fmt.Println("num1: ", num1, ". num2: ", num2);
}

func returnTwoNums() (int, int) {
	return 1, 2;
}