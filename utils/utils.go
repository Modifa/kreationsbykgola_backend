package utils

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

func MustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("%s environment variable not set.", k)
	}
	return v
}

func FormatMobileNumberToClient(message string) (string, error) {

	countrycode := message[:len(message)-9]
	fmt.Println(countrycode)

	if len(message) == 11 && countrycode == "27" {
		return "0" + message[2:11], nil
	} else if len(message) == 10 && countrycode == "0" {

		return message, nil
	} else {
		err1 := errors.New("invalid mobile number")
		return "", err1
	}
}

func FormatMobileNumber(message string) (string, error) {

	countrycode := message[:len(message)-9]
	fmt.Println(countrycode)

	if len(message) == 11 && countrycode == "27" {
		return message, nil
	} else if len(message) == 10 && countrycode == "0" {

		message = "27" + message[1:10]
		return message, nil
	} else {
		err1 := errors.New("invalid mobile number")
		return "", err1
	}
}

func IsFormatMobileNumber(value interface{}) error {

	s, _ := value.(string)
	_, err := FormatMobileNumber(s)

	if err != nil {
		return errors.New("Invalid Mobile Number")
	}
	return nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func DecryptString(ed string) string {

	dbytes, err := hex.DecodeString(ed)

	CheckError(err)

	return string(dbytes)
}

func Encrypt(data json.RawMessage) string {

	// return hex string
	return hex.EncodeToString(data)
}
