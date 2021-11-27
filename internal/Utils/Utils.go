package Utils

import (
	errors "CentralBankTask/internal/Middleware/Error"
	"io"
	"net/http"
	"os"
	"strconv"
)

// DownloadFile func to download file by filepath and string
func DownloadFile(filepath string, url string) error {
	out, err := os.Create("files/" + filepath)
	if err != nil {
		return &errors.Errors{
			Alias: errors.ErrCreate,
		}
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return &errors.Errors{
			Alias: errors.ErrGet,
		}
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return &errors.Errors{
			Alias: errors.ErrCopy,
		}
	}

	return nil
}

// InterfaceConvertInt converting data to int or string if possible
func InterfaceConvertInt(value interface{}) (int, error) {
	var intConvert int
	var errorConvert error
	switch value.(type) {
	case string:
		intConvert, errorConvert = strconv.Atoi(value.(string))
		if errorConvert != nil {
			return errors.IntNil, &errors.Errors{
				Alias: errors.ErrAtoi,
			}
		}
		return intConvert, nil
	case int:
		intConvert = value.(int)
		return intConvert, nil
	default:
		return errors.IntNil, &errors.Errors{
			Alias: errors.ErrNotStringAndInt,
		}
	}
}
