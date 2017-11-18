package config

//DB Config
const (

	// DBMS add database management service
	DBMS = "mysql"

	// DBNAME add database name
	DBNAME = ""

	// DBUSER add database user
	DBUSER = "root"

	// DBPASS add database pass
	DBPASS = ""

	//DBPROTOCOL add database protocol
	DBPROTOCOL = "tcp(192.168.33.10:3306)"

	//DBOPTIONS add database options
	DBOPTIONS = "charset=utf8&parseTime=true&loc=Local"

	//DBDSN add DB source name
	DBDSN = DBUSER + ":" + DBPASS + "@" + DBPROTOCOL + "/" + DBNAME + "?" + DBOPTIONS
)
