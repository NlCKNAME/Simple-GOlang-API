package main

//import of used packages
import (
	"net/http"
	"encoding/json"
	"bytes"
	"fmt"

	_ "io"
	"io/ioutil"
	_ "net/http"
	_"net/http/httptest"
)

type user struct {
	ID   string   `json:"ID"`
    Nom string `json:"Nom"`
}

type version struct {
	Version string `json:"Version"`
}

func main() {

	//Every function's tests
	fmt.Println(ShVersion())

	User := &user{Nom: "Toto", ID: "1"}
  
	fmt.Println(AddUser(User))
  
	fmt.Println(GetUsers())

	User = &user{Nom: "Toto", ID: "2"}

	fmt.Println(AddUser(User))

	fmt.Println(GetUsers())

	fmt.Println(DeleteUser("2"))

	fmt.Println(GetUsers())

	fmt.Println(DeleteUser("2"))

	fmt.Println(DeleteUser("1"))
}

func GetUsers() []user {
	resp, err := http.Get("http://localhost:3334/users")
	if err != nil {
		panic(err)
	}

	//Read the Response's body
	body, err := ioutil.ReadAll(resp.Body)

	//Create new tab of users
	var Users []user

	//Parse the JSON data
	err = json.Unmarshal(body, &Users)

	if err != nil {
		panic(err)
	}

	//Close the request
	defer resp.Body.Close()

	//Return Users's tab
	return Users
}

func AddUser(User *user) user {
	//Encode User in JSON
	encodedJson, err := json.Marshal(User)
  
	if err != nil {
	  panic(err)
	}
	resp, err := http.Post("http://localhost:3334/users/NewUser/user=" + User.Nom  + "/id=" + User.ID, "application/json", bytes.NewBuffer(encodedJson))
  
	if err != nil {
	  panic(err)
	}
  
	//Read the Response's body
	body, err := ioutil.ReadAll(resp.Body)
  
	var returnedUser user

	//Parse the JSON data
	err = json.Unmarshal(body, &returnedUser)

	if err != nil {
	  panic(err)
	}

	//Close the request
	defer resp.Body.Close()
  
	//Return the created User
	return returnedUser
}

func DeleteUser(id string) int{

    //Create the http client
    client := &http.Client{}

    //Create the DELETE request whith the specific ID
    req, err := http.NewRequest("DELETE", "http://localhost:3334/users/delete/" + id, nil)
    if err != nil {
        fmt.Println(err)
        return 1
    }

    //Get Request's responce
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return 2
    }
    defer resp.Body.Close()

    //Read Response Body
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        return 3
    }

    // Display Results
    fmt.Println("response Code : ", resp.Status)
    fmt.Println("response Headers : ", resp.Header)
    fmt.Println("response Content : ", string(respBody))
	
	return 0
}

func ShVersion() version{
	resp, err := http.Get("http://localhost:3334/version")
	if err != nil {
		panic(err)
	}

	//Read the Response's body
	body, err := ioutil.ReadAll(resp.Body)

	//Create var Version
	var Ver version

	//Parse the JSON data
	err = json.Unmarshal(body, &Ver)

	if err != nil {
		panic(err)
	}

	//Close the request
	defer resp.Body.Close()

	//Return Version
	return Ver
}