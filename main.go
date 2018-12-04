package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	"gopkg.in/yaml.v2"
)

// configuration

type configuracio struct {
	P        string `yaml:"p"`
	Virginia struct {
		Db   string `yaml:"db"`
		User string `yaml:"user"`
	} `yaml:"virginia"`
	Irlanda struct {
		Db   string `yaml:"db"`
		User string `yaml:"user"`
	} `yaml:"irlanda"`
}

func ReadConfig() *configuracio {
	var config configuracio

	source, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}

func connection() {

	var (
		password string
		port     = 1433
	)
	c := ReadConfig()
	password = c.P

	//CONN
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()
}

func main() {
	connection()
	/*flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}
	//QUERIES

	showDB, err := conn.Query("SELECT name FROM master.sys.databases where database_id > '5'")
	if err != nil {
		log.Fatal("Query show failed:", err.Error())
	}
	defer showDB.Close()

	lockDB, err := conn.Query("SELECT sqltext.TEXT, DB_NAME(req.database_id) as database_name, USER_NAME(req.user_id) as user_name, req.session_id, req.status, req.command, req.blocking_session_id, req.wait_time, CONVERT(VARCHAR(50), req.start_time, 20) as start_time, req.cpu_time, req.total_elapsed_time FROM sys.dm_exec_requests req CROSS APPLY sys.dm_exec_sql_text(sql_handle) AS sqltext")
	if err != nil {
		log.Fatal("Query lock failed:", err.Error())
	}
	defer lockDB.Close()

	//RESULTS
	for showDB.Next() {
		var name string
		err := showDB.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}

	for lockDB.Next() {
		var text string
		var datab string
		var username string
		var sessionID int
		var requestStatus string
		var commanda string
		var blocksession string
		var waitTime string
		var startTime string
		var cpu int
		var elapsedTime string
		err := lockDB.Scan(&text, &datab, &username, &sessionID, &requestStatus, &commanda, &blocksession, &waitTime, &startTime, &cpu, &elapsedTime)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(datab, username, sessionID, requestStatus, commanda, waitTime)
	}
	err = lockDB.Err()
	if err != nil {
		log.Fatal(err)
	}

	//show, err := conn.Query(SELECT resource_database_id, request_type, request_status, request_lifetime, request_session_id, request_owner_id FROM sys.dm_tran_locks WHERE resource_database_id = DB_ID() AND resource_associated_entity_id = OBJECT_ID(N'dbo.$TABLE_NAME'))
	*/
}
