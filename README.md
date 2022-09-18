# [Simple Server in GO]

Simple server made in GO.

When the server is started, it run as http://localhost:3334

The possible path was :
- `/version/` : Display the current version of the server
- `/users/` : Display the current users in the database
- `/users/{id}` : Display the specific user's data
- `/users/NewUser/user={UserName}/id={UserID}` : Create new user
- `/users/delete/{UserID}` : Delete a user

`API_Client` is used to test the server's fonctionnalities.