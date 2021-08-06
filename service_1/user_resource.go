package main

import (
	"fmt"
	"github.com/emicklei/go-restful/v3"
)

type UserResource struct {

}

func (u *UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")

	_, err := response.Write([]byte(fmt.Sprintf("User %s", id)))
	if err != nil {
		return
	}

}