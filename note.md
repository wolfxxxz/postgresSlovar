// "postgres://localhost:5435/postgres?sslmode=disable&user=postgres&password=1"

type ConfigPostgress struct {
	Token             string `env:"TOKEN"`
	LogLevel          string `env:"LOGGER_LEVEL"`
	SqlHost           string `env:"SQLHost"`   //localhost
	SqlPort           string `env:"SQL_PORT"`  //5435
	SqlType           string `env:"SQL_TYPE"`  //postgres
	SqlMode           string `env:"SQL_MODE"`  //disable
	UserName          string `env:"USER_NAME"` // postgres
	Password          string `env:"PASSWORD"`  //1
	DBName            string `env:"DB_NAME"`
	TimeoutMongoQuery string `env:"TIMEOUT_MONGO_QUERY"`
}