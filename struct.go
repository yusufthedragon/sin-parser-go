package sinparser

// SIN struct contains data of Indonesian Identity Card.
type SIN struct {
	Age          int
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
	Zodiac       string
}
