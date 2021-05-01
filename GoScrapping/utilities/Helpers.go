package utilities

import "errors"

func MakingDays(yearInit, monthInit, dayInit, yearEnd, monthEnd, dayEnd int) ([]string, error) {

	dateFormated := make([]string, 0)

	if yearInit == yearEnd {
		if monthInit < monthEnd {
			return nil, errors.New("the month init cannot be less than month end in the same year")
		}
	}

	return dateFormated, nil

}
