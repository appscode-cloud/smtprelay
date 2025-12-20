package main

import (
	"fmt"
	"net/url"
	"os"

	"golang.org/x/crypto/bcrypt"
)

// nolint: unused
func main_() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: hasher PASSWORD")
		os.Exit(1)
	}
	password := os.Args[1]

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error generating hash: %s\n", err)
	}
	fmt.Println(string(hash))
}

func main() {
	var u url.URL
	u.Scheme = "starttls"
	u.User = url.UserPassword("emailapikey", "admin")
	u.Host = "smtp.zeptomail.com:587"
	qr := url.Values{}
	qr.Add("auth", "login")
	u.RawQuery = qr.Encode()
	fmt.Println(u.String())
}
