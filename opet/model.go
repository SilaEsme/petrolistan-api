package opet

type OpetDto []struct {
	ProvinceCode int
	ProvinceName string
	DistrictCode string
	DistrictName string
	Prices       OpetPriceDto
}

type OpetPriceDto []struct {
	ID               string
	ProductName      string
	ProductShortName string
	Amount           float32
	ProductCode      string
}