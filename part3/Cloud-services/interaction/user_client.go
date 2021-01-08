package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	pb "./userpb"
)

func main() {
	res,err:=http.Get("http://localhost:8080")
	if err != nil {
		return
	}
	defer res.Body.Close()
	b,err:=ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	var u pb.User
	err= proto.Unmarshal(b,&u)
	if err != nil {
		return
	}
	fmt.Println(u.GetName())
	fmt.Println(u.GetId())
	fmt.Println(u.GetEmail())
}
