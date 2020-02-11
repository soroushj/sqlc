package ast

type Node interface {
	Pos() int
}

type Statement struct {
	Raw *RawStmt
}

type RawStmt struct {
	Stmt Node
}

func (n *RawStmt) Pos() int {
	return 0
}

type TableName struct {
	Catalog string
	Schema  string
	Name    string
}

type CreateTableStmt struct {
	IfNotExists bool
	Name        TableName
	Cols        []*ColumnDef
}

func (n *CreateTableStmt) Pos() int {
	return 0
}

type ColumnDef struct {
	Colname  string
	TypeName *TypeName
}

func (n *ColumnDef) Pos() int {
	return 0
}

type TypeName struct {
	Name string
}

func (n *TypeName) Pos() int {
	return 0
}
