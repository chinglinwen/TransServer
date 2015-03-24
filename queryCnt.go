package main

func (r *Record) queryCnt(conditionIndexes *[]int) (int64, error) {

	table := r.Table
	stmt := "SELECT count(*) FROM " + table + " WHERE "

	for i, v := range *conditionIndexes {
		if i == 0 {
			stmt += " " + r.Columns[v] + "=" + "\"" + r.Values[v] + "\""
			continue
		}
		stmt += " AND " + " " + r.Columns[v] + "=" + "\"" + r.Values[v] + "\""
	}

	var cnt int64
	err := db.QueryRow(stmt).Scan(&cnt)
	if err != nil {
		return 0, err
	}

	if cnt == 0 {
		return 0, nil
	} else {
		return cnt, nil
	}
}
