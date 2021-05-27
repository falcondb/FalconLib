package exercise

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

//func main() {
//
//	// http example
//	http.HandleFunc("/hello", hello)
//	http.HandleFunc("/headers", headers)
//
//	http.ListenAndServe(":8090", nil)
//
//	resp, _ := http.Get("localhost:8090/hello")
//	defer resp.Body.Close()
//	body, _ := io.ReadAll(resp.Body)
//	resp.Body.Close()
//	fmt.Printf("%s", body)
//
//	// file example
//	file, _ := os.Open("sample.txt")
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	scanner.Split(bufio.ScanLines)
//
//	for scanner.Scan() {
//		fmt.Printf("%s", scanner.Text())
//	}
//
//	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
//	scanner = bufio.NewScanner(strings.NewReader(input))
//	scanner.Split(bufio.ScanWords)
//	// Count the words.
//	count := 0
//	for scanner.Scan() {
//		count++
//	}
//
//	w := bufio.NewWriter(os.Stdout)
//	defer w.Flush()
//	fmt.Fprint(w, "Hello, ")
//
//}
