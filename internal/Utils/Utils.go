package Utils

import (
	errors "CentralBankTask/internal/Middleware/Error"
	"CentralBankTask/internal/domain"
	"strconv"
	"strings"
	"time"
)

func ConvertTimeToString(date time.Time, separator string) string {
	splitedTime := strings.Split(date.String(), " ")
	formatedDate := strings.Split(splitedTime[0], "-")
	return formatedDate[2] + separator + formatedDate[1] + separator + formatedDate[0]
}

func GetDate(date time.Time) domain.DateInterval {
	var dateFormat domain.DateInterval
	dateFormat.BegD = date
	dateFormat.EndD = date.AddDate(0, 0, -DaysAmount)
	return dateFormat
}

func ConvertStringToTime(date string) time.Time {
	dateSlice := strings.Split(date, ".")
	DD, _ := InterfaceConvertInt(dateSlice[0])
	MM, _ := InterfaceConvertInt(dateSlice[1])
	YY, _ := InterfaceConvertInt(dateSlice[2])
	convertedTime := time.Date(YY, time.Month(MM), DD, 0, 0, 0, 0, time.UTC)
	return convertedTime
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
