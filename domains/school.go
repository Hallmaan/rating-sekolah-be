package domains

import "context"

type School struct {
	Id              string   `json:"id"`
	Name         string  `json:"name"`
	//Npsn            int64   `json:"npsn"`
	//Bentuk          string  `json:"bentuk"`
	//Status          string  `json:"status"`
	//Province        string  `json:"propinsi"`
	//ProvinceCode    int64   `json:"kode_prop"`
	//City            string  `json:"kabupaten_kota"`
	//CityCode        int64   `json:"kode_kab_kota"`
	//SubDistrictCode int64   `json:"kode_kec"`
	//SubDistrict     string  `json:"kecamatan"`
	//Address         string  `json:"alamat_jalan"`
	//Latitude        float64 `json:"lintang"`
	//Longitude       float64 `json:"bujur"`
}

type SchoolDistrict struct {
	DistrictCode        int64   `json:"kode_kab_kota"`
	District            string  `json:"kabupaten_kota"`
}

type SchoolUseCase interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]School, error)
	GetByID(ctx context.Context, id string) (School, error)
	//GetDistrict(ctx context.Context) (SchoolDistrict, error)
	//Update(ctx context.Context, ar *School) error
	//GetByTitle(ctx context.Context, title string) (School, error)
	//Store(context.Context, *School) error
	//Delete(ctx context.Context, id int64) error
}

type SchoolRepository interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]School, error)
	GetByID(ctx context.Context, id string) (School, error)
	//GetDistrict(ctx context.Context) (SchoolDistrict, error)
	//Update(ctx context.Context, ar *School) error
	//GetByTitle(ctx context.Context, title string) (School, error)
	//Store(context.Context, *School) error
	//Delete(ctx context.Context, id int64) error
}
