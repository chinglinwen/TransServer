package main

import (
	"errors"
)

func (r1 *Record) compare(r2 *Record, columnsIndexes *[]int) (bool, error) {
	if r2 == nil {
		return false, errors.New("Compare: Record r2 is empty.")
	}

	if columnsIndexes == nil {
		return false, errors.New("Compare: columnsIndexes is empty.")
	}

	if len(r1.Values) != len(r2.Values) {
		return false, errors.New("Compare: length of Values not the same.")
	}

	for _, v := range *columnsIndexes {

		if v >= len(r1.Values) {
			return false, errors.New("Compare: columnsIndexes value out of range")
		}

		if r1.Columns[v] == r2.Columns[v] {

			//Skip the date time columns.
			switch r1.Columns[v] {
			case "id", "date", "time", "ts":
				continue
			}

			if r1.Values[v] == r2.Values[v] {
				continue
			} else {
				return false, nil
			}
		} else {
			return false, errors.New("Compare: Columns are not alignment")
		}
	}
	return true, nil
}
