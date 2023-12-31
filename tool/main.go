package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Print("Usage: generate-ast <output directory>")
		os.Exit(64)
	}

	outputDir := args[0]

	defineAst(outputDir, "Expr", []string{
		"Binary : left Expr, operator token.Token, right Expr",
		"Grouping : exprssion Expr",
		"Literal : value interface{}",
		"Unary : operator token.Token, right Expr",
	})
}

func defineAst(outputDir string, baseName string, types []string) {
	path := outputDir + "/" + strings.ToLower(baseName) + ".go"
	file, err := os.Create(path)
	if err != nil {
		log.Fatal("couldn't create a file")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	writer.WriteString("package ast\n")
	writer.WriteByte('\n')
	writer.WriteString("import \"p-friday/glox/token\"\n")
	writer.WriteByte('\n')
	writer.WriteString("type " + baseName + " interface {\n")
	writer.WriteString("	accept(Visitor) interface{}\n")
	writer.WriteString("}\n")
	writer.WriteByte('\n')
	
	defineVisitor(writer, baseName, types)

	writer.WriteByte('\n')
	for _, tp := range types {
		typeName := strings.TrimSpace(strings.Split(tp, ":")[0])
		fields := strings.TrimSpace(strings.Split(tp, ":")[1])
		defineType(writer, baseName, typeName, fields)
	}

	writer.Flush()
}

func defineType(writer *bufio.Writer, baseName string, typeName string, fieldList string) {
	writer.WriteString("type " + typeName + " struct {\n")
	fields := strings.Split(fieldList, ", ")
	for _, field := range fields {
		writer.WriteByte('\t')
		writer.WriteString(field)
		writer.WriteByte('\n')
	}
	writer.WriteString("}\n")
	writer.WriteByte('\n')

	// constructor
	//writer.WriteString("func New" + typeName + baseName + "()")
	
	// accept method override
	writer.WriteString("func (" + strings.ToLower(typeName) + " *" + typeName + ") accept(visitor Visitor) interface{} {\n")
	writer.WriteString("\treturn visitor.visit" + typeName + baseName + "(" + strings.ToLower(typeName) + ")\n")
	writer.WriteString("}\n")
	writer.WriteByte('\n')
}

func defineVisitor(writer *bufio.Writer, baseName string, types []string) {
	writer.WriteString("type Visitor interface {\n")
	for _, tp := range types {
		typeName := strings.TrimSpace(strings.Split(tp, ":")[0])
		writer.WriteString("	visit" + typeName + baseName + "(*" + typeName + ") interface{}\n")
	}
	writer.WriteString("}\n")
}
