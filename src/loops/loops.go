package loops;

import (
	"fmt";
	div "divider";
);

func LoopsEx() {
	defer div.PrintDivider();
	for i := 0; i <= 10; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("BuzzFizz");
		} else if i % 3 == 0 {
			fmt.Println("Buzz");
		} else if i % 5 == 0 {
			fmt.Println("Fizz");
		} else {
			fmt.Println(i);
		}
	}
}
