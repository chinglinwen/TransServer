package main

import (
	"fmt"
	"net/http"
	"strings"
)

func insertHandle(w http.ResponseWriter, req *http.Request) {
	table := req.FormValue("table")
	columns := strings.Split(req.FormValue("columns"), ",")
	values := strings.Split(req.FormValue("values"), ",")
	condition := strings.Split(req.FormValue("condition"), ",")
	way := req.FormValue("way")

	if table == "" {
		fmt.Fprintf(w, "Table name is missing.\n")
		return
	}
	if fmt.Sprintf("%v", columns) == "[]" || fmt.Sprintf("%v", values) == "[]" {
		fmt.Fprintf(w, "Columns or Values string it is not present.\n")
		return
	}
	if len(columns) != len(values) {
		fmt.Fprintf(w, "", columns, values)
		fmt.Fprintf(w, "Columns and Values it is not alignment.\n")
		return
	}

	record := &Record{table, columns, values, condition, way}
	//logger.Printf("Got record: %v\n", record)

	switch table {
	case "ip_core", "ip_extra":
		doCompareInsertUpdate(w, record)
	case "ip_conn", "syscheck_result":
		doDirectInsert(w, record)
	case "FST_TAB_cnt":
		doInsertUpdate(w, record)
	default:
		otherWay(w, record)
	}
}

func otherWay(w http.ResponseWriter, r *Record) {
	switch r.Way {
	case "direct_insert":
		doDirectInsert(w, r)
	case "insert_update":
		doInsertUpdate(w, r)
	case "compare_insert_update":
		doCompareInsertUpdate(w, r)
	default:
		doDirectInsert(w, r)
	}
}

func doCompareInsertUpdate(w http.ResponseWriter, r *Record) {
	if fmt.Sprintf("%v", r.Condition) == "[]" {
		fmt.Fprintf(w, "Condition is mising.\n")
		return
	}

	if ok, err := compareInsertUpdate(r); err != nil {
		fmt.Fprintf(w, "compareInsertUpdate error: %v\n", err)
	} else if ok {
		fmt.Fprintf(w, "okay\n")
		logger.Printf("Table: %v, Localid: %v is processed okay.\n", r.Table, r.Values[0])
	} else {
		fmt.Fprintf(w, "There is no error, but it is not ok, Should never get here.\n")
		logger.Printf("There is no error, but it is not ok, Should never get here.\n")
	}
}

func doDirectInsert(w http.ResponseWriter, r *Record) {
	if ok, err := directInsert(r); err != nil {
		fmt.Fprintf(w, "directInsert error: %v\n", err)
	} else if ok {
		fmt.Fprintf(w, "okay\n")
		logger.Printf("Table: %v, Localid: %v,round: %v is processed okay.\n",
			r.Table, r.Values[0], r.Values[1])
	} else {
		fmt.Fprintf(w, "There is no error, but it is not ok, Should never get here.\n")
		logger.Printf("There is no error, but it is not ok, Should never get here.\n")
	}
}

func doInsertUpdate(w http.ResponseWriter, r *Record) {
	if fmt.Sprintf("%v", r.Condition) == "[]" {
		fmt.Fprintf(w, "Condition is mising.\n")
		return
	}

	if ok, err := insertUpdate(r); err != nil {
		fmt.Fprintf(w, "%v\n", err)
	} else if ok {
		fmt.Fprintf(w, "okay\n")
		logger.Printf("Table: %v, Localid: %v is processed okay.\n", r.Table, r.Values[0])
	} else {
		fmt.Fprintf(w, "There is no error, but it is not ok, Should never get here.\n")
		logger.Printf("There is no error, but it is not ok, Should never get here.\n")
	}
}

func queryHandle(w http.ResponseWriter, req *http.Request) { 
	table := req.FormValue("table")
	columns := strings.Split(req.FormValue("columns"), ",")
	values := strings.Split(req.FormValue("values"), ",")
	condition := strings.Split(req.FormValue("condition"), ",")
	way := req.FormValue("way")

	if table == "" {
		fmt.Fprintf(w, "Table name is missing.\n")
		return
	}
	if fmt.Sprintf("%v", columns) == "[]" || fmt.Sprintf("%v", values) == "[]" {
		fmt.Fprintf(w, "Columns or Values string it is not present.\n")
		return
	}
	if len(columns) != len(values) {
		fmt.Fprintf(w, "", columns, values)
		fmt.Fprintf(w, "Columns and Values it is not alignment.\n")
		return
	}

	record := &Record{table, columns, values, condition, way}
	//logger.Printf("Got record: %v\n", record)

	records,err := record.doQuery()
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
	} else {
		fmt.Fprintf(w, "%v\n", *records )
		//logger.Printf("Table: %v, Localid: %v is processed okay.\n", r.Table, r.Values[0])
	}
	
}

func (r *Record) doQeury ()  (*[]Record, error) {
	conditionLen := len(r.Condition)
	conditionIndexes := make([]int, conditionLen)
	for i, err := 0, errors.New(""); i < conditionLen; i++ {
		conditionIndexes[i], err = strconv.Atoi(r.Condition[i])
		if err != nil {
			return false, err
		}
		if conditionIndexes[i] >= len(r.Columns) {
			return false, errors.New("insertUpdate: ConditionIndexes out of range.")
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

	records, err := r.query(&columnsIndexes, &conditionIndexes)
	if err != nil {
		return false, err
	}	
	return records, nil
}