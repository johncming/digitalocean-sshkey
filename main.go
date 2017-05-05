package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

var do_token string

func init() {
	flag.StringVar(&do_token, "token", "", "digital ocean token")
	flag.Parse()
}

func main() {
	url := "https://api.digitalocean.com/v2/account/keys"

	req, err := http.NewRequest("GET", url, http.NoBody)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", do_token))

	cli := &http.Client{}

	res, err := cli.Do(req)

	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		fmt.Println("fail to get token")
	}

	dec := json.NewDecoder(res.Body)

	type Msg struct {
		Key   string
		Value string
	}

	// open bracket
	_, err = dec.Token()
	if err != nil {
		panic(err)
	}

	for dec.More() {
		var m Msg
		err := dec.Decode(&m)
		if err != nil {
			panic(err)
		}

		if m.Key == "id" {
			fmt.Printf("DIGITAL OCEAN TOKEN: %s", m.Value)
		}

	}

	// close bracket
	_, err = dec.Token()
	if err != nil {
		panic(err)
	}

}
