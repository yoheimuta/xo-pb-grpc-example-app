package expdep_test

import (
	"database/sql"
	"time"

	"github.com/yoheimuta/xo-example-app/infra/exptime"

	"github.com/yoheimuta/xo-example-app/app/userapp"
	"github.com/yoheimuta/xo-example-app/infra/expmysql/expfixture_test"
	"github.com/yoheimuta/xo-example-app/repository/expdb/expmysql"
)

// Dep represents dependencies used from tests.
type Dep struct {
	dbDataSource *expfixture_test.DataSource
	clock        *exptime.Clock
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

	clock := exptime.NewClock()

	userApp := userapp.NewApp(
		db,
		clock,
	)

	return &Dep{
		dbDataSource: dataSource,
		clock:        clock,
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

// Now returns a current time.
func (d *Dep) Now() time.Time {
	return d.clock.Now()
}

// UserApp returns an user application.
func (d *Dep) UserApp() *userapp.App {
	return d.userApp
}
