package licence

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Please implement a scruct that can produce driving licence numbers using TDD.
//
// A drving licence number it produced by taking the applicants initials, date of birth (YYYYMMDD)
// and a 3 digit random number.
//
// i.e
//     Mark David Bradley DOB: 12/05/1997 driving licence number could be MDB19970512999
//     Harry Jim James Smith DOB 09/10/1985 driving licence number could be HJJS19851009999
//     Jane Bond 01/01/2001 driving licence number could be JB20010101999
//
// You will need to create a stub test double of the RandomNumberGenerator interface to ensure
// tests always pass

type mockRandomNumberGenerator struct{}

func (r mockRandomNumberGenerator) createRandomNumber(numberOfDigits int) string {
	return "999"
}

func TestItCanGenerateLicenseNumberForMark(t *testing.T) {
	mockRNG := mockRandomNumberGenerator{}
	var testCases = []struct {
		licenseNumberGenerator LicenseNumberGenerator
		licenseNumber          string
	}{
		{
			licenseNumberGenerator: LicenseNumberGenerator{
				Person{"Mark David Bradley", time.Date(1997, 05, 12, 0, 0, 0, 0, time.Local)},
				mockRNG,
			},
			licenseNumber: "MDB19970512999",
		},
		{
			licenseNumberGenerator: LicenseNumberGenerator{
				p:            Person{"Harry Jim James Smith", time.Date(1985, 10, 9, 0, 0, 0, 0, time.Local)},
				randomNumber: mockRNG,
			},
			licenseNumber: "HJJS19851009999",
		},
		{
			licenseNumberGenerator: LicenseNumberGenerator{
				p:            Person{"Jane Bond", time.Date(2001, 01, 01, 0, 0, 0, 0, time.Local)},
				randomNumber: mockRNG,
			},
			licenseNumber: "JB20010101999",
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, testCase.licenseNumber, testCase.licenseNumberGenerator.GenerateLicenseNumber())
	}
}
