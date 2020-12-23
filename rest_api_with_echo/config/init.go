package config

import (
	"context"
	"log"

	"os"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

type (
	// Connection struct is will contain two driver database 1 for live 1 for log
	Connection struct {
		DBLive driver.Database
		DBLog  driver.Database
	}
)

var (
	// Instance for addressed connection will be call after create a connection
	Instance *Connection
	// DBURL database url
	DBURL string = os.Getenv("DBURL")
	// DBURL database log url
	DBLOGURL string = os.Getenv("DBLOGURL")
	// DBUSER database username
	DBUSERNAME string = os.Getenv("DBUSERNAME")
	// DBPASS database password
	DBPASSWORD string = os.Getenv("DBPASSWORD")
	// DBNAME database name
	DBNAME string = os.Getenv("DBNAME")
	// DBLogNAME database log
	DBLOGNAME string = os.Getenv("DBLOGNAME")
	// STAGEENV stage environment
	STAGEENV = os.Getenv("STAGEENV")
)

// User context will be use on route
type UserContext string

func init() {
	// Create new connection with endpoint at env DBURL
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBURL},
	})

	if err != nil {
		panic(err)
	}
	// Create a new client with basic authentication, username and password at env DBUSERNAME and DBPASSWORD
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(DBUSERNAME, DBPASSWORD),
	})

	if err != nil {
		panic(err)
	}

	// context for your incoming request
	ctx := context.Background()
	db, err := client.Database(ctx, DBNAME)
	if err != nil {
		log.Printf("Error connecting to database, cause: %+v \n", err)
		log.Printf("URL: %s, Name: %s, User: %s, Pass: %s\n", DBURL, DBNAME, DBUSERNAME, DBPASSWORD)
		panic(err)
	}

	//DB Log
	connLog, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{DBLOGURL},
	})

	if err != nil {
		panic(err)
	}

	clientLog, err := driver.NewClient(driver.ClientConfig{
		Connection:     connLog,
		Authentication: driver.BasicAuthentication(DBUSERNAME, DBPASSWORD),
	})

	if err != nil {
		panic(err)
	}

	dbLog, err := clientLog.Database(ctx, DBLOGNAME)
	if err != nil {
		log.Printf("Error connecting to database, cause: %+v \n", err)
		log.Printf("URL: %s, Name: %s, User: %s, Pass: %s\n", DBURL, DBNAME, DBUSERNAME, DBPASSWORD)
		panic(err)
	}
	// so driver database at struct connection will be contain connection we create
	Instance = &Connection{
		DBLive: db,
		DBLog:  dbLog,
	}
}

// will be implement address connection to get connection
func GetInstance() *Connection {
	return Instance
}
