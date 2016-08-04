package main

import (
	"fmt"

	"github.com/itsyouonline/identityserver/clients/go/itsyouonline"
)

const (
	appID  = "..."
	apiKey = "..."
)

func main() {
	klien := itsyouonline.NewItsyouonline()
	token, err := klien.LoginWithClientCredentials(appID, apiKey)
	if err != nil {
		fmt.Printf("get token failed:%v\n", err)
		return
	}
	fmt.Printf("token = %v\n", token)

	jwt, err := klien.CreateJWTToken([]string{}, []string{})
	fmt.Printf("err = %v, jwt=%v\n", err, jwt)

	gr := Newgoramldir()
	gr.AuthHeader = "token " + jwt

	_, _, err = gr.UsersGet(map[string]interface{}{}, map[string]interface{}{})
	fmt.Printf("err=%v\n", err)
}
