package expfixture_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	"os/exec"
	"time"

	"github.com/yoheimuta/xo-example-app/infra/exppath_test"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// DataSource represents a data source for test.
type DataSource struct {
	database string
	user     string
	password string
	host     string
}

// NewDataSource creates a new DataSource.
func NewDataSource(
	user string,
	password string,
	host string,
) *DataSource {
	database := fmt.Sprintf("test-%v", rand.Int63())
	return &DataSource{
		database: database,
		user:     user,
		password: password,
		host:     host,
	}
}

// Setup creates a new database and returns a dataSourceName.
func (d DataSource) Setup() (string, error) {
	out, err := exec.Command(
		"bash",
		exppath_test.ScriptRootPath("create_mysql_test_db.sh"),
		d.database,
		d.user,
		d.password,
		d.host,
	).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("err=%v, output=%v", err, string(out))
	}
	return d.dataSourceName(), nil
}

// OpenDB opens a new database connection.
func (d DataSource) OpenDB() (*sql.DB, error) {
	return sql.Open("mysql", d.dataSourceName())
}

// Close runs teardown operations.
func (d DataSource) Close() {
	// TODO: remove the test database.
}

func (d DataSource) dataSourceName() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?parseTime=true",
		d.user,
		d.password,
		d.host,
		d.database,
	)
}
