package Error

import (
	"fmt"
	"net/http"
)

// CheckErrorsBank can be used to handle errors, best practice - add switch over errors for every single sequence
func (c CheckError) CheckErrorsBank(errIn error) int {
	if errIn == nil {
		return http.StatusOK
	}
	fmt.Println(errIn)
	c.Logger.Errorf("", errIn.Error(), errIn)
	return http.StatusInternalServerError
}
