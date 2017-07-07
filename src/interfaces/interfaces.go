package interfaces;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
	"math";
);

func InterfacesEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Interfaces");

	rect := Rect{10,10};
	c := Circle{4};

	fmt.Println("Rectangle Radius: ", getArea(rect));
	fmt.Println("Circle Radius: ", getArea(c));
}

type Shape interface {
	area() float64;
}

type Rect struct {
	width int;
	height int;
}

type Circle struct {
	diameter int;
}

func (rect Rect) area() float64 {
	return float64(rect.width * rect.height);
}

func (c Circle) area() float64 {
	return float64(math.Pi * float64(c.diameter/2));
}

func getArea(shape Shape) float64 {
	return shape.area();
}
