package listen

import (
	"encoding/json"
	"github.com/antlr4-go/antlr/v4"
	"testing"

	"blingmoon.com/curd-tool/parser"
	"github.com/stretchr/testify/require"
)

func obj2str(obj interface{}) string {
	objByte, _ := json.Marshal(obj)
	return string(objByte)
}

func Test_listen(t *testing.T) {

	ddlFile, err := antlr.NewFileStream("creat_sql.sql")
	require.Nil(t, err)

	// Create the Lexer
	lexer := parser.NewMySQLCreateTableLexer(ddlFile)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewMySQLCreateTableParser(stream)

	// Finally walk the tree
	l := NewSqlListen2()
	create_table_stmt := p.Create_table_stmt()
	antlr.ParseTreeWalkerDefault.Walk(l, create_table_stmt)

	t.Logf("tablename: %s", l.tableName)
	t.Logf("fields: %v", obj2str(l.fields))
}

func Test_listen2(t *testing.T) {

	ddlFile, err := antlr.NewFileStream("test2_sql.sql")
	require.Nil(t, err)

	// Create the Lexer
	lexer := parser.NewMySQLCreateTableLexer(ddlFile)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewMySQLCreateTableParser(stream)

	// Finally walk the tree
	l := NewSqlListen2()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Create_table_stmt())

	t.Logf("tablename: %s", l.tableName)
	t.Logf("fields: %v", obj2str(l.fields))
}
