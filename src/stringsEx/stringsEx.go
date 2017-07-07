package stringsEx;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
);

func StringsEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Strings");

	var myName string = "Javin Ambridge";
	fmt.Println("[1]: ", myName, " and the length: ", len(myName));
	myName += " this is a test";
	fmt.Println("[2]: ", myName, " and the length: ", len(myName));
}
