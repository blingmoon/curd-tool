package listen

import (
	"blingmoon.com/curd-tool/parser"
)

type Field struct {
	GoType  string `json:"go_type"`
	SqlType string `json:"sql_type"`
	Name    string `json:"name"`
}
type SQLListen2 struct {
	*parser.BaseMySQLCreateTableListener
	tableName    string
	fields       []*Field
	currentField *Field
	// 会在EnterColumnDefinition 置为true, 在ExitColumnDefinition 置为false
	isBuildField bool //当前是否在build field,true 为当前是在build field,false 为当前不是在build field
}

func NewSqlListen2() *SQLListen2 {
	return &SQLListen2{
		fields: make([]*Field, 0),
	}
}

func (s *SQLListen2) ExitTable_name(c *parser.Table_nameContext) {
	s.tableName = c.GetText() // 获取表名
}
func (s *SQLListen2) ExitColumn_name(c *parser.Column_nameContext) {
	if s.isBuildField {
		s.currentField.Name = c.GetText()
	}
}

func (s *SQLListen2) EnterColumn_def(c *parser.Column_defContext) {
	// 每次进入一个字段定义时，都要初始化一个新的currentField
	s.currentField = &Field{}
	s.isBuildField = true
}
func (s *SQLListen2) ExitColumn_def(c *parser.Column_defContext) {
	// 每次退出一个字段定义时，都要把currentField加入到fields中
	s.fields = append(s.fields, s.currentField)
	s.isBuildField = false
}

// EnterIntegerType is called when entering the IntegerType production.
func (s *SQLListen2) EnterIntegerType(c *parser.IntegerTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "int64"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterBigintType is called when entering the BigintType production.
func (s *SQLListen2) EnterBigintType(c *parser.BigintTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "int64"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterTinyintType is called when entering the TinyintType production.
func (s *SQLListen2) EnterTinyintType(c *parser.TinyintTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "int64"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterVarcharType is called when entering the VarcharType production.
func (s *SQLListen2) EnterVarcharType(c *parser.VarcharTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "string"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterCharType is called when entering the CharType production.
func (s *SQLListen2) EnterCharType(c *parser.CharTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "string"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterDecimalType is called when entering the DecimalType production.
func (s *SQLListen2) EnterDecimalType(c *parser.DecimalTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "float64"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterDatetimeType is called when entering the DatetimeType production.
func (s *SQLListen2) EnterDatetimeType(c *parser.DatetimeTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "time.Time"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterDateType is called when entering the DateType production.
func (s *SQLListen2) EnterDateType(c *parser.DateTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "time.Time"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterTimestampType is called when entering the TimestampType production.
func (s *SQLListen2) EnterTimestampType(c *parser.TimestampTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "time.Time"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterTextType is called when entering the TextType production.
func (s *SQLListen2) EnterTextType(c *parser.TextTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "string"
		s.currentField.SqlType = c.GetText()
	}
}

// EnterJsonType is called when entering the JsonType production.
func (s *SQLListen2) EnterJsonType(c *parser.JsonTypeContext) {
	if s.isBuildField {
		s.currentField.GoType = "[]byte"
		s.currentField.SqlType = c.GetText()
	}
}
