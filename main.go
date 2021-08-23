package main

import(
    "net/http"
    "fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello world")
}

func main() {
    http.HandleFunc("/", hello);
    err := http.ListenAndServe(":8000", nil);
    if err != nil {
            
    }
}


// func main(){
// 	var r = 0

// 	startTime := time.Now().UnixNano()
//     for i := 0;i < 1000000000;i++{
// 		r += i
// 	}
// 	endTime := time.Now().UnixNano()
// 	Milliseconds:= float64((endTime - startTime) / 1e6)
// 	fmt.Println(Milliseconds)

// }