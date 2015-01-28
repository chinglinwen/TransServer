package main

import (
	"fmt"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, req *http.Request) {
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
		doCompareInsertUpdate(w,record)
	case "ip_conn", "syscheck_result":
		doDirectInsert(w,record)
	case "FST_TAB_cnt":
		doInsertUpdate(w,record)
	default:
		otherWay(w,record)
	}
}

func otherWay (w http.ResponseWriter,r *Record) {
	switch r.way {
	case "direct_insert" :
		doDirectInsert(w,r)
	case "insert_update" :
		doInsertUpdate(w,r)
	case "compare_insert_update" :
		doCompareInsertUpdate(w,r)
	cases default:
		doDirectInsert(w,r)
	}
}

func doCompareInsertUpdate ( w http.ResponseWriter,r *Record) {
		if fmt.Sprintf("%v", r.condition) == "[]" {
			fmt.Fprintf(w, "Condition is mising.\n")
			return
		}

		if ok, err := compareInsertUpdate(r); err != nil {
			fmt.Fprintf(w, "compareInsertUpdate error: %v\n", err)
		} else if ok {
			fmt.Fprintf(w, "okay\n")
			logger.Printf("Table: %v, Localid: %v is processed okay.\n", table, r.Values[0])
		} else {
			fmt.Fprintf(w, "There is no error, but it is not ok, Should never get here.\n")
			logger.Printf("There is no error, but it is not ok, Should never get here.\n")
		}
}

func doDirectInsert ( w http.ResponseWriter,r *Record) {
		if ok, err := directInsert(r); err != nil {
			fmt.Fprintf(w, "directInsert error: %v\n", err)
		} else if ok {
			fmt.Fprintf(w, "okay\n")
			logger.Printf("Table: %v, Localid: %v,round: %v is processed okay.\n",
				table, r.Values[0], r.Values[1])
		} else {
			fmt.Fprintf(w, "There is no error, but it is not ok, Should never get here.\n")
			logger.Printf("There is no error, but it is not ok, Should never get here.\n")
		}
}

func doInsertUpdate ( w http.ResponseWriter,r *Record) {
		if fmt.Sprintf("%v", condition) == "[]" {
			fmt.Fprintf(w, "Condition is mising.\n")
			return
		}

		if ok, err := insertUpdate(r); err != nil {
			fmt.Fprintf(w, "%v\n", err)
		} else if ok {
			fmt.Fprintf(w, "okay\n")
			logger.Printf("Table: %v, Localid: %v is processed okay.\n", table, r.Values[0])
		} else {
			fmt.Fprintf(w, "There is no error, but it is not ok, Should never get here.\n")
			logger.Printf("There is no error, but it is not ok, Should never get here.\n")
		}
}