package stringsEx;

import (
	"fmt";
	div "divider";
);

func StringsEx() {
	defer div.PrintDivider();
	var myName string = "Javin Ambridge";
	fmt.Println("[1]: ", myName, " and the length: ", len(myName));
	myName += " this is a test";
	fmt.Println("[2]: ", myName, " and the length: ", len(myName));
}
