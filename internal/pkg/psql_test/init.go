package testpsql

import (
	"context"
	"os"
	"testing"

	"github.com/ilfey/devback/internal/pkg/models"
	"github.com/ilfey/devback/internal/pkg/psql"
	"github.com/ilfey/devback/internal/pkg/store"
	"github.com/ilfey/devback/internal/pkg/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	Store *store.Store
)

func init() {
	// change pwd to root
	os.Chdir("../../..")

	// load env
	godotenv.Load()

	// get vars from env
	databaseUrl := utils.GetEnv("TEST_DATABASE_URL", "postgresql://ilfey:QWEasd123@localhost:5432/test")
	schemePath := utils.GetEnv("DATABASE_SCHEME", "./scheme.sql")

	// create logger instance
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,

		DisableQuote:     true,
		DisableTimestamp: true,
	})

	// connect to database and store instance
	var err error
	Store, err = initPostgres(logger, databaseUrl, schemePath)
	if err != nil {
		logger.Fatal(err)
	}
}

func initPostgres(logger *logrus.Logger, databaseUrl, schemePath string) (*store.Store, error) {
	db, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		logger.Errorf("error connecting to database: %s", databaseUrl)
		return nil, err
	}

	// read database scheme
	scheme, err := utils.ReadScheme(schemePath)
	if err != nil {
		logger.Errorf("error reading scheme of database: %s", databaseUrl)
		return nil, err
	}

	// execute database scheme
	_, err = db.Exec(context.Background(), scheme)
	if err != nil {
		logger.Errorf("error exec scheme of database: %s", databaseUrl)
		return nil, err
	}

	// create store instance
	store := psql.NewStore(db, logger)

	return store, nil
}

func bgCtx() context.Context {
	return context.Background()
}

func getUser(t *testing.T) *models.User {
	u := models.TestUser(t)

	// find user
	user, err := Store.User.Find(bgCtx(), u.Username)
	if err != nil {
		switch err.Type() {
		// user not found
		case store.StoreNotFound:
			// create it
			user, err = Store.User.Create(bgCtx(), u)
		}
	}

	assert.NoError(t, err)

	return user
}

func removeUser(t *testing.T) {
	u := models.TestUser(t)

	// find user
	user, err := Store.User.Find(bgCtx(), u.Username)
	if err != nil {
		switch err.Type() {
		// user not found
		case store.StoreNotFound:
			// create it
			user, err = Store.User.Create(bgCtx(), u)
		}
	}

	assert.NoError(t, err)

	err = Store.User.DeletePermanently(bgCtx(), user.Username)

	assert.NoError(t, err)
}
