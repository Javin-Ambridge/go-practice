package structs;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
);

func StructsEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Structs");

	basicStruct();
}

func basicStruct() {
	tmp := coolThingy {
		name:"test", 
		arr: [5]int{1,2,3,4,5},
		slice: []bool{true},
	};
	fmt.Println("struct: ", tmp.name);
	tmp.arr[1] = 10;
	fmt.Println("struct.arr: ", tmp.arr);
	tmp.slice = append(tmp.slice, false, true);
	fmt.Println("struct.slice: ", tmp.slice);
	tmp2 := coolThingy2{
		slice: []coolThingy{
			coolThingy{
				name: "hoep this works",
				arr: [5]int{1,2},
				slice: []bool{true},
			},
		},
	};
	fmt.Println("tmp2.slice", tmp2.slice);
	newAdding := coolThingy {
		name: "1",
		arr: [5]int{100},
		slice: []bool{true, true, false},
	};
	tmp2.addOne(&newAdding);
	fmt.Println("tmp2.slice", tmp2.slice);
}

type coolThingy struct {
	name string;
	arr [5]int;
	slice []bool;
}

type coolThingy2 struct {
	slice []coolThingy;
}

//func (item *whatItsAttachedTo) functionName(params) returnType {}
func (item *coolThingy2) addOne(newOne *coolThingy) {
	item.slice = append(item.slice, *newOne);
}
