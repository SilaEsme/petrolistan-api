package main

import (
	"petrolistan.com/opet"
)


// func opetToData(opetResponse opet.OpetDto) []DataModel{

// 	var list []DataModel
// 	for _ ,v := range opetResponse {		
// 		for _, z := range v.Prices {
// 			list = append(list, DataModel{
// 				BrandName: "Opet",
// 				BrandCode: "OPT", 
// 				ProvinceName: v.ProvinceName, 
// 				ProvinceCode: v.ProvinceCode,
// 				ProductName: z.ProductName,
// 				Amount: z.Amount})
// 		}
// 	}
// 	return list
// }

func OpetToData(opetResponse opet.OpetDto) []DataModel2{

	var list []DataModel2
	for _ ,v := range opetResponse {
		data := DataModel2 {
			ProvinceName: v.ProvinceName,
			ProvinceCode: v.ProvinceCode,
			Brands: []Brand{
				{
					BrandName: "Opet",
					BrandCode: "OPT",
					Products: make([]Product, 0),
				},
			},
		}
		for _, z := range v.Prices {
			data.Brands[0].Products = append(data.Brands[0].Products, Product{ProductName: z.ProductName, Amount: z.Amount})			
		}
		list = append(list, data);
	}
	return list;
}