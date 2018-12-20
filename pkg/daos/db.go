package daos

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alamin-mahamud/gapi/pkg/models"
	"github.com/alamin-mahamud/gapi/pkg/utils"
	_ "github.com/go-sql-driver/mysql"
)

// DB is the base type which will be derived By Implementation Logics and Drivers
type DB interface {
	Open(*DBConfig) error
	Fetch(ctx context.Context, num int64) ([]*models.Post, error)
	GetByID(ctx context.Context, id int64) (*models.Post, error)
	Create(ctx context.Context, p *models.Post) (int64, error)
	Update(ctx context.Context, p *models.Post) (*models.Post, error)
	Delete(ctx context.Context, id int64) (bool, error)
	// Close()
}

type MySQL struct {
	Conn *sql.DB
}

func NewMySQL() DB {
	return &MySQL{}
}

// Open - implements the mysql driver open
func (m *MySQL) Open(dbConfig *DBConfig) error {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		dbConfig.DbUser,
		dbConfig.DbPass,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbName,
	)
	d, err := sql.Open("mysql", dbSource)
	utils.CheckErr(err)
	m.Conn = d
	return err
}

func (m *MySQL) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := m.Conn.QueryContext(
		ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payload := make([]*models.Post, 0)

	for rows.Next() {
		data := new(models.Post)
		err := rows.Scan(&data.ID, &data.Title, &data.Content)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}

	return payload, nil
}

func (m *MySQL) Fetch(ctx context.Context, num int64) ([]*models.Post, error) {
	if num == 0 {
		num = 10
	}

	query := "Select id, title, content from posts limit ?"

	return m.fetch(ctx, query, num)
}

func (m *MySQL) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := "Select id, title, content From posts where id=?"
	rows, err := m.fetch(ctx, query, id)

	if err != nil {
		return nil, err
	}

	payload := &models.Post{}

	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *MySQL) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "Insert posts SET title=?, content=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)

	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Title, p.Content)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *MySQL) Update(ctx context.Context, p *models.Post) (*models.Post, error) {
	query := "Update posts set title=?, content=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Content,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *MySQL) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From posts Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
