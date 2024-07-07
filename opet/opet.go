package opet

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoadOpetHandler(w http.ResponseWriter) OpetDto{
	opetUrl := "https://api.opet.com.tr/api/fuelprices/allprices"
	resp, err := http.Get(opetUrl)
	if err != nil {
		fmt.Errorf("ooopsss an error occurred, please try again")
	}
	defer resp.Body.Close()

	var opetResponse OpetDto

	if err := json.NewDecoder(resp.Body).Decode(&opetResponse); err != nil {
		fmt.Errorf("ooopsss! an error occurred, please try again")
	}
	
	// for _, or := range opetResponse {
	// 	for _, p := range or.Prices {
	// 		fmt.Println(w,"ProvinceName: %s\nDistrictName : %s\nProductName: %s\nAmount: %f\n\n", or.ProvinceName, or.DistrictName, p.ProductName, p.Amount)
	// 	}
	// }
	return opetResponse;
}