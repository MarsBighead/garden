package page

import (
	"html/template"
	"log"
	"net/http"
)

// HomePage  home/index web page
func (s *Service) HomePage(w http.ResponseWriter, r *http.Request) {
	log.Printf("Template directory: %s\n", s.Environment["TEMPLATE"])

	tpl, err := template.ParseFiles(s.Environment["TEMPLATE"] + "/home.htm")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
