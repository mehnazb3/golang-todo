package db

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
)

// PgConfig holds the configuration of Postgres
type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Ssl      string
}

// Info returns the PG configuration in string
func (pc *PgConfig) Info() string {
	res := ""

	if len(pc.Host) > 0 {
		res += fmt.Sprintf(" host=%s", pc.Host)
	}

	if len(pc.Port) > 0 {
		res += fmt.Sprintf(" port=%s", pc.Port)
	}

	if len(pc.User) > 0 {
		res += fmt.Sprintf(" user=%s", pc.User)
	}

	if len(pc.Password) > 0 {
		res += fmt.Sprintf(" password=%s", pc.Password)
	}

	if len(pc.DBName) > 0 {
		res += fmt.Sprintf(" dbname=%s", pc.DBName)
	}

	if pc.Ssl == "false" {
		res += " sslmode=disable"
	}

	return res[1:]
}

// PgClient helps to connect to the Postgres
type PgClient struct {
	Config *PgConfig

	db *gorm.DB
}

// Db returns the direct connection to the DB
func (cli *PgClient) Db() (*sql.DB, error) {
	return cli.db.DB()
}

// Orm returns an interface to use GORM
func (cli *PgClient) Orm() *gorm.DB {
	return cli.db
}

// Connect to the Postgres
func (cli *PgClient) Connect() error {
	db, err := gorm.Open(postgres.Open(cli.Config.Info()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	cli.db = db
	log.Info("Connected to Postgres successfully")

	return nil
}
