package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// ServiceConfig defines the service configuration parameters
type ServiceConfig struct {
	InQueueName       string // SQS queue name
	MessageBucketName string // message bucket name (for oversize messages)
	PollTimeOut       int64  // queue wait time
	Workers           int    // number of workers
	WorkerQueueSize   int    // worker queue depth

	// database attributes
	DbHost string // database hostname
	DbPort int    // database port
	DbUser string // database user
	DbPass string // database password
	DbName string // database name

	// database insertion attributes
	DbInsertStatement string // the database insert statement
	DbInsertFields    string // the database field names
}

func ensureSet(env string) string {
	val, set := os.LookupEnv(env)

	if set == false {
		log.Printf("FATAL: environment variable not set: [%s]", env)
		os.Exit(1)
	}

	return val
}

func ensureSetAndNonEmpty(env string) string {
	val := ensureSet(env)

	if val == "" {
		log.Printf("FATAL: environment variable set but empty: [%s]", env)
		os.Exit(1)
	}

	return val
}

func envToInt(env string) int {
	number := ensureSetAndNonEmpty(env)

	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	return n
}

// LoadConfiguration will load the service configuration from env/cmdline
// and return a pointer to it. Any failures are fatal.
func LoadConfiguration() *ServiceConfig {

	var cfg ServiceConfig

	cfg.InQueueName = ensureSetAndNonEmpty("SQS_DB_INSERTER_IN_QUEUE")
	cfg.MessageBucketName = ensureSetAndNonEmpty("SQS_MESSAGE_BUCKET")
	cfg.PollTimeOut = int64(envToInt("SQS_DB_INSERTER_POLL_TIMEOUT"))
	cfg.Workers = envToInt("SQS_DB_INSERTER_WORKERS")
	cfg.WorkerQueueSize = envToInt("SQS_DB_INSERTER_WORKER_QUEUE_SIZE")

	// database attributes
	cfg.DbHost = ensureSetAndNonEmpty("SQS_DB_INSERTER_DB_HOST")
	cfg.DbPort = envToInt("SQS_DB_INSERTER_DB_PORT")
	cfg.DbUser = ensureSetAndNonEmpty("SQS_DB_INSERTER_DB_USER")
	cfg.DbPass = ensureSetAndNonEmpty("SQS_DB_INSERTER_DB_PASS")
	cfg.DbName = ensureSetAndNonEmpty("SQS_DB_INSERTER_DB_NAME")

	// database insertion attributes
	cfg.DbInsertStatement = ensureSetAndNonEmpty("SQS_DB_INSERTER_DB_INSERT_STATEMENT")
	cfg.DbInsertFields = ensureSetAndNonEmpty("SQS_DB_INSERTER_DB_INSERT_FIELDS")

	// handle a special case for AWS deployment
	cfg.DbInsertStatement = strings.Replace(cfg.DbInsertStatement, "$$", "$", -1)

	log.Printf("[config] InQueueName       = [%s]", cfg.InQueueName)
	log.Printf("[config] MessageBucketName = [%s]", cfg.MessageBucketName)
	log.Printf("[config] PollTimeOut       = [%d]", cfg.PollTimeOut)
	log.Printf("[config] Workers           = [%d]", cfg.Workers)
	log.Printf("[config] WorkerQueueSize   = [%d]", cfg.WorkerQueueSize)

	log.Printf("[config] DbHost            = [%s]", cfg.DbHost)
	log.Printf("[config] DbPort            = [%d]", cfg.DbPort)
	log.Printf("[config] DbUser            = [%s]", cfg.DbUser)
	log.Printf("[config] DbPass            = [REDACTED]")
	log.Printf("[config] DbName            = [%s]", cfg.DbName)

	// database insertion attributes
	log.Printf("[config] DbInsertStatement = [%s]", cfg.DbInsertStatement)
	log.Printf("[config] DbInsertFields    = [%s]", cfg.DbInsertFields)

	return &cfg
}

//
// end of file
//
