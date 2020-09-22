package lexer

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/yagihash/monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
`
	cases := []struct {
		name string
		want token.Token
	}{
		{
			name: "let five = 5;",
			want: token.Token{Type: token.LET, Literal: "let"},
		},
		{
			name: "let five = 5;",
			want: token.Token{Type: token.IDENT, Literal: "five"},
		},
		{
			name: "let five = 5;",
			want: token.Token{Type: token.ASSIGN, Literal: "="},
		},
		{
			name: "let five = 5;",
			want: token.Token{Type: token.INT, Literal: "5"},
		},
		{
			name: "let five = 5;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "let ten = 10;",
			want: token.Token{Type: token.LET, Literal: "let"},
		},
		{
			name: "let ten = 10;",
			want: token.Token{Type: token.IDENT, Literal: "ten"},
		},
		{
			name: "let ten = 10;",
			want: token.Token{Type: token.ASSIGN, Literal: "="},
		},
		{
			name: "let ten = 10;",
			want: token.Token{Type: token.INT, Literal: "10"},
		},
		{
			name: "let ten = 10;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.LET, Literal: "let"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.IDENT, Literal: "add"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.ASSIGN, Literal: "="},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.FUNCTION, Literal: "fn"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.LPAREN, Literal: "("},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.IDENT, Literal: "x"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.COMMA, Literal: ","},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.IDENT, Literal: "y"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.RPAREN, Literal: ")"},
		},
		{
			name: "let add = fn(x, y) {",
			want: token.Token{Type: token.LBRACE, Literal: "{"},
		},
		{
			name: "x + y;",
			want: token.Token{Type: token.IDENT, Literal: "x"},
		},
		{
			name: "x + y;",
			want: token.Token{Type: token.PLUS, Literal: "+"},
		},
		{
			name: "x + y;",
			want: token.Token{Type: token.IDENT, Literal: "y"},
		},
		{
			name: "x + y;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "};",
			want: token.Token{Type: token.RBRACE, Literal: "}"},
		},
		{
			name: "};",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.LET, Literal: "let"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.IDENT, Literal: "result"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.ASSIGN, Literal: "="},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.IDENT, Literal: "add"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.LPAREN, Literal: "("},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.IDENT, Literal: "five"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.COMMA, Literal: ","},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.IDENT, Literal: "ten"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.RPAREN, Literal: ")"},
		},
		{
			name: "let result = add(five, ten);",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "!-/*5;",
			want: token.Token{Type: token.NOT, Literal: "!"},
		},
		{
			name: "!-/*5;",
			want: token.Token{Type: token.MINUS, Literal: "-"},
		},
		{
			name: "!-/*5;",
			want: token.Token{Type: token.SLASH, Literal: "/"},
		},
		{
			name: "!-/*5;",
			want: token.Token{Type: token.ASTERISK, Literal: "*"},
		},
		{
			name: "!-/*5;",
			want: token.Token{Type: token.INT, Literal: "5"},
		},
		{
			name: "!-/*5;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "5 < 10 > 5;",
			want: token.Token{Type: token.INT, Literal: "5"},
		},
		{
			name: "5 < 10 > 5;",
			want: token.Token{Type: token.LT, Literal: "<"},
		},
		{
			name: "5 < 10 > 5;",
			want: token.Token{Type: token.INT, Literal: "10"},
		},
		{
			name: "5 < 10 > 5;",
			want: token.Token{Type: token.GT, Literal: ">"},
		},
		{
			name: "5 < 10 > 5;",
			want: token.Token{Type: token.INT, Literal: "5"},
		},
		{
			name: "5 < 10 > 5;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.IF, Literal: "if"},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.LPAREN, Literal: "("},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.INT, Literal: "5"},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.LT, Literal: "<"},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.INT, Literal: "10"},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.RPAREN, Literal: ")"},
		},
		{
			name: "if (5 < 10) {",
			want: token.Token{Type: token.LBRACE, Literal: "{"},
		},
		{
			name: "return true;",
			want: token.Token{Type: token.RETURN, Literal: "return"},
		},
		{
			name: "return true;",
			want: token.Token{Type: token.TRUE, Literal: "true"},
		},
		{
			name: "return true;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "} else {",
			want: token.Token{Type: token.RBRACE, Literal: "}"},
		},
		{
			name: "} else {",
			want: token.Token{Type: token.ELSE, Literal: "else"},
		},
		{
			name: "} else {",
			want: token.Token{Type: token.LBRACE, Literal: "{"},
		},
		{
			name: "return false;",
			want: token.Token{Type: token.RETURN, Literal: "return"},
		},
		{
			name: "return false;",
			want: token.Token{Type: token.FALSE, Literal: "false"},
		},
		{
			name: "return false;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "}",
			want: token.Token{Type: token.RBRACE, Literal: "}"},
		},
		{
			name: "10 == 10;",
			want: token.Token{Type: token.INT, Literal: "10"},
		},
		{
			name: "10 == 10;",
			want: token.Token{Type: token.EQ, Literal: "=="},
		},
		{
			name: "10 == 10;",
			want: token.Token{Type: token.INT, Literal: "10"},
		},
		{
			name: "10 == 10;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "10 != 9;",
			want: token.Token{Type: token.INT, Literal: "10"},
		},
		{
			name: "10 != 9;",
			want: token.Token{Type: token.NOT_EQ, Literal: "!="},
		},
		{
			name: "10 != 9;",
			want: token.Token{Type: token.INT, Literal: "9"},
		},
		{
			name: "10 != 9;",
			want: token.Token{Type: token.SEMICOLON, Literal: ";"},
		},
		{
			name: "foobar",
			want: token.Token{Type: token.STRING, Literal: "foobar"},
		},
		{
			name: "foo bar",
			want: token.Token{Type: token.STRING, Literal: "foo bar"},
		},
		{
			name: "EOF",
			want: token.Token{Type: token.EOF, Literal: ""},
		},
	}

	l := New(input)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := l.NextToken()

			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("unexpected return value\n%s", diff)
			}
		})
	}
}
