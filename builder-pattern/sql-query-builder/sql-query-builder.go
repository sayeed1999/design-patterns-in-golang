package sqlquerybuilder

import (
	"fmt"
	"strings"
)

type sqlQueryBuilder struct {
	query string
}

func NewQueryBuilder() *sqlQueryBuilder {
	return &sqlQueryBuilder{query: ""}
}

func (qb *sqlQueryBuilder) Select(fields ...string) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("SELECT %v ", strings.Join(fields, ", "))
	return qb
}

func (qb *sqlQueryBuilder) From(table string) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("FROM %v\n", table)
	return qb
}

func (qb *sqlQueryBuilder) Distinct() *sqlQueryBuilder {
	qb.query = strings.Replace(qb.query, "SELECT", "SELECT DISTINCT", 1)
	return qb
}

func (qb *sqlQueryBuilder) Where(condition string) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("WHERE %v\n", condition)
	return qb
}

func (qb *sqlQueryBuilder) OrderBy(fields ...string) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("ORDER BY %v\n", strings.Join(fields, ", "))
	return qb
}

func (qb *sqlQueryBuilder) OrderByDescending(fields ...string) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("ORDER BY %v DESC\n", strings.Join(fields, ", "))
	return qb
}

func (qb *sqlQueryBuilder) Limit(limit int) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("LIMIT %v\n", limit)
	return qb
}

func (qb *sqlQueryBuilder) Offset(offset int) *sqlQueryBuilder {
	qb.query += fmt.Sprintf("OFFSET %v\n", offset)
	return qb
}

func (qb *sqlQueryBuilder) Build() string {
	return qb.query
}
