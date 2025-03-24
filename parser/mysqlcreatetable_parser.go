// Code generated from MySQLCreateTable.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MySQLCreateTable

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MySQLCreateTableParser struct {
	*antlr.BaseParser
}

var MySQLCreateTableParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func mysqlcreatetableParserInit() {
	staticData := &MySQLCreateTableParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "','", "'('", "')'", "';'", "'.'",
	}
	staticData.SymbolicNames = []string{
		"", "CREATE", "TEMPORARY", "TABLE", "IF", "NOT", "EXISTS", "PRIMARY",
		"KEY", "FOREIGN", "REFERENCES", "INDEX", "UNIQUE", "CONSTRAINT", "CHECK",
		"DEFAULT", "AUTO_INCREMENT", "ENGINE", "CHARSET", "COLLATE", "COMMENT",
		"NULL", "CURRENT_TIMESTAMP", "ON", "UPDATE", "CHARACTER", "SET", "USING",
		"JSON", "AND", "OR", "INT", "BIGINT", "TINYINT", "VARCHAR", "CHAR",
		"DECIMAL", "DATETIME", "DATE", "TIMESTAMP", "TEXT", "UNSIGNED", "OPERATOR",
		"COMMA", "LPAREN", "RPAREN", "SEMI", "DOT", "DELETE", "RESTRICT", "CASCADE",
		"NO", "ACTION", "IDENTIFIER", "STRING_LITERAL", "NUMERIC_LITERAL", "WS",
		"LINE_COMMENT", "BLOCK_COMMENT",
	}
	staticData.RuleNames = []string{
		"create_table_stmt", "table_element", "column_def", "data_type", "table_constraint",
		"action", "table_option", "table_name", "column_name", "column_list",
		"engine_name", "charset_name", "collation_name", "literal", "identifier",
		"index_type", "expr", "logical_expr", "comparison_expr", "primary_expr",
		"function_call",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 359, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 3, 0, 45, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 51, 8, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 58, 8, 0, 10, 0, 12, 0, 61, 9, 0, 1, 0, 1,
		0, 5, 0, 65, 8, 0, 10, 0, 12, 0, 68, 9, 0, 1, 0, 3, 0, 71, 8, 0, 1, 1,
		1, 1, 3, 1, 75, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 82, 8, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 90, 8, 2, 3, 2, 92, 8, 2, 3, 2,
		94, 8, 2, 1, 2, 3, 2, 97, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 103, 8, 2,
		3, 2, 105, 8, 2, 1, 2, 1, 2, 3, 2, 109, 8, 2, 1, 2, 1, 2, 1, 2, 3, 2, 114,
		8, 2, 1, 2, 1, 2, 3, 2, 118, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 5, 2, 128, 8, 2, 10, 2, 12, 2, 131, 9, 2, 3, 2, 133, 8, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 140, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 146, 8, 3, 1, 3, 3, 3, 149, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 155,
		8, 3, 1, 3, 3, 3, 158, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 164, 8, 3, 1,
		3, 3, 3, 167, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 182, 8, 3, 1, 3, 3, 3, 185, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 192, 8, 3, 1, 4, 1, 4, 3, 4, 196, 8, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 217, 8, 4, 10, 4, 12, 4, 220,
		9, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 229, 8, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 236, 8, 4, 1, 4, 1, 4, 3, 4, 240, 8, 4, 1,
		4, 3, 4, 243, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 250, 8, 4, 3, 4,
		252, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 260, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 266, 8, 6, 1, 6, 1, 6, 1, 6, 3, 6, 271, 8, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 276, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 287, 8, 6, 1, 7, 1, 7, 1, 7, 3, 7, 292, 8, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 9, 5, 9, 299, 8, 9, 10, 9, 12, 9, 302, 9, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 5, 17, 324, 8,
		17, 10, 17, 12, 17, 327, 9, 17, 1, 18, 1, 18, 1, 18, 3, 18, 332, 8, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 338, 8, 18, 1, 19, 1, 19, 1, 19, 3,
		19, 343, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 350, 8, 20, 10,
		20, 12, 20, 353, 9, 20, 3, 20, 355, 8, 20, 1, 20, 1, 20, 1, 20, 0, 1, 34,
		21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 0, 4, 2, 0, 24, 24, 48, 48, 2, 0, 8, 8, 11, 11, 2, 0, 21, 22,
		54, 55, 1, 0, 29, 30, 406, 0, 42, 1, 0, 0, 0, 2, 74, 1, 0, 0, 0, 4, 76,
		1, 0, 0, 0, 6, 191, 1, 0, 0, 0, 8, 195, 1, 0, 0, 0, 10, 259, 1, 0, 0, 0,
		12, 286, 1, 0, 0, 0, 14, 288, 1, 0, 0, 0, 16, 293, 1, 0, 0, 0, 18, 295,
		1, 0, 0, 0, 20, 303, 1, 0, 0, 0, 22, 305, 1, 0, 0, 0, 24, 307, 1, 0, 0,
		0, 26, 309, 1, 0, 0, 0, 28, 311, 1, 0, 0, 0, 30, 313, 1, 0, 0, 0, 32, 315,
		1, 0, 0, 0, 34, 317, 1, 0, 0, 0, 36, 337, 1, 0, 0, 0, 38, 342, 1, 0, 0,
		0, 40, 344, 1, 0, 0, 0, 42, 44, 5, 1, 0, 0, 43, 45, 5, 2, 0, 0, 44, 43,
		1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 50, 5, 3, 0, 0,
		47, 48, 5, 4, 0, 0, 48, 49, 5, 5, 0, 0, 49, 51, 5, 6, 0, 0, 50, 47, 1,
		0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 3, 14, 7, 0, 53,
		54, 5, 44, 0, 0, 54, 59, 3, 2, 1, 0, 55, 56, 5, 43, 0, 0, 56, 58, 3, 2,
		1, 0, 57, 55, 1, 0, 0, 0, 58, 61, 1, 0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60,
		1, 0, 0, 0, 60, 62, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 62, 66, 5, 45, 0, 0,
		63, 65, 3, 12, 6, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 70, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69,
		71, 5, 46, 0, 0, 70, 69, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 1, 1, 0, 0,
		0, 72, 75, 3, 4, 2, 0, 73, 75, 3, 8, 4, 0, 74, 72, 1, 0, 0, 0, 74, 73,
		1, 0, 0, 0, 75, 3, 1, 0, 0, 0, 76, 77, 3, 16, 8, 0, 77, 81, 3, 6, 3, 0,
		78, 79, 5, 5, 0, 0, 79, 82, 5, 21, 0, 0, 80, 82, 5, 21, 0, 0, 81, 78, 1,
		0, 0, 0, 81, 80, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 93, 1, 0, 0, 0, 83,
		91, 5, 15, 0, 0, 84, 92, 3, 26, 13, 0, 85, 89, 5, 22, 0, 0, 86, 87, 5,
		23, 0, 0, 87, 88, 5, 24, 0, 0, 88, 90, 5, 22, 0, 0, 89, 86, 1, 0, 0, 0,
		89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 84, 1, 0, 0, 0, 91, 85, 1,
		0, 0, 0, 92, 94, 1, 0, 0, 0, 93, 83, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94,
		96, 1, 0, 0, 0, 95, 97, 5, 16, 0, 0, 96, 95, 1, 0, 0, 0, 96, 97, 1, 0,
		0, 0, 97, 104, 1, 0, 0, 0, 98, 99, 5, 7, 0, 0, 99, 105, 5, 8, 0, 0, 100,
		102, 5, 12, 0, 0, 101, 103, 5, 8, 0, 0, 102, 101, 1, 0, 0, 0, 102, 103,
		1, 0, 0, 0, 103, 105, 1, 0, 0, 0, 104, 98, 1, 0, 0, 0, 104, 100, 1, 0,
		0, 0, 104, 105, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 107, 5, 20, 0, 0,
		107, 109, 5, 54, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		113, 1, 0, 0, 0, 110, 111, 5, 25, 0, 0, 111, 112, 5, 26, 0, 0, 112, 114,
		3, 22, 11, 0, 113, 110, 1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 117, 1,
		0, 0, 0, 115, 116, 5, 19, 0, 0, 116, 118, 3, 24, 12, 0, 117, 115, 1, 0,
		0, 0, 117, 118, 1, 0, 0, 0, 118, 132, 1, 0, 0, 0, 119, 120, 5, 10, 0, 0,
		120, 121, 3, 14, 7, 0, 121, 122, 5, 44, 0, 0, 122, 123, 3, 16, 8, 0, 123,
		129, 5, 45, 0, 0, 124, 125, 5, 23, 0, 0, 125, 126, 7, 0, 0, 0, 126, 128,
		3, 10, 5, 0, 127, 124, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0,
		0, 0, 129, 130, 1, 0, 0, 0, 130, 133, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0,
		132, 119, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 139, 1, 0, 0, 0, 134,
		135, 5, 14, 0, 0, 135, 136, 5, 44, 0, 0, 136, 137, 3, 32, 16, 0, 137, 138,
		5, 45, 0, 0, 138, 140, 1, 0, 0, 0, 139, 134, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 5, 1, 0, 0, 0, 141, 145, 5, 31, 0, 0, 142, 143, 5, 44, 0, 0,
		143, 144, 5, 55, 0, 0, 144, 146, 5, 45, 0, 0, 145, 142, 1, 0, 0, 0, 145,
		146, 1, 0, 0, 0, 146, 148, 1, 0, 0, 0, 147, 149, 5, 41, 0, 0, 148, 147,
		1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 192, 1, 0, 0, 0, 150, 154, 5, 32,
		0, 0, 151, 152, 5, 44, 0, 0, 152, 153, 5, 55, 0, 0, 153, 155, 5, 45, 0,
		0, 154, 151, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 157, 1, 0, 0, 0, 156,
		158, 5, 41, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 192,
		1, 0, 0, 0, 159, 163, 5, 33, 0, 0, 160, 161, 5, 44, 0, 0, 161, 162, 5,
		55, 0, 0, 162, 164, 5, 45, 0, 0, 163, 160, 1, 0, 0, 0, 163, 164, 1, 0,
		0, 0, 164, 166, 1, 0, 0, 0, 165, 167, 5, 41, 0, 0, 166, 165, 1, 0, 0, 0,
		166, 167, 1, 0, 0, 0, 167, 192, 1, 0, 0, 0, 168, 169, 5, 34, 0, 0, 169,
		170, 5, 44, 0, 0, 170, 171, 5, 55, 0, 0, 171, 192, 5, 45, 0, 0, 172, 173,
		5, 35, 0, 0, 173, 174, 5, 44, 0, 0, 174, 175, 5, 55, 0, 0, 175, 192, 5,
		45, 0, 0, 176, 184, 5, 36, 0, 0, 177, 178, 5, 44, 0, 0, 178, 181, 5, 55,
		0, 0, 179, 180, 5, 43, 0, 0, 180, 182, 5, 55, 0, 0, 181, 179, 1, 0, 0,
		0, 181, 182, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 185, 5, 45, 0, 0, 184,
		177, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 192, 1, 0, 0, 0, 186, 192,
		5, 37, 0, 0, 187, 192, 5, 38, 0, 0, 188, 192, 5, 39, 0, 0, 189, 192, 5,
		40, 0, 0, 190, 192, 5, 28, 0, 0, 191, 141, 1, 0, 0, 0, 191, 150, 1, 0,
		0, 0, 191, 159, 1, 0, 0, 0, 191, 168, 1, 0, 0, 0, 191, 172, 1, 0, 0, 0,
		191, 176, 1, 0, 0, 0, 191, 186, 1, 0, 0, 0, 191, 187, 1, 0, 0, 0, 191,
		188, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 191, 190, 1, 0, 0, 0, 192, 7, 1,
		0, 0, 0, 193, 194, 5, 13, 0, 0, 194, 196, 3, 28, 14, 0, 195, 193, 1, 0,
		0, 0, 195, 196, 1, 0, 0, 0, 196, 251, 1, 0, 0, 0, 197, 198, 5, 7, 0, 0,
		198, 199, 5, 8, 0, 0, 199, 200, 5, 44, 0, 0, 200, 201, 3, 18, 9, 0, 201,
		202, 5, 45, 0, 0, 202, 252, 1, 0, 0, 0, 203, 204, 5, 9, 0, 0, 204, 205,
		5, 8, 0, 0, 205, 206, 5, 44, 0, 0, 206, 207, 3, 18, 9, 0, 207, 208, 5,
		45, 0, 0, 208, 209, 5, 10, 0, 0, 209, 210, 3, 14, 7, 0, 210, 211, 5, 44,
		0, 0, 211, 212, 3, 18, 9, 0, 212, 218, 5, 45, 0, 0, 213, 214, 5, 23, 0,
		0, 214, 215, 7, 0, 0, 0, 215, 217, 3, 10, 5, 0, 216, 213, 1, 0, 0, 0, 217,
		220, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 252,
		1, 0, 0, 0, 220, 218, 1, 0, 0, 0, 221, 222, 5, 14, 0, 0, 222, 223, 5, 44,
		0, 0, 223, 224, 3, 32, 16, 0, 224, 225, 5, 45, 0, 0, 225, 252, 1, 0, 0,
		0, 226, 228, 7, 1, 0, 0, 227, 229, 3, 28, 14, 0, 228, 227, 1, 0, 0, 0,
		228, 229, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 5, 44, 0, 0, 231,
		232, 3, 18, 9, 0, 232, 235, 5, 45, 0, 0, 233, 234, 5, 27, 0, 0, 234, 236,
		3, 30, 15, 0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 252, 1,
		0, 0, 0, 237, 239, 5, 12, 0, 0, 238, 240, 7, 1, 0, 0, 239, 238, 1, 0, 0,
		0, 239, 240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 243, 3, 28, 14, 0,
		242, 241, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244,
		245, 5, 44, 0, 0, 245, 246, 3, 18, 9, 0, 246, 249, 5, 45, 0, 0, 247, 248,
		5, 27, 0, 0, 248, 250, 3, 30, 15, 0, 249, 247, 1, 0, 0, 0, 249, 250, 1,
		0, 0, 0, 250, 252, 1, 0, 0, 0, 251, 197, 1, 0, 0, 0, 251, 203, 1, 0, 0,
		0, 251, 221, 1, 0, 0, 0, 251, 226, 1, 0, 0, 0, 251, 237, 1, 0, 0, 0, 252,
		9, 1, 0, 0, 0, 253, 260, 5, 49, 0, 0, 254, 260, 5, 50, 0, 0, 255, 256,
		5, 26, 0, 0, 256, 260, 5, 21, 0, 0, 257, 258, 5, 51, 0, 0, 258, 260, 5,
		52, 0, 0, 259, 253, 1, 0, 0, 0, 259, 254, 1, 0, 0, 0, 259, 255, 1, 0, 0,
		0, 259, 257, 1, 0, 0, 0, 260, 11, 1, 0, 0, 0, 261, 262, 5, 17, 0, 0, 262,
		263, 5, 42, 0, 0, 263, 287, 3, 20, 10, 0, 264, 266, 5, 15, 0, 0, 265, 264,
		1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 270, 1, 0, 0, 0, 267, 271, 5, 18,
		0, 0, 268, 269, 5, 25, 0, 0, 269, 271, 5, 26, 0, 0, 270, 267, 1, 0, 0,
		0, 270, 268, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 273, 5, 42, 0, 0, 273,
		287, 3, 22, 11, 0, 274, 276, 5, 15, 0, 0, 275, 274, 1, 0, 0, 0, 275, 276,
		1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 278, 5, 19, 0, 0, 278, 279, 5, 42,
		0, 0, 279, 287, 3, 24, 12, 0, 280, 281, 5, 16, 0, 0, 281, 282, 5, 42, 0,
		0, 282, 287, 5, 55, 0, 0, 283, 284, 5, 20, 0, 0, 284, 285, 5, 42, 0, 0,
		285, 287, 5, 54, 0, 0, 286, 261, 1, 0, 0, 0, 286, 265, 1, 0, 0, 0, 286,
		275, 1, 0, 0, 0, 286, 280, 1, 0, 0, 0, 286, 283, 1, 0, 0, 0, 287, 13, 1,
		0, 0, 0, 288, 291, 5, 53, 0, 0, 289, 290, 5, 47, 0, 0, 290, 292, 5, 53,
		0, 0, 291, 289, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 15, 1, 0, 0, 0,
		293, 294, 5, 53, 0, 0, 294, 17, 1, 0, 0, 0, 295, 300, 3, 16, 8, 0, 296,
		297, 5, 43, 0, 0, 297, 299, 3, 16, 8, 0, 298, 296, 1, 0, 0, 0, 299, 302,
		1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 19, 1, 0,
		0, 0, 302, 300, 1, 0, 0, 0, 303, 304, 5, 53, 0, 0, 304, 21, 1, 0, 0, 0,
		305, 306, 5, 53, 0, 0, 306, 23, 1, 0, 0, 0, 307, 308, 5, 53, 0, 0, 308,
		25, 1, 0, 0, 0, 309, 310, 7, 2, 0, 0, 310, 27, 1, 0, 0, 0, 311, 312, 5,
		53, 0, 0, 312, 29, 1, 0, 0, 0, 313, 314, 5, 53, 0, 0, 314, 31, 1, 0, 0,
		0, 315, 316, 3, 34, 17, 0, 316, 33, 1, 0, 0, 0, 317, 318, 6, 17, -1, 0,
		318, 319, 3, 36, 18, 0, 319, 325, 1, 0, 0, 0, 320, 321, 10, 2, 0, 0, 321,
		322, 7, 3, 0, 0, 322, 324, 3, 36, 18, 0, 323, 320, 1, 0, 0, 0, 324, 327,
		1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 35, 1, 0,
		0, 0, 327, 325, 1, 0, 0, 0, 328, 331, 3, 38, 19, 0, 329, 330, 5, 42, 0,
		0, 330, 332, 3, 38, 19, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0,
		332, 338, 1, 0, 0, 0, 333, 334, 5, 44, 0, 0, 334, 335, 3, 32, 16, 0, 335,
		336, 5, 45, 0, 0, 336, 338, 1, 0, 0, 0, 337, 328, 1, 0, 0, 0, 337, 333,
		1, 0, 0, 0, 338, 37, 1, 0, 0, 0, 339, 343, 3, 16, 8, 0, 340, 343, 3, 26,
		13, 0, 341, 343, 3, 40, 20, 0, 342, 339, 1, 0, 0, 0, 342, 340, 1, 0, 0,
		0, 342, 341, 1, 0, 0, 0, 343, 39, 1, 0, 0, 0, 344, 345, 5, 53, 0, 0, 345,
		354, 5, 44, 0, 0, 346, 351, 3, 32, 16, 0, 347, 348, 5, 43, 0, 0, 348, 350,
		3, 32, 16, 0, 349, 347, 1, 0, 0, 0, 350, 353, 1, 0, 0, 0, 351, 349, 1,
		0, 0, 0, 351, 352, 1, 0, 0, 0, 352, 355, 1, 0, 0, 0, 353, 351, 1, 0, 0,
		0, 354, 346, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356,
		357, 5, 45, 0, 0, 357, 41, 1, 0, 0, 0, 49, 44, 50, 59, 66, 70, 74, 81,
		89, 91, 93, 96, 102, 104, 108, 113, 117, 129, 132, 139, 145, 148, 154,
		157, 163, 166, 181, 184, 191, 195, 218, 228, 235, 239, 242, 249, 251, 259,
		265, 270, 275, 286, 291, 300, 325, 331, 337, 342, 351, 354,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MySQLCreateTableParserInit initializes any static state used to implement MySQLCreateTableParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMySQLCreateTableParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MySQLCreateTableParserInit() {
	staticData := &MySQLCreateTableParserStaticData
	staticData.once.Do(mysqlcreatetableParserInit)
}

// NewMySQLCreateTableParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMySQLCreateTableParser(input antlr.TokenStream) *MySQLCreateTableParser {
	MySQLCreateTableParserInit()
	this := new(MySQLCreateTableParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &MySQLCreateTableParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "MySQLCreateTable.g4"

	return this
}

// MySQLCreateTableParser tokens.
const (
	MySQLCreateTableParserEOF               = antlr.TokenEOF
	MySQLCreateTableParserCREATE            = 1
	MySQLCreateTableParserTEMPORARY         = 2
	MySQLCreateTableParserTABLE             = 3
	MySQLCreateTableParserIF                = 4
	MySQLCreateTableParserNOT               = 5
	MySQLCreateTableParserEXISTS            = 6
	MySQLCreateTableParserPRIMARY           = 7
	MySQLCreateTableParserKEY               = 8
	MySQLCreateTableParserFOREIGN           = 9
	MySQLCreateTableParserREFERENCES        = 10
	MySQLCreateTableParserINDEX             = 11
	MySQLCreateTableParserUNIQUE            = 12
	MySQLCreateTableParserCONSTRAINT        = 13
	MySQLCreateTableParserCHECK             = 14
	MySQLCreateTableParserDEFAULT           = 15
	MySQLCreateTableParserAUTO_INCREMENT    = 16
	MySQLCreateTableParserENGINE            = 17
	MySQLCreateTableParserCHARSET           = 18
	MySQLCreateTableParserCOLLATE           = 19
	MySQLCreateTableParserCOMMENT           = 20
	MySQLCreateTableParserNULL              = 21
	MySQLCreateTableParserCURRENT_TIMESTAMP = 22
	MySQLCreateTableParserON                = 23
	MySQLCreateTableParserUPDATE            = 24
	MySQLCreateTableParserCHARACTER         = 25
	MySQLCreateTableParserSET               = 26
	MySQLCreateTableParserUSING             = 27
	MySQLCreateTableParserJSON              = 28
	MySQLCreateTableParserAND               = 29
	MySQLCreateTableParserOR                = 30
	MySQLCreateTableParserINT               = 31
	MySQLCreateTableParserBIGINT            = 32
	MySQLCreateTableParserTINYINT           = 33
	MySQLCreateTableParserVARCHAR           = 34
	MySQLCreateTableParserCHAR              = 35
	MySQLCreateTableParserDECIMAL           = 36
	MySQLCreateTableParserDATETIME          = 37
	MySQLCreateTableParserDATE              = 38
	MySQLCreateTableParserTIMESTAMP         = 39
	MySQLCreateTableParserTEXT              = 40
	MySQLCreateTableParserUNSIGNED          = 41
	MySQLCreateTableParserOPERATOR          = 42
	MySQLCreateTableParserCOMMA             = 43
	MySQLCreateTableParserLPAREN            = 44
	MySQLCreateTableParserRPAREN            = 45
	MySQLCreateTableParserSEMI              = 46
	MySQLCreateTableParserDOT               = 47
	MySQLCreateTableParserDELETE            = 48
	MySQLCreateTableParserRESTRICT          = 49
	MySQLCreateTableParserCASCADE           = 50
	MySQLCreateTableParserNO                = 51
	MySQLCreateTableParserACTION            = 52
	MySQLCreateTableParserIDENTIFIER        = 53
	MySQLCreateTableParserSTRING_LITERAL    = 54
	MySQLCreateTableParserNUMERIC_LITERAL   = 55
	MySQLCreateTableParserWS                = 56
	MySQLCreateTableParserLINE_COMMENT      = 57
	MySQLCreateTableParserBLOCK_COMMENT     = 58
)

// MySQLCreateTableParser rules.
const (
	MySQLCreateTableParserRULE_create_table_stmt = 0
	MySQLCreateTableParserRULE_table_element     = 1
	MySQLCreateTableParserRULE_column_def        = 2
	MySQLCreateTableParserRULE_data_type         = 3
	MySQLCreateTableParserRULE_table_constraint  = 4
	MySQLCreateTableParserRULE_action            = 5
	MySQLCreateTableParserRULE_table_option      = 6
	MySQLCreateTableParserRULE_table_name        = 7
	MySQLCreateTableParserRULE_column_name       = 8
	MySQLCreateTableParserRULE_column_list       = 9
	MySQLCreateTableParserRULE_engine_name       = 10
	MySQLCreateTableParserRULE_charset_name      = 11
	MySQLCreateTableParserRULE_collation_name    = 12
	MySQLCreateTableParserRULE_literal           = 13
	MySQLCreateTableParserRULE_identifier        = 14
	MySQLCreateTableParserRULE_index_type        = 15
	MySQLCreateTableParserRULE_expr              = 16
	MySQLCreateTableParserRULE_logical_expr      = 17
	MySQLCreateTableParserRULE_comparison_expr   = 18
	MySQLCreateTableParserRULE_primary_expr      = 19
	MySQLCreateTableParserRULE_function_call     = 20
)

// ICreate_table_stmtContext is an interface to support dynamic dispatch.
type ICreate_table_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	Table_name() ITable_nameContext
	LPAREN() antlr.TerminalNode
	AllTable_element() []ITable_elementContext
	Table_element(i int) ITable_elementContext
	RPAREN() antlr.TerminalNode
	TEMPORARY() antlr.TerminalNode
	IF() antlr.TerminalNode
	NOT() antlr.TerminalNode
	EXISTS() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllTable_option() []ITable_optionContext
	Table_option(i int) ITable_optionContext
	SEMI() antlr.TerminalNode

	// IsCreate_table_stmtContext differentiates from other interfaces.
	IsCreate_table_stmtContext()
}

type Create_table_stmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreate_table_stmtContext() *Create_table_stmtContext {
	var p = new(Create_table_stmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_create_table_stmt
	return p
}

func InitEmptyCreate_table_stmtContext(p *Create_table_stmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_create_table_stmt
}

func (*Create_table_stmtContext) IsCreate_table_stmtContext() {}

func NewCreate_table_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Create_table_stmtContext {
	var p = new(Create_table_stmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_create_table_stmt

	return p
}

func (s *Create_table_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Create_table_stmtContext) CREATE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCREATE, 0)
}

func (s *Create_table_stmtContext) TABLE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserTABLE, 0)
}

func (s *Create_table_stmtContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Create_table_stmtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *Create_table_stmtContext) AllTable_element() []ITable_elementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITable_elementContext); ok {
			len++
		}
	}

	tst := make([]ITable_elementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITable_elementContext); ok {
			tst[i] = t.(ITable_elementContext)
			i++
		}
	}

	return tst
}

func (s *Create_table_stmtContext) Table_element(i int) ITable_elementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_elementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_elementContext)
}

func (s *Create_table_stmtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *Create_table_stmtContext) TEMPORARY() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserTEMPORARY, 0)
}

func (s *Create_table_stmtContext) IF() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIF, 0)
}

func (s *Create_table_stmtContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNOT, 0)
}

func (s *Create_table_stmtContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserEXISTS, 0)
}

func (s *Create_table_stmtContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserCOMMA)
}

func (s *Create_table_stmtContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOMMA, i)
}

func (s *Create_table_stmtContext) AllTable_option() []ITable_optionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITable_optionContext); ok {
			len++
		}
	}

	tst := make([]ITable_optionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITable_optionContext); ok {
			tst[i] = t.(ITable_optionContext)
			i++
		}
	}

	return tst
}

func (s *Create_table_stmtContext) Table_option(i int) ITable_optionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_optionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_optionContext)
}

func (s *Create_table_stmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSEMI, 0)
}

func (s *Create_table_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Create_table_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Create_table_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterCreate_table_stmt(s)
	}
}

func (s *Create_table_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitCreate_table_stmt(s)
	}
}

func (p *MySQLCreateTableParser) Create_table_stmt() (localctx ICreate_table_stmtContext) {
	localctx = NewCreate_table_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MySQLCreateTableParserRULE_create_table_stmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(MySQLCreateTableParserCREATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserTEMPORARY {
		{
			p.SetState(43)
			p.Match(MySQLCreateTableParserTEMPORARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(46)
		p.Match(MySQLCreateTableParserTABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserIF {
		{
			p.SetState(47)
			p.Match(MySQLCreateTableParserIF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(48)
			p.Match(MySQLCreateTableParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(49)
			p.Match(MySQLCreateTableParserEXISTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(52)
		p.Table_name()
	}
	{
		p.SetState(53)
		p.Match(MySQLCreateTableParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.Table_element()
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MySQLCreateTableParserCOMMA {
		{
			p.SetState(55)
			p.Match(MySQLCreateTableParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(56)
			p.Table_element()
		}

		p.SetState(61)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(62)
		p.Match(MySQLCreateTableParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&35618816) != 0 {
		{
			p.SetState(63)
			p.Table_option()
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserSEMI {
		{
			p.SetState(69)
			p.Match(MySQLCreateTableParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITable_elementContext is an interface to support dynamic dispatch.
type ITable_elementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Column_def() IColumn_defContext
	Table_constraint() ITable_constraintContext

	// IsTable_elementContext differentiates from other interfaces.
	IsTable_elementContext()
}

type Table_elementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_elementContext() *Table_elementContext {
	var p = new(Table_elementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_element
	return p
}

func InitEmptyTable_elementContext(p *Table_elementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_element
}

func (*Table_elementContext) IsTable_elementContext() {}

func NewTable_elementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_elementContext {
	var p = new(Table_elementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_table_element

	return p
}

func (s *Table_elementContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_elementContext) Column_def() IColumn_defContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_defContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_defContext)
}

func (s *Table_elementContext) Table_constraint() ITable_constraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_constraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_constraintContext)
}

func (s *Table_elementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_elementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_elementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTable_element(s)
	}
}

func (s *Table_elementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTable_element(s)
	}
}

func (p *MySQLCreateTableParser) Table_element() (localctx ITable_elementContext) {
	localctx = NewTable_elementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MySQLCreateTableParserRULE_table_element)
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Column_def()
		}

	case MySQLCreateTableParserPRIMARY, MySQLCreateTableParserKEY, MySQLCreateTableParserFOREIGN, MySQLCreateTableParserINDEX, MySQLCreateTableParserUNIQUE, MySQLCreateTableParserCONSTRAINT, MySQLCreateTableParserCHECK:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)
			p.Table_constraint()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumn_defContext is an interface to support dynamic dispatch.
type IColumn_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	Data_type() IData_typeContext
	NOT() antlr.TerminalNode
	NULL() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	AUTO_INCREMENT() antlr.TerminalNode
	PRIMARY() antlr.TerminalNode
	KEY() antlr.TerminalNode
	UNIQUE() antlr.TerminalNode
	COMMENT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode
	SET() antlr.TerminalNode
	Charset_name() ICharset_nameContext
	COLLATE() antlr.TerminalNode
	Collation_name() ICollation_nameContext
	REFERENCES() antlr.TerminalNode
	Table_name() ITable_nameContext
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	CHECK() antlr.TerminalNode
	Expr() IExprContext
	Literal() ILiteralContext
	AllCURRENT_TIMESTAMP() []antlr.TerminalNode
	CURRENT_TIMESTAMP(i int) antlr.TerminalNode
	AllON() []antlr.TerminalNode
	ON(i int) antlr.TerminalNode
	AllAction_() []IActionContext
	Action_(i int) IActionContext
	AllDELETE() []antlr.TerminalNode
	DELETE(i int) antlr.TerminalNode
	AllUPDATE() []antlr.TerminalNode
	UPDATE(i int) antlr.TerminalNode

	// IsColumn_defContext differentiates from other interfaces.
	IsColumn_defContext()
}

type Column_defContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_defContext() *Column_defContext {
	var p = new(Column_defContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_column_def
	return p
}

func InitEmptyColumn_defContext(p *Column_defContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_column_def
}

func (*Column_defContext) IsColumn_defContext() {}

func NewColumn_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_defContext {
	var p = new(Column_defContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_column_def

	return p
}

func (s *Column_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_defContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Column_defContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_defContext) Data_type() IData_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IData_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IData_typeContext)
}

func (s *Column_defContext) NOT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNOT, 0)
}

func (s *Column_defContext) NULL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNULL, 0)
}

func (s *Column_defContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDEFAULT, 0)
}

func (s *Column_defContext) AUTO_INCREMENT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserAUTO_INCREMENT, 0)
}

func (s *Column_defContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserPRIMARY, 0)
}

func (s *Column_defContext) KEY() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserKEY, 0)
}

func (s *Column_defContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUNIQUE, 0)
}

func (s *Column_defContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOMMENT, 0)
}

func (s *Column_defContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSTRING_LITERAL, 0)
}

func (s *Column_defContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCHARACTER, 0)
}

func (s *Column_defContext) SET() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSET, 0)
}

func (s *Column_defContext) Charset_name() ICharset_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharset_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharset_nameContext)
}

func (s *Column_defContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOLLATE, 0)
}

func (s *Column_defContext) Collation_name() ICollation_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollation_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollation_nameContext)
}

func (s *Column_defContext) REFERENCES() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserREFERENCES, 0)
}

func (s *Column_defContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Column_defContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserLPAREN)
}

func (s *Column_defContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, i)
}

func (s *Column_defContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserRPAREN)
}

func (s *Column_defContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, i)
}

func (s *Column_defContext) CHECK() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCHECK, 0)
}

func (s *Column_defContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Column_defContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Column_defContext) AllCURRENT_TIMESTAMP() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserCURRENT_TIMESTAMP)
}

func (s *Column_defContext) CURRENT_TIMESTAMP(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCURRENT_TIMESTAMP, i)
}

func (s *Column_defContext) AllON() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserON)
}

func (s *Column_defContext) ON(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserON, i)
}

func (s *Column_defContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *Column_defContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *Column_defContext) AllDELETE() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserDELETE)
}

func (s *Column_defContext) DELETE(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDELETE, i)
}

func (s *Column_defContext) AllUPDATE() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserUPDATE)
}

func (s *Column_defContext) UPDATE(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUPDATE, i)
}

func (s *Column_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterColumn_def(s)
	}
}

func (s *Column_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitColumn_def(s)
	}
}

func (p *MySQLCreateTableParser) Column_def() (localctx IColumn_defContext) {
	localctx = NewColumn_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MySQLCreateTableParserRULE_column_def)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Column_name()
	}
	{
		p.SetState(77)
		p.Data_type()
	}
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserNOT:
		{
			p.SetState(78)
			p.Match(MySQLCreateTableParserNOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(79)
			p.Match(MySQLCreateTableParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserNULL:
		{
			p.SetState(80)
			p.Match(MySQLCreateTableParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserPRIMARY, MySQLCreateTableParserREFERENCES, MySQLCreateTableParserUNIQUE, MySQLCreateTableParserCHECK, MySQLCreateTableParserDEFAULT, MySQLCreateTableParserAUTO_INCREMENT, MySQLCreateTableParserCOLLATE, MySQLCreateTableParserCOMMENT, MySQLCreateTableParserCHARACTER, MySQLCreateTableParserCOMMA, MySQLCreateTableParserRPAREN:

	default:
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserDEFAULT {
		{
			p.SetState(83)
			p.Match(MySQLCreateTableParserDEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(84)
				p.Literal()
			}

		case 2:
			{
				p.SetState(85)
				p.Match(MySQLCreateTableParserCURRENT_TIMESTAMP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MySQLCreateTableParserON {
				{
					p.SetState(86)
					p.Match(MySQLCreateTableParserON)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(87)
					p.Match(MySQLCreateTableParserUPDATE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(88)
					p.Match(MySQLCreateTableParserCURRENT_TIMESTAMP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserAUTO_INCREMENT {
		{
			p.SetState(95)
			p.Match(MySQLCreateTableParserAUTO_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserPRIMARY:
		{
			p.SetState(98)
			p.Match(MySQLCreateTableParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Match(MySQLCreateTableParserKEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserUNIQUE:
		{
			p.SetState(100)
			p.Match(MySQLCreateTableParserUNIQUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserKEY {
			{
				p.SetState(101)
				p.Match(MySQLCreateTableParserKEY)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MySQLCreateTableParserREFERENCES, MySQLCreateTableParserCHECK, MySQLCreateTableParserCOLLATE, MySQLCreateTableParserCOMMENT, MySQLCreateTableParserCHARACTER, MySQLCreateTableParserCOMMA, MySQLCreateTableParserRPAREN:

	default:
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserCOMMENT {
		{
			p.SetState(106)
			p.Match(MySQLCreateTableParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.Match(MySQLCreateTableParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserCHARACTER {
		{
			p.SetState(110)
			p.Match(MySQLCreateTableParserCHARACTER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Match(MySQLCreateTableParserSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Charset_name()
		}

	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserCOLLATE {
		{
			p.SetState(115)
			p.Match(MySQLCreateTableParserCOLLATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Collation_name()
		}

	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserREFERENCES {
		{
			p.SetState(119)
			p.Match(MySQLCreateTableParserREFERENCES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Table_name()
		}
		{
			p.SetState(121)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.Column_name()
		}
		{
			p.SetState(123)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MySQLCreateTableParserON {
			{
				p.SetState(124)
				p.Match(MySQLCreateTableParserON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(125)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLCreateTableParserUPDATE || _la == MySQLCreateTableParserDELETE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(126)
				p.Action_()
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserCHECK {
		{
			p.SetState(134)
			p.Match(MySQLCreateTableParserCHECK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Expr()
		}
		{
			p.SetState(137)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IData_typeContext is an interface to support dynamic dispatch.
type IData_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsData_typeContext differentiates from other interfaces.
	IsData_typeContext()
}

type Data_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_typeContext() *Data_typeContext {
	var p = new(Data_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_data_type
	return p
}

func InitEmptyData_typeContext(p *Data_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_data_type
}

func (*Data_typeContext) IsData_typeContext() {}

func NewData_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_typeContext {
	var p = new(Data_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_data_type

	return p
}

func (s *Data_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_typeContext) CopyAll(ctx *Data_typeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Data_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntegerTypeContext struct {
	Data_typeContext
}

func NewIntegerTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerTypeContext {
	var p = new(IntegerTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *IntegerTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserINT, 0)
}

func (s *IntegerTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *IntegerTypeContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *IntegerTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *IntegerTypeContext) UNSIGNED() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUNSIGNED, 0)
}

func (s *IntegerTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterIntegerType(s)
	}
}

func (s *IntegerTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitIntegerType(s)
	}
}

type JsonTypeContext struct {
	Data_typeContext
}

func NewJsonTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonTypeContext {
	var p = new(JsonTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *JsonTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonTypeContext) JSON() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserJSON, 0)
}

func (s *JsonTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterJsonType(s)
	}
}

func (s *JsonTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitJsonType(s)
	}
}

type VarcharTypeContext struct {
	Data_typeContext
}

func NewVarcharTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarcharTypeContext {
	var p = new(VarcharTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *VarcharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarcharTypeContext) VARCHAR() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserVARCHAR, 0)
}

func (s *VarcharTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *VarcharTypeContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *VarcharTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *VarcharTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterVarcharType(s)
	}
}

func (s *VarcharTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitVarcharType(s)
	}
}

type TinyintTypeContext struct {
	Data_typeContext
}

func NewTinyintTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TinyintTypeContext {
	var p = new(TinyintTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *TinyintTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TinyintTypeContext) TINYINT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserTINYINT, 0)
}

func (s *TinyintTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *TinyintTypeContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *TinyintTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *TinyintTypeContext) UNSIGNED() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUNSIGNED, 0)
}

func (s *TinyintTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTinyintType(s)
	}
}

func (s *TinyintTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTinyintType(s)
	}
}

type DecimalTypeContext struct {
	Data_typeContext
}

func NewDecimalTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DecimalTypeContext {
	var p = new(DecimalTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *DecimalTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalTypeContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDECIMAL, 0)
}

func (s *DecimalTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *DecimalTypeContext) AllNUMERIC_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserNUMERIC_LITERAL)
}

func (s *DecimalTypeContext) NUMERIC_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, i)
}

func (s *DecimalTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *DecimalTypeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOMMA, 0)
}

func (s *DecimalTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterDecimalType(s)
	}
}

func (s *DecimalTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitDecimalType(s)
	}
}

type DatetimeTypeContext struct {
	Data_typeContext
}

func NewDatetimeTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatetimeTypeContext {
	var p = new(DatetimeTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *DatetimeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeTypeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDATETIME, 0)
}

func (s *DatetimeTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterDatetimeType(s)
	}
}

func (s *DatetimeTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitDatetimeType(s)
	}
}

type TimestampTypeContext struct {
	Data_typeContext
}

func NewTimestampTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimestampTypeContext {
	var p = new(TimestampTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *TimestampTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimestampTypeContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserTIMESTAMP, 0)
}

func (s *TimestampTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTimestampType(s)
	}
}

func (s *TimestampTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTimestampType(s)
	}
}

type CharTypeContext struct {
	Data_typeContext
}

func NewCharTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharTypeContext {
	var p = new(CharTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *CharTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharTypeContext) CHAR() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCHAR, 0)
}

func (s *CharTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *CharTypeContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *CharTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *CharTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterCharType(s)
	}
}

func (s *CharTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitCharType(s)
	}
}

type DateTypeContext struct {
	Data_typeContext
}

func NewDateTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateTypeContext {
	var p = new(DateTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *DateTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateTypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDATE, 0)
}

func (s *DateTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterDateType(s)
	}
}

func (s *DateTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitDateType(s)
	}
}

type BigintTypeContext struct {
	Data_typeContext
}

func NewBigintTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BigintTypeContext {
	var p = new(BigintTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *BigintTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BigintTypeContext) BIGINT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserBIGINT, 0)
}

func (s *BigintTypeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *BigintTypeContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *BigintTypeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *BigintTypeContext) UNSIGNED() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUNSIGNED, 0)
}

func (s *BigintTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterBigintType(s)
	}
}

func (s *BigintTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitBigintType(s)
	}
}

type TextTypeContext struct {
	Data_typeContext
}

func NewTextTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TextTypeContext {
	var p = new(TextTypeContext)

	InitEmptyData_typeContext(&p.Data_typeContext)
	p.parser = parser
	p.CopyAll(ctx.(*Data_typeContext))

	return p
}

func (s *TextTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextTypeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserTEXT, 0)
}

func (s *TextTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTextType(s)
	}
}

func (s *TextTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTextType(s)
	}
}

func (p *MySQLCreateTableParser) Data_type() (localctx IData_typeContext) {
	localctx = NewData_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MySQLCreateTableParserRULE_data_type)
	var _la int

	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserINT:
		localctx = NewIntegerTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(MySQLCreateTableParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserLPAREN {
			{
				p.SetState(142)
				p.Match(MySQLCreateTableParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(143)
				p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(144)
				p.Match(MySQLCreateTableParserRPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserUNSIGNED {
			{
				p.SetState(147)
				p.Match(MySQLCreateTableParserUNSIGNED)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MySQLCreateTableParserBIGINT:
		localctx = NewBigintTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.Match(MySQLCreateTableParserBIGINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserLPAREN {
			{
				p.SetState(151)
				p.Match(MySQLCreateTableParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(152)
				p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(153)
				p.Match(MySQLCreateTableParserRPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserUNSIGNED {
			{
				p.SetState(156)
				p.Match(MySQLCreateTableParserUNSIGNED)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MySQLCreateTableParserTINYINT:
		localctx = NewTinyintTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Match(MySQLCreateTableParserTINYINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserLPAREN {
			{
				p.SetState(160)
				p.Match(MySQLCreateTableParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(161)
				p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(162)
				p.Match(MySQLCreateTableParserRPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserUNSIGNED {
			{
				p.SetState(165)
				p.Match(MySQLCreateTableParserUNSIGNED)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MySQLCreateTableParserVARCHAR:
		localctx = NewVarcharTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(168)
			p.Match(MySQLCreateTableParserVARCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
			p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserCHAR:
		localctx = NewCharTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(172)
			p.Match(MySQLCreateTableParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserDECIMAL:
		localctx = NewDecimalTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(176)
			p.Match(MySQLCreateTableParserDECIMAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserLPAREN {
			{
				p.SetState(177)
				p.Match(MySQLCreateTableParserLPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(178)
				p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(181)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == MySQLCreateTableParserCOMMA {
				{
					p.SetState(179)
					p.Match(MySQLCreateTableParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(180)
					p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			{
				p.SetState(183)
				p.Match(MySQLCreateTableParserRPAREN)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case MySQLCreateTableParserDATETIME:
		localctx = NewDatetimeTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(186)
			p.Match(MySQLCreateTableParserDATETIME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserDATE:
		localctx = NewDateTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(187)
			p.Match(MySQLCreateTableParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserTIMESTAMP:
		localctx = NewTimestampTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(188)
			p.Match(MySQLCreateTableParserTIMESTAMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserTEXT:
		localctx = NewTextTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(189)
			p.Match(MySQLCreateTableParserTEXT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserJSON:
		localctx = NewJsonTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(190)
			p.Match(MySQLCreateTableParserJSON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITable_constraintContext is an interface to support dynamic dispatch.
type ITable_constraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY() antlr.TerminalNode
	KEY() antlr.TerminalNode
	AllLPAREN() []antlr.TerminalNode
	LPAREN(i int) antlr.TerminalNode
	AllColumn_list() []IColumn_listContext
	Column_list(i int) IColumn_listContext
	AllRPAREN() []antlr.TerminalNode
	RPAREN(i int) antlr.TerminalNode
	FOREIGN() antlr.TerminalNode
	REFERENCES() antlr.TerminalNode
	Table_name() ITable_nameContext
	CHECK() antlr.TerminalNode
	Expr() IExprContext
	UNIQUE() antlr.TerminalNode
	CONSTRAINT() antlr.TerminalNode
	AllIdentifier() []IIdentifierContext
	Identifier(i int) IIdentifierContext
	INDEX() antlr.TerminalNode
	AllON() []antlr.TerminalNode
	ON(i int) antlr.TerminalNode
	AllAction_() []IActionContext
	Action_(i int) IActionContext
	USING() antlr.TerminalNode
	Index_type() IIndex_typeContext
	AllDELETE() []antlr.TerminalNode
	DELETE(i int) antlr.TerminalNode
	AllUPDATE() []antlr.TerminalNode
	UPDATE(i int) antlr.TerminalNode

	// IsTable_constraintContext differentiates from other interfaces.
	IsTable_constraintContext()
}

type Table_constraintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_constraintContext() *Table_constraintContext {
	var p = new(Table_constraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_constraint
	return p
}

func InitEmptyTable_constraintContext(p *Table_constraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_constraint
}

func (*Table_constraintContext) IsTable_constraintContext() {}

func NewTable_constraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_constraintContext {
	var p = new(Table_constraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_table_constraint

	return p
}

func (s *Table_constraintContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_constraintContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserPRIMARY, 0)
}

func (s *Table_constraintContext) KEY() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserKEY, 0)
}

func (s *Table_constraintContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserLPAREN)
}

func (s *Table_constraintContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, i)
}

func (s *Table_constraintContext) AllColumn_list() []IColumn_listContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_listContext); ok {
			len++
		}
	}

	tst := make([]IColumn_listContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_listContext); ok {
			tst[i] = t.(IColumn_listContext)
			i++
		}
	}

	return tst
}

func (s *Table_constraintContext) Column_list(i int) IColumn_listContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_listContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_listContext)
}

func (s *Table_constraintContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserRPAREN)
}

func (s *Table_constraintContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, i)
}

func (s *Table_constraintContext) FOREIGN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserFOREIGN, 0)
}

func (s *Table_constraintContext) REFERENCES() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserREFERENCES, 0)
}

func (s *Table_constraintContext) Table_name() ITable_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITable_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITable_nameContext)
}

func (s *Table_constraintContext) CHECK() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCHECK, 0)
}

func (s *Table_constraintContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Table_constraintContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUNIQUE, 0)
}

func (s *Table_constraintContext) CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCONSTRAINT, 0)
}

func (s *Table_constraintContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *Table_constraintContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Table_constraintContext) INDEX() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserINDEX, 0)
}

func (s *Table_constraintContext) AllON() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserON)
}

func (s *Table_constraintContext) ON(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserON, i)
}

func (s *Table_constraintContext) AllAction_() []IActionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IActionContext); ok {
			len++
		}
	}

	tst := make([]IActionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IActionContext); ok {
			tst[i] = t.(IActionContext)
			i++
		}
	}

	return tst
}

func (s *Table_constraintContext) Action_(i int) IActionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IActionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IActionContext)
}

func (s *Table_constraintContext) USING() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUSING, 0)
}

func (s *Table_constraintContext) Index_type() IIndex_typeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndex_typeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndex_typeContext)
}

func (s *Table_constraintContext) AllDELETE() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserDELETE)
}

func (s *Table_constraintContext) DELETE(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDELETE, i)
}

func (s *Table_constraintContext) AllUPDATE() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserUPDATE)
}

func (s *Table_constraintContext) UPDATE(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserUPDATE, i)
}

func (s *Table_constraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_constraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_constraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTable_constraint(s)
	}
}

func (s *Table_constraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTable_constraint(s)
	}
}

func (p *MySQLCreateTableParser) Table_constraint() (localctx ITable_constraintContext) {
	localctx = NewTable_constraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MySQLCreateTableParserRULE_table_constraint)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserCONSTRAINT {
		{
			p.SetState(193)
			p.Match(MySQLCreateTableParserCONSTRAINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(194)
			p.Identifier()
		}

	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserPRIMARY:
		{
			p.SetState(197)
			p.Match(MySQLCreateTableParserPRIMARY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.Match(MySQLCreateTableParserKEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(200)
			p.Column_list()
		}
		{
			p.SetState(201)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserFOREIGN:
		{
			p.SetState(203)
			p.Match(MySQLCreateTableParserFOREIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(MySQLCreateTableParserKEY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(206)
			p.Column_list()
		}
		{
			p.SetState(207)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(208)
			p.Match(MySQLCreateTableParserREFERENCES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.Table_name()
		}
		{
			p.SetState(210)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Column_list()
		}
		{
			p.SetState(212)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MySQLCreateTableParserON {
			{
				p.SetState(213)
				p.Match(MySQLCreateTableParserON)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(214)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLCreateTableParserUPDATE || _la == MySQLCreateTableParserDELETE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(215)
				p.Action_()
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case MySQLCreateTableParserCHECK:
		{
			p.SetState(221)
			p.Match(MySQLCreateTableParserCHECK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(223)
			p.Expr()
		}
		{
			p.SetState(224)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserKEY, MySQLCreateTableParserINDEX:
		{
			p.SetState(226)
			_la = p.GetTokenStream().LA(1)

			if !(_la == MySQLCreateTableParserKEY || _la == MySQLCreateTableParserINDEX) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserIDENTIFIER {
			{
				p.SetState(227)
				p.Identifier()
			}

		}
		{
			p.SetState(230)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(231)
			p.Column_list()
		}
		{
			p.SetState(232)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserUSING {
			{
				p.SetState(233)
				p.Match(MySQLCreateTableParserUSING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(234)
				p.Index_type()
			}

		}

	case MySQLCreateTableParserUNIQUE:
		{
			p.SetState(237)
			p.Match(MySQLCreateTableParserUNIQUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserKEY || _la == MySQLCreateTableParserINDEX {
			{
				p.SetState(238)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLCreateTableParserKEY || _la == MySQLCreateTableParserINDEX) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserIDENTIFIER {
			{
				p.SetState(241)
				p.Identifier()
			}

		}
		{
			p.SetState(244)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(245)
			p.Column_list()
		}
		{
			p.SetState(246)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserUSING {
			{
				p.SetState(247)
				p.Match(MySQLCreateTableParserUSING)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(248)
				p.Index_type()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IActionContext is an interface to support dynamic dispatch.
type IActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESTRICT() antlr.TerminalNode
	CASCADE() antlr.TerminalNode
	SET() antlr.TerminalNode
	NULL() antlr.TerminalNode
	NO() antlr.TerminalNode
	ACTION() antlr.TerminalNode

	// IsActionContext differentiates from other interfaces.
	IsActionContext()
}

type ActionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyActionContext() *ActionContext {
	var p = new(ActionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_action
	return p
}

func InitEmptyActionContext(p *ActionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_action
}

func (*ActionContext) IsActionContext() {}

func NewActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ActionContext {
	var p = new(ActionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_action

	return p
}

func (s *ActionContext) GetParser() antlr.Parser { return s.parser }

func (s *ActionContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRESTRICT, 0)
}

func (s *ActionContext) CASCADE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCASCADE, 0)
}

func (s *ActionContext) SET() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSET, 0)
}

func (s *ActionContext) NULL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNULL, 0)
}

func (s *ActionContext) NO() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNO, 0)
}

func (s *ActionContext) ACTION() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserACTION, 0)
}

func (s *ActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterAction(s)
	}
}

func (s *ActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitAction(s)
	}
}

func (p *MySQLCreateTableParser) Action_() (localctx IActionContext) {
	localctx = NewActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MySQLCreateTableParserRULE_action)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserRESTRICT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.Match(MySQLCreateTableParserRESTRICT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserCASCADE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(MySQLCreateTableParserCASCADE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserSET:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Match(MySQLCreateTableParserSET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.Match(MySQLCreateTableParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case MySQLCreateTableParserNO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(257)
			p.Match(MySQLCreateTableParserNO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(258)
			p.Match(MySQLCreateTableParserACTION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITable_optionContext is an interface to support dynamic dispatch.
type ITable_optionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENGINE() antlr.TerminalNode
	OPERATOR() antlr.TerminalNode
	Engine_name() IEngine_nameContext
	Charset_name() ICharset_nameContext
	CHARSET() antlr.TerminalNode
	CHARACTER() antlr.TerminalNode
	SET() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	COLLATE() antlr.TerminalNode
	Collation_name() ICollation_nameContext
	AUTO_INCREMENT() antlr.TerminalNode
	NUMERIC_LITERAL() antlr.TerminalNode
	COMMENT() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode

	// IsTable_optionContext differentiates from other interfaces.
	IsTable_optionContext()
}

type Table_optionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_optionContext() *Table_optionContext {
	var p = new(Table_optionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_option
	return p
}

func InitEmptyTable_optionContext(p *Table_optionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_option
}

func (*Table_optionContext) IsTable_optionContext() {}

func NewTable_optionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_optionContext {
	var p = new(Table_optionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_table_option

	return p
}

func (s *Table_optionContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_optionContext) ENGINE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserENGINE, 0)
}

func (s *Table_optionContext) OPERATOR() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserOPERATOR, 0)
}

func (s *Table_optionContext) Engine_name() IEngine_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEngine_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEngine_nameContext)
}

func (s *Table_optionContext) Charset_name() ICharset_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICharset_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICharset_nameContext)
}

func (s *Table_optionContext) CHARSET() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCHARSET, 0)
}

func (s *Table_optionContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCHARACTER, 0)
}

func (s *Table_optionContext) SET() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSET, 0)
}

func (s *Table_optionContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDEFAULT, 0)
}

func (s *Table_optionContext) COLLATE() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOLLATE, 0)
}

func (s *Table_optionContext) Collation_name() ICollation_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollation_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollation_nameContext)
}

func (s *Table_optionContext) AUTO_INCREMENT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserAUTO_INCREMENT, 0)
}

func (s *Table_optionContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *Table_optionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOMMENT, 0)
}

func (s *Table_optionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSTRING_LITERAL, 0)
}

func (s *Table_optionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_optionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_optionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTable_option(s)
	}
}

func (s *Table_optionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTable_option(s)
	}
}

func (p *MySQLCreateTableParser) Table_option() (localctx ITable_optionContext) {
	localctx = NewTable_optionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MySQLCreateTableParserRULE_table_option)
	var _la int

	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(261)
			p.Match(MySQLCreateTableParserENGINE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.Match(MySQLCreateTableParserOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(263)
			p.Engine_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserDEFAULT {
			{
				p.SetState(264)
				p.Match(MySQLCreateTableParserDEFAULT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case MySQLCreateTableParserCHARSET:
			{
				p.SetState(267)
				p.Match(MySQLCreateTableParserCHARSET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case MySQLCreateTableParserCHARACTER:
			{
				p.SetState(268)
				p.Match(MySQLCreateTableParserCHARACTER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(269)
				p.Match(MySQLCreateTableParserSET)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		{
			p.SetState(272)
			p.Match(MySQLCreateTableParserOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Charset_name()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == MySQLCreateTableParserDEFAULT {
			{
				p.SetState(274)
				p.Match(MySQLCreateTableParserDEFAULT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(277)
			p.Match(MySQLCreateTableParserCOLLATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(278)
			p.Match(MySQLCreateTableParserOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Collation_name()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(280)
			p.Match(MySQLCreateTableParserAUTO_INCREMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(281)
			p.Match(MySQLCreateTableParserOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(282)
			p.Match(MySQLCreateTableParserNUMERIC_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(283)
			p.Match(MySQLCreateTableParserCOMMENT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Match(MySQLCreateTableParserOPERATOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(285)
			p.Match(MySQLCreateTableParserSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITable_nameContext is an interface to support dynamic dispatch.
type ITable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsTable_nameContext differentiates from other interfaces.
	IsTable_nameContext()
}

type Table_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_nameContext() *Table_nameContext {
	var p = new(Table_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_name
	return p
}

func InitEmptyTable_nameContext(p *Table_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_table_name
}

func (*Table_nameContext) IsTable_nameContext() {}

func NewTable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_nameContext {
	var p = new(Table_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_table_name

	return p
}

func (s *Table_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_nameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserIDENTIFIER)
}

func (s *Table_nameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, i)
}

func (s *Table_nameContext) DOT() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserDOT, 0)
}

func (s *Table_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterTable_name(s)
	}
}

func (s *Table_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitTable_name(s)
	}
}

func (p *MySQLCreateTableParser) Table_name() (localctx ITable_nameContext) {
	localctx = NewTable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MySQLCreateTableParserRULE_table_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == MySQLCreateTableParserDOT {
		{
			p.SetState(289)
			p.Match(MySQLCreateTableParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(290)
			p.Match(MySQLCreateTableParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumn_nameContext is an interface to support dynamic dispatch.
type IColumn_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsColumn_nameContext differentiates from other interfaces.
	IsColumn_nameContext()
}

type Column_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_nameContext() *Column_nameContext {
	var p = new(Column_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_column_name
	return p
}

func InitEmptyColumn_nameContext(p *Column_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_column_name
}

func (*Column_nameContext) IsColumn_nameContext() {}

func NewColumn_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_nameContext {
	var p = new(Column_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_column_name

	return p
}

func (s *Column_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *Column_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterColumn_name(s)
	}
}

func (s *Column_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitColumn_name(s)
	}
}

func (p *MySQLCreateTableParser) Column_name() (localctx IColumn_nameContext) {
	localctx = NewColumn_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MySQLCreateTableParserRULE_column_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IColumn_listContext is an interface to support dynamic dispatch.
type IColumn_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllColumn_name() []IColumn_nameContext
	Column_name(i int) IColumn_nameContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsColumn_listContext differentiates from other interfaces.
	IsColumn_listContext()
}

type Column_listContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_listContext() *Column_listContext {
	var p = new(Column_listContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_column_list
	return p
}

func InitEmptyColumn_listContext(p *Column_listContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_column_list
}

func (*Column_listContext) IsColumn_listContext() {}

func NewColumn_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_listContext {
	var p = new(Column_listContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_column_list

	return p
}

func (s *Column_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_listContext) AllColumn_name() []IColumn_nameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumn_nameContext); ok {
			len++
		}
	}

	tst := make([]IColumn_nameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumn_nameContext); ok {
			tst[i] = t.(IColumn_nameContext)
			i++
		}
	}

	return tst
}

func (s *Column_listContext) Column_name(i int) IColumn_nameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *Column_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserCOMMA)
}

func (s *Column_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOMMA, i)
}

func (s *Column_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterColumn_list(s)
	}
}

func (s *Column_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitColumn_list(s)
	}
}

func (p *MySQLCreateTableParser) Column_list() (localctx IColumn_listContext) {
	localctx = NewColumn_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MySQLCreateTableParserRULE_column_list)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Column_name()
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == MySQLCreateTableParserCOMMA {
		{
			p.SetState(296)
			p.Match(MySQLCreateTableParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(297)
			p.Column_name()
		}

		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEngine_nameContext is an interface to support dynamic dispatch.
type IEngine_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsEngine_nameContext differentiates from other interfaces.
	IsEngine_nameContext()
}

type Engine_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEngine_nameContext() *Engine_nameContext {
	var p = new(Engine_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_engine_name
	return p
}

func InitEmptyEngine_nameContext(p *Engine_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_engine_name
}

func (*Engine_nameContext) IsEngine_nameContext() {}

func NewEngine_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Engine_nameContext {
	var p = new(Engine_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_engine_name

	return p
}

func (s *Engine_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Engine_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *Engine_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Engine_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Engine_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterEngine_name(s)
	}
}

func (s *Engine_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitEngine_name(s)
	}
}

func (p *MySQLCreateTableParser) Engine_name() (localctx IEngine_nameContext) {
	localctx = NewEngine_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MySQLCreateTableParserRULE_engine_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICharset_nameContext is an interface to support dynamic dispatch.
type ICharset_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsCharset_nameContext differentiates from other interfaces.
	IsCharset_nameContext()
}

type Charset_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharset_nameContext() *Charset_nameContext {
	var p = new(Charset_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_charset_name
	return p
}

func InitEmptyCharset_nameContext(p *Charset_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_charset_name
}

func (*Charset_nameContext) IsCharset_nameContext() {}

func NewCharset_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Charset_nameContext {
	var p = new(Charset_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_charset_name

	return p
}

func (s *Charset_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Charset_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *Charset_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Charset_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Charset_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterCharset_name(s)
	}
}

func (s *Charset_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitCharset_name(s)
	}
}

func (p *MySQLCreateTableParser) Charset_name() (localctx ICharset_nameContext) {
	localctx = NewCharset_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MySQLCreateTableParserRULE_charset_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollation_nameContext is an interface to support dynamic dispatch.
type ICollation_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsCollation_nameContext differentiates from other interfaces.
	IsCollation_nameContext()
}

type Collation_nameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollation_nameContext() *Collation_nameContext {
	var p = new(Collation_nameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_collation_name
	return p
}

func InitEmptyCollation_nameContext(p *Collation_nameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_collation_name
}

func (*Collation_nameContext) IsCollation_nameContext() {}

func NewCollation_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Collation_nameContext {
	var p = new(Collation_nameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_collation_name

	return p
}

func (s *Collation_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Collation_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *Collation_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Collation_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Collation_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterCollation_name(s)
	}
}

func (s *Collation_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitCollation_name(s)
	}
}

func (p *MySQLCreateTableParser) Collation_name() (localctx ICollation_nameContext) {
	localctx = NewCollation_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MySQLCreateTableParserRULE_collation_name)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	NUMERIC_LITERAL() antlr.TerminalNode
	NULL() antlr.TerminalNode
	CURRENT_TIMESTAMP() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) NUMERIC_LITERAL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNUMERIC_LITERAL, 0)
}

func (s *LiteralContext) NULL() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserNULL, 0)
}

func (s *LiteralContext) CURRENT_TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCURRENT_TIMESTAMP, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *MySQLCreateTableParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MySQLCreateTableParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&54043195534737408) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *MySQLCreateTableParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, MySQLCreateTableParserRULE_identifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIndex_typeContext is an interface to support dynamic dispatch.
type IIndex_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsIndex_typeContext differentiates from other interfaces.
	IsIndex_typeContext()
}

type Index_typeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndex_typeContext() *Index_typeContext {
	var p = new(Index_typeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_index_type
	return p
}

func InitEmptyIndex_typeContext(p *Index_typeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_index_type
}

func (*Index_typeContext) IsIndex_typeContext() {}

func NewIndex_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Index_typeContext {
	var p = new(Index_typeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_index_type

	return p
}

func (s *Index_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Index_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *Index_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Index_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Index_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterIndex_type(s)
	}
}

func (s *Index_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitIndex_type(s)
	}
}

func (p *MySQLCreateTableParser) Index_type() (localctx IIndex_typeContext) {
	localctx = NewIndex_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, MySQLCreateTableParserRULE_index_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Logical_expr() ILogical_exprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Logical_expr() ILogical_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *MySQLCreateTableParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, MySQLCreateTableParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.logical_expr(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILogical_exprContext is an interface to support dynamic dispatch.
type ILogical_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLogical_exprContext differentiates from other interfaces.
	IsLogical_exprContext()
}

type Logical_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_exprContext() *Logical_exprContext {
	var p = new(Logical_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_logical_expr
	return p
}

func InitEmptyLogical_exprContext(p *Logical_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_logical_expr
}

func (*Logical_exprContext) IsLogical_exprContext() {}

func NewLogical_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_exprContext {
	var p = new(Logical_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_logical_expr

	return p
}

func (s *Logical_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_exprContext) CopyAll(ctx *Logical_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Logical_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SingleComparisonContext struct {
	Logical_exprContext
}

func NewSingleComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SingleComparisonContext {
	var p = new(SingleComparisonContext)

	InitEmptyLogical_exprContext(&p.Logical_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Logical_exprContext))

	return p
}

func (s *SingleComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleComparisonContext) Comparison_expr() IComparison_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_exprContext)
}

func (s *SingleComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterSingleComparison(s)
	}
}

func (s *SingleComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitSingleComparison(s)
	}
}

type LogicalOpContext struct {
	Logical_exprContext
}

func NewLogicalOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOpContext {
	var p = new(LogicalOpContext)

	InitEmptyLogical_exprContext(&p.Logical_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Logical_exprContext))

	return p
}

func (s *LogicalOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOpContext) Logical_expr() ILogical_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogical_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogical_exprContext)
}

func (s *LogicalOpContext) Comparison_expr() IComparison_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparison_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparison_exprContext)
}

func (s *LogicalOpContext) AND() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserAND, 0)
}

func (s *LogicalOpContext) OR() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserOR, 0)
}

func (s *LogicalOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterLogicalOp(s)
	}
}

func (s *LogicalOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitLogicalOp(s)
	}
}

func (p *MySQLCreateTableParser) Logical_expr() (localctx ILogical_exprContext) {
	return p.logical_expr(0)
}

func (p *MySQLCreateTableParser) logical_expr(_p int) (localctx ILogical_exprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewLogical_exprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_exprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, MySQLCreateTableParserRULE_logical_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewSingleComparisonContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(318)
		p.Comparison_expr()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalOpContext(p, NewLogical_exprContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, MySQLCreateTableParserRULE_logical_expr)
			p.SetState(320)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(321)
				_la = p.GetTokenStream().LA(1)

				if !(_la == MySQLCreateTableParserAND || _la == MySQLCreateTableParserOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(322)
				p.Comparison_expr()
			}

		}
		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IComparison_exprContext is an interface to support dynamic dispatch.
type IComparison_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsComparison_exprContext differentiates from other interfaces.
	IsComparison_exprContext()
}

type Comparison_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_exprContext() *Comparison_exprContext {
	var p = new(Comparison_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_comparison_expr
	return p
}

func InitEmptyComparison_exprContext(p *Comparison_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_comparison_expr
}

func (*Comparison_exprContext) IsComparison_exprContext() {}

func NewComparison_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_exprContext {
	var p = new(Comparison_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_comparison_expr

	return p
}

func (s *Comparison_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparison_exprContext) CopyAll(ctx *Comparison_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Comparison_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ComparisonOpContext struct {
	Comparison_exprContext
}

func NewComparisonOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonOpContext {
	var p = new(ComparisonOpContext)

	InitEmptyComparison_exprContext(&p.Comparison_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Comparison_exprContext))

	return p
}

func (s *ComparisonOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOpContext) AllPrimary_expr() []IPrimary_exprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			len++
		}
	}

	tst := make([]IPrimary_exprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrimary_exprContext); ok {
			tst[i] = t.(IPrimary_exprContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonOpContext) Primary_expr(i int) IPrimary_exprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimary_exprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimary_exprContext)
}

func (s *ComparisonOpContext) OPERATOR() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserOPERATOR, 0)
}

func (s *ComparisonOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterComparisonOp(s)
	}
}

func (s *ComparisonOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitComparisonOp(s)
	}
}

type ParenExprContext struct {
	Comparison_exprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyComparison_exprContext(&p.Comparison_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Comparison_exprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (p *MySQLCreateTableParser) Comparison_expr() (localctx IComparison_exprContext) {
	localctx = NewComparison_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, MySQLCreateTableParserRULE_comparison_expr)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case MySQLCreateTableParserNULL, MySQLCreateTableParserCURRENT_TIMESTAMP, MySQLCreateTableParserIDENTIFIER, MySQLCreateTableParserSTRING_LITERAL, MySQLCreateTableParserNUMERIC_LITERAL:
		localctx = NewComparisonOpContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Primary_expr()
		}
		p.SetState(331)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(329)
				p.Match(MySQLCreateTableParserOPERATOR)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(330)
				p.Primary_expr()
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case MySQLCreateTableParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Match(MySQLCreateTableParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Expr()
		}
		{
			p.SetState(335)
			p.Match(MySQLCreateTableParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimary_exprContext is an interface to support dynamic dispatch.
type IPrimary_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrimary_exprContext differentiates from other interfaces.
	IsPrimary_exprContext()
}

type Primary_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_exprContext() *Primary_exprContext {
	var p = new(Primary_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_primary_expr
	return p
}

func InitEmptyPrimary_exprContext(p *Primary_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_primary_expr
}

func (*Primary_exprContext) IsPrimary_exprContext() {}

func NewPrimary_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_exprContext {
	var p = new(Primary_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_primary_expr

	return p
}

func (s *Primary_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_exprContext) CopyAll(ctx *Primary_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Primary_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctionExprContext struct {
	Primary_exprContext
}

func NewFunctionExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExprContext {
	var p = new(FunctionExprContext)

	InitEmptyPrimary_exprContext(&p.Primary_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_exprContext))

	return p
}

func (s *FunctionExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExprContext) Function_call() IFunction_callContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunction_callContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *FunctionExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterFunctionExpr(s)
	}
}

func (s *FunctionExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitFunctionExpr(s)
	}
}

type LiteralValueContext struct {
	Primary_exprContext
}

func NewLiteralValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralValueContext {
	var p = new(LiteralValueContext)

	InitEmptyPrimary_exprContext(&p.Primary_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_exprContext))

	return p
}

func (s *LiteralValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralValueContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterLiteralValue(s)
	}
}

func (s *LiteralValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitLiteralValue(s)
	}
}

type ColumnRefContext struct {
	Primary_exprContext
}

func NewColumnRefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnRefContext {
	var p = new(ColumnRefContext)

	InitEmptyPrimary_exprContext(&p.Primary_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Primary_exprContext))

	return p
}

func (s *ColumnRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnRefContext) Column_name() IColumn_nameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumn_nameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumn_nameContext)
}

func (s *ColumnRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterColumnRef(s)
	}
}

func (s *ColumnRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitColumnRef(s)
	}
}

func (p *MySQLCreateTableParser) Primary_expr() (localctx IPrimary_exprContext) {
	localctx = NewPrimary_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, MySQLCreateTableParserRULE_primary_expr)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		localctx = NewColumnRefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.Column_name()
		}

	case 2:
		localctx = NewLiteralValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(340)
			p.Literal()
		}

	case 3:
		localctx = NewFunctionExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(341)
			p.Function_call()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_function_call
	return p
}

func InitEmptyFunction_callContext(p *Function_callContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = MySQLCreateTableParserRULE_function_call
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = MySQLCreateTableParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserIDENTIFIER, 0)
}

func (s *Function_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserLPAREN, 0)
}

func (s *Function_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserRPAREN, 0)
}

func (s *Function_callContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *Function_callContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Function_callContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(MySQLCreateTableParserCOMMA)
}

func (s *Function_callContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(MySQLCreateTableParserCOMMA, i)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MySQLCreateTableListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (p *MySQLCreateTableParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, MySQLCreateTableParserRULE_function_call)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(344)
		p.Match(MySQLCreateTableParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(345)
		p.Match(MySQLCreateTableParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&63067986975522816) != 0 {
		{
			p.SetState(346)
			p.Expr()
		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == MySQLCreateTableParserCOMMA {
			{
				p.SetState(347)
				p.Match(MySQLCreateTableParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(348)
				p.Expr()
			}

			p.SetState(353)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(356)
		p.Match(MySQLCreateTableParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *MySQLCreateTableParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *Logical_exprContext = nil
		if localctx != nil {
			t = localctx.(*Logical_exprContext)
		}
		return p.Logical_expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *MySQLCreateTableParser) Logical_expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
