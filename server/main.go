package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)


type Message struct {
	Text string `json:"text"`
}

type MessageList struct {
	Messages []Message `json:"messages"`
}

var msgUrl, _= url.Parse("/get-messages")
var rootUrl, _ = url.Parse("/")
var filePath = "messanges.json"
func handleConn(w http.ResponseWriter, r *http.Request) {

	content, _ := ioutil.ReadFile(filePath)
	var messageList MessageList
	json.Unmarshal(content, &messageList)
	switch r.Method {

	case "GET": 
		http.ServeFile(w, r, "./static/index.html")


	
	case "POST": 

		t := fmt.Sprintf("[%02d:%02d:%02d] ", time.Now().Hour(), time.Now().Minute(), time.Now().Second())
		
		if strings.Trim(r.FormValue("username"), " ") == "" || strings.Trim(r.FormValue("text"), " ") == "" {
			return 
		}
		content := fmt.Sprintf("%s %s: %s", t , r.FormValue("username"),  r.FormValue("text"))

		

		newMassage := Message{Text: content}
		messageList.Messages = append(messageList.Messages, newMassage)
		
		update, _ := json.MarshalIndent(messageList, "", " ")
		
		ioutil.WriteFile(filePath, update, 0644)

		
	}

}

func handleMsg(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Open(filePath)	
	w.Header().Set("ContentType", "application/json")
	io.Copy(w, f)

}
func main() {


	http.HandleFunc("/", handleConn)
	http.HandleFunc("/get-messages", handleMsg)
	
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }

}
