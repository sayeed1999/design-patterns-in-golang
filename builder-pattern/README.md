# **Builder Pattern in Golang**  

This sub-repository demonstrates the **Builder Design Pattern** in Golang using an SQL Query Builder as an example.  

## **Why Should You Care?**  

Ever built SQL queries dynamically and ended up with **string concatenation nightmares**?  

- Adding conditions (`WHERE`, `ORDER BY`) makes the query **messy**.  
- Handling optional filters leads to **tons of if-else statements**.  
- Forgetting to escape values properly can **break the query**.  

Using the **Builder Pattern**, we can construct complex queries **step by step** while keeping the code **clean, flexible, and readable**.  

## **The Problem**  

Say you need to generate SQL queries dynamically:  

1. Sometimes you need a `WHERE` clause, sometimes you don’t.  
2. Sorting (`ORDER BY`) should be optional.  
3. Filtering by multiple fields should be **composable**.  

Hardcoding these conditions in string concatenation quickly **gets out of control**.  

## **How the Builder Pattern Helps**  

1. **Encapsulates query-building logic** in a structured way.  
2. **Allows method chaining** for an intuitive API.  
3. **Prevents syntax errors** by handling query construction internally.  

### **Example Usage**  

```go
query := NewSQLBuilder().
    Select("id", "name").
    From("users").
    Where("age > ?", 18).
    OrderBy("created_at DESC").
    Build()

fmt.Println(query) 
// Output: SELECT id, name FROM users WHERE age > ? ORDER BY created_at DESC;
```

## References

Read my article on Medium on Builder Design Pattern: -

- [Builder Design Pattern: Build your own ORM that generates SQL queries](https://medium.com/@sayeedrahman_67698/builder-design-pattern-build-your-custom-orm-that-generates-sql-queries-8cc7b7d94a9f)
