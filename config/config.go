package config

import (
	"errors"
	"os"
)

func IsServerMode() (string, bool, error) {
	args := os.Args[1:]

	switch len(args) {
	case 0:
		return ":8989", true, nil

	case 1:
		port := args[0]
		if len(port) == 5 && port[0] == ':' && isNumeric(port[1:]) {
			return port, true, nil
		} else {
			return port, true, errors.New("not a port")
		}

	case 2:
		flag := args[0]
		if flag != "-c" {
			return ":0000", false, errors.New("wrong usage")
		} else {
			port := args[1]
			if len(port) == 5 && port[0] == ':' && isNumeric(port[1:]) {
				return port, false, nil
			} else {
				return port, false, errors.New("not a port")
			}
		}
	}
	return ":0000", false, errors.New("wrong usage")
}

func isNumeric(input string) bool {
	for _, symb := range input {
		if symb < '0' || symb > '9' {
			return false
		}
	}
	return true
}
