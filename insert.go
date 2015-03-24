package main

// Do the insert for the record.
//
func (r *Record) insert() (int64, error) {
	table := r.Table

	var cs, vs string
	for i, v := range r.Columns {
		if i == 0 {
			cs = r.Columns[0]
			vs = "\"" + r.Values[0] + "\""
			continue
		}
		cs += "," + v
		vs += "," + string('"') + r.Values[i] + string('"')
	}

	stmt := "INSERT INTO " + table + " ( " + cs + " ) " + " VALUES " + " ( " + vs + " ) "

	result, err := db.Exec(stmt)
	if err != nil {
		return 0, err
	}

	//return two value: affectedCnt, err
	return result.RowsAffected()
}
