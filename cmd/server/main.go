package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/app/server"
	"github.com/ilfey/devback/internal/pkg/psql"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	logLevel      string
	address       string
	databaseUrl   string
	schemePath    string
	key           string
	adminPath     string
	adminUsername string
	lifeSpan      int
)

func main() {
	godotenv.Load()

	flag.StringVar(&logLevel, "ll", getEnv("LOGLEVEL", "info"), "LogLevel")
	flag.StringVar(&databaseUrl, "du", getEnv("DATABASE_URL", "postgresql://ilfey:QWEasd123@localhost:5432/devpage"), "PostgreSQL database url")
	flag.StringVar(&schemePath, "df", getEnv("DATABASE_SCHEME", "./scheme.sql"), "Scheme database file")
	flag.StringVar(&address, "a", getEnv("ADDRESS", "0.0.0.0:8080"), "Address")
	flag.StringVar(&adminPath, "ap", getEnv("ADMIN_PATH", "/admin"), "Admin path")
	flag.StringVar(&adminUsername, "au", getEnv("ADMIN_USERNAME", "admin"), "Admin username")
	flag.StringVar(&key, "jk", getEnv("JWT_KEY", "secret"), "JWT key")
	flag.IntVar(&lifeSpan, "jls", getEnvInt("JWT_LIFE_SPAN", 24), "JWT life span (in hours)")

	flag.Parse()

	conf := new(config.Config)

	conf.Addr = address
	conf.LifeSpan = lifeSpan
	conf.AdminPath = adminPath
	conf.AdminUsername = adminUsername
	conf.Key = key

	log, err := createLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}

	store, err := initPostgres(log, databaseUrl, schemePath)
	if err != nil {
		log.Fatal(err)
	}

	serv := server.NewServer(conf, log, store)

	err = serv.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func readScheme(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func initPostgres(logger *logrus.Logger, databaseUrl, schemePath string) (*store.Store, error) {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		return nil, err
	}

	scheme, err := readScheme(schemePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(context.Background(), scheme)
	if err != nil {
		return nil, err
	}

	store := psql.NewStore(db, logger)

	return store, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if s, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		return value
	}
	return fallback
}

func createLogger(level string) (*logrus.Logger, error) {
	logger := logrus.New()
	lvl, err := logrus.ParseLevel(level)

	if err != nil {
		return nil, err
	}

	logger.SetLevel(lvl)

	return logger, nil
}
