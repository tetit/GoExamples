package mobile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParsePhoneNumberSuccess(t *testing.T) {

	actual, err := ParsePhoneNumber("+38012333345", "UA")
	assert.Equal(t, "(UA)012333345", actual)
	assert.NoError(t, err)
}
func Test_ParsePhoneNumberInvalidCountry(t *testing.T) {

	actual, err := ParsePhoneNumber("+38012333345", "HY")
	assert.Equal(t, "", actual)
	assert.EqualError(t, err, "invalid country")

	// got, _ := ParsePhoneNumber("+38012333345", "UA")
	// t.Log(got)
	// exp := "(UA)012333345"
	// got1, _ := ParsePhoneNumber("+3", "UK")
	// assert.Equal(t, exp, got)
	// assert.Equal(t, "", got1)

	// if got != exp {
	// 	t.FailNow()
	// }
}

func TestParsePohoneNuber(t *testing.T) {
	tests := map[string]struct {
		number  string
		country string
		exp     string
		expErr  string
	}{
		"success UA": {
			number:  "+38012333345",
			country: "UA",
			exp:     "(UA)012333345",
			expErr:  "",
		},
		"invalid country": {number: "+38012333345", country: "NY", exp: "", expErr: "invalid country"},
		"invalid lenght": {number: "+3", country: "NY", exp: "", expErr: "too short number"},
		"succeess UK": {number: "+28012333345", country: "UK", exp: "(UK)012333345", expErr: ""},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {

			actual, err := ParsePhoneNumber(tt.number, tt.country)
			assert.Equal(t, tt.exp, actual)

			if tt.expErr != "" {
				assert.EqualError(t, err, tt.expErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
