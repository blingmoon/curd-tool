// Code generated from MySQLCreateTable.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MySQLCreateTable

import "github.com/antlr4-go/antlr/v4"

// MySQLCreateTableListener is a complete listener for a parse tree produced by MySQLCreateTableParser.
type MySQLCreateTableListener interface {
	antlr.ParseTreeListener

	// EnterCreate_table_stmt is called when entering the create_table_stmt production.
	EnterCreate_table_stmt(c *Create_table_stmtContext)

	// EnterTable_element is called when entering the table_element production.
	EnterTable_element(c *Table_elementContext)

	// EnterColumn_def is called when entering the column_def production.
	EnterColumn_def(c *Column_defContext)

	// EnterIntegerType is called when entering the IntegerType production.
	EnterIntegerType(c *IntegerTypeContext)

	// EnterBigintType is called when entering the BigintType production.
	EnterBigintType(c *BigintTypeContext)

	// EnterTinyintType is called when entering the TinyintType production.
	EnterTinyintType(c *TinyintTypeContext)

	// EnterVarcharType is called when entering the VarcharType production.
	EnterVarcharType(c *VarcharTypeContext)

	// EnterCharType is called when entering the CharType production.
	EnterCharType(c *CharTypeContext)

	// EnterDecimalType is called when entering the DecimalType production.
	EnterDecimalType(c *DecimalTypeContext)

	// EnterDatetimeType is called when entering the DatetimeType production.
	EnterDatetimeType(c *DatetimeTypeContext)

	// EnterDateType is called when entering the DateType production.
	EnterDateType(c *DateTypeContext)

	// EnterTimestampType is called when entering the TimestampType production.
	EnterTimestampType(c *TimestampTypeContext)

	// EnterTextType is called when entering the TextType production.
	EnterTextType(c *TextTypeContext)

	// EnterJsonType is called when entering the JsonType production.
	EnterJsonType(c *JsonTypeContext)

	// EnterTable_constraint is called when entering the table_constraint production.
	EnterTable_constraint(c *Table_constraintContext)

	// EnterAction is called when entering the action production.
	EnterAction(c *ActionContext)

	// EnterTable_option is called when entering the table_option production.
	EnterTable_option(c *Table_optionContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterEngine_name is called when entering the engine_name production.
	EnterEngine_name(c *Engine_nameContext)

	// EnterCharset_name is called when entering the charset_name production.
	EnterCharset_name(c *Charset_nameContext)

	// EnterCollation_name is called when entering the collation_name production.
	EnterCollation_name(c *Collation_nameContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIndex_type is called when entering the index_type production.
	EnterIndex_type(c *Index_typeContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterSingleComparison is called when entering the SingleComparison production.
	EnterSingleComparison(c *SingleComparisonContext)

	// EnterLogicalOp is called when entering the LogicalOp production.
	EnterLogicalOp(c *LogicalOpContext)

	// EnterComparisonOp is called when entering the ComparisonOp production.
	EnterComparisonOp(c *ComparisonOpContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterColumnRef is called when entering the ColumnRef production.
	EnterColumnRef(c *ColumnRefContext)

	// EnterLiteralValue is called when entering the LiteralValue production.
	EnterLiteralValue(c *LiteralValueContext)

	// EnterFunctionExpr is called when entering the FunctionExpr production.
	EnterFunctionExpr(c *FunctionExprContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// ExitCreate_table_stmt is called when exiting the create_table_stmt production.
	ExitCreate_table_stmt(c *Create_table_stmtContext)

	// ExitTable_element is called when exiting the table_element production.
	ExitTable_element(c *Table_elementContext)

	// ExitColumn_def is called when exiting the column_def production.
	ExitColumn_def(c *Column_defContext)

	// ExitIntegerType is called when exiting the IntegerType production.
	ExitIntegerType(c *IntegerTypeContext)

	// ExitBigintType is called when exiting the BigintType production.
	ExitBigintType(c *BigintTypeContext)

	// ExitTinyintType is called when exiting the TinyintType production.
	ExitTinyintType(c *TinyintTypeContext)

	// ExitVarcharType is called when exiting the VarcharType production.
	ExitVarcharType(c *VarcharTypeContext)

	// ExitCharType is called when exiting the CharType production.
	ExitCharType(c *CharTypeContext)

	// ExitDecimalType is called when exiting the DecimalType production.
	ExitDecimalType(c *DecimalTypeContext)

	// ExitDatetimeType is called when exiting the DatetimeType production.
	ExitDatetimeType(c *DatetimeTypeContext)

	// ExitDateType is called when exiting the DateType production.
	ExitDateType(c *DateTypeContext)

	// ExitTimestampType is called when exiting the TimestampType production.
	ExitTimestampType(c *TimestampTypeContext)

	// ExitTextType is called when exiting the TextType production.
	ExitTextType(c *TextTypeContext)

	// ExitJsonType is called when exiting the JsonType production.
	ExitJsonType(c *JsonTypeContext)

	// ExitTable_constraint is called when exiting the table_constraint production.
	ExitTable_constraint(c *Table_constraintContext)

	// ExitAction is called when exiting the action production.
	ExitAction(c *ActionContext)

	// ExitTable_option is called when exiting the table_option production.
	ExitTable_option(c *Table_optionContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitEngine_name is called when exiting the engine_name production.
	ExitEngine_name(c *Engine_nameContext)

	// ExitCharset_name is called when exiting the charset_name production.
	ExitCharset_name(c *Charset_nameContext)

	// ExitCollation_name is called when exiting the collation_name production.
	ExitCollation_name(c *Collation_nameContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIndex_type is called when exiting the index_type production.
	ExitIndex_type(c *Index_typeContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitSingleComparison is called when exiting the SingleComparison production.
	ExitSingleComparison(c *SingleComparisonContext)

	// ExitLogicalOp is called when exiting the LogicalOp production.
	ExitLogicalOp(c *LogicalOpContext)

	// ExitComparisonOp is called when exiting the ComparisonOp production.
	ExitComparisonOp(c *ComparisonOpContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitColumnRef is called when exiting the ColumnRef production.
	ExitColumnRef(c *ColumnRefContext)

	// ExitLiteralValue is called when exiting the LiteralValue production.
	ExitLiteralValue(c *LiteralValueContext)

	// ExitFunctionExpr is called when exiting the FunctionExpr production.
	ExitFunctionExpr(c *FunctionExprContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)
}
