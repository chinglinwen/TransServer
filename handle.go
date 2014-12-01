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

	if table == "" {
		fmt.Fprintf(w, "Table name is missing.\n")
		return
	}
	if fmt.Sprintf("%v", columns) == "[]" || fmt.Sprintf("%v", values) == "[]" {
		fmt.Fprintf(w, "Columns or Values string is not present.\n")
		return
	}
	if len(columns) != len(values) {
		fmt.Fprintf(w, "", columns, values)
		fmt.Fprintf(w, "Columns and Values is not alignment.\n")
		return
	}

	record := &Record{table, columns, values, condition}
	//logger.Printf("Got record: %v\n", record)

	switch table {
	case "ip_core", "ip_extra" :

		if fmt.Sprintf("%v", condition) == "[]" {
			fmt.Fprintf(w, "Condition is mising.\n")
			return
		}

		if ok, err := processIpCoreExtra(record); err != nil {
			fmt.Fprintf(w, "processIpCoreExtra error: %v\n", err)
		} else if ok {
			fmt.Fprintf(w, "okay\n")
			logger.Printf("Table: %v, Localid: %v is processed okay.\n", table, record.Values[0])
		} else {
			fmt.Fprintf(w, "There no error, but is not ok, Should never get here.\n")
			logger.Printf("There no error, but is not ok, Should never get here.\n")
		}
	case "ip_conn", "syscheck_result":

		if ok, err := processSyscheckResult(record); err != nil {
			fmt.Fprintf(w, "processSyscheckResult error: %v\n", err)
		} else if ok {
			fmt.Fprintf(w, "okay\n")
			logger.Printf("Table: %v, Localid: %v,round: %v is processed okay.\n",
				table, record.Values[0], record.Values[1])
		} else {
			fmt.Fprintf(w, "There no error, but is not ok, Should never get here.\n")
			logger.Printf("There no error, but is not ok, Should never get here.\n")
		}
	case "FST_TAB_cnt" :
	
		if fmt.Sprintf("%v", condition) == "[]" {
			fmt.Fprintf(w, "Condition is mising.\n")
			return
		}

		if ok, err := processOnlyOneEntry(record); err != nil {
			fmt.Fprintf(w, "processOnlyOneEntry error: %v\n", err)
		} else if ok {
			fmt.Fprintf(w, "okay\n")
			logger.Printf("Table: %v, Localid: %v is processed okay.\n", table, record.Values[0])
		} else {
			fmt.Fprintf(w, "There no error, but is not ok, Should never get here.\n")
			logger.Printf("There no error, but is not ok, Should never get here.\n")
		}
	
	default:
		//processDefault(record)
		fmt.Fprintf(w, "Unknow table name.\n")
		return
	}
}
