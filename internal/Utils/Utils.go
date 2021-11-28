package Utils

import (
	errors "CentralBankTask/internal/Middleware/Error"
	"CentralBankTask/internal/domain"
	"strconv"
	"strings"
	"time"
)

func ConvertDateToString(date domain.Date) string {
	return strconv.Itoa(date.DD) + "-" + strconv.Itoa(date.MM) + "-" + strconv.Itoa(date.YY)
}

func ConvertTimeToString(date time.Time) string {
	formatedDate := strings.Split(strings.Split(date.String(), " ")[0], "-")
	return formatedDate[2] + "/" + formatedDate[1] + "/" + formatedDate[0]
}

func GetDate(date string) domain.DateInterval {
	var dateFormat domain.DateInterval
	dateSlice := strings.Split(date, "-")
	dateFormat.Beg.DD, _ = InterfaceConvertInt(dateSlice[0])
	dateFormat.Beg.MM, _ = InterfaceConvertInt(dateSlice[1])
	dateFormat.Beg.YY, _ = InterfaceConvertInt(dateSlice[2])
	dateFormat.BegD = time.Date(dateFormat.Beg.YY, time.Month(dateFormat.Beg.MM), dateFormat.Beg.DD, 0, 0, 0, 0, time.UTC)

	dateOtherFormat := strconv.Itoa(dateFormat.Beg.YY) + "-" + strconv.Itoa(dateFormat.Beg.MM) + "-" + strconv.Itoa(dateFormat.Beg.DD)
	parsedDate, _ := time.Parse("2006-01-02", dateOtherFormat)
	endYear, endMonth, endDay := parsedDate.AddDate(0, 0, -DaysAmount).Date()
	dateFormat.End.DD, _ = InterfaceConvertInt(endDay)
	dateFormat.End.MM = int(endMonth)
	dateFormat.End.YY, _ = InterfaceConvertInt(endYear)
	dateFormat.EndD = time.Date(dateFormat.End.YY, time.Month(dateFormat.End.MM), dateFormat.End.DD, 0, 0, 0, 0, time.UTC)
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
