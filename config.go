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
	// MaxIdleConnections set the maximum idle connections that are held open. Default 2.
	MaxIdleConnections int `mapstructure:"max_idle_connections"`
	// MaxConnectionLifetimeMinutes ensures that connections are recycled after a certain amount of time,
	// preventing stale connections from accumulating. Default 5 minutes.
	MaxConnectionLifetimeMinutes int `mapstructure:"max_connection_lifetime_minutes"`
}
