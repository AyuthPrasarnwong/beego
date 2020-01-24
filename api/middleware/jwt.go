package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"reflect"
	"strings"

	"github.com/astaxie/beego/context"
	jwt "github.com/dgrijalva/jwt-go"
)

func init() {
}

const (
	pubKeyPath  = "storage/oauth-public.key" // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

type Response struct {
	StatusCode  int `json:"status_code"`
	Message string `json:"message"`
}


func Jwt(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json")
	var uri string = ctx.Input.URI()
	if uri == "/v1/jwt" {
		return
	}

	authString := ctx.Input.Header("Authorization")

	kv := strings.Split(authString, " ")

	if len(kv) != 2 || kv[0] != "Bearer" {
		ctx.Output.SetStatus(401)

		res := map[string]Response{"errors": {
			StatusCode: 401,
			Message:  "Token not provided."}}

		resBody, _ := json.Marshal(res)
		ctx.Output.Body(resBody)
		return
	}

	tokenString := kv[1]

	//fmt.Println(tokenString)


	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 必要的验证 RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		verifyBytes, err := ioutil.ReadFile(pubKeyPath)

		if err != nil {
			return nil, err
			//log.Fatal(err)
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

		if err != nil {

			return nil, err
		}

		return verifyKey, nil
	})



	if err != nil {

		ctx.Output.SetStatus(401)

		res := map[string]Response{"errors": {
			StatusCode: 401,
			Message:  err.Error()}}

		resBody, _ := json.Marshal(res)
		ctx.Output.Body(resBody)

	}

	if token != nil {
		if token.Valid {

			//fmt.Println("token.Valid")
			//"You look nice today"

			claims, _ := token.Claims.(jwt.MapClaims)
			//var user string = claims["username"].(string)

			//fmt.Println("TypeOf", reflect.TypeOf(claims["permissions"]))
			fmt.Println("permissions", claims["permissions"])

			return

		} else if ve, ok := err.(*jwt.ValidationError); ok {

			if ve.Errors&jwt.ValidationErrorMalformed != 0 {

				fmt.Println("Token is invalid.")

			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {

				fmt.Println("Provided token is expired.")

			} else {
				//"Couldn't handle this token:"

			}
		} else {
			//"Couldn't handle this token:"

		}
	}

}
