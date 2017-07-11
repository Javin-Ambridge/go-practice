package server;

import (
	"fmt";
	div "divider";
	tp "titlePrint";
	"net/http";
	"github.com/justinas/alice";
	"log";
	"time";
);

type returnType struct {
	name string;
	slice []string;
}

func ServerEx() {
	defer div.PrintDivider();
	tp.TitlePrint("HTTP");
	fmt.Println("Server starting...");
	fmt.Println("Middleware added:");
	fmt.Println("loggingHandler");
	div.PrintDivider();

	//Adding some middleware
	middleWare := alice.New(loggingHandler);

	http.Handle("/", middleWare.ThenFunc(defaultRoute));
	http.Handle("/test", middleWare.ThenFunc(testRoute));

	//Start the server on port 8080
	http.ListenAndServe(":8080", nil);
}

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	tmp := returnType {
		name: "temp return",
		slice: []string{"a", "b"},
	};
	fmt.Fprint(w, tmp);
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	tmp := returnType {
		name: "test return",
		slice: []string{"c", "d"},
	};
	fmt.Fprint(w, tmp);
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//log.Println("Request: ", r);
		t1 := time.Now();
		next.ServeHTTP(w, r);
		go func() {
			t2 := time.Now();
			log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1));
		}();
	}
	return http.HandlerFunc(fn);
}
