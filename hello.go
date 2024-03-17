package main
import "fmt"
import "net/http"
import "io"
import "strings"

func main() {
	http.HandleFunc("/google_search", func(w http.ResponseWriter, r *http.Request) {
        resp, err := http.Get("https://www.google.com")
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("error in requesting google")
			fmt.Fprintf(w, "failed to fetch request from google")
			return 
		}

		buf := new(strings.Builder)
		_, err = io.Copy(buf, resp.Body)
		fmt.Fprintf(w, buf.String())
    })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
    })
    http.ListenAndServe(":80", nil)
}
