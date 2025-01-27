package querybuilder

import (
	"fmt"
	"strings"
)

type queryBuilder struct {
	query string
}

func NewQueryBuilder() *queryBuilder {
	return &queryBuilder{query: ""}
}

func (qb *queryBuilder) Select(fields ...string) *queryBuilder {
	qb.query += fmt.Sprintf("SELECT %v ", strings.Join(fields, ", "))
	return qb
}

func (qb *queryBuilder) From(table string) *queryBuilder {
	qb.query += fmt.Sprintf("FROM %v\n", table)
	return qb
}

func (qb *queryBuilder) Distinct() *queryBuilder {
	qb.query = strings.Replace(qb.query, "SELECT", "SELECT DISTINCT", 1)
	return qb
}

func (qb *queryBuilder) Where(condition string) *queryBuilder {
	qb.query += fmt.Sprintf("WHERE %v\n", condition)
	return qb
}

func (qb *queryBuilder) OrderBy(fields ...string) *queryBuilder {
	qb.query += fmt.Sprintf("ORDER BY %v\n", strings.Join(fields, ", "))
	return qb
}

func (qb *queryBuilder) OrderByDescending(fields ...string) *queryBuilder {
	qb.query += fmt.Sprintf("ORDER BY %v DESC\n", strings.Join(fields, ", "))
	return qb
}

func (qb *queryBuilder) Limit(limit int) *queryBuilder {
	qb.query += fmt.Sprintf("LIMIT %v\n", limit)
	return qb
}

func (qb *queryBuilder) Offset(offset int) *queryBuilder {
	qb.query += fmt.Sprintf("OFFSET %v\n", offset)
	return qb
}

func (qb *queryBuilder) Build() string {
	return qb.query
}
