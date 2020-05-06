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
			name: "EOF",
			want: token.Token{Type: token.EOF, Literal: ""},
		},
	}

	l := New(input)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := l.NextToken()

			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("unexpected returned value\n%s", diff)
			}
		})
	}
}
