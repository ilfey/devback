package main

import (
	"context"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilfey/devback/internal/app/config"
	"github.com/ilfey/devback/internal/app/server"
	"github.com/ilfey/devback/internal/pkg/psql"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/ttys3/rotatefilehook"
)

var (
	logLevel      string
	address       string
	port          string
	databaseUrl   string
	schemePath    string
	key           string
	apiPath       string
	adminPath     string
	adminUsername string
	lifeSpan      int
)

func main() {
	godotenv.Load()

	flag.StringVar(&logLevel, "ll", getEnv("LOGLEVEL", "info"), "LogLevel")
	flag.StringVar(&databaseUrl, "du", getEnv("DATABASE_URL", "postgresql://ilfey:QWEasd123@localhost:5432/devpage"), "PostgreSQL database url")
	flag.StringVar(&schemePath, "df", getEnv("DATABASE_SCHEME", "./scheme.sql"), "Scheme database file")
	flag.StringVar(&address, "a", getEnv("ADDRESS", "0.0.0.0"), "Address")
	flag.StringVar(&port, "p", getEnv("PORT", "8080"), "Port")
	flag.StringVar(&apiPath, "api", getEnv("API_PATH", "/api"), "Api path")
	flag.StringVar(&adminPath, "ap", getEnv("ADMIN_PATH", "/admin"), "Admin path")
	flag.StringVar(&adminUsername, "au", getEnv("ADMIN_USERNAME", "admin"), "Admin username")
	flag.StringVar(&key, "jk", getEnv("JWT_KEY", "secret"), "JWT key")
	flag.IntVar(&lifeSpan, "jls", getEnvInt("JWT_LIFE_SPAN", 24), "JWT life span (in hours)")

	flag.Parse()

	conf := new(config.Config)

	conf.Addr = address + ":" + port
	conf.LifeSpan = lifeSpan
	conf.ApiPath = apiPath
	conf.AdminPath = adminPath
	conf.AdminUsername = adminUsername
	conf.Key = key
	conf.StartTime = time.Now()

	log, err := createLogger(logLevel)
	if err != nil {
		log.Fatal(err)
	}

	store, err := initPostgres(log, databaseUrl, schemePath)
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode("release")

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

	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		ForceQuote:      true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/file.log",
		MaxSize:    50, // Mb
		MaxBackups: 3,
		MaxAge:     7, // Days
		Level:      lvl,
		Formatter: &logrus.TextFormatter{

			FullTimestamp:   true,
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	logger.AddHook(rotateFileHook)

	return logger, nil
}
