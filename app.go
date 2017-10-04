package main

import (
    "bytes"
    "encoding/json"
    "io"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

type Classification struct {
    Class string
    Score float32
}

type App struct {
    Router   *mux.Router
    TFClient InceptionClient
}

func (a *App) Initialize(tfAddress string) {
    a.Router = mux.NewRouter()
    a.TFClient = InceptionClient{tfAddress}
    a.InitializeRoutes()
}

func (a *App) Run(addr string) {
    log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) InitializeRoutes() {
    a.Router.HandleFunc("/classify", a.classifyImage).Methods("POST")
}

func (a *App) classifyImage(w http.ResponseWriter, r *http.Request) {
    var buf bytes.Buffer
    file, _, err := r.FormFile("file")
    if err != nil {
        log.Fatalln(err)
        return
    }

    defer file.Close()
    // Copy the file data to my buffer
    io.Copy(&buf, file)

    resp, err := a.TFClient.Predict(buf.Bytes())

    // scrape classification data
    scores := resp.Outputs["scores"].GetFloatVal()
    var classifications []Classification
    for index, element := range resp.Outputs["classes"].GetStringVal() {
        score := scores[index]

        classifications = append(classifications, Classification{string(element), score})
    }

    log.Println(classifications)

    response, _ := json.Marshal(classifications)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(response)
}

func main() {
    a := App{}
    a.Initialize(os.Getenv("TF_ADDRESS"))

    a.Run(":8080")
}
