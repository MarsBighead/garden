package page

import (
	"fmt"
	"garden/util"
	"html/template"
	"log"
	"net/http"
)

//AES list all handle in the top path
func AES(w http.ResponseWriter, r *http.Request) {
	log.Print("Running http handle model.AES!")
	tpl, err := template.ParseFiles("/template/aes.htm")
	if err != nil {
		log.Fatal(err)
	}
	if r.Method == "GET" {
		err = tpl.Execute(w, "")
		if err != nil {
			log.Fatal(err)
		}
	} else if r.Method == "POST" {
		err = r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		if slice, found := r.Form["Val"]; found && len(slice) > 0 {
			src := slice[0]
			crypted := util.EncryptAES(src)
			fmt.Printf("slice %v, found %v\n", src, crypted)
		}
	}
	tpl.Execute(w, nil)
}
