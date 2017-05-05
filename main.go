package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

var pat string

func init() {
	flag.StringVar(&pat, "token", "", "digital ocean token")
	flag.Parse()
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {

	tokenSource := &TokenSource{
		AccessToken: pat,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)

	ctx := context.TODO()

	keys, _, err := client.Keys.List(ctx, nil)
	if err != nil {
		panic(err)
	}

	for _, k := range keys {
		fmt.Printf("%+v\n", k.ID) // output for debug
	}

}
