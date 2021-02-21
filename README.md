# SIN Parser Go

Library for get data of Single Identity Number (No. KTP) from applications written with Go (requires Go v1.14 or newer).

- [Installation](#installation)
- [Usage Example](#usage-example)
- [Test](#test)
- [Contributing](#contributing)

---

## Installation

Install sin-parser-go using Go Module by following command:

```bash
go get github.com/yusufthedragon/sin-parser-go
```

Then you import it by following code:

```go
import sinparser "github.com/yusufthedragon/sin-parser-go"
```

## Usage Example

```go
var parsedData, err = sinparser.ParseSIN("3204110609970001")

if err != nil {
    panic(err.Error())
}

fmt.Printf("Parsed Data: %+v\n", parsedData)

// Result:
// {
//     "Age": 23,
//     "BornDate": "06 September 1997",
//     "CityID": 3204,
//     "CityName": "KAB. BANDUNG",
//     "DistrictID": 320411,
//     "DistrictName": "KATAPANG",
//     "Gender": "Male",
//     "IsValid": true,
//     "PostalCode": "40921",
//     "ProvinceID": 32,
//     "ProvinceName": "JAWA BARAT",
//     "UniqueCode": "0001",
//     "Zodiac": "Virgo"
// }
```

## Test

You can run the test by following command:

```bash
go test -v
```

## Contributing

For any requests, bugs, or comments, please open an [issue](github.com/yusufthedragon/sin-parser-go/issues) or [submit a pull request](github.com/yusufthedragon/sin-parser-go/pulls).
