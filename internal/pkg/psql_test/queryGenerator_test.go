package testpsql

import (
	"strings"
	"testing"

	"github.com/ilfey/devback/internal/pkg/psql"
	"github.com/stretchr/testify/assert"
)

var gen *psql.QueryGenerator

func init() {
	gen = psql.NewQueryGenerator(
		"users",
		[]string{
			"user_id",
			"password",
			"is_deleted",
			"created_at",
			"modified_at",
		},
	)
}

func TestQueryGenerator_Insert(t *testing.T) {
	query := gen.Insert([]string{
		"user_id",
		"password",
	})

	assert.True(
		t,
		strings.ToLower(query) == "insert into users(user_id, password) values ($1, $2) returning user_id, password, is_deleted, created_at, modified_at;",
	)
}

func TestQueryGenerator_Update(t *testing.T) {
	query := gen.Update([]string{
		"user_id",
		"password",
	},
		"user_id = $$",
	)

	assert.True(
		t,
		strings.ToLower(query) == "update users set user_id = $1, password = $2 where user_id = $3 returning user_id, password, is_deleted, created_at, modified_at;",
	)
}

func TestQueryGenerator_Delete(t *testing.T) {
	query := gen.Delete("user_id = $$")

	assert.True(
		t,
		strings.ToLower(query) == "delete from users where user_id = $1;",
	)
}

func TestQueryGenerator_SelectWithEmptyConfig(t *testing.T) {
	query := gen.Select(psql.SelectConfig{})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, password, is_deleted, created_at, modified_at from users;",
	)
}

func TestQueryGenerator_SelectWithAttrs(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"password",
			"is_deleted",
		},
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, password, is_deleted from users;",
	)
}

func TestQueryGenerator_SelectWithJoins(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"password",
			"is_deleted",
			"content",
		},
		Joins: []psql.Join{
			{
				Join:      psql.LEFT,
				Table:     "messages",
				Condition: "fk_user_id = user_id",
			},
		},
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, password, is_deleted, content from users left join messages on fk_user_id = user_id;",
	)
}

func TestQueryGenerator_SelectWithCondition(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"password",
			"is_deleted",
			"content",
		},
		Joins: []psql.Join{
			{
				Join:      psql.LEFT,
				Table:     "messages",
				Condition: "fk_user_id = user_id",
			},
		},
		Condition: "length(content) < 10",
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, password, is_deleted, content from users left join messages on fk_user_id = user_id where length(content) < 10;",
	)
}

func TestQueryGenerator_SelectWithGroup(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"count(*)",
		},
		Joins: []psql.Join{
			{
				Join:      psql.LEFT,
				Table:     "messages",
				Condition: "fk_user_id = user_id",
			},
		},
		Condition: "length(content) < 10",
		GroupBy: []string{
			"user_id",
		},
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, count(*) from users left join messages on fk_user_id = user_id where length(content) < 10 group by user_id;",
	)
}

func TestQueryGenerator_SelectWithHaving(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"count(*)",
		},
		Joins: []psql.Join{
			{
				Join:      psql.LEFT,
				Table:     "messages",
				Condition: "fk_user_id = user_id",
			},
		},
		Condition: "length(content) < 10",
		GroupBy: []string{
			"user_id",
		},
		Having: "count(*) > 10",
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, count(*) from users left join messages on fk_user_id = user_id where length(content) < 10 group by user_id having count(*) > 10;",
	)
}

func TestQueryGenerator_SelectWithOrder(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"count(*)",
		},
		Joins: []psql.Join{
			{
				Join:      psql.LEFT,
				Table:     "messages",
				Condition: "fk_user_id = user_id",
			},
		},
		Condition: "length(content) < 10",
		GroupBy: []string{
			"user_id",
		},
		Having: "count(*) > 10",
		OrderBy: []psql.Order{
			{
				Attr: "count(*)",
				Desc: false,
			},
		},
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, count(*) from users left join messages on fk_user_id = user_id where length(content) < 10 group by user_id having count(*) > 10 order by count(*) asc;",
	)
}

func TestQueryGenerator_SelectWithLimit(t *testing.T) {
	query := gen.Select(psql.SelectConfig{
		Attrs: []string{
			"user_id",
			"count(*)",
		},
		Joins: []psql.Join{
			{
				Join:      psql.LEFT,
				Table:     "messages",
				Condition: "fk_user_id = user_id",
			},
		},
		Condition: "length(content) < 10",
		GroupBy: []string{
			"user_id",
		},
		Having: "count(*) > 10",
		OrderBy: []psql.Order{
			{
				Attr: "count(*)",
				Desc: false,
			},
		},
		Limit: 10,
	})

	assert.True(
		t,
		strings.ToLower(query) == "select user_id, count(*) from users left join messages on fk_user_id = user_id where length(content) < 10 group by user_id having count(*) > 10 order by count(*) asc limit 10;",
	)
}
