package ast

import (
	"strings"
	"sync"

	"github.com/Neftik/project3/internal/models"
)

var (
	mu sync.Mutex
)

func Build(expression string) (*models.AstNode, error) {
	mu.Lock()
	defer mu.Unlock()
	expression = strings.ReplaceAll(expression, " ", "")

	err := expErr(expression)
	if err != nil {
		return nil, err
	}

	tokens := tokens(expression)

	rpn, err := rpn(tokens)
	if err != nil {
		return nil, err
	}

	astRoot, err := ast(rpn)
	if err != nil {
		return nil, err
	}

	return astRoot, nil
}
