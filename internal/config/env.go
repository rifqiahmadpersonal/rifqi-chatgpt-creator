package config

import (
	"time"

	"github.com/spf13/viper"
)

var _ = viper.GetString

// EnvConfig holds all application configuration loaded from environment
type EnvConfig struct {
	App           AppConfig
	DB            DatabaseConfig
	Redis         RedisConfig
	Frontend      FrontendConfig
	Registration  RegistrationConfig
	EmailDomains  EmailDomainsConfig
	Logging       LoggingConfig
	API           APIConfig
	WebSocket     WebSocketConfig
}

type AppConfig struct {
	Env  string `mapstructure:"APP_ENV"`
	Name string `mapstructure:"APP_NAME"`
	Port int    `mapstructure:"APP_PORT"`
	Mode string `mapstructure:"APP_MODE"`
}

type DatabaseConfig struct {
	Host            string        `mapstructure:"DB_HOST"`
	Port            int           `mapstructure:"DB_PORT"`
	User            string        `mapstructure:"DB_USER"`
	Password        string        `mapstructure:"DB_PASSWORD"`
	Name            string        `mapstructure:"DB_NAME"`
	SSLMode         string        `mapstructure:"DB_SSL_MODE"`
	MaxOpenConns    int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `mapstructure:"DB_CONN_MAX_LIFETIME"`
}

type RedisConfig struct {
	Enabled  bool   `mapstructure:"REDIS_ENABLED"`
	Host     string `mapstructure:"REDIS_HOST"`
	Port     int    `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	DB       int    `mapstructure:"REDIS_DB"`
}

type FrontendConfig struct {
	URL         string `mapstructure:"FRONTEND_URL"`
	CORSOrigins string `mapstructure:"CORS_ORIGINS"`
}

type RegistrationConfig struct {
	DefaultProxy    string        `mapstructure:"DEFAULT_PROXY"`
	DefaultPassword string        `mapstructure:"DEFAULT_PASSWORD"`
	DefaultDomain   string        `mapstructure:"DEFAULT_DOMAIN"`
	WorkerPoolSize  int           `mapstructure:"WORKER_POOL_SIZE"`
	Timeout         time.Duration `mapstructure:"REGISTRATION_TIMEOUT"`
	MaxRetries      int           `mapstructure:"MAX_RETRIES"`
}

type EmailDomainsConfig struct {
	HealthCheckEnabled  bool          `mapstructure:"EMAIL_DOMAINS_HEALTH_CHECK_ENABLED"`
	HealthCheckInterval time.Duration `mapstructure:"EMAIL_DOMAINS_HEALTH_CHECK_INTERVAL"`
	DefaultSource       string        `mapstructure:"EMAIL_DOMAINS_DEFAULT_SOURCE"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"LOG_LEVEL"`
	Format string `mapstructure:"LOG_FORMAT"`
	Output string `mapstructure:"LOG_OUTPUT"`
}

type APIConfig struct {
	RateLimitEnabled  bool          `mapstructure:"API_RATE_LIMIT_ENABLED"`
	RateLimitRequests int           `mapstructure:"API_RATE_LIMIT_REQUESTS"`
	RateLimitDuration time.Duration `mapstructure:"API_RATE_LIMIT_DURATION"`
	ReadTimeout       time.Duration `mapstructure:"API_READ_TIMEOUT"`
	WriteTimeout      time.Duration `mapstructure:"API_WRITE_TIMEOUT"`
	ShutdownTimeout   time.Duration `mapstructure:"API_SHUTDOWN_TIMEOUT"`
}

type WebSocketConfig struct {
	Enabled        bool          `mapstructure:"WS_ENABLED"`
	PingInterval   time.Duration `mapstructure:"WS_PING_INTERVAL"`
	PongWait       time.Duration `mapstructure:"WS_PONG_WAIT"`
	MaxMessageSize int64         `mapstructure:"WS_MAX_MESSAGE_SIZE"`
}

// LoadEnv loads configuration from environment variables
func LoadEnv() (*EnvConfig, error) {
	viper.SetEnvPrefix("")
	viper.AutomaticEnv()

	// Bind env vars explicitly AFTER AutomaticEnv
	viper.BindEnv("APP_ENV")
	viper.BindEnv("APP_NAME")
	viper.BindEnv("APP_PORT")
	viper.BindEnv("APP_MODE")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASSWORD")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("DB_SSL_MODE")
	viper.BindEnv("DB_MAX_OPEN_CONNS")
	viper.BindEnv("DB_MAX_IDLE_CONNS")
	viper.BindEnv("DB_CONN_MAX_LIFETIME")
	viper.BindEnv("REDIS_ENABLED")
	viper.BindEnv("REDIS_HOST")
	viper.BindEnv("REDIS_PORT")
	viper.BindEnv("REDIS_PASSWORD")
	viper.BindEnv("REDIS_DB")
	viper.BindEnv("FRONTEND_URL")
	viper.BindEnv("CORS_ORIGINS")
	viper.BindEnv("DEFAULT_PROXY")
	viper.BindEnv("DEFAULT_PASSWORD")
	viper.BindEnv("DEFAULT_DOMAIN")
	viper.BindEnv("WORKER_POOL_SIZE")
	viper.BindEnv("REGISTRATION_TIMEOUT")
	viper.BindEnv("MAX_RETRIES")
	viper.BindEnv("EMAIL_DOMAINS_HEALTH_CHECK_ENABLED")
	viper.BindEnv("EMAIL_DOMAINS_HEALTH_CHECK_INTERVAL")
	viper.BindEnv("EMAIL_DOMAINS_DEFAULT_SOURCE")
	viper.BindEnv("LOG_LEVEL")
	viper.BindEnv("LOG_FORMAT")
	viper.BindEnv("LOG_OUTPUT")
	viper.BindEnv("API_RATE_LIMIT_ENABLED")
	viper.BindEnv("API_RATE_LIMIT_REQUESTS")
	viper.BindEnv("API_RATE_LIMIT_DURATION")
	viper.BindEnv("API_READ_TIMEOUT")
	viper.BindEnv("API_WRITE_TIMEOUT")
	viper.BindEnv("API_SHUTDOWN_TIMEOUT")
	viper.BindEnv("WS_ENABLED")
	viper.BindEnv("WS_PING_INTERVAL")
	viper.BindEnv("WS_PONG_WAIT")
	viper.BindEnv("WS_MAX_MESSAGE_SIZE")

// Set defaults AFTER binding so env vars take precedence
	setDefaults()

	// Manual override for critical values - viper string->int conversion is unreliable
	cfg := EnvConfig{
		App: AppConfig{
			Env:  viper.GetString("APP_ENV"),
			Name: viper.GetString("APP_NAME"),
			Port: viper.GetInt("APP_PORT"),
			Mode: viper.GetString("APP_MODE"),
		},
		DB: DatabaseConfig{
			Host:            viper.GetString("DB_HOST"),
			Port:            viper.GetInt("DB_PORT"),
			User:            viper.GetString("DB_USER"),
			Password:        viper.GetString("DB_PASSWORD"),
			Name:            viper.GetString("DB_NAME"),
			SSLMode:         viper.GetString("DB_SSL_MODE"),
			MaxOpenConns:    viper.GetInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns:    viper.GetInt("DB_MAX_IDLE_CONNS"),
			ConnMaxLifetime: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
		},
		Redis: RedisConfig{
			Enabled:  viper.GetBool("REDIS_ENABLED"),
			Host:     viper.GetString("REDIS_HOST"),
			Port:     viper.GetInt("REDIS_PORT"),
			Password: viper.GetString("REDIS_PASSWORD"),
			DB:       viper.GetInt("REDIS_DB"),
		},
		Frontend: FrontendConfig{
			URL:         viper.GetString("FRONTEND_URL"),
			CORSOrigins: viper.GetString("CORS_ORIGINS"),
		},
		Registration: RegistrationConfig{
			DefaultProxy:    viper.GetString("DEFAULT_PROXY"),
			DefaultPassword: viper.GetString("DEFAULT_PASSWORD"),
			DefaultDomain:   viper.GetString("DEFAULT_DOMAIN"),
			WorkerPoolSize:  viper.GetInt("WORKER_POOL_SIZE"),
			Timeout:         viper.GetDuration("REGISTRATION_TIMEOUT"),
			MaxRetries:      viper.GetInt("MAX_RETRIES"),
		},
		EmailDomains: EmailDomainsConfig{
			HealthCheckEnabled:  viper.GetBool("EMAIL_DOMAINS_HEALTH_CHECK_ENABLED"),
			HealthCheckInterval: viper.GetDuration("EMAIL_DOMAINS_HEALTH_CHECK_INTERVAL"),
			DefaultSource:       viper.GetString("EMAIL_DOMAINS_DEFAULT_SOURCE"),
		},
		Logging: LoggingConfig{
			Level:  viper.GetString("LOG_LEVEL"),
			Format: viper.GetString("LOG_FORMAT"),
			Output: viper.GetString("LOG_OUTPUT"),
		},
		API: APIConfig{
			RateLimitEnabled:  viper.GetBool("API_RATE_LIMIT_ENABLED"),
			RateLimitRequests: viper.GetInt("API_RATE_LIMIT_REQUESTS"),
			RateLimitDuration: viper.GetDuration("API_RATE_LIMIT_DURATION"),
			ReadTimeout:       viper.GetDuration("API_READ_TIMEOUT"),
			WriteTimeout:      viper.GetDuration("API_WRITE_TIMEOUT"),
			ShutdownTimeout:   viper.GetDuration("API_SHUTDOWN_TIMEOUT"),
		},
		WebSocket: WebSocketConfig{
			Enabled:        viper.GetBool("WS_ENABLED"),
			PingInterval:   viper.GetDuration("WS_PING_INTERVAL"),
			PongWait:       viper.GetDuration("WS_PONG_WAIT"),
			MaxMessageSize: viper.GetInt64("WS_MAX_MESSAGE_SIZE"),
		},
	}

	return &cfg, nil
}

func setDefaults() {
	// App defaults
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_NAME", "chatgpt-creator")
	viper.SetDefault("APP_PORT", 8080)
	viper.SetDefault("APP_MODE", "release")

	// Database defaults
	viper.SetDefault("DB_HOST", "postgres")
	viper.SetDefault("DB_PORT", 5432)
	viper.SetDefault("DB_USER", "chatgpt")
	viper.SetDefault("DB_PASSWORD", "chatgpt_secret")
	viper.SetDefault("DB_NAME", "chatgpt_creator")
	viper.SetDefault("DB_SSL_MODE", "disable")
	viper.SetDefault("DB_MAX_OPEN_CONNS", 25)
	viper.SetDefault("DB_MAX_IDLE_CONNS", 25)
	viper.SetDefault("DB_CONN_MAX_LIFETIME", "5m")

	// Redis defaults
	viper.SetDefault("REDIS_ENABLED", false)
	viper.SetDefault("REDIS_HOST", "localhost")
	viper.SetDefault("REDIS_PORT", 6379)

	// Frontend defaults
	viper.SetDefault("FRONTEND_URL", "http://localhost:3000")
	viper.SetDefault("CORS_ORIGINS", "http://localhost:3000")

	// Registration defaults
	viper.SetDefault("WORKER_POOL_SIZE", 5)
	viper.SetDefault("REGISTRATION_TIMEOUT", "10m")
	viper.SetDefault("MAX_RETRIES", 3)

	// Email domains defaults
	viper.SetDefault("EMAIL_DOMAINS_HEALTH_CHECK_ENABLED", true)
	viper.SetDefault("EMAIL_DOMAINS_HEALTH_CHECK_INTERVAL", "5m")
	viper.SetDefault("EMAIL_DOMAINS_DEFAULT_SOURCE", "generator")

	// Logging defaults
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("LOG_FORMAT", "json")

	// API defaults
	viper.SetDefault("API_RATE_LIMIT_ENABLED", true)
	viper.SetDefault("API_RATE_LIMIT_REQUESTS", 100)
	viper.SetDefault("API_RATE_LIMIT_DURATION", "1m")
	viper.SetDefault("API_READ_TIMEOUT", "30s")
	viper.SetDefault("API_WRITE_TIMEOUT", "30s")
	viper.SetDefault("API_SHUTDOWN_TIMEOUT", "10s")

	// WebSocket defaults
	viper.SetDefault("WS_ENABLED", true)
	viper.SetDefault("WS_PING_INTERVAL", "30s")
	viper.SetDefault("WS_PONG_WAIT", "60s")
	viper.SetDefault("WS_MAX_MESSAGE_SIZE", 4096)
}
