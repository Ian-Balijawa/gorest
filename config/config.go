package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/pilinux/gorest/lib/middleware"
	log "github.com/sirupsen/logrus"
)

// Activated - "yes" keyword to activate a service
const Activated string = "yes"

// Configuration - server and db configuration variables
type Configuration struct {
	Database   DatabaseConfig
	Logger     LoggerConfig
	Server     ServerConfig
	Security   SecurityConfig
	ViewConfig ViewConfig
}

// env - load the configurations from .env
func env() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Panic("panic code: 101")
	}
}

// Config - load all the configurations
func Config() Configuration {
	var configuration Configuration

	configuration.Database = Database()
	configuration.Logger = Logger()
	configuration.Security = Security()
	configuration.Server = Server()
	configuration.ViewConfig = View()

	return configuration
}

// Database - all DB variables
func Database() DatabaseConfig {
	var databaseConfig DatabaseConfig

	// Load environment variables
	env()

	// RDBMS
	activateRDBMS := os.Getenv("ACTIVATE_RDBMS")
	if activateRDBMS == "yes" {
		databaseConfig.RDBMS = DatabaseRDBMS().RDBMS
	}
	databaseConfig.RDBMS.Activate = activateRDBMS

	// REDIS
	activateRedis := os.Getenv("ACTIVATE_REDIS")
	if activateRedis == "yes" {
		databaseConfig.REDIS = DatabaseRedis().REDIS
	}
	databaseConfig.REDIS.Activate = activateRedis

	// MongoDB
	activateMongo := os.Getenv("ACTIVATE_MONGO")
	if activateMongo == "yes" {
		databaseConfig.MongoDB = DatabaseMongo().MongoDB
	}
	databaseConfig.MongoDB.Activate = activateMongo

	return databaseConfig
}

// DatabaseRDBMS - all RDBMS variables
func DatabaseRDBMS() DatabaseConfig {
	var databaseConfig DatabaseConfig
	var err error

	// Load environment variables
	env()

	// Env
	databaseConfig.RDBMS.Env.Driver = os.Getenv("DBDRIVER")
	databaseConfig.RDBMS.Env.Host = os.Getenv("DBHOST")
	databaseConfig.RDBMS.Env.Port = os.Getenv("DBPORT")
	databaseConfig.RDBMS.Env.TimeZone = os.Getenv("DBTIMEZONE")
	// Access
	databaseConfig.RDBMS.Access.DbName = os.Getenv("DBNAME")
	databaseConfig.RDBMS.Access.User = os.Getenv("DBUSER")
	databaseConfig.RDBMS.Access.Pass = os.Getenv("DBPASS")
	// SSL
	databaseConfig.RDBMS.Ssl.Sslmode = os.Getenv("DBSSLMODE")
	// Conn
	dbMaxIdleConns := os.Getenv("DBMAXIDLECONNS")
	dbMaxOpenConns := os.Getenv("DBMAXOPENCONNS")
	dbConnMaxLifetime := os.Getenv("DBCONNMAXLIFETIME")
	databaseConfig.RDBMS.Conn.MaxIdleConns, err = strconv.Atoi(dbMaxIdleConns)
	if err != nil {
		log.WithError(err).Panic("panic code: 131")
	}
	databaseConfig.RDBMS.Conn.MaxOpenConns, err = strconv.Atoi(dbMaxOpenConns)
	if err != nil {
		log.WithError(err).Panic("panic code: 132")
	}
	databaseConfig.RDBMS.Conn.ConnMaxLifetime, err = time.ParseDuration(dbConnMaxLifetime)
	if err != nil {
		log.WithError(err).Panic("panic code: 133")
	}

	// Logger
	dbLogLevel := os.Getenv("DBLOGLEVEL")
	databaseConfig.RDBMS.Log.LogLevel, err = strconv.Atoi(dbLogLevel)
	if err != nil {
		log.WithError(err).Panic("panic code: 134")
	}

	return databaseConfig
}

// DatabaseRedis - all REDIS DB variables
func DatabaseRedis() DatabaseConfig {
	var databaseConfig DatabaseConfig

	// Load environment variables
	env()

	// REDIS
	poolSize, err := strconv.Atoi(os.Getenv("POOLSIZE"))
	if err != nil {
		log.WithError(err).Panic("panic code: 135")
	}
	connTTL, err := strconv.Atoi(os.Getenv("CONNTTL"))
	if err != nil {
		log.WithError(err).Panic("panic code: 136")
	}

	databaseConfig.REDIS.Env.Host = os.Getenv("REDISHOST")
	databaseConfig.REDIS.Env.Port = os.Getenv("REDISPORT")
	databaseConfig.REDIS.Conn.PoolSize = poolSize
	databaseConfig.REDIS.Conn.ConnTTL = connTTL

	return databaseConfig
}

// DatabaseMongo - all MongoDB variables
func DatabaseMongo() DatabaseConfig {
	var databaseConfig DatabaseConfig

	// Load environment variables
	env()

	// MongoDB
	poolSize, err := strconv.ParseUint(os.Getenv("MONGO_POOLSIZE"), 10, 64)
	if err != nil {
		log.WithError(err).Panic("panic code: 137")
	}
	connTTL, err := strconv.Atoi(os.Getenv("MONGO_CONNTTL"))
	if err != nil {
		log.WithError(err).Panic("panic code: 138")
	}

	databaseConfig.MongoDB.Env.URI = os.Getenv("MONGO_URI")
	databaseConfig.MongoDB.Env.AppName = os.Getenv("MONGO_APP")
	databaseConfig.MongoDB.Env.PoolSize = poolSize
	databaseConfig.MongoDB.Env.PoolMon = os.Getenv("MONGO_MONITOR_POOL")
	databaseConfig.MongoDB.Env.ConnTTL = connTTL

	return databaseConfig
}

// Logger ...
func Logger() LoggerConfig {
	var loggerConfig LoggerConfig

	// Load environment variables
	env()

	loggerSentryDsn := os.Getenv("SentryDSN")
	loggerConfig.SentryDsn = loggerSentryDsn

	return loggerConfig
}

// Security - configs for generating tokens and hashes
func Security() SecurityConfig {
	var securityConfig SecurityConfig

	// Load environment variables
	env()

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	// JWT
	SecurityConfigAll.MustJWT = os.Getenv("ACTIVATE_JWT")
	if SecurityConfigAll.MustJWT == Activated {
		securityConfig.JWT = readEnvJWT()

		// set params globally
		setParamsJWT(securityConfig.JWT)
	}

	hashPassMemory64, err := strconv.ParseUint((os.Getenv("HASHPASSMEMORY")), 10, 32)
	if err != nil {
		log.WithError(err).Panic("panic code: 121")
	}
	hashPassIterations64, err := strconv.ParseUint((os.Getenv("HASHPASSITERATIONS")), 10, 32)
	if err != nil {
		log.WithError(err).Panic("panic code: 122")
	}
	hashPassParallelism64, err := strconv.ParseUint((os.Getenv("HASHPASSPARALLELISM")), 10, 8)
	if err != nil {
		log.WithError(err).Panic("panic code: 123")
	}
	hashPassSaltLength64, err := strconv.ParseUint((os.Getenv("HASHPASSSALTLENGTH")), 10, 32)
	if err != nil {
		log.WithError(err).Panic("panic code: 124")
	}
	hashPassKeyLength64, err := strconv.ParseUint((os.Getenv("HASHPASSKEYLENGTH")), 10, 32)
	if err != nil {
		log.WithError(err).Panic("panic code: 125")
	}
	hashPassMemory := uint32(hashPassMemory64)
	hashPassIterations := uint32(hashPassIterations64)
	hashPassParallelism := uint8(hashPassParallelism64)
	hashPassSaltLength := uint32(hashPassSaltLength64)
	hashPassKeyLength := uint32(hashPassKeyLength64)

	listType := os.Getenv("LISTTYPE")
	ip := os.Getenv("IP")

	securityConfig.BasicAuth.Username = username
	securityConfig.BasicAuth.Password = password

	securityConfig.HashPass.Memory = hashPassMemory
	securityConfig.HashPass.Iterations = hashPassIterations
	securityConfig.HashPass.Parallelism = hashPassParallelism
	securityConfig.HashPass.SaltLength = hashPassSaltLength
	securityConfig.HashPass.KeyLength = hashPassKeyLength

	securityConfig.Firewall.ListType = listType
	securityConfig.Firewall.IP = ip

	securityConfig.CORS.Origin = os.Getenv("CORS_ORIGIN")
	securityConfig.CORS.Credentials = os.Getenv("CORS_CREDENTIALS")
	securityConfig.CORS.Headers = os.Getenv("CORS_HEADERS")
	securityConfig.CORS.Methods = os.Getenv("CORS_METHODS")
	securityConfig.CORS.MaxAge = os.Getenv("CORS_MAXAGE")

	securityConfig.TrustedPlatform = os.Getenv("TRUSTED_PLATFORM")

	return securityConfig
}

// Server - port and env
func Server() ServerConfig {
	var serverConfig ServerConfig

	// Load environment variables
	env()

	serverConfig.ServerPort = os.Getenv("APP_PORT")
	serverConfig.ServerEnv = os.Getenv("APP_ENV")

	return serverConfig
}

// View - HTML renderer
func View() ViewConfig {
	var viewConfig ViewConfig

	// Load environment variables
	env()

	viewConfig.Dir = os.Getenv("TEMPLATE_DIR")

	return viewConfig
}

// readEnvJWT - read parameters from env
func readEnvJWT() middleware.JWTParameters {
	env()

	params := middleware.JWTParameters{}
	var err error

	params.AccessKey = []byte(os.Getenv("ACCESS_KEY"))
	params.AccessKeyTTL, err = strconv.Atoi(os.Getenv("ACCESS_KEY_TTL"))
	if err != nil {
		log.WithError(err).Panic("panic code: 111")
	}
	params.RefreshKey = []byte(os.Getenv("REFRESH_KEY"))
	params.RefreshKeyTTL, err = strconv.Atoi(os.Getenv("REFRESH_KEY_TTL"))
	if err != nil {
		log.WithError(err).Panic("panic code: 112")
	}
	params.Audience = os.Getenv("AUDIENCE")
	params.Issuer = os.Getenv("ISSUER")
	params.AccNbf, err = strconv.Atoi(os.Getenv("NOT_BEFORE_ACC"))
	if err != nil {
		log.WithError(err).Panic("panic code: 113")
	}
	params.RefNbf, err = strconv.Atoi(os.Getenv("NOT_BEFORE_REF"))
	if err != nil {
		log.WithError(err).Panic("panic code: 114")
	}
	params.Subject = os.Getenv("SUBJECT")

	return params
}

// setParamsJWT - set parameters for JWT
func setParamsJWT(c middleware.JWTParameters) {
	middleware.JWTParams.AccessKey = c.AccessKey
	middleware.JWTParams.AccessKeyTTL = c.AccessKeyTTL
	middleware.JWTParams.RefreshKey = c.RefreshKey
	middleware.JWTParams.RefreshKeyTTL = c.RefreshKeyTTL
	middleware.JWTParams.Audience = c.Audience
	middleware.JWTParams.Issuer = c.Issuer
	middleware.JWTParams.AccNbf = c.AccNbf
	middleware.JWTParams.RefNbf = c.RefNbf
	middleware.JWTParams.Subject = c.Subject
}
