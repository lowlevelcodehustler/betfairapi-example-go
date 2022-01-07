package main

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

func main() {
	BFlogin()
	DumbFunc()
}

func BFlogin() {

	type AuthSuccess struct {
		sessionToken string
		loginStatus  string
	}
	type AuthError struct {
		loginStatus string
	}

	client := resty.New()
	cert, err := tls.LoadX509KeyPair("client-2048.crt", "client-2048.key")
	if err != nil {
		log.Fatalf("ERROR client certificate: %s", err)
	}
	client.SetCertificates(cert)
	res, err := client.R().
		SetHeader("X-Application", "YourAppHere").
		ExpectContentType("application/x-www-form-urlencoded").
		SetResult(&AuthSuccess{}).
		SetError(&AuthError{}).
		Post("https://identitysso-cert.betfair.com/api/certlogin?username={username}&password={password}")
	fmt.Printf("%v", res)
	fmt.Printf("%v", err)
}

func DumbFunc() {
	client := resty.New()

	res, err := client.R().Get("https://api.betfair.com/exchange/betting/rest/v1.0/listEvents")
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", res.StatusCode())
	fmt.Printf("\nResponse Status: %v", res.Status())
	fmt.Printf("\nResponse Body: %v", res)
	fmt.Printf("\nResponse Time: %v", res.Time())
	fmt.Printf("\nResponse Received At: %v", res.ReceivedAt())

}
