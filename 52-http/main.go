package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

// 00000000
// 00000001
// 00000010
// 00000100
// 00000111

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//log.Println("This log message includes date, time, and file information.")
	var (
		PORT string
	)
	PORT = os.Getenv("PORT")

	if PORT == "" {
		flag.StringVar(&PORT, "port", "9091", "--port=9091")
	}
	flag.Parse()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		//w.Write([]byte("pong"))
		fmt.Fprintln(w, "pong")
	})
	http.HandleFunc("/health", Health)

	http.HandleFunc("/user", CreateUser)

	log.Println("Started server on port:", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		runtime.Goexit()
		log.Fatalln(err.Error())
	}
}

// func fatal(msg ...any) {
// 	fmt.Println(msg)
// 	os.Exit(1)
// }

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("pong"))
	fmt.Fprintln(w, "ok")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte("somethig went wrong with the request"))
			return
		}

		user := new(User)

		err = json.Unmarshal(bytes, user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte("somethig went wrong with the request"))
			return
		}

		err = user.Validate()
		if err != nil {
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
		}
		user.LastModified = time.Now().Unix()
		_, err = WriteUserToFile("data.txt", user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte("somethig went wrong with the request"))
			return
		}

		w.WriteHeader(201)
		w.Write([]byte("User successfully created in the file"))

	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	LastModified int64  `json:"last_modified"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("empty or invalid name")
	}
	if u.Email == "" {
		return errors.New("empty or invalid email")
	}
	return nil
}

func WriteUserToFile(fineName string, user *User) (int, error) {
	f, err := os.OpenFile(fineName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	bytes, err := json.Marshal(user)
	if err != nil {
		return 0, err
	}
	return f.Write(bytes)
}
