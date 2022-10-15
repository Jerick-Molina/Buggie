package helper

import "math/rand"

const CodeSearch = `
select CompanyId from Company where CompanyCode = ?
`

func CreateCompanyCode() string {
	codeLength := 8

	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I",
		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	code := ""

	for i := 0; i < codeLength; i++ {

		code = code + alphabet[rand.Intn(len(alphabet))]
	}

	return code
}
