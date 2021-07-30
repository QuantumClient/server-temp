package util

import "regexp"

func IsUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func Alphanumeric3p(s string) bool {
	r := regexp.MustCompile(`^[a-zA-Z0-9_]{3,20}$`)
	return r.MatchString(s)
}
