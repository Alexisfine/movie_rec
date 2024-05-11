package main
import "fmt"
import "net/http"
import "io"
import "strings"
import "os"

func main() {
	ip := os.Args[1]
	/* TODO: Validate IP */
	
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

	http.HandleFunc("/check_es", func(w http.ResponseWriter, r *http.Request) {
		// send http request to es 
		resp, err := http.Get("http://" + ip + ":8080" + "/search/movie_catalog/_search")

		if err != nil {
			fmt.Printf("failed to get movie_catalog from elasticsearch {}", err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "server internal error encountered, please try again one or two minutes later or contact admin")
			return 
		}
		// handle result 
		defer resp.Body.Close()
		buf := new(strings.Builder)
		_, err = io.Copy(buf, resp.Body)
		if err != nil {
			fmt.Printf("failed to convert elasticsearch http response body to a local string. The actual error is:{}", err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "server internal error encountered, please try again one or two minutes later or contact admin")
			return 
		}
		fmt.Fprintf(w, buf.String())
    })

    http.ListenAndServe(":80", nil)
}
