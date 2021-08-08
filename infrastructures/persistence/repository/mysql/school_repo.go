package mysql

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	"rating-sekolah/domains"
	"rating-sekolah/helpers"
)

type mysqlSchoolRepository struct {
	Conn *sql.DB
}

func NewMysqlSchoolRepository(Conn *sql.DB) domains.SchoolRepository {
	return &mysqlSchoolRepository{Conn}
}

func (m *mysqlSchoolRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domains.School, err error) {
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

	result = make([]domains.School, 0)
	for rows.Next() {
		t := domains.School{}
		err = rows.Scan(
			&t.Id,
			&t.Name,
			//&t.Npsn,
			//&t.Bentuk,
			//&t.Status,
			//&t.Province,
			//&t.ProvinceCode,
			//&t.City,
			//&t.CityCode,
			//&t.SubDistrictCode,
			//&t.SubDistrict,
			//&t.Address,
			//&t.Latitude,
			//&t.Longitude,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *mysqlSchoolRepository) Fetch(ctx context.Context, limit int64, offset int64) (result []domains.School, err error) {
	query := `SELECT id,name FROM school LIMIT ? OFFSET ? `

	//query := `SELECT id,sekolah FROM sekolah limit 1`

	result, err = m.fetch(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlSchoolRepository) GetByID(ctx context.Context, id string) (result domains.School, err error) {
	query := `SELECT id,name
  						FROM school WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domains.School{}, err
	}

	if len(list) > 0 {
		result = list[0]
	} else {
		return result, helpers.ErrNotFound
	}

	return
}