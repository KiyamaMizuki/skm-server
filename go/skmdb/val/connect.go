package val

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	// MySQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Config DBに関する情報
type Config struct {
	Dbms   string `required:"true"`
	Dbuser string `required:"true"`
	Pass   string `required:"true"`
	Dbname string `required:"true"`
}
type MysqlConfig struct {
	Dbms     string `required:"true"`
	Dbuser   string `required:"true"`
	Pass     string `required:"true"`
	Protocol string `required:"true"`
	Dbname   string `required:"true"`
	Time     string `required:"true"`
}

//GormConnect データベースの情報を記載/接続
func MysqlConnect() *gorm.DB {
	doterr := godotenv.Load()
	if doterr != nil {
		log.Fatal("Error loading .env file")
	}

	var config MysqlConfig
	config_err := envconfig.Process("mysql", &config)
	if config_err != nil {
		log.Fatal(config_err.Error())
	}
	CONNECT := config.Dbuser + ":" + config.Pass + "@" + config.Protocol + "/" + config.Dbname + "?" + config.Time
	//データベースに接続
	db, err := gorm.Open(config.Dbms, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func PostgresConnect() *gorm.DB {
	doterr := godotenv.Load()
	if doterr != nil {
		log.Fatal("Error loading .env file")
	}

	var config Config
	config_err := envconfig.Process("", &config)
	if config_err != nil {
		log.Fatal(config_err.Error())
	}

	CONNECT := "host=localhost port=5432 user=" + config.Dbuser + " dbname=" + config.Dbname + " password=" + config.Pass + " sslmode=disable"
	//データベースに接続
	db, err := gorm.Open(config.Dbms, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func GoConnect() *sql.DB {
	doterr := godotenv.Load()
	if doterr != nil {
		log.Fatal("Error loading .env file")
	}

	var config Config
	config_err := envconfig.Process("", &config)
	if config_err != nil {
		log.Fatal(config_err.Error())
	}
	CONNECT := "host=localhost port=5432 user=" + config.Dbuser + " dbname=" + config.Dbname + " password=" + config.Pass + " sslmode=disable"
	db, err := sql.Open(config.Dbms, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}
