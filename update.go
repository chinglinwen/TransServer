package main

import (
	"errors"
)

// Update specific columns
//
//	setCols:=&[]string{"status"}
//	setVals:=&[]string{"Invalid"}
//
//	conditionCols:=&[]string{"localid","status"}
//	conditionVals:=&[]string{"7320","Valid"}
//
//	update(r.table,setCols,setVals,conditionCols,conditionVals)
//
// Update whole record
//	update(r.table,r.Columns,r.Values,conditionCols,conditionVals)
//
func update(table string, setCols, setVals *[]string,
	conditionCols, conditionVals *[]string) (int64, error) {

	if table == "" {
		return 0, errors.New("Update error: table name is empty.")
	}

	if setCols == nil || setVals == nil {
		return 0, errors.New("Update error: set columns is empty.")
	}

	if conditionCols == nil || conditionVals == nil {
		return 0, errors.New("Update error: condition columns is empty.")
	}

	if len(*setCols) != len(*setVals) {
		return 0, errors.New("Update error: length of setCols and setVals is not equal.")
	}

	if len(*conditionCols) != len(*conditionVals) {
		return 0, errors.New("Update error: length of conditionCols and conditionVals is not equal.")
	}

	var setStmt string
	for i, v := range *setCols {
		if i == 0 {
			setStmt = v + "=" + "\"" + (*setVals)[i] + "\""
			continue
		}
		setStmt += "," + v + "=" + "\"" + (*setVals)[i] + "\""
	}

	var conditionStmt string
	for i, v := range *conditionCols {
		if i == 0 {
			conditionStmt = v + "=" + "\"" + (*conditionVals)[i] + "\""
			continue
		}
		conditionStmt += " AND " + v + "=" + "\"" + (*conditionVals)[i] + "\""
	}

	stmt := "UPDATE " + table + " SET " + setStmt + " WHERE " + conditionStmt

	result, err := db.Exec(stmt)
	if err != nil {
		return 0, err
	}

	affectedCnt, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedCnt, nil
}
