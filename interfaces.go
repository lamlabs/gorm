package gorm

import (
	"context"
	"database/sql"

	"github.com/lamlabs/gorm/clause"
	"github.com/lamlabs/gorm/schema"
)

// Dialector GORM database dialector
type Dialector interface {
	Name() string
	Initialize(*DB) error
	Migrator(db *DB) Migrator
	DataTypeOf(*schema.Field) string
	DefaultValueOf(*schema.Field) clause.Expression
	BindVarTo(writer clause.Writer, stmt *Statement, v interface{})
	QuoteTo(clause.Writer, string)
	Explain(sql string, vars ...interface{}) string
}

// Plugin GORM plugin interface
type Plugin interface {
	Name() string
	Initialize(*DB) error
}

// ConnPool db conns pool interface
type ConnPool interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// SavePointerDialectorInterface save pointer interface
type SavePointerDialectorInterface interface {
	SavePoint(tx *DB, name string) error
	RollbackTo(tx *DB, name string) error
}

// TxBeginner tx beginner
type TxBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// ConnPoolBeginner conn pool beginner
type ConnPoolBeginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (ConnPool, error)
}

// TxCommitter tx committer
type TxCommitter interface {
	Commit() error
	Rollback() error
}

// Tx sql.Tx interface
type Tx interface {
	ConnPool
	TxCommitter
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
}

// Valuer gorm valuer interface
type Valuer interface {
	GormValue(context.Context, *DB) clause.Expr
}

// GetDBConnector SQL db connector
type GetDBConnector interface {
	GetDBConn() (*sql.DB, error)
}

type Db interface {
	Session(config *Session) *DB
	WithContext(ctx context.Context) *DB
	Debug() (tx *DB)
	Set(key string, value interface{}) *DB
	Get(key string) (interface{}, bool)
	InstanceSet(key string, value interface{}) *DB
	InstanceGet(key string) (interface{}, bool)
	Callback() *callbacks
	AddError(err error) error
	DB() (*sql.DB, error)
	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
	Use(plugin Plugin) error
	ToSQL(queryFn func(tx *DB) *DB) string
	Association(column string) *Association
	Create(value interface{}) (tx *DB)
	CreateInBatches(value interface{}, batchSize int) (tx *DB)
	Save(value interface{}) (tx *DB)
	First(dest interface{}, conds ...interface{}) (tx *DB)
	Take(dest interface{}, conds ...interface{}) (tx *DB)
	Last(dest interface{}, conds ...interface{}) (tx *DB)
	Find(dest interface{}, conds ...interface{}) (tx *DB)
	FindInBatches(dest interface{}, batchSize int, fc func(tx *DB, batch int) error) *DB
	FirstOrInit(dest interface{}, conds ...interface{}) (tx *DB)
	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *DB)
	Update(column string, value interface{}) (tx *DB)
	Updates(values interface{}) (tx *DB)
	UpdateColumn(column string, value interface{}) (tx *DB)
	UpdateColumns(values interface{}) (tx *DB)
	Delete(value interface{}, conds ...interface{}) (tx *DB)
	Count(count *int64) (tx *DB)
	Row() *sql.Row
	Rows() (*sql.Rows, error)
	Scan(dest interface{}) (tx *DB)
	Pluck(column string, dest interface{}) (tx *DB)
	ScanRows(rows *sql.Rows, dest interface{}) error
	Connection(fc func(tx *DB) error) (err error)
	Transaction(fc func(tx *DB) error, opts ...*sql.TxOptions) (err error)
	Begin(opts ...*sql.TxOptions) *DB
	Commit() *DB
	Rollback() *DB
	SavePoint(name string) *DB
	RollbackTo(name string) *DB
	Exec(sql string, values ...interface{}) (tx *DB)
	Migrator() Migrator
	AutoMigrate(dst ...interface{}) error
	Model(value interface{}) (tx *DB)
	Clauses(conds ...clause.Expression) (tx *DB)
	Table(name string, args ...interface{}) (tx *DB)
	Distinct(args ...interface{}) (tx *DB)
	Select(query interface{}, args ...interface{}) (tx *DB)
	Omit(columns ...string) (tx *DB)
	Where(query interface{}, args ...interface{}) (tx *DB)
	Not(query interface{}, args ...interface{}) (tx *DB)
	Or(query interface{}, args ...interface{}) (tx *DB)
	Joins(query string, args ...interface{}) (tx *DB)
	Group(name string) (tx *DB)
	Having(query interface{}, args ...interface{}) (tx *DB)
	Order(value interface{}) (tx *DB)
	Limit(limit int) (tx *DB)
	Offset(offset int) (tx *DB)
	Scopes(funcs ...func(*DB) *DB) (tx *DB)
	Preload(query string, args ...interface{}) (tx *DB)
	Attrs(attrs ...interface{}) (tx *DB)
	Assign(attrs ...interface{}) (tx *DB)
	Unscoped() (tx *DB)
	Raw(sql string, values ...interface{}) (tx *DB)
}
