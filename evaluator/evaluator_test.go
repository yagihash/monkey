package evaluator

import (
	"testing"

	"github.com/yagihash/monkey/lexer"
	"github.com/yagihash/monkey/parser"

	"github.com/yagihash/monkey/object"
)

func TestEval(t *testing.T) {
	t.Run("IntegerExpression", func(t *testing.T) {
		cases := []struct {
			input    string
			expected int64
		}{
			{"5", 5},
			{"10", 10},
			{"-5", -5},
			{"-10", -10},
			{"5 + 5 + 5 + 5 - 10", 10},
			{"2 * 2 * 2 * 2 * 2", 32},
			{"-50 + 100 + -50", 0},
			{"5 * 2 + 10", 20},
			{"5 + 2 * 10", 25},
			{"20 + 2 * -10", 0},
			{"50 / 2 * 2 + 10", 60},
			{"2 * (5 + 10)", 30},
			{"3 * 3 * 3 + 10", 37},
			{"3 * (3 * 3) + 10", 37},
			{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)
				testIntegerObject(t, evaluated, c.expected)
			})
		}
	})

	t.Run("BooleanExpression", func(t *testing.T) {
		cases := []struct {
			input    string
			expected bool
		}{
			{"true", true},
			{"false", false},
			{"1 < 2", true},
			{"1 > 2", false},
			{"1 < 1", false},
			{"1 > 1", false},
			{"1 == 1", true},
			{"1 != 1", false},
			{"1 == 2", false},
			{"1 != 2", true},
			{"true == true", true},
			{"false == false", true},
			{"true == false", false},
			{"true != false", true},
			{"false != true", true},
			{"(1 < 2) == true", true},
			{"(1 < 2) == false", false},
			{"(1 > 2) == true", false},
			{"(1 > 2) == false", true},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)
				testBooleanObject(t, evaluated, c.expected)
			})
		}
	})

	t.Run("NotOperator", func(t *testing.T) {
		cases := []struct {
			input    string
			expected bool
		}{
			{"!true", false},
			{"!false", true},
			{"!5", false},
			{"!!true", true},
			{"!!false", false},
			{"!!5", true},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)
				testBooleanObject(t, evaluated, c.expected)
			})
		}
	})

	t.Run("IfElseExpressions", func(t *testing.T) {
		cases := []struct {
			input    string
			expected interface{}
		}{
			{"if (true) { 10 }", 10},
			{"if (false) { 10 }", nil},
			{"if (1) { 10 }", 10},
			{"if (1 < 2) { 10 }", 10},
			{"if (1 > 2) { 10 }", nil},
			{"if (1 > 2) { 10 } else { 20 }", 20},
			{"if (1 < 2) { 10 } else { 20 }", 10},
		}

		for _, c := range cases {
			evaluated := testEval(t, c.input)
			integer, ok := c.expected.(int)
			if ok {
				testIntegerObject(t, evaluated, int64(integer))
			} else {
				testNullObject(t, evaluated)
			}
		}
	})

	t.Run("ReturnStatements", func(t *testing.T) {
		cases := []struct {
			input    string
			expected int64
		}{
			{"return 10;", 10},
			{"return 10; 9;", 10},
			{"return 2 * 5; 9;", 10},
			{"9; return 2 * 5; 9;", 10},
			{"if (10 > 1) { return 10; }", 10},
			{
				`
if (10 > 1) {
  if (10 > 1) {
    return 10;
  }

  return 1;
}
`,
				10,
			},
			{
				`
let f = fn(x) {
  return x;
  x + 10;
};
f(10);`,
				10,
			},
			{
				`
let f = fn(x) {
   let result = x + 10;
   return result;
   return 10;
};
f(10);`,
				20,
			},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)
				testIntegerObject(t, evaluated, c.expected)
			})
		}
	})

	t.Run("ErrorHandling", func(t *testing.T) {
		cases := []struct {
			input           string
			expectedMessage string
		}{
			{
				"5 + true;",
				"type mismatch: INTEGER + BOOLEAN",
			},
			{
				"5 + true; 5;",
				"type mismatch: INTEGER + BOOLEAN",
			},
			{
				"-true",
				"unknown operator: -BOOLEAN",
			},
			{
				"true + false;",
				"unknown operator: BOOLEAN + BOOLEAN",
			},
			{
				"true + false + true + false;",
				"unknown operator: BOOLEAN + BOOLEAN",
			},
			{
				"5; true + false; 5",
				"unknown operator: BOOLEAN + BOOLEAN",
			},
			{
				"if (10 > 1) { true + false; }",
				"unknown operator: BOOLEAN + BOOLEAN",
			},
			{
				`
if (10 > 1) {
  if (10 > 1) {
    return true + false;
  }

  return 1;
}
`,
				"unknown operator: BOOLEAN + BOOLEAN",
			},
			{
				"foobar",
				"identifier not found: foobar",
			},
			{
				`"Hello" - "World"`,
				"unknown operator: STRING - STRING",
			},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)

				errObj, ok := evaluated.(*object.Error)
				if !ok {
					t.Errorf("no error object returned. got=%T(%+v)", evaluated, evaluated)
					return
				}

				if errObj.Message != c.expectedMessage {
					t.Errorf("wrong error message. expected=%q, got=%q", c.expectedMessage, errObj.Message)
				}
			})
		}
	})

	t.Run("LetStatements", func(t *testing.T) {
		cases := []struct {
			input    string
			expected int64
		}{
			{"let a = 5; a;", 5},
			{"let a = 5 * 5; a;", 25},
			{"let a = 5; let b = a; b;", 5},
			{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				testIntegerObject(t, testEval(t, c.input), c.expected)
			})
		}
	})

	t.Run("FunctionObject", func(t *testing.T) {
		input := "fn(x) { x + 2; };"
		evaluated := testEval(t, input)
		fn, ok := evaluated.(*object.Function)
		if !ok {
			t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
		}

		if len(fn.Parameters) != 1 {
			t.Fatalf("function has wrong parameers. Parameters=%+v", fn.Parameters)
		}

		if fn.Parameters[0].String() != "x" {
			t.Fatalf("parameter is not 'x'. got=%q", fn.Parameters[0])
		}

		expectedBody := "(x + 2)"

		if fn.Body.String() != expectedBody {
			t.Fatalf("body is not %q. got=%q", expectedBody, fn.Body.String())
		}
	})

	t.Run("FunctionCall", func(t *testing.T) {
		cases := []struct {
			input    string
			expected int64
		}{
			{"let identity = fn(x) { x; }; identity(5);", 5},
			{"let identity = fn(x) { return x; }; identity(5);", 5},
			{"let double = fn(x) { x * 2; }; double(5);", 10},
			{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
			{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
			{"fn(x) { x; }(5)", 5},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				testIntegerObject(t, testEval(t, c.input), c.expected)
			})
		}
	})

	t.Run("StringLiteral", func(t *testing.T) {
		input := `"Hello World!"`

		evaluated := testEval(t, input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T", evaluated)
		}

		if str.Value != "Hello World!" {
			t.Errorf("String has wrong value. got=%q", str.Value)
		}
	})

	t.Run("StringCOncatenation", func(t *testing.T) {
		input := `"Hello" + " " + "World!"`

		evaluated := testEval(t, input)
		str, ok := evaluated.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", evaluated, evaluated)
		}

		if str.Value != "Hello World!" {
			t.Errorf("String has wrong value. got=%q", str.Value)
		}
	})

	t.Run("StringComparison", func(t *testing.T) {
		cases := []struct {
			input string
			want  bool
		}{
			{`"aaa" == "aaa"`, true},
			{`"aaa" == "bbb"`, false},
			{`"aaa" != "aaa"`, false},
			{`"aaa" != "bbb"`, true},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)
				result, ok := evaluated.(*object.Boolean)
				if !ok {
					t.Fatalf("object is not Boolean. got=%T (%+v)", evaluated, evaluated)
				}

				if result.Value != c.want {
					t.Errorf("Result is not true. got=%t", result.Value)
				}
			})
		}
	})

	t.Run("BuiltinFunctions", func(t *testing.T) {
		cases := []struct {
			input string
			want  interface{}
		}{
			{`len("")`, 0},
			{`len("four")`, 4},
			{`len("hello world")`, 11},
			{`len(1)`, "argument to `len` not supported, got INTEGER"},
			{`len("one", "two")`, "wrong number of arguments. got=2, want=1"},
		}

		for _, c := range cases {
			t.Run(c.input, func(t *testing.T) {
				evaluated := testEval(t, c.input)

				switch want := c.want.(type) {
				case int:
					testIntegerObject(t, evaluated, int64(want))
				case string:
					errObj, ok := evaluated.(*object.Error)
					if !ok {
						t.Errorf("object is not Error. got=%T (%+v)", evaluated, evaluated)
						return
					}
					if errObj.Message != want {
						t.Errorf("wrong error message. expected=%q, got=%q",
							want, errObj.Message)
					}
				}
			})
		}
	})
}

func testNullObject(t *testing.T, obj object.Object) bool {
	t.Helper()

	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return false
	}

	return true
}

func testEval(t *testing.T, input string) object.Object {
	t.Helper()

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	t.Helper()

	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong avlue. got=%d, want=%d", result.Value, expected)
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	t.Helper()

	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, want=%t", result.Value, expected)
		return false
	}

	return true
}
