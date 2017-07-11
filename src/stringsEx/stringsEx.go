package stringsEx;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
	"strings";
	"sort";
	"strconv";
);

func StringsEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Strings");

	var myName string = "Javin Ambridge";
	fmt.Println("[1]: ", myName, " and the length: ", len(myName));
	myName += " this is a test";
	fmt.Println("[2]: ", myName, " and the length: ", len(myName));
	builtInStringFnc();
	sortEx();
}

func builtInStringFnc() {
	str := "Hello World";
	fmt.Println(strings.Contains(str, "Hello"));
	fmt.Println(strings.Index(str, "hello"));
	fmt.Println(strings.Count(str, "hello"));

	//replace hello with test for first 3 instances
	fmt.Println(strings.Replace(str, "hello", "test", 3));

	splitEx := "hello this is a test";
	splitExArr := strings.Split(splitEx, " ");
	fmt.Println("splitExArr: ", splitExArr, "and", splitExArr[0]);
	fmt.Println("join ex: ", strings.Join([]string {"this", "is", "a", "test"}, " "));

	//string + number
	num := 1;
	str2 := "string #" + strconv.Itoa(num);
	fmt.Println("string: ", str2);
}

func sortEx() {
	strArr := []string {"a", "d", "c"};
	sort.Strings(strArr);
	fmt.Println(strArr);
}
