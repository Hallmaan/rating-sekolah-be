package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"rating-sekolah/domains"
	"rating-sekolah/helpers"
)

type mysqlProvinceRepository struct {
	Conn *sql.DB
}

func NewMysqlProvinceRepository(Conn *sql.DB) domains.ProvinceRepository {
	return &mysqlProvinceRepository{Conn}
}

func (m *mysqlProvinceRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domains.Province, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	//rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domains.Province, 0)
	for rows.Next() {
		t := domains.Province{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlProvinceRepository) Fetch(ctx context.Context, limit int64, offset int64) (result []domains.Province, err error) {
	query := `SELECT id,name FROM province LIMIT ? OFFSET ? `

	//query := `SELECT id,sekolah FROM sekolah limit 1`

	result, err = m.fetch(ctx, query, limit, offset)
	fmt.Println(result, "result from")
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlProvinceRepository) GetByID(ctx context.Context, id string) (result domains.Province, err error) {
	query := `SELECT id,name
  						FROM province WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domains.Province{}, err
	}

	if len(list) > 0 {
		result = list[0]
	} else {
		return result, helpers.ErrNotFound
	}

	fmt.Println(result)

	return
}