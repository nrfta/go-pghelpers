package pghelpers

// PostgresConfig holds the information for connecting with a postgres database
type PostgresConfig struct {
	ApplicationName string `mapstructure:"application_name"`
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	SSLEnabled      bool   `mapstructure:"ssl_enabled"`
	MigrationsTable string `mapstructure:"migrations_table"`
	// MaxOpenConnections sets the maximum size of the connection pool. Default 10.
	MaxOpenConnections int `mapstructure:"max_open_connections"`
}
