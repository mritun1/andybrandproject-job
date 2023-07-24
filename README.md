# Introduction
Here you will learn how to use AndyBrandProject Golang API.

# Installation
#### Get the source code
1. First, Install Golang on your Server or PC.
2. Then download the source code from our GitHub. https://github.com/mritun1/andybrandproject-job
#### Connect with Database (MongoDB)
1. Inside the project, open the file /db/connect.go and get Change the value variable MongoURI with your MongoDB URI. And change the value of variable dbName with your Database name.
```
const dbName = "andybrandproject"
const MongoURI = "mongodb+srv://user:password@cluster.tte3tkm.mongodb.net/" + dbName
```
<br/>

# Usage
This is a REST API Golang source code for creating, reading, updating, and deleting user's data from the MongoDB database.
<br/><br/>
Here are some of the APIs that are available in the source code.
<br/>
1. GET: /users :- This call returns the available lists of users data from the MongoDB collections named Users.
2. POST: /users :- This call inserts new data into the collections.
3. PUT: /users/:id :- This call update data into the collections.
4. DELETE: /users/:id :- This call delete data from the collections.
<br/>

If you want to go more in depth about how to use the REST API, then please go through this link: https://documenter.getpostman.com/view/20669824/2s946mbqso 