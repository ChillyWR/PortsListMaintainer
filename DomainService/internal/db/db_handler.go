package db

import (
	"github.com/ChillyWR/PortsListMaintainer/proto"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Port struct {
	ID         int64
	Name       string
	IsActive   bool
	Company    string
	Email      string
	Phone      string
	Address    string
	About      string
	Registered string
	Latitude   float32
	Longitude  float32
}

type Repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *Repo {
	dsn := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	return &Repo{db: db}
}

func ConnectToDBWithGorm() *Repo {
	dsn := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Printf("Connected to %v", dbname)
	return NewRepo(db)
}

func (r *Repo) ListPorts(lowerId int64, upperId int64, filterString string) *proto.Ports {
	var ports []Port
	result := r.db
	result = r.applyFilters(result, lowerId, upperId, filterString)
	result = result.Find(&ports)
	if result.Error != nil {
		log.Fatal("Failed to find table that fits given structure")
	}
	log.Println("Got ports from db")
	protoPorts := make([]*proto.Port, len(ports))
	for i, port := range ports {
		protoPorts[i] = convertToProtoPortStruct(&port)
	}

	return &proto.Ports{Ports: protoPorts}
}

func (r *Repo) applyFilters(view *gorm.DB, lowerId int64, upperId int64, filterString string) *gorm.DB {
	if lowerId != -1 || upperId != -1 {
		if lowerId != -1 && upperId != -1 && lowerId < upperId {
			view = view.Where("id BETWEEN ? AND ?", lowerId, upperId)
		} else if lowerId != -1 {
			view = view.Where("id > ?", lowerId)
		} else {
			view = view.Where("id < ?", upperId)
		}
	}
	if filterString != "" {
		view = view.Where("port LIKE '%?%'", filterString)
		log.Println("Applied filters")
	}
	return view
}

func (r *Repo) UpsertPorts(ports *proto.Ports) {
	for _, protoPort := range ports.GetPorts() {
		port := convertToPortStruct(protoPort)
		r.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(port)
		log.Printf("Saved port \n%+v\n", *port)
	}
}

func (r *Repo) CloseConnectionWithDB() {
	db, err := r.db.DB()
	if err != nil {
		log.Fatalf("Failed to get connection to DB instance from grom.DB: %v", err)
	}
	err = db.Close()
	if err != nil {
		log.Fatalf("Failed to close connection to DB: %v", err)
	}
}

//func ConnectToDBWithMigrations() {
//	connectionMessage := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
//	db, err := sql.Open(databaseDriver, connectionMessage)
//	if err != nil {
//		log.Fatalf("Failed to connect to DB: %v", err)
//	} else {
//		log.Println("Connected to DB")
//	}
//	err = db.Ping()
//	if err != nil {
//		log.Fatalf("Failed to ping db: %v", err)
//	}
//	driver, err := postgres.WithInstance(db, &postgres.Config{})
//	m, err := migrate.NewWithDatabaseInstance(
//		"file:///migrations",
//		dbname, driver)
//	err = m.Up()
//	if err != nil {
//		log.Fatalf("Failed to migarte up: %v", err)
//		return
//	}
//
//}
