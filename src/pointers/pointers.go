package pointers;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
);

func PointerEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Pointers");

	simpleSHit();
}

func simpleSHit() {
	x := 1;
	calculator(&x, "+");
	fmt.Println("x: ", x);
}

func calculator(x *int, math string) {
	switch math {
		case "+":
			*x += 2;
			break;
	}
}