package main

import "errors"

// Get multiple columns or one columns
// Get multiple rows or one row.
//
func (r *Record) query(columnsIndexes, conditionIndexes *[]int) (*[]Record, error) {
	cnt, err := r.queryCnt(conditionIndexes)
	if err != nil {
		return nil, err
	}
	if cnt <= 0 {
		return nil, errors.New("There is no result")
	}

	// Zero records for now, Later it will append.
	records := make([]Record, 0)

	table := r.Table

	var columns string
	for i, v := range *columnsIndexes {
		if i == 0 {
			columns = r.Columns[v]
			continue
		}
		columns += "," + r.Columns[v]
	}

	var condition string
	for i, v := range *conditionIndexes {
		if i == 0 {
			condition += r.Columns[v] + "=" + "\"" + r.Values[v] + "\""
			continue
		}
		condition += " AND " + " " + r.Columns[v] + "=" + "\"" + r.Values[v] + "\""
	}

	stmt := "SELECT " + columns + " FROM " + table + " WHERE " + condition
	//fmt.Println("stmt is: ", stmt)

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rawValue := make([][]byte, len(cols))
	value := make([]string, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawValue {
		dest[i] = &rawValue[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err := rows.Scan(dest...)
		if err != nil {
			//fmt.Println("Failed to scan row", err)
			return nil, err
		}
		for i, raw := range rawValue {
			if raw == nil {
				value[i] = "NULL"
			} else {
				value[i] = string(raw)
			}
		}

		record := Record{table, cols, value, nil}
		records = append(records, record)
	}
	return &records, nil
}
