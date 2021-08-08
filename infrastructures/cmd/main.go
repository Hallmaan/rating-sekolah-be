package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"rating-sekolah/handlers"
	repoMysql "rating-sekolah/infrastructures/persistence/repository/mysql"
	"rating-sekolah/usecases"
	"time"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	router := gin.Default()
	router.Use(cors.Default())

	/* SCHOOL ADAPTER */
	schoolRepo := repoMysql.NewMysqlSchoolRepository(dbConn)
	schoolUsecase := usecases.NewSchoolUsecase(schoolRepo, timeoutContext)
	schoolHandler := handlers.NewSchoolHandler(schoolUsecase)

	/* PROVINCE ADAPTER */
	provinceRepo := repoMysql.NewMysqlProvinceRepository(dbConn)
	provinceUsecase := usecases.NewProvinceUsecase(provinceRepo, timeoutContext)
	provinceHandler := handlers.NewProvinceHandler(provinceUsecase)

	/* DISTRICT ADAPTER */
	districtRepo := repoMysql.NewMysqlDistrictRepository(dbConn)
	districtUsecase := usecases.NewDistrictUsecase(districtRepo, timeoutContext)
	districtHandler := handlers.NewDistrictHandler(districtUsecase)

	/* ANY ADAPTER */

	api := router.Group("/api/v1")
	api.GET("/school", schoolHandler.FetchSchool)
	api.GET("/school/:id", schoolHandler.GetSchoolById)
	api.GET("/school/province", provinceHandler.FetchProvince) //provinsi
	//api.GET("/school/province/:id", provinceHandler.GetById) //provinsi
	api.GET("/school/district", districtHandler.FetchDistrict) //kabupaten
	//api.GET("/school/district/:id", districtHandler.GetById) //kabupaten
	//api.GET("/school/level", schoolHandler.FetchSchool) //tingkat
	//api.GET("/school/category", schoolHandler.FetchSchool) //kategori

	router.Run()
}