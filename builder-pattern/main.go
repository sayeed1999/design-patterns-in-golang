package main

import (
	"fmt"

	sqlquerybuilder "github.com/sayeed1999/design-patterns-in-golang/builder-pattern/sql-query-builder"
)

func main() {
	queryBuilder := sqlquerybuilder.NewQueryBuilder()

	query := queryBuilder.
		Select("salary", "name").
		Distinct().
		From("employees").
		OrderByDescending("salary").
		Limit(1).
		Offset(1).Build()

	fmt.Println("Query to return the second highest salary in postgresql:")
	fmt.Println()
	fmt.Println(query)
}
