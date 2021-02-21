package sinparser

// SIN struct contains data of Indonesian Identity Card.
type SIN struct {
	BornDate     string
	CityID       int
	CityName     string
	DistrictID   int
	DistrictName string
	Gender       string
	IsValid      bool
	PostalCode   string
	ProvinceID   int
	ProvinceName string
	UniqueCode   string
}
