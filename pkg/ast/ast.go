package ast

import (
	"github.com/Neftik/project3/internal/models"
)

var (
	id int = 0
)

func priority(op string) (int, error) {
	switch {
	case op == "/" || op == "*":
		return 3, nil
	case op == "+" || op == "-":
		return 2, nil
	case op == "(":
		return 1, nil
	default:
		return 0, ErrUnknownOperator
	}
}

func ast(tokens []*token) (*models.AstNode, error) {
	var stack []*models.AstNode

	for _, tok := range tokens {
		switch tok.t {
		case operand:
			node := &models.AstNode{
				ID:      id,
				AstType: "number",
				Value:   tok.val,
			}
			stack = append(stack, node)
			id++

		case operator:
			if len(stack) < 2 {
				return nil, ErrInvalidExpression
			}

			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			node := &models.AstNode{
				ID:      id,
				AstType: "operation",
				Value:   tok.val,
				Left:    left,
				Right:   right,
			}
			stack = append(stack, node)
			id++

		default:
			return nil, ErrWrongCharacter
		}
	}

	if len(stack) != 1 {
		return nil, ErrInvalidExpression
	}

	return stack[0], nil
}
