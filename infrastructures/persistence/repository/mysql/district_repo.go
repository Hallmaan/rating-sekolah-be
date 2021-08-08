package mysql

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"rating-sekolah/domains"
	"rating-sekolah/helpers"
)

type mysqlDistrictRepository struct {
	Conn *sql.DB
}

func NewMysqlDistrictRepository(Conn *sql.DB) domains.DistrictRepository {
	return &mysqlDistrictRepository{Conn}
}

func (m *mysqlDistrictRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domains.District, err error) {
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

	result = make([]domains.District, 0)
	for rows.Next() {
		t := domains.District{}
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

func (m *mysqlDistrictRepository) Fetch(ctx context.Context, limit int64, offset int64) (result []domains.District, err error) {

	if(limit > 0){
		query := `SELECT id,name FROM district LIMIT ? OFFSET ? `
		result, err = m.fetch(ctx, query, limit, offset)
	} else {
		query := `SELECT id,name FROM district`
		result, err = m.fetch(ctx, query)
	}

	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlDistrictRepository) GetByID(ctx context.Context, id string) (result domains.District, err error) {
	query := `SELECT id,name
  						FROM district WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domains.District{}, err
	}

	if len(list) > 0 {
		result = list[0]
	} else {
		return result, helpers.ErrNotFound
	}

	return
}
