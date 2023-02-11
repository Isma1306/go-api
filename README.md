# go-api
go restful api to get to learn about it. Followed this guide https://tutorialedge.net/golang/creating-restful-api-with-golang/ and then went crazy.

## how to use it
start the api put your username and password of mongo db atlas and the api will start.
You can change the behavior from ./utils/utils.go and comment the Prompt function if you dont want to type each time you start the server.

### working paths
"/" home
"/all/" return all articles
"/article/" POST article
"/article/{id}", POST, UPDATE, DELETE

### interface

{
Content":string,
"Desc":string,
"Title":string,
"_id":string
}

enjoy!
