package garden

// Router garden router http service
type Router struct {
	Environment map[string]string
	Pattern     map[string]*Coordinate
}

//Coordinate URL mapping filesystem directory
type Coordinate struct {
	URI       string
	Directory string
}

//AddCoordinate garden route struct add Coordinate
func (r *Router) AddCoordinate() {
	home := &Coordinate{
		URI:       "/home",
		Directory: r.Environment["HOME"],
	}
	r.Pattern = map[string]*Coordinate{
		"home": home,
	}
	/*

		http.HandleFunc("/list", model.HomeList)
		http.HandleFunc("/pbt", model.Pbt)
		http.HandleFunc("/aes", model.AES)
		http.HandleFunc("/reproto", model.RebuildPbt)
		http.HandleFunc("/api/protobuf/xiaodu", model.FromXiaodu)
		http.HandleFunc("/json", model.ProtocolJSON)
		http.HandleFunc("/statistic", model.HomeStatistic)
		http.HandleFunc("/statistics", model.AdvancedStatistic)
		http.HandleFunc("/test/protocol", model.ProtocalHTTP)
		http.HandleFunc("/tpl", model.HomeTemplate)
		http.HandleFunc("/", model.Home)
	*/

}
