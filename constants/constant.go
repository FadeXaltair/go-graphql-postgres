package constants

import (
	"os"
)

// export local=true
//  export host="localhost"
//  export dbname="postgres"
//  export user="postgres"
//  export port=5432
//  export password="postgres"

var (
	Host     = os.Getenv("host")
	Port     = os.Getenv("port")
	Dbname   = os.Getenv("dbname")
	User     = os.Getenv("user")
	Password = os.Getenv("password")
)
