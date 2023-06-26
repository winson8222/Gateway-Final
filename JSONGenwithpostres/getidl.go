package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// The gorm model for Microservice table from postgres database
type Microservice struct {
	BaseModel
	Svcname  string    `gorm:"unique;column:svcname"`
	Versions []Version `gorm:"foreignKey:MicroserviceId"`
}

func (Microservice) TableName() string {
	return "Microservice"
}

// The gorm model for Version table from postgres database
type Version struct {
	BaseModel
	Vname          string       `gorm:"column:vname"`
	Upstreamurl    string       `gorm:"column:upstreamurl"`
	Idlfile        []byte       `gorm:"column:idlfile"`
	Idlname        string       `gorm:"column:idlname"`
	MicroserviceId string       `gorm:"column:microserviceId"`
	Microservice   Microservice `gorm:"foreignKey:MicroserviceId"`
}

func (Version) TableName() string {
	return "Version"
}

// Download IDL files and retrieve the information of services shwoon
func GetIDL() (GatewayInfo, []ServiceInfo) {
	dsn := "host=127.0.0.1 user=postgres password=mysecretpassword dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	//Access the underlying sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get *sql.DB instance: " + err.Error())
	}

	var microservices []Microservice

	// Retrieve microservices from the database
	result := db.Preload("Versions").Find(&microservices)

	if result.Error != nil {
		panic("Failed to retrieve Microservice entries: " + result.Error.Error())
	}

	for _, microservice := range microservices {
		fmt.Printf("ID: %s, Svcname: %s\n", microservice.BaseModel.ID, microservice.Svcname)
		for _, version := range microservice.Versions {
			fmt.Printf("Version ID: %s, Vname: %s\n", version.BaseModel.ID, version.Vname)
			data := version.Idlfile
			err := ioutil.WriteFile("idl/"+version.Idlname, data, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = sqlDB.Close()
	if err != nil {
		panic("Failed to close database connection: " + err.Error())
	}

	gatewayinfo := GatewayInfo{
		GatewayURL:  "0.0.0.0:8888",
		ETCD_URL:    "0.0.0.0:20000",
		GatewayName: "example",
	}

	services := []ServiceInfo{}

	for i, service := range microservices {
		index := i + 1
		fmt.Print("IDL " + fmt.Sprint(index) + ": " + service.Versions[len(service.Versions)-1].Idlname + "\n")
		toadd := ServiceInfo{
			IDLName:             service.Versions[len(service.Versions)-1].Idlname, //to be replaced by data from postgres
			ServiceUpstreamURL:  service.Versions[len(service.Versions)-1].Upstreamurl,
			Load_Balancing_Type: "ROUND_ROBIN",
		}

		services = append(services, toadd)
	}

	return gatewayinfo, services
}
