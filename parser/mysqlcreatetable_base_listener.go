// Code generated from MySQLCreateTable.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MySQLCreateTable

import "github.com/antlr4-go/antlr/v4"

// BaseMySQLCreateTableListener is a complete listener for a parse tree produced by MySQLCreateTableParser.
type BaseMySQLCreateTableListener struct{}

var _ MySQLCreateTableListener = &BaseMySQLCreateTableListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMySQLCreateTableListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMySQLCreateTableListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMySQLCreateTableListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMySQLCreateTableListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterCreate_table_stmt is called when production create_table_stmt is entered.
func (s *BaseMySQLCreateTableListener) EnterCreate_table_stmt(ctx *Create_table_stmtContext) {}

// ExitCreate_table_stmt is called when production create_table_stmt is exited.
func (s *BaseMySQLCreateTableListener) ExitCreate_table_stmt(ctx *Create_table_stmtContext) {}

// EnterTable_element is called when production table_element is entered.
func (s *BaseMySQLCreateTableListener) EnterTable_element(ctx *Table_elementContext) {}

// ExitTable_element is called when production table_element is exited.
func (s *BaseMySQLCreateTableListener) ExitTable_element(ctx *Table_elementContext) {}

// EnterColumn_def is called when production column_def is entered.
func (s *BaseMySQLCreateTableListener) EnterColumn_def(ctx *Column_defContext) {}

// ExitColumn_def is called when production column_def is exited.
func (s *BaseMySQLCreateTableListener) ExitColumn_def(ctx *Column_defContext) {}

// EnterIntegerType is called when production IntegerType is entered.
func (s *BaseMySQLCreateTableListener) EnterIntegerType(ctx *IntegerTypeContext) {}

// ExitIntegerType is called when production IntegerType is exited.
func (s *BaseMySQLCreateTableListener) ExitIntegerType(ctx *IntegerTypeContext) {}

// EnterBigintType is called when production BigintType is entered.
func (s *BaseMySQLCreateTableListener) EnterBigintType(ctx *BigintTypeContext) {}

// ExitBigintType is called when production BigintType is exited.
func (s *BaseMySQLCreateTableListener) ExitBigintType(ctx *BigintTypeContext) {}

// EnterTinyintType is called when production TinyintType is entered.
func (s *BaseMySQLCreateTableListener) EnterTinyintType(ctx *TinyintTypeContext) {}

// ExitTinyintType is called when production TinyintType is exited.
func (s *BaseMySQLCreateTableListener) ExitTinyintType(ctx *TinyintTypeContext) {}

// EnterVarcharType is called when production VarcharType is entered.
func (s *BaseMySQLCreateTableListener) EnterVarcharType(ctx *VarcharTypeContext) {}

// ExitVarcharType is called when production VarcharType is exited.
func (s *BaseMySQLCreateTableListener) ExitVarcharType(ctx *VarcharTypeContext) {}

// EnterCharType is called when production CharType is entered.
func (s *BaseMySQLCreateTableListener) EnterCharType(ctx *CharTypeContext) {}

// ExitCharType is called when production CharType is exited.
func (s *BaseMySQLCreateTableListener) ExitCharType(ctx *CharTypeContext) {}

// EnterDecimalType is called when production DecimalType is entered.
func (s *BaseMySQLCreateTableListener) EnterDecimalType(ctx *DecimalTypeContext) {}

// ExitDecimalType is called when production DecimalType is exited.
func (s *BaseMySQLCreateTableListener) ExitDecimalType(ctx *DecimalTypeContext) {}

// EnterDatetimeType is called when production DatetimeType is entered.
func (s *BaseMySQLCreateTableListener) EnterDatetimeType(ctx *DatetimeTypeContext) {}

// ExitDatetimeType is called when production DatetimeType is exited.
func (s *BaseMySQLCreateTableListener) ExitDatetimeType(ctx *DatetimeTypeContext) {}

// EnterDateType is called when production DateType is entered.
func (s *BaseMySQLCreateTableListener) EnterDateType(ctx *DateTypeContext) {}

// ExitDateType is called when production DateType is exited.
func (s *BaseMySQLCreateTableListener) ExitDateType(ctx *DateTypeContext) {}

// EnterTimestampType is called when production TimestampType is entered.
func (s *BaseMySQLCreateTableListener) EnterTimestampType(ctx *TimestampTypeContext) {}

// ExitTimestampType is called when production TimestampType is exited.
func (s *BaseMySQLCreateTableListener) ExitTimestampType(ctx *TimestampTypeContext) {}

// EnterTextType is called when production TextType is entered.
func (s *BaseMySQLCreateTableListener) EnterTextType(ctx *TextTypeContext) {}

// ExitTextType is called when production TextType is exited.
func (s *BaseMySQLCreateTableListener) ExitTextType(ctx *TextTypeContext) {}

// EnterJsonType is called when production JsonType is entered.
func (s *BaseMySQLCreateTableListener) EnterJsonType(ctx *JsonTypeContext) {}

// ExitJsonType is called when production JsonType is exited.
func (s *BaseMySQLCreateTableListener) ExitJsonType(ctx *JsonTypeContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BaseMySQLCreateTableListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BaseMySQLCreateTableListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterAction is called when production action is entered.
func (s *BaseMySQLCreateTableListener) EnterAction(ctx *ActionContext) {}

// ExitAction is called when production action is exited.
func (s *BaseMySQLCreateTableListener) ExitAction(ctx *ActionContext) {}

// EnterTable_option is called when production table_option is entered.
func (s *BaseMySQLCreateTableListener) EnterTable_option(ctx *Table_optionContext) {}

// ExitTable_option is called when production table_option is exited.
func (s *BaseMySQLCreateTableListener) ExitTable_option(ctx *Table_optionContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseMySQLCreateTableListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseMySQLCreateTableListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseMySQLCreateTableListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseMySQLCreateTableListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterColumn_list is called when production column_list is entered.
func (s *BaseMySQLCreateTableListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BaseMySQLCreateTableListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterEngine_name is called when production engine_name is entered.
func (s *BaseMySQLCreateTableListener) EnterEngine_name(ctx *Engine_nameContext) {}

// ExitEngine_name is called when production engine_name is exited.
func (s *BaseMySQLCreateTableListener) ExitEngine_name(ctx *Engine_nameContext) {}

// EnterCharset_name is called when production charset_name is entered.
func (s *BaseMySQLCreateTableListener) EnterCharset_name(ctx *Charset_nameContext) {}

// ExitCharset_name is called when production charset_name is exited.
func (s *BaseMySQLCreateTableListener) ExitCharset_name(ctx *Charset_nameContext) {}

// EnterCollation_name is called when production collation_name is entered.
func (s *BaseMySQLCreateTableListener) EnterCollation_name(ctx *Collation_nameContext) {}

// ExitCollation_name is called when production collation_name is exited.
func (s *BaseMySQLCreateTableListener) ExitCollation_name(ctx *Collation_nameContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseMySQLCreateTableListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseMySQLCreateTableListener) ExitLiteral(ctx *LiteralContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseMySQLCreateTableListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseMySQLCreateTableListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIndex_type is called when production index_type is entered.
func (s *BaseMySQLCreateTableListener) EnterIndex_type(ctx *Index_typeContext) {}

// ExitIndex_type is called when production index_type is exited.
func (s *BaseMySQLCreateTableListener) ExitIndex_type(ctx *Index_typeContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseMySQLCreateTableListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseMySQLCreateTableListener) ExitExpr(ctx *ExprContext) {}

// EnterSingleComparison is called when production SingleComparison is entered.
func (s *BaseMySQLCreateTableListener) EnterSingleComparison(ctx *SingleComparisonContext) {}

// ExitSingleComparison is called when production SingleComparison is exited.
func (s *BaseMySQLCreateTableListener) ExitSingleComparison(ctx *SingleComparisonContext) {}

// EnterLogicalOp is called when production LogicalOp is entered.
func (s *BaseMySQLCreateTableListener) EnterLogicalOp(ctx *LogicalOpContext) {}

// ExitLogicalOp is called when production LogicalOp is exited.
func (s *BaseMySQLCreateTableListener) ExitLogicalOp(ctx *LogicalOpContext) {}

// EnterComparisonOp is called when production ComparisonOp is entered.
func (s *BaseMySQLCreateTableListener) EnterComparisonOp(ctx *ComparisonOpContext) {}

// ExitComparisonOp is called when production ComparisonOp is exited.
func (s *BaseMySQLCreateTableListener) ExitComparisonOp(ctx *ComparisonOpContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseMySQLCreateTableListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseMySQLCreateTableListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterColumnRef is called when production ColumnRef is entered.
func (s *BaseMySQLCreateTableListener) EnterColumnRef(ctx *ColumnRefContext) {}

// ExitColumnRef is called when production ColumnRef is exited.
func (s *BaseMySQLCreateTableListener) ExitColumnRef(ctx *ColumnRefContext) {}

// EnterLiteralValue is called when production LiteralValue is entered.
func (s *BaseMySQLCreateTableListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production LiteralValue is exited.
func (s *BaseMySQLCreateTableListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterFunctionExpr is called when production FunctionExpr is entered.
func (s *BaseMySQLCreateTableListener) EnterFunctionExpr(ctx *FunctionExprContext) {}

// ExitFunctionExpr is called when production FunctionExpr is exited.
func (s *BaseMySQLCreateTableListener) ExitFunctionExpr(ctx *FunctionExprContext) {}

// EnterFunction_call is called when production function_call is entered.
func (s *BaseMySQLCreateTableListener) EnterFunction_call(ctx *Function_callContext) {}

// ExitFunction_call is called when production function_call is exited.
func (s *BaseMySQLCreateTableListener) ExitFunction_call(ctx *Function_callContext) {}
