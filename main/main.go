package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"petrolistan.com/opet"
	u "petrolistan.com/utils"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/allPrices", loadHandler()).Methods("GET") //.BuildOnly().Host("petrolistan.com");

	http.ListenAndServe(":8080", r)
}

func loadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w);
		var response []DataModel2
		// Get Opet Api
		opetResponse := opet.LoadOpetHandler(w)
		response = append(response, opetToData(opetResponse)...)

		resp := u.Message(true, "başarılı")
		resp["data"] = response
		u.Respond(w, resp)
		// Get PO Api
		// Get Shell Api
	}
}

type DataModel struct {
	BrandName    string  `json:"branch_name"`
	BrandCode    string  `json:"branch_code"`
	ProductName  string  `json:"product_name"`
	Amount       float32 `json:"amount"`
	ProvinceName string  `json:"province_name"`
	ProvinceCode int     `json:"province_code"`
}

type DataModel2 struct {
	ProvinceName string  `json:"province_name"`
	ProvinceCode int     `json:"province_code"`
	Brands       []Brand `json:"brands"`
}

type Brand struct {
	BrandName string    `json:"brand_name"`
	BrandCode string    `json:"brand_code"`
	Products  []Product `json:"products"`
}
type Product struct {
	ProductName string  `json:"product_name"`
	Amount      float32 `json:"amount"`
}
