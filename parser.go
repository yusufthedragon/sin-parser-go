package sinparser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	goage "github.com/bearbin/go-age"
)

// montList contains list of months.
var monthList = map[string]string{
	"01": "January",
	"02": "February",
	"03": "March",
	"04": "April",
	"05": "May",
	"06": "June",
	"07": "July",
	"08": "August",
	"09": "September",
	"10": "October",
	"11": "November",
	"12": "December",
}

// region contains data from parsed region.json.
var region map[string]map[string]string

// isFemale contains status gender from integer date.
var isFemale bool

// dateInt contains integer of date.
var dateInt int

// dateInt contains integer of month.
var monthInt int

func init() {
	setRegion()
}

// ParseSIN function gets the data from the given Single Identity Number.
func ParseSIN(sin string) (*SIN, error) {
	if len(sin) != 16 {
		return nil, errors.New("Invalid Single Identity Number Format")
	}

	gender, err := getGender(sin[6:10])
	if err != nil {
		return nil, err
	}

	bornDate, err := getBornDate(sin[6:12])
	if err != nil {
		return nil, err
	}

	age, err := getAge(bornDate)
	if err != nil {
		return nil, err
	}

	provinceID, err := strconv.ParseInt(sin[0:2], 10, 64)
	if err != nil {
		return nil, err
	}

	provinceName := getRegionName("province", int(provinceID))

	cityID, err := strconv.ParseInt(sin[0:4], 10, 64)
	if err != nil {
		return nil, err
	}

	cityName := getRegionName("city", int(cityID))

	districtID, err := strconv.ParseInt(sin[0:6], 10, 64)
	if err != nil {
		return nil, err
	}

	district := getRegionName("district", int(districtID))
	splitDistrict := strings.Split(district, " -- ")
	districtName := ""
	postalCode := ""

	if len(splitDistrict) == 2 {
		districtName = splitDistrict[0]
		postalCode = splitDistrict[1]
	}

	zodiac := getZodiac()

	parsedData := SIN{
		Age:          age,
		BornDate:     bornDate,
		CityID:       int(cityID),
		CityName:     cityName,
		DistrictID:   int(districtID),
		DistrictName: districtName,
		Gender:       gender,
		PostalCode:   postalCode,
		ProvinceID:   int(provinceID),
		ProvinceName: provinceName,
		UniqueCode:   sin[12:16],
		Zodiac:       zodiac,
	}

	isValid := validateSIN(parsedData)
	parsedData.IsValid = isValid

	return &parsedData, nil
}

// getAge function gets the age based on current date and born date.
func getAge(date string) (int, error) {
	// currentTime := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	// currentTime := time.Now().UTC()
	bornTime, err := time.Parse("02 January 2006", date)

	if err != nil {
		return 0, err
	}

	// interval := currentTime.Sub(bornTime)
	// age := math.Floor(interval.Hours() / 24 / 365)

	age := goage.Age(bornTime)

	return int(age), nil
}

// getBornDate function gets the born date from the given serial number.
func getBornDate(bornDate string) (string, error) {
	date := strconv.Itoa(dateInt)

	if dateInt < 10 {
		date = "0" + date
	}

	month := bornDate[2:4]
	year := bornDate[4:6]

	if year > time.Now().Format("06") {
		year = "19" + year
	} else {
		year = "20" + year
	}

	return date + " " + monthList[month] + " " + year, nil
}

// getGender function determines the gender from the given date.
func getGender(date string) (string, error) {
	err := setBornInt(date)
	if err != nil {
		return "", err
	}

	gender := "Male"

	if isFemale {
		gender = "Female"
	}

	return gender, nil
}

// getRegionName function gets the province/city/district name from region.
func getRegionName(regionType string, regionID int) string {
	ID := strconv.Itoa(regionID)

	return region[regionType][ID]
}

// getZodiac function gets the zodiac name based on month and date.
func getZodiac() string {
	if (monthInt == 1 && dateInt >= 20) || (monthInt == 2 && dateInt < 19) {
		return "Aquarius"
	} else if (monthInt == 2 && dateInt >= 19) || (monthInt == 3 && dateInt < 21) {
		return "Pisces"
	} else if (monthInt == 3 && dateInt >= 21) || (monthInt == 4 && dateInt < 20) {
		return "Aries"
	} else if (monthInt == 4 && dateInt >= 20) || (monthInt == 5 && dateInt < 21) {
		return "Taurus"
	} else if (monthInt == 5 && dateInt >= 21) || (monthInt == 6 && dateInt < 21) {
		return "Gemini"
	} else if (monthInt == 6 && dateInt >= 21) || (monthInt == 7 && dateInt < 23) {
		return "Cancer"
	} else if (monthInt == 7 && dateInt >= 23) || (monthInt == 8 && dateInt < 23) {
		return "Leo"
	} else if (monthInt == 8 && dateInt >= 23) || (monthInt == 9 && dateInt < 23) {
		return "Virgo"
	} else if (monthInt == 9 && dateInt >= 23) || (monthInt == 10 && dateInt < 24) {
		return "Libra"
	} else if (monthInt == 10 && dateInt >= 24) || (monthInt == 11 && dateInt < 23) {
		return "Scorpio"
	} else if (monthInt == 11 && dateInt >= 23) || (monthInt == 12 && dateInt < 22) {
		return "Sagitarius"
	} else if (monthInt == 12 && dateInt >= 22) || (monthInt == 1 && dateInt < 20) {
		return "Capricorn"
	}

	return "Zodiak Not Found"
}

// setBornInt function sets the date to integer to determine it's a female or not.
func setBornInt(date string) error {
	dateInteger, err := strconv.ParseInt(date[0:2], 10, 64)
	if err != nil {
		return err
	}

	if int(dateInteger) > 40 {
		dateInteger -= 40
		isFemale = true
	}

	monthInteger, err := strconv.ParseInt(date[2:4], 10, 64)
	if err != nil {
		return err
	}

	dateInt = int(dateInteger)
	monthInt = int(monthInteger)

	return nil
}

// setRegion function reads the region.json file and store the data to variable.
func setRegion() {
	jsonFile, err := os.Open("region.json")

	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &region)
}

// validateSIN function determines the result is valid or not.
func validateSIN(data SIN) bool {
	return data.Age != 0 && data.ProvinceName != "" && data.CityName != "" && data.DistrictName != "" && data.PostalCode != ""
}
