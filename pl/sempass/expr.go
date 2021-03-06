package sempass

import (
	"shanhu.io/smlvm/pl/ast"
	"shanhu.io/smlvm/pl/tast"
	"shanhu.io/smlvm/pl/types"
)

func buildExpr(b *builder, expr ast.Expr) tast.Expr {
	if expr == nil {
		panic("bug")
	}

	switch expr := expr.(type) {
	case *ast.Operand:
		return buildOperand(b, expr)
	case *ast.ParenExpr:
		return buildExpr(b, expr.Expr)
	case *ast.MemberExpr:
		return buildMember(b, expr)
	case *ast.OpExpr:
		return buildOpExpr(b, expr)
	case *ast.StarExpr:
		return buildStarExpr(b, expr)
	case *ast.IndexExpr:
		return buildIndexExpr(b, expr)
	case *ast.CallExpr:
		return buildCallExpr(b, expr)
	case *ast.ArrayTypeExpr:
		t := b.buildType(expr)
		if t == nil {
			return nil
		}
		return tast.NewType(t)
	case *ast.FuncTypeExpr:
		t := b.buildType(expr)
		if t == nil {
			return nil
		}
		return tast.NewType(t)
	case *ast.ExprList:
		return buildExprList(b, expr)
	case *ast.ArrayLiteral:
		return buildArrayLit(b, expr)
	}

	b.Errorf(ast.ExprPos(expr), "invalid or not implemented: %T", expr)
	return nil
}

func buildConstExpr(b *builder, expr ast.Expr) tast.Expr {
	if expr == nil {
		panic("bug")
	}

	switch expr := expr.(type) {
	case *ast.ParenExpr:
		return buildConstExpr(b, expr.Expr)
	case *ast.Operand:
		return buildConstOperand(b, expr)
	case *ast.MemberExpr:
		return buildConstMember(b, expr)
	case *ast.OpExpr:
		return buildConstOpExpr(b, expr)
	case *ast.CallExpr:
		f := b.buildExpr(expr.Func)
		if f == nil {
			return nil
		}
		fref := f.R()
		// expr.Func must a Type
		if t, ok := fref.T.(*types.Type); ok {
			return buildConstCast(b, expr, t.T)
		}
	}
	b.CodeErrorf(
		ast.ExprPos(expr), "pl.expectConstExpr",
		"const var can only define by a const",
	)
	return nil
}

func buildExprStmt(b *builder, expr ast.Expr) tast.Stmt {
	if e, ok := expr.(*ast.CallExpr); ok {
		ret := buildExpr(b, e)
		if ret == nil {
			return nil
		}
		return &tast.ExprStmt{Expr: ret}
	}

	b.CodeErrorf(ast.ExprPos(expr), "pl.invalidExprStmt",
		"invalid expression statement")
	return nil
}
