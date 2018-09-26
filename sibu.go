package sibu

const (
	// OpEqual represents the equal operator
	OpEqual = iota
)

// TableName represents a table name, plus it's alias. For example, "users",
// "users AS u", "users u" are all valid table names
type TableName string

// Condition represents a condition, typically used in a Where clause
type Condition struct {
	// Operand typically represents the name of the column
	Operand string
	// A constant used to represent the operator. See constants (they start
	// with Op)
	Operator int
	// The value to which it is compared
	Value interface{}
}

// Select is used to build a SELECT request
type Select struct {
	// Fields represent the fields that are going to be selected, comma
	// seperated
	Fields string

	// From creates the From clause. It can contain the alias.
	// "users", "users AS u" and "users u" are all valid values
	From TableName

	// Where represents a condition
	Where Condition

	// InnerJoin represents the table name to join to. Same rules apply
	InnerJoin TableName
}
