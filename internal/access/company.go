package access

// func InsertCompany(compName string) (int, int) {

// 	var companyId int
// 	sqlExc := "insert into Company (Name,CompanyCode) values (?,?)"
// 	sqlQuery := "select CompanyId from Company where CompanyCode = ?"
// 	compCode := CreateCompanyCode()
// 	result, err := db.Exec(sqlExc, compName, compCode)
// 	if err != nil {
// 		fmt.Println(err.Error())

// 		return 0, http.StatusInternalServerError
// 	}
// 	fmt.Println(result)

// 	if err := db.QueryRow(sqlQuery, compCode).Scan(&companyId); err != nil {
// 		if err == sql.ErrConnDone {
// 			return 0, http.StatusUnauthorized
// 		}
// 	}

// 	return companyId, http.StatusOK
// }

// func CreateCompanyCode() string {
// 	codeLength := 8

// 	alphabet := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I",
// 		"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

// 	code := ""

// 	for i := 0; i < codeLength; i++ {

// 		code = code + alphabet[rand.Intn(len(alphabet))]
// 	}

// 	return code
// }
