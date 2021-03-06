package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/yagihash/monkey/ast"
)

type ObjectType string

const (
	IntegerObj     = "INTEGER"
	StringObj      = "STRING"
	BoolenaObj     = "BOOLEAN"
	NullObj        = "NULL"
	ErrObj         = "ERROR"
	ReturnValueObj = "RETURN_VALUE"
	FunctionObj    = "FUNCTION"
	BuiltinObj     = "BUILTIN"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i Integer) Type() ObjectType {
	return IntegerObj
}

func (i Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Boolean struct {
	Value bool
}

func (b Boolean) Type() ObjectType {
	return BoolenaObj
}

func (b Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Null struct{}

func (n Null) Type() ObjectType {
	return NullObj
}

func (n Null) Inspect() string {
	return "null"
}

type Error struct {
	Message string
}

func (e Error) Type() ObjectType {
	return ErrObj
}

func (e Error) Inspect() string {
	return "ERROR: " + e.Message
}

type ReturnValue struct {
	Value Object
}

func (rv ReturnValue) Type() ObjectType {
	return ReturnValueObj
}

func (rv ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f Function) Type() ObjectType {
	return FunctionObj
}

func (f Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s String) Type() ObjectType {
	return StringObj
}

func (s String) Inspect() string {
	return s.Value
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b Builtin) Type() ObjectType {
	return BuiltinObj
}

func (b Builtin) Inspect() string {
	return "fn() { builtin function }"
}
