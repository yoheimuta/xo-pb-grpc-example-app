package expdep_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/yoheimuta/xo-pb-example-app/app/userproductapp"

	"github.com/yoheimuta/xo-pb-example-app/infra/exptime"

	"github.com/yoheimuta/xo-pb-example-app/app/userapp"
	"github.com/yoheimuta/xo-pb-example-app/infra/expmysql/expfixture_test"
	"github.com/yoheimuta/xo-pb-example-app/repository/expdb/expmysql"
)

// Dep represents dependencies used from tests.
type Dep struct {
	dbDataSource   *expfixture_test.DataSource
	clock          *exptime.Clock
	userApp        *userapp.App
	userProductApp *userproductapp.App

	rawDB *sql.DB
}

// NewDep creates a new Dep.
func NewDep() (*Dep, error) {
	dataSource := expfixture_test.NewDataSource(
		"root",
		"my-pw",
		"127.0.0.1",
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
	userProductApp := userproductapp.NewApp(
		db,
	)

	rawDB, err := dataSource.OpenDB()
	if err != nil {
		return nil, err
	}

	return &Dep{
		dbDataSource:   dataSource,
		clock:          clock,
		userApp:        userApp,
		userProductApp: userProductApp,
		rawDB:          rawDB,
	}, nil
}

// Close destroys some dependencies.
func (d *Dep) Close(t *testing.T) {
	d.dbDataSource.Close()
	err := d.rawDB.Close()
	if err != nil {
		t.Errorf("failed rawDB.Close(), err=%v", err)
	}
}

// RawDB returns a raw database connection.
func (d *Dep) RawDB() *sql.DB {
	return d.rawDB
}

// Now returns a current time.
func (d *Dep) Now() time.Time {
	return d.clock.Now()
}

// UserApp returns an user application.
func (d *Dep) UserApp() *userapp.App {
	return d.userApp
}

// UserProductApp returns an userProduct application.
func (d *Dep) UserProductApp() *userproductapp.App {
	return d.userProductApp
}
