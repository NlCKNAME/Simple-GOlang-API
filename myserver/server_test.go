package myserver

//import of used packages
import (
	"fmt"

	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/require"
)

// TestServer Do the HTTP tests and test our server
func TestServer(t *testing.T) {
	require := require.New(t)

	// Create our service
	server, err := NewServer()
	require.NoError(err)

	// Start this service in a http test server
	ts := httptest.NewServer(server)
	defer ts.Close()

	// Print URL of our service
	fmt.Println(ts.URL)

	// Get version
	res, err := http.Get(ts.URL + "/version")
	require.NoError(err)
	require.Equal(200, res.StatusCode)
	// TODO : Maye I should test if i got the expected version (variable Version)

	res, err = http.Get(ts.URL + "/users")
	require.NoError(err)
	require.Equal(200, res.StatusCode)
	// Well request is oK... but do i have a list of users?
	// Yes :) 

	res, err = http.Post(ts.URL+"/NewUser/user=Toto/id=1/", "application/json", nil)
	require.NoError(err)
	require.Equal(200, res.StatusCode)
	fmt.Println(res.StatusCode)

	//Create user which already exist
	res, err = http.Post(ts.URL+"/NewUser/user=Toto/id=1/", "application/json", nil)
	require.NoError(err)
	require.Equal(403, res.StatusCode)
	fmt.Println(res.StatusCode)

	//Create user which already exist
	res, err = http.MethodDelete(ts.URL+"/delete/1", "application/json", nil)
	require.NoError(err)
	require.Equal(200, res.StatusCode)
	fmt.Println(res.StatusCode)

	//Create user without id specified
	res, err = http.Post(ts.URL+"/NewUser/user=Titi/id=/", "application/json", nil)
	require.NoError(err)
	require.Equal(400, res.StatusCode)
	fmt.Println(res.StatusCode)
}

func TestOldFunctions(t *testing.T) {
	require := require.New(t)

	// Set our users to a defined state
	users = []User{User{ID: 1, Nom: "test"}, User{ID: 2, Nom: "test"}}

	// we should have 2 users
	require.Len(users, 2)

	// We remove One...
	remove(1)

	// We should have one now
	require.Len(users, 1)

}

func TestFunctions(t *testing.T) {
	require := require.New(t)

	// Create our service
	server, err := NewServer()
	require.NoError(err)
	// Note : we doesn't start any HTTP server !

	// Se should have 0 users
	require.Len(server.users, 0)
	// We are asking for the number of users IN our service

	err = server.AddUser(User{ID: 1, Nom: "ToTo"})
	require.NoError(err)

	// So, was the user created ?
	require.Len(server.users, 1, "I should have at least one user !")
	require.Equal("ToTo", server.users[0].Nom)

	// Try add one more
	err = server.AddUser(User{ID: 2, Nom: "TaTa"})
	require.NoError(err)
	require.Len(server.users, 2, "I should have 2 user !")

	// Now delete !
	err = server.DeleteUser(1)
	require.NoError(err)
	require.Len(server.users, 1, "I should have 1 user !")

	// We try to delete a non existing user
	// How can I test that ??
	err = server.DeleteUser(1)
	require.NoError(err)
	require.Len(server.users, 1, "I should have 1 user again !")
	// Maybe like that ? ¯\_(ツ)_/¯

}