package main	

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
	"regexp"
	"strconv"
	"time"
	"crypto/md5"
	"io"
	"os"
)

func helloWorld( w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

// logic part of log in
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method: ", r.Method) // get request method
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))  // Length of the token
		token := fmt.Sprintf("%x", h.Sum(nil))


		t, _ := template.ParseFiles("main.html")
		t.Execute(w, token)
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("Checking Token")
		} else {
			fmt.Println("No Token")
		}

		// validate that user has been entered
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("Need To Enter Username")
		}
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))

		// validate that an @ is used for email
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("Need @ Sign")
		}

		// Protext again cross site scripting
		fmt.Println("email:", template.HTMLEscapeString(r.Form.Get("email")))

		// validate the inputs so people don't make up stuff
		yearSlice := []string{"freshmen", "sophmore", "junior", "senior"}

		for _, v := range yearSlice {
			if v == r.Form.Get("year") {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		}

		fmt.Println("year:", template.HTMLEscapeString(r.Form.Get("year")))

		fmt.Println("password:", r.Form["password"])

		template.HTMLEscape(w, []byte(r.Form.Get("username"))) // responded to clients
	}
}

// upload logic
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":8080", nil)
	
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}