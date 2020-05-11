package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func jsonPrettyPrint(in []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, in, "", "  ")
	if err != nil {
		return string(in)
	}
	return out.String()
}

var HASURA_GRAPHQL_ADDRESS = os.Getenv("HASURA_GRAPHQL_ADDRESS")
var HASURA_GRAPHQL_ADMIN_SECRET = os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")

type GraphqlQuery struct {
	Data interface{}
}

var client = &http.Client{Transport: http.DefaultTransport}

func hasura(query string, variables interface{}, data interface{}) (err error) {
	variablesBytes, err := json.Marshal(variables)
	if err != nil {
		return
	}

	v := string(variablesBytes)
	requestBody := []byte(`{"query":"` + query + `","variables":` + v + `}`)
	requestBytes := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", HASURA_GRAPHQL_ADDRESS, requestBytes)
	if err != nil {
		return
	}

	req.Header.Add("X-Hasura-Admin-Secret", HASURA_GRAPHQL_ADMIN_SECRET)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// loggin the answer for debugging purposes
	fmt.Println(jsonPrettyPrint(b))

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return errors.New(http.StatusText(resp.StatusCode))
	}

	return json.Unmarshal(b, &GraphqlQuery{Data: data})
}

type User struct {
	GithubLogin string
}

const userQuery = `
query {
	user {
		githubLogin
	}
}`

func getUsers() (users []User, err error) {
	var data map[string][]User
	err = hasura(userQuery, nil, &data)
	return data["user"], err
}

func main() {
	fmt.Println(getUsers())
}

// HASURA_GRAPHQL_ADMIN_SECRET=VERYVERYSECRET HASURA_GRAPHQL_ADDRESS=http://localhost/graphql-engine/v1alpha1/graphql go run call-graphql.go
