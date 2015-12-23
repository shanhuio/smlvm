package sempass

import (
	"e8vm.io/e8vm/g8/ast"
	"e8vm.io/e8vm/g8/tast"
)

func buildExpr(b *builder, expr ast.Expr) tast.Expr {
	if expr == nil {
		panic("bug")
	}

	switch expr := expr.(type) {
	case *ast.Operand:
		return buildOperand(b, expr)
	default:
		b.Errorf(ast.ExprPos(expr), "invalid or not implemented: %T", expr)
		return nil
	}

	panic("todo")
}