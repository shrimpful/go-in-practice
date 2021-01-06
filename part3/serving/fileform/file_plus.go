package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",fileForm)
	http.ListenAndServe(":8080",nil)
}

func fileForm(w http.ResponseWriter, r *http.Request) {
	if r.Method==http.MethodGet{
		t,_:=template.ParseFiles("file_plus.html")
		t.Execute(w,nil)
	}else {
		mr,err:=r.MultipartReader()
		if err != nil {
			panic("Failed to read multipart message")
		}
		values:=make(map[string][]string)
		maxValueBytes:=int64(10<<20)
		for{
			part,err:=mr.NextPart()
			if err ==io.EOF {
				break
			}

			name:=part.FormName()
			if name==""{
				continue
			}
			filename:=part.FileName()
			var b bytes.Buffer
			if filename==""{
				n,err:=io.CopyN(&b,part,maxValueBytes)
				if err!=nil &&err!=io.EOF{
					fmt.Fprint(w,"Error processing form")
					return
				}
				maxValueBytes-=n

				if maxValueBytes==0{
					msg:="multipart message too large"
					fmt.Fprint(w,msg)
					return
				}
				values[name]=append(values[name],b.String())
				continue
			}
			dst,err:=os.Create("/tmp/"+filename)
			defer dst.Close()
			if err != nil {
				return
			}
			for{
				buffer:=make([]byte,100000)
				cBytes,err:=part.Read(buffer)
				if err ==io.EOF {
					break
				}
				dst.Write(buffer[0:cBytes])
			}
		}
		fmt.Fprint(w,"uploaded")
	}
}