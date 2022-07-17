package mobile

import "errors"

const (
	UK = "UK"
	UA = "UA"
)

func ParsePhoneNumber(number, country string) (string, error) {
	if len(number) < 3 {
		return "", errors.New("too short number")
	}
	switch country {
	case UA:
		return "(" + UA + ")" + number[3:], nil
	case UK:
		return "(" + UK + ")" + number[3:], nil
	}
	return "", errors.New("invalid country")
}

//go build .
//go test -v -count=1
// go test -v
// go test -v -count=1  -covermode=atomic
// go test -v -count=1  -covermode=atomic --coverprofile=coverage.out
// go tool cover -func coverage.out
