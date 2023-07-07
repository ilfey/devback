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
	// Change pwd to root
	os.Chdir("../../..")

	// Load env
	godotenv.Load()

	// Get vars from env
	databaseUrl := utils.GetEnv("TEST_DATABASE_URL", "postgresql://ilfey:QWEasd123@localhost:5432/test")
	schemePath := utils.GetEnv("DATABASE_SCHEME", "./scheme.sql")

	// Create logger instance
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,

		DisableQuote:     true,
		DisableTimestamp: true,
	})

	// Connect to database and create store instance
	var err error
	Store, err = initPostgres(logger, databaseUrl, schemePath)
	if err != nil {
		logger.Fatal(err)
	}
}

// Initializes the database connection and returns the store
func initPostgres(logger *logrus.Logger, databaseUrl, schemePath string) (*store.Store, error) {
	db, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		logger.Errorf("error connecting to database: %s", databaseUrl)
		return nil, err
	}

	// Read database scheme
	scheme, err := utils.ReadScheme(schemePath)
	if err != nil {
		logger.Errorf("error reading scheme of database: %s", databaseUrl)
		return nil, err
	}

	// Execute database scheme
	_, err = db.Exec(context.Background(), scheme)
	if err != nil {
		logger.Errorf("error exec scheme of database: %s", databaseUrl)
		return nil, err
	}

	// Create store instance
	store := psql.NewStore(db, logger)

	return store, nil
}

// bgCtx() is a shortcut for context.Background()
func bgCtx() context.Context {
	return context.Background()
}

// Finds or creates test user
func getUser(t *testing.T) *models.User {
	u := models.TestUser(t)

	// Find user
	user, err := Store.User.Find(bgCtx(), u.Username)
	if err != nil {
		switch err.Type() {
		// User not found
		case store.StoreNotFound:
			// Create it
			user, err = Store.User.Create(bgCtx(), u)
		}
	}

	assert.NoError(t, err)

	return user
}

// Finds or creates many test users
func getManyUsers(t *testing.T) []*models.User {
	users := models.TestManyUsers(t)

	for i, u := range users {
		// Find user
		temp, err := Store.User.Find(bgCtx(), u.Username)
		if err != nil {
			switch err.Type() {
			// If user not found
			case store.StoreNotFound:
				// Create it
				temp, err = Store.User.Create(bgCtx(), u)
			}
		}

		assert.NoError(t, err)

		users[i] = temp
	}

	return users
}

// Removes test user if his exist
func removeUser(t *testing.T) {
	u := models.TestUser(t)

	// Find user
	user, err := Store.User.Find(bgCtx(), u.Username)
	if err != nil {
		switch err.Type() {
		// User not found
		case store.StoreNotFound:
			return
		}
	}

	assert.NoError(t, err)

	// Delete user
	err = Store.User.DeletePermanently(bgCtx(), user.Username)

	assert.NoError(t, err)
}

// Removes many test users if they exist
func removeManyUsers(t *testing.T) {
	users := models.TestManyUsers(t)

	for _, u := range users {
		// Remove user
		err := Store.User.DeletePermanently(bgCtx(), u.Username)

		assert.NoError(t, err)
	}
}
