package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"rating-sekolah/domains"
)

type mysqlSchoolRepository struct {
	Conn *sql.DB
}

func NewMysqlSchoolRepository(Conn *sql.DB) domains.SchoolRepository {
	return &mysqlSchoolRepository{Conn}
}

func (m *mysqlSchoolRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domains.School, err error) {
	//rows, err := m.Conn.QueryContext(ctx, query, args...)
	rows, err := m.Conn.QueryContext(ctx, query)
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
			&t.Sekolah,
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
	//query := `SELECT * FROM sekolah LIMIT ? OFFSET ? `

	query := `SELECT id,sekolah FROM sekolah limit 1`

	result, err = m.fetch(ctx, query, limit, offset)
	fmt.Println(result, "result from")
	if err != nil {
		return nil, err
	}

	return
}