package functions;

import "fmt";
import div "divider";

func FunctionEx() {
	defer div.PrintDivider();
	//multiple return values.
	num1, num2 := returnTwoNums();
	fmt.Println("num1: ", num1, ". num2: ", num2);
	mySlice := []int64 {12, -10, 100000, 12314, -9999, 9897612};
	fmt.Println(minAndMax(mySlice...));
	closureEx();
	fmt.Println(recoverEx(3, 0));
	fmt.Println(recoverEx(3, 1));
	panicEx();
}

func returnTwoNums() (int, int) {
	return 1, 2;
}

//Returns min and max from numbers
func minAndMax(args ...int64) (int64, int64) {
	if len(args) == 0 {
		return 0, 0;
	}

	var min int64 = 100000000000;
	var max int64 = -100000000000;
	for i := 0; i < len(args); i++ {
		if (args[i] > max) {
			max = args[i];
		}
		if (args[i] < min) {
			min = args[i];
		}
	}
	return min, max;
}

func closureEx() {
	num1 := 1;
	doubleNum := func() int {
		num1 *= 2;
		return num1;
	}
	fmt.Println(doubleNum(), doubleNum());
}

func recoverEx(num1, num2 int) int {
	defer func() {
		recover(); //doesnt crash
	}();
	solution := num1 / num2;
	return solution;
}

func panicEx() {
	defer func() {
		fmt.Println(recover());
	}();
	panic("FAKE PANIC ERROR1!");
}
