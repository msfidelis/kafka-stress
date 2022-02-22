package fakejson

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"encoding/json"
)

type fake struct {
	UserName           string  `faker:"username" json:"username"`
	PhoneNumber        string  `faker:"phone_number" json:"phone_number"`
	IPV4               string  `faker:"ipv4" json:"ipv4"`
	IPV6               string  `faker:"ipv6" json:"ipv6"`
	MacAddress         string  `faker:"mac_address" json:"mac_address"`
	URL                string  `faker:"url" json: "url"`
	DayOfWeek          string  `faker:"day_of_week" json: "day_of_week"`
	DayOfMonth         string  `faker:"day_of_month" json: "day_of_month"`
	Timestamp          string  `faker:"timestamp" json: "timestamp"`
	Century            string  `faker:"century" json: "century"`
	TimeZone           string  `faker:"timezone", json:"timezone"`
	TimePeriod         string  `faker:"time_period" json:"time_period"`
	Word               string  `faker:"word" json:"word"`
	Sentence           string  `faker:"sentence" json:"sentence"`
	Paragraph          string  `faker:"paragraph" json:"paragraph"`
	Currency           string  `faker:"currency" json:"currency"`
	Amount             float64 `faker:"amount" json:"amount" `
	AmountWithCurrency string  `faker:"amount_with_currency" json:"amount_with_currency"`
	UUIDHypenated      string  `faker:"uuid_hyphenated" json:"uuid_hyphenated"`
	UUID               string  `faker:"uuid_digit" json:"uuid_digit"`
	PaymentMethod      string  `faker:"oneof: cc, paypal, check, money order"`
}

// RandJSONPayload Gerenerate random json with fake data
func RandJSONPayload() string {

	a := fake{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
		return "{}"
	}

	b, err := json.Marshal(a)
    if err != nil {
        fmt.Println(err)
        return "{}"
    }
	return string(b)

}