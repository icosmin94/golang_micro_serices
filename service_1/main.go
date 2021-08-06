package main

import (
	"github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
)

type User struct {
	name string
	age int
}

func main() {

	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	var userResource UserResource

	ws.Route(ws.GET("/{user-id}").To(userResource.findUser).
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(User{}))

	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
