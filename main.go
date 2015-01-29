package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Record struct {
	Table     string
	Columns   []string
	Values    []string
	Condition []string
	Way       string
}

var (
	logger *log.Logger
	port   string
	db     *sql.DB
)

func main() {
	version := flag.Bool("v", false, "Show version.")
	author := flag.Bool("author", false, "Show author.")

	flag.StringVar(&port, "port", "3000", "Port number.")
	//flag.StringVar(&path, "path", ".", "File server path.")
	//flag.StringVar(&filter, "filter", "", "A string, If matched, will service matched file only.")

	sqlDrive := flag.String("sql-drive", "mysql", "The database drive name.")
	dbuser := flag.String("dbuser", "sysCheckV2", "Database username.")
	dbpass := flag.String("dbpass", "sysCheckV2123", "Database password.")
	dbip := flag.String("dbip", "10.100.2.108", "Database ip address.")
	dbport := flag.String("dbport", "3307", "Database port number.")
	dbname := flag.String("dbname", "M", "Which database to be use.")

	flag.Parse()

	//Display version info.
	if *version {
		fmt.Println("TransServer version=1.0.3, Date:2015-1-29")
		os.Exit(0)
	}

	//Display author info.
	if *author {
		fmt.Println("Author is: Wen Zhenglin")
		os.Exit(0)
	}

	//Removed os.O_APPEND for log file size concern.
	logfile, err := os.OpenFile("TransServer.log", os.O_RDWR|os.O_CREATE, 0666)
	//logfile, err := os.OpenFile(os.Stdout, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer logfile.Close()

	logger = log.New(logfile, "", log.LstdFlags)
	//logger = log.New(os.Stdout, "", log.LstdFlags)

	//db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/test?charset=utf8")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", *dbuser, *dbpass, *dbip, *dbport, *dbname)
	db, err = sql.Open(*sqlDrive, dsn)
	if err != nil {
		logger.Panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Panic(err)
	}
	logger.Print("db ping ok.")

	serv()

	//fmt.Println("test")
}
