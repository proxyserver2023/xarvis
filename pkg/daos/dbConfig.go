package daos

// DBConfig holds all the vars that should be responsible for data source.
type DBConfig struct {
	DbName, DbUser, DbPass, DbHost, DbPort string
}
