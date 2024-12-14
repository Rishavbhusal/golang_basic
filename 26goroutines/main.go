





package main



import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("HELLO")
	// greeter("World")

	webSitelist := []string{

		"https://go.dev",
		"https://google.com",
		"https://fb.com",
	}
	for _, web := range webSitelist {
		wg.Add(1)
		getStatusCode(web)

	}
	wg.Wait()
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2 * time.Second)
// 		fmt.Println(s)
// 	}

// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oppss wrong ")
	} else {
		mut.Lock()

		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code fo r %s \n", result.StatusCode, endpoint)
	}
}
