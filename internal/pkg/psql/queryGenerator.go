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
	attrs     []string
	joins     []join
	condition string
	groupBy   []string
	having    string
	orderBy   []order
	firstRow  int
	lastRow   int
}

type join struct {
	join      string
	table     string
	condition string
}

type order struct {
	attr string
	desc bool
}

func (g *QueryGenerator) Select(config SelectConfig) string {

	joins := ""
	if len(config.joins) != 0 {
		for _, join := range config.joins {
			joins += fmt.Sprintf("%s %s ON %s", join.join, join.table, join.condition)
		}
	}

	where := ""
	if config.condition != "" {
		where = fmt.Sprintf("WHERE %s", config.condition)
	}

	groupBy := ""
	if len(config.groupBy) != 0 {
		groupBy = fmt.Sprintf("GROUP BY %s", strings.Join(config.groupBy, ", "))
	}

	orderBy := ""
	if len(config.orderBy) != 0 {
		var orders []string
		for _, order := range config.orderBy {
			sort := ""
			if order.desc {
				sort = "DESC"
			}

			orders = append(orders, fmt.Sprintf("%s %s", order.attr, sort))
		}
		orderBy = fmt.Sprintf("ORDER BY %s", strings.Join(orders, ", "))
	}

	limit := ""
	if config.firstRow != 0 {
		if config.lastRow != 0 {
			limit = fmt.Sprintf("LIMIT %d, %d", config.firstRow, config.lastRow+1)
		} else {
			limit = fmt.Sprintf("LIMIT %d", config.firstRow+1)
		}
	}

	sql := fmt.Sprintf(
		"SELECT %s FROM %s %s %s %s %s %s %s;",
		strings.Join(config.attrs, ", "),
		g.Table,
		joins,
		where,
		groupBy,
		config.having,
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
