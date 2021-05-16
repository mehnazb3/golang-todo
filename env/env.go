package env

import "os"

/////////////////////////////////////////
//         Postgres Configurations     //
/////////////////////////////////////////

// PostgresHost is the address of the PG server
// 127.0.0.1 - localhost by default
func PostgresHost() string {
	return get("POSTGRES_HOST", "127.0.0.1")
}

// PostgresPort is the address of the PG server
func PostgresPort() string {
	return get("POSTGRES_PORT", "5432")
}

// PostgresUser is the user of the PG DB
func PostgresUser() string {
	return get("POSTGRES_USER", "postgres")
}

// PostgresPassword is the password of the PG user
func PostgresPassword() string {
	return get("POSTGRES_PASSWORD", "postgres")
}

// PostgresDBName is the database name in PG
func PostgresDBName() string {
	return get("POSTGRES_DB_NAME", "golang-interview-challenge")
}

// PostgresSSL is whether SSL is enabled in Postgres.
// It should be false on local
func PostgresSSL() string {
	return get("POSTGRES_SSL", "false")
}

func get(name string, defaultVal string) string {
	val := os.Getenv(name)

	if len(val) > 0 {
		return val
	}
	return defaultVal
}
