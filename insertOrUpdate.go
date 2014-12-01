package main

import (
	"errors"
	"strconv"
)

// Do the insert for the record.
// If exists then update.
//
func insertOrUpdate(r *Record) (bool, error) {

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
			return false, errors.New("insertOrUpdate: ConditionIndexes out of range.")
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
			return false, errors.New("insertOrUpdate: Insert affectedCnt is not equal 1")
		}
		return true, nil
	} else {
		conditionCols := make([]string, conditionLen)
		conditionVals := make([]string, conditionLen)
		for i, v := range conditionIndexes {
			conditionCols[i] = r.Columns[v]
			conditionVals[i] = r.Values[v] //need to be valid
		}
	
		affectedCnt, err := update(r.Table, &(r.Columns), &(r.Values), &conditionCols, &conditionVals)
		if err != nil {
			return false, err
		}
		if affectedCnt != 1 {
			return false, errors.New("insertOrUpdate: Update affectedCnt is not equal 1, But its okay.")
		}
		return true, nil
	}
}
