package intsAndFloats;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
);

func IntsAndFloats() {
	defer div.PrintDivider();
	tp.TitlePrint("Ints and Floats");

	// An int is a positive or negative number without decimals
	// Versions
	// uint8 : unsigned  8-bit integers (0 to 255)
	// uint16 : unsigned 16-bit integers (0 to 65535)
	// uint32 : unsigned 32-bit integers (0 to 4294967295)
	// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
	// int8 : signed  8-bit integers (-128 to 127)
	// int16 : signed 16-bit integers (-32768 to 32767)
	// int32 : signed 32-bit integers (-2147483648 to 2147483647)
	// int64 : signed 64-bit integers (-9223372036854775808 to
	// 9223372036854775807)

	var age int64 = 40;

	// A float is a number with decimals
	// Versions : float32, float64
	var favNum float64 = 1.61803398875;
	fmt.Println(age, " and then its favNum: ", favNum);

	//constant
	const pi float32 = 3.14159;

	//multiple declarations
	var (
		num1 int8 = 120;
		num2 int8 = 127;
	)

	fmt.Println("num1: ", num1, " and num2: ", num2);

	var tmpBool bool = true;
	fmt.Println("boolean: ", tmpBool);
}
