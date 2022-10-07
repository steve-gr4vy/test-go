//https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service
//https://tutorialedge.net/golang/creating-restful-api-with-golang/
//https://eli.thegreenplace.net/2019/on-concurrency-in-go-http-servers/#:~:text=Go's%20built%2Din%20net%2Fhttp,also%20lead%20to%20some%20gotchas.
//https://go.dev/doc/tutorial/compile-install#:~:text=Add%20the%20Go%20install%20directory,specifying%20where%20the%20executable%20is.&text=Once%20you've%20updated%20the,compile%20and%20install%20the%20package.&text=Run%20your%20application%20by%20simply%20typing%20its%20name.
package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
        log.Print("starting server...")
        http.HandleFunc("/", handler)
        http.HandleFunc("/test", test)

        // Determine port for HTTP service.
        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
                log.Printf("defaulting to port %s", port)
        }

        // Start HTTP server.
        log.Printf("listening on port %s", port)
        if err := http.ListenAndServe(":"+port, nil); err != nil {
                log.Fatal(err)
        }
}

func handler(w http.ResponseWriter, r *http.Request) {
        name := os.Getenv("NAME")
        if name == "" {
                name = "World"
        }
        fmt.Fprintf(w, "Hello %s!\n", name)
}


func test(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello %s!\n", "steve")
}