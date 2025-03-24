package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xwb1989/sqlparser"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "curd-tool",
		Short: "A tool to generate Gorm PO from MySQL DDL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please provide a MySQL DDL file.")
		},
	}

	var ddlFile string

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate Gorm PO from MySQL DDL",
		Run: func(cmd *cobra.Command, args []string) {
			if ddlFile == "" {
				fmt.Println("Please specify a DDL file using --file flag.")
				return
			}
			// 调用生成Gorm PO的函数
			generateGormPO(ddlFile)
		},
	}

	generateCmd.Flags().StringVarP(&ddlFile, "file", "f", "", "MySQL DDL file")
	rootCmd.AddCommand(generateCmd)
	rootCmd.Execute()
}

func generateGormPO(ddlFile string) {
	// 读取DDL文件内容
	content, err := ioutil.ReadFile(ddlFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", ddlFile, err)
		return
	}

	// 解析DDL文件内容并生成Gorm PO
	ddl := string(content)
	poContent := parseDDLToGormPO(ddl)

	// 获取输出文件路径
	outputFile := strings.Replace(ddlFile, ".sql", ".go", 1)

	// 将生成的Gorm PO写入文件
	err = ioutil.WriteFile(outputFile, []byte(poContent), 0644)
	if err != nil {
		fmt.Printf("Error writing file %s: %v\n", outputFile, err)
		return
	}

	fmt.Printf("Generated Gorm PO file: %s\n", outputFile)
}

func parseDDLToGormPO(ddl string) string {
	stmt, err := sqlparser.Parse(ddl)
	if err != nil {
		fmt.Printf("Error parsing DDL: %v\n", err)
		return ""
	}

	// 解析DDL并生成Gorm PO
	switch stmt := stmt.(type) {
	case *sqlparser.DDL:
		if stmt.Action == "create" {
			return generateGormStruct(stmt)
		}
	}

	return ""
}

func generateGormStruct(ddl *sqlparser.DDL) string {
	structName := ddl.NewName.Name.String()
	var fields []string

	for _, col := range ddl.TableSpec.Columns {
		field := fmt.Sprintf("%s %s `gorm:\"column:%s\"`", col.Name.String(), sqlTypeToGoType(col.Type.Type), col.Name.String())
		fields = append(fields, field)
	}

	return fmt.Sprintf(`package example

import (
	"github.com/jinzhu/gorm"
)

type %s struct {
	gorm.Model
	%s
}`, structName, strings.Join(fields, "\n\t"))
}

func sqlTypeToGoType(sqlType string) string {
	switch sqlType {
	case "int", "bigint":
		return "int"
	case "varchar", "text":
		return "string"
	// 添加更多类型映射
	default:
		return "string"
	}
}
