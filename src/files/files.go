package files;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
	"os";
	"log";
	"io/ioutil";
);

func FilesEx() {
	defer div.PrintDivider();
	tp.TitlePrint("Files");
	createFile();
	readFile();
}

func createFile() {
	file, err := os.Create("samp.txt");
	if err != nil {
		log.Fatal(err);
	}

	file.WriteString("This is some radnom text");
	file.Close();
}

func readFile() {
	stream, err := ioutil.ReadFile("samp.txt");
	if err != nil {
		log.Fatal(err);
	}
	readString := string(stream);
	fmt.Println("readstring: ", readString);
}
