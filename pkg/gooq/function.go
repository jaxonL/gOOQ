package gooq

func Count(
	expr Expression,
) Expression {
	return NewExpressionFunction("COUNT", expr)
}

///////////////////////////////////////////////////////////////////////////////
// functions for expression
///////////////////////////////////////////////////////////////////////////////

type aliasFunction struct {
	expressionImpl
	expression Expression
	alias      string
}

func newAliasFunction(
	expression Expression, alias string,
) Expression {
	function := &aliasFunction{expression: expression, alias: alias}
	function.expressionImpl.initFunctionExpression(function)
	return function
}

func (expr *aliasFunction) Render(
	builder *Builder,
) {
	builder.RenderExpression(expr.expression)
	builder.Printf(" AS %s", expr.alias)
}

type filterWhereFunction struct {
	expressionImpl
	expression Expression
}

func newFilterWhereFunction(
	expression Expression, arguments ...Expression,
) Expression {
	function := &filterWhereFunction{expression: expression}
	function.expressionImpl.initFunctionExpression(function, arguments...)
	return function
}

func (expr *filterWhereFunction) Render(
	builder *Builder,
) {
	builder.RenderExpression(expr.expression)
	builder.Print(" FILTER (WHERE ")
	builder.RenderConditions(expr.expressions)
	builder.Printf(")")
}

///////////////////////////////////////////////////////////////////////////////
// Function Expression
///////////////////////////////////////////////////////////////////////////////

type expressionFunctionImpl struct {
	expressionImpl
	name string
}

func NewExpressionFunction(
	name string, arguments ...Expression,
) Expression {
	function := &expressionFunctionImpl{name: name}
	function.expressionImpl.initFunctionExpression(function, arguments...)
	return function
}

func (expr *expressionFunctionImpl) Render(
	builder *Builder,
) {
	builder.Printf("%s(", expr.name)
	for index, argument := range expr.expressions {
		argument.Render(builder)
		if index != len(expr.expressions)-1 {
			builder.Print(", ")
		}
	}
	builder.Printf(")")
}

///////////////////////////////////////////////////////////////////////////////
// Table 9.3. Comparison Functions
// https://www.postgresql.org/docs/11/functions-comparison.html
// [Good First Issue][Help Wanted] TODO: implement remaining
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Table 9.5. Mathematical Functions
// Table 9.6. Random Functions
// Table 9.7. Trigonometric Functions
// https://www.postgresql.org/docs/11/functions-math.html
// [Good First Issue][Help Wanted] TODO: implement remaining
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Table 9.8. SQL String Functions and Operators
// Table 9.9. Other String Functions
// https://www.postgresql.org/docs/11/functions-string.html
// [Good First Issue][Help Wanted] TODO: implement remaining functions (not operators)
///////////////////////////////////////////////////////////////////////////////

func Ascii(
	input StringExpression,
) Expression {
	return NewExpressionFunction("ASCII", input)
}

func BTrim(
	source StringExpression, characters ...StringExpression,
) Expression {
	expressions := []Expression{source}
	if characters != nil {
		expressions = append(expressions, characters[0])
	}
	return NewExpressionFunction("BTRIM", expressions...)
}

func LTrim(
	source StringExpression, characters ...StringExpression,
) Expression {
	expressions := []Expression{source}
	if characters != nil {
		expressions = append(expressions, characters[0])
	}
	return NewExpressionFunction("LTRIM", expressions...)
}

func RTrim(
	source StringExpression, characters ...StringExpression,
) Expression {
	expressions := []Expression{source}
	if characters != nil {
		expressions = append(expressions, characters[0])
	}
	return NewExpressionFunction("RTRIM", expressions...)
}

///////////////////////////////////////////////////////////////////////////////
// Table 9.11. SQL Binary String Functions and Operators
// Table 9.12. Other Binary String Functions
// https://www.postgresql.org/docs/11/functions-binarystring.html
// [Good First Issue][Help Wanted] TODO: implement remaining functions (not operators)
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Table 9.23. Formatting Functions
// https://www.postgresql.org/docs/11/functions-formatting.html
// [Good First Issue][Help Wanted] TODO: implement remaining functions
////////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Table 9.30. Date/Time Functions
// https://www.postgresql.org/docs/11/functions-datetime.html
// [Good First Issue][Help Wanted] TODO: implement remaining functions
////////////////////////////////////////////////////////////////////////////////

func DateTrunc(
	text string, timestamp DateTimeExpression,
) Expression {
	expressions := []Expression{String(text), timestamp}
	return NewExpressionFunction("DATE_TRUNC", expressions...)
}

func Greatest(
	expr Expression, rests ...Expression,
) Expression {
	expressions := append([]Expression{expr}, rests...)
	return NewExpressionFunction("GREATEST", expressions...)
}

func Least(
	expr Expression, rests ...Expression,
) Expression {
	expressions := append([]Expression{expr}, rests...)
	return NewExpressionFunction("LEAST", expressions...)
}

// TODO(Peter): implement Case When

func Coalesce(
	expr Expression, rests ...Expression,
) Expression {
	expressions := append([]Expression{expr}, rests...)
	return NewExpressionFunction("COALESCE", expressions...)
}

func NullIf(
	value1, value2 Expression,
) Expression {
	expressions := []Expression{value1, value2}
	return NewExpressionFunction("NULLIF", expressions...)
}

///////////////////////////////////////////////////////////////////////////////
// Array Functions and Operators
// https://www.postgresql.org/docs/11/functions-array.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Range Functions and Operators
// https://www.postgresql.org/docs/11/functions-range.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Aggregate Functions
// https://www.postgresql.org/docs/11/functions-aggregate.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Subquery Expressions
// https://www.postgresql.org/docs/11/functions-subquery.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Window Functions
// https://www.postgresql.org/docs/11/functions-window.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Set Returning Functions
// https://www.postgresql.org/docs/11/functions-srf.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// Type Conversion - Functions
// https://www.postgresql.org/docs/11/typeconv-func.html
// [Help Wanted] TODO: implement remaining functions
///////////////////////////////////////////////////////////////////////////////
