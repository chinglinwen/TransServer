package main

import (
	"errors"
	"strconv"
)

func processIpCoreExtra(r *Record) (bool, error) {

	//Suppose to get from curl url
	//	0 = localid, 9= status
	//	conditionIndexes := &[]int{0, 9}
	//
	conditionLen := len(r.Condition)
	conditionIndexes := make([]int, conditionLen)
	for i, err := 0, errors.New(""); i < conditionLen; i++ {
		conditionIndexes[i], err = strconv.Atoi(r.Condition[i])
		if err != nil {
			return false, err
		}
		if conditionIndexes[i] >= len(r.Columns) {
			return false, errors.New("ConditionIndexes out of range.")
		}
	}

	//Default columnsIndexes is for all columns.
	//
	columnsLen := len(r.Columns)
	columnsIndexes := make([]int, columnsLen)
	for i, _ := range r.Columns {
		columnsIndexes[i] = i
	}

	cnt, err := r.queryCnt(&conditionIndexes)
	if err != nil {
		return false, err
	}

	if cnt == 0 {
		affectedCnt, err := r.insert()
		if err != nil {
			return false, err
		}
		if affectedCnt != 1 {
			return false, errors.New("Insert affectedCnt is not equal 1")
		}
		return true, nil
	} else if cnt == 1 {
		records, err := r.query(&columnsIndexes, &conditionIndexes)
		if err != nil {
			return false, err
		}

		recordsCnt := len(*records)

		if recordsCnt == 0 {
			return false, errors.New("Query get zero record, should never get here.")
		} else if recordsCnt == 1 {
			isSame, err := r.compare(&(*records)[0], &columnsIndexes)
			if err != nil {
				return false, err
			}
			if isSame {
				return false, errors.New("No value is changed, Its okay.")
			}
			setCols := &[]string{"status"}
			setVals := &[]string{"invalid"}

			conditionCols := make([]string, conditionLen)
			conditionVals := make([]string, conditionLen)
			for i, v := range conditionIndexes {
				conditionCols[i] = r.Columns[v]
				conditionVals[i] = r.Values[v] //need to be valid
			}

			affectedCnt, err := update(r.Table, setCols, setVals, &conditionCols, &conditionVals)
			if err != nil {
				return false, err
			}
			if affectedCnt != 1 {
				return false, errors.New("Update affectedCnt is not equal 1")
			}

			affectedCnt, err = r.insert()
			if err != nil {
				return false, err
			}
			if affectedCnt != 1 {
				return false, errors.New("After update, insert affectedCnt is not equal 1")
			}
			return true, nil
		} else {
			return false, errors.New("Can't compare with Multiple records. Ignore this record.")
		}
	} else {
		return false, errors.New("Multiple valid localid exists. Ignore this record.")
	}

}

func processSyscheckResult(r *Record) (bool, error) {
	affectedCnt, err := r.insert()
	if err != nil {
		return false, err
	}
	if affectedCnt != 1 {
		return false, errors.New("Insert affectedCnt is not equal 1")
	}
	return true, nil
}

func processDefault(r *Record) (bool, error) {
	return processSyscheckResult(r)
}
