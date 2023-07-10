package tasks

import (
	"errors"
	"fmt"
	"strings"
)

// AddUniqueUserToMap validates and filters of an input user data, and save the user if
// the user is unique, otherwise return an error
func AddUniqueUserToMap(userData map[string]string, username, phone string) error {
	// check userData on empty
	if len(userData) == 0 {
		return errors.New("error, userData is empty")
	}

	// Validation of input data (username and phone)
	username, phone, err := validateNameAndPhoneNumber(username, phone)
	if err != nil { // error checking
		return err
	}

	// filter of input data (username and phone),
	username, phone, err = filterUsernameAndPhoneNumber(username, phone)
	if err != nil {
		return err
	}

	// save the unique user to map
	userData, err = saveUserToMap(userData, username, phone)
	if err != nil {
		return err
	}

	// print the user data on console
	PrintUserData(userData)

	return nil
}

// PrintUserData on console
func PrintUserData(userData map[string]string) {
	for username, phone := range userData {
		fmt.Printf("User, userusername: %s, phone: %s\n", username, phone)
	}
}

func saveUserToMap(userData map[string]string, username, phone string) (map[string]string, error) {
	// check unique values
	if _, ok := userData[username]; ok {
		return nil, errors.New("error, is not unique username or phone number")
	}

	userData[username] = phone

	return userData, nil
}

// validateNameAndPhoneNumber input data (checking for the presence of emptiness, incorrect length of the name must be
// more than 4 letters, number phone number more than 6 digits, phone number format)
func validateNameAndPhoneNumber(username, number string) (isValidUsername, isValidPhone string, err error) {
	// validate username
	if len(username) == 0 || len(username) < 4 {
		return "", "", errors.New("error, invalid username of the user")
	}
	isValidUsername = username

	// validate phone number
	if len(number) == 0 || len(number) < 6 {
		return "", "", errors.New("error, invalid phone number")
	}
	isValidPhone = number

	return isValidUsername, isValidPhone, nil
}

// filterUsernameAndPhoneNumber, removes extra characters, and brings them into one form, otherwise an error
func filterUsernameAndPhoneNumber(username string, phone string) (string, string, error) {
	filteredUsername := filterData(username)
	filteredPhoneNumber := filterData(phone)

	return filteredUsername, filteredPhoneNumber, nil
}

// filterData removes extra characters, and brings them into one form
func filterData(data string) (filteredData string) {
	data = strings.ReplaceAll(data, "-", "")
	data = strings.ReplaceAll(data, "+", "")
	data = strings.ReplaceAll(data, ")", "")
	data = strings.ReplaceAll(data, "(", "")
	data = strings.ReplaceAll(data, " ", "")
	data = strings.ReplaceAll(data, "_", "")
	data = strings.ToLower(strings.TrimSpace(data))

	return data
}
