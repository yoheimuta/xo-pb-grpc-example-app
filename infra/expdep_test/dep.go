package expdep_test

import (
	"database/sql"

	"github.com/yoheimuta/xo-example-app/app/userapp"
	"github.com/yoheimuta/xo-example-app/infra/expmysql/expfixture_test"
	"github.com/yoheimuta/xo-example-app/repository/expdb/expmysql"
)

// Dep represents dependencies used from tests.
type Dep struct {
	dbDataSource *expfixture_test.DataSource
	userApp      *userapp.App
}

// NewDep creates a new Dep.
func NewDep() (*Dep, error) {
	dataSource := expfixture_test.NewDataSource(
		"root",
		"my-pw",
		"0.0.0.0",
	)
	dataSourceName, err := dataSource.Setup()
	if err != nil {
		return nil, err
	}
	db, err := expmysql.NewClient(dataSourceName)
	if err != nil {
		return nil, err
	}

	userApp := userapp.NewApp(
		db,
	)

	return &Dep{
		dbDataSource: dataSource,
		userApp:      userApp,
	}, nil
}

// Close destroys some dependencies.
func (d *Dep) Close() {
	d.dbDataSource.Close()
}

// OpenRawDB opens a new database connection.
func (d *Dep) OpenRawDB() (*sql.DB, error) {
	return d.dbDataSource.OpenDB()
}

// UserApp returns an user application.
func (d *Dep) UserApp() *userapp.App {
	return d.userApp
}
