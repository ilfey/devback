package psql

import (
	"fmt"
	"strings"
)

type QueryGenerator struct {
	Table           string
	Attributes      []string
	EscapeCharacter string
}

func NewQueryGenerator(table string, attrs []string) *QueryGenerator {
	return &QueryGenerator{
		Table:           table,
		Attributes:      attrs,
		EscapeCharacter: "$",
	}
}

type SelectConfig struct {
	Attrs     []string
	Joins     []Join
	Condition string
	GroupBy   []string
	Having    string
	OrderBy   []Order
	Limit     int
}

type Join struct {
	Join      JoinType
	Table     string
	Condition string
}

type JoinType string

const (
	INNER JoinType = "INNER"
	LEFT  JoinType = "LEFT"
	RIGHT JoinType = "RIGHT"
	FULL  JoinType = "FULL"
)

type Order struct {
	Attr string
	Desc bool
}

func (g *QueryGenerator) Select(config SelectConfig) string {

	var attrs string
	if len(config.Attrs) != 0 {
		attrs = strings.Join(config.Attrs, ", ")
	} else {
		attrs = strings.Join(g.Attributes, ", ")
	}

	joins := ""
	if len(config.Joins) != 0 {
		for _, join := range config.Joins {
			joins += fmt.Sprintf(" %s JOIN %s ON %s", join.Join, join.Table, join.Condition)
		}
	}

	where := ""
	if config.Condition != "" {
		where = fmt.Sprintf(" WHERE %s", config.Condition)
	}

	groupBy := ""
	if len(config.GroupBy) != 0 {
		groupBy = fmt.Sprintf(" GROUP BY %s", strings.Join(config.GroupBy, ", "))
	}

	having := ""
	if config.Having != "" {
		having = fmt.Sprintf(" HAVING %s", config.Having)
	}

	orderBy := ""
	if len(config.OrderBy) != 0 {
		var orders []string
		for _, order := range config.OrderBy {
			sort := "ASC"
			if order.Desc {
				sort = "DESC"
			}

			orders = append(orders, fmt.Sprintf("%s %s", order.Attr, sort))
		}
		orderBy = fmt.Sprintf(" ORDER BY %s", strings.Join(orders, ", "))
	}

	limit := ""
	if config.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", config.Limit+1)
	}

	sql := fmt.Sprintf(
		"SELECT %s FROM %s%s%s%s%s%s%s;",
		attrs,
		g.Table,
		joins,
		where,
		groupBy,
		having,
		orderBy,
		limit,
	)

	return g.escapeQuery(sql)
}

func (g *QueryGenerator) Insert(attrs []string) string {
	sql := fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES (%s) RETURNING %s;",
		g.Table,
		strings.Join(attrs, ", "),
		strings.Join(g.makeFilledArray(len(attrs)), ", "),
		strings.Join(g.Attributes, ", "),
	)

	return g.escapeQuery(sql)
}

func (g *QueryGenerator) InsertMany(attrs []string, count int) string {

	var values []string

	for i := 1; i <= count; i++ {
		value := fmt.Sprintf("(%s)", strings.Join(g.makeFilledArray(len(attrs)), ", "))
		values = append(values, value)
	}

	sql := fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES %s RETURNING %s;",
		g.Table,
		strings.Join(attrs, ", "),
		strings.Join(values, ", "),
		strings.Join(g.Attributes, ", "),
	)
	return g.escapeQuery(sql)
}

func (g *QueryGenerator) Update(attrs []string, condition string) string {

	var equalses []string

	for _, attr := range attrs {
		equalses = append(equalses, fmt.Sprintf("%s = %s", attr, g.EscapeCharacter+g.EscapeCharacter))
	}

	sql := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s RETURNING %s;",
		g.Table,
		strings.Join(equalses, ", "),
		condition,
		strings.Join(g.Attributes, ", "),
	)

	return g.escapeQuery(sql)
}

func (g *QueryGenerator) Delete(condition string) string {
	sql := fmt.Sprintf(
		"DELETE FROM %s WHERE %s;",
		g.Table,
		condition,
	)

	return g.escapeQuery(sql)
}

// replace "$$" to "$x"
func (g *QueryGenerator) escapeQuery(condition string) string {
	split := strings.Split(condition, g.EscapeCharacter+g.EscapeCharacter)

	if len(split) == 1 {
		return condition
	}

	for i := 1; i < len(split); i++ {
		split[i-1] += fmt.Sprintf("%s%d", g.EscapeCharacter, i)
	}

	return strings.Join(
		split,
		"",
	)
}

func (g *QueryGenerator) makeFilledArray(lenght int) []string {
	arr := make([]string, lenght)

	for i := range arr {
		arr[i] = g.EscapeCharacter + g.EscapeCharacter
	}

	return arr
}
