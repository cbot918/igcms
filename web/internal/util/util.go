package util

import (
	"regexp"
	"strings"
)


func GetJwtToken(header string)string {
	bearer := regexp.MustCompile(`Bearer.*`).FindString(header)
	token := strings.Trim(bearer, "Bearer")
	return token
}