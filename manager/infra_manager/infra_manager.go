package inframanager

import (
	"api-with-sql-nativ/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type InfraManager interface {
	Connect() *sql.DB
}

type infraManager struct {
	cfg *config.Config
	db  *sql.DB
}

func (i *infraManager) openConn() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", i.cfg.DB_HOST, i.cfg.DB_USER, i.cfg.DB_PASS, i.cfg.DB_NAME, i.cfg.DB_PORT)

	db, err := sql.Open(i.cfg.DB_DRIVER, dsn)

	if err != nil {
		panic("failed connect db")
	}

	db.Ping()

	i.db = db
}

// Connect implements InfraManager.
func (i *infraManager) Connect() *sql.DB {
	return i.db
}

func NewInfraManager(cfg *config.Config) InfraManager {
	conn := &infraManager{cfg: cfg}

	conn.openConn()

	return conn
}
