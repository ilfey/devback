package testpsql

import (
	"context"
	"os"

	"github.com/ilfey/devback/internal/pkg/psql"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/ilfey/devback/internal/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	Store *store.Store
)

func init() {
	os.Chdir("../../..") // change pwd to root

	godotenv.Load()

	databaseUrl := utils.GetEnv("TEST_DATABASE_URL", "postgresql://ilfey:QWEasd123@localhost:5432/test")
	schemePath := utils.GetEnv("DATABASE_SCHEME", "./scheme.sql")

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	var err error
	Store, err = initPostgres(logger, databaseUrl, schemePath)
	if err != nil {
		logger.Fatal(err)
	}
}

func initPostgres(logger *logrus.Logger, databaseUrl, schemePath string) (*store.Store, error) {
	db, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		logger.Errorf("error connecting to database: %s", databaseUrl)
		return nil, err
	}

	scheme, err := utils.ReadScheme(schemePath)
	if err != nil {
		logger.Errorf("error reading scheme of database: %s", databaseUrl)
		return nil, err
	}

	_, err = db.Exec(context.Background(), scheme)
	if err != nil {
		logger.Errorf("error exec scheme of database: %s", databaseUrl)
		return nil, err
	}

	store := psql.NewStore(db, logger)

	return store, nil
}

func bgCtx() context.Context {
	return context.Background()
}
