package routine;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
	"time";
);

func routineEx() {
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i);

		time.Sleep(time.Millisecond * 500);
	}
}

func RoutineEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Routines");
	//Routines are things that run concurrently

	go routineEx();
	fmt.Println("This will be printed inside somehwere");
}