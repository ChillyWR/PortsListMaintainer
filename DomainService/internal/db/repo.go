package db

import (
	"github.com/ChillyWR/PortsListMaintainer/DomainService/pkg"
	"github.com/ChillyWR/PortsListMaintainer/config"
	"github.com/ChillyWR/PortsListMaintainer/proto"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type Repo struct {
	db *gorm.DB
}

func NewRepo() (*Repo, error) {
	dsn := "user=" + config.User + " dbname=" + config.DBName + " password=" + config.Password + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Printf("Connected to %v", config.DBName)
	return &Repo{db: db}, nil
}

func (r *Repo) ListPorts(id int64, idLower int64, idUpper int64, filterString *string) (*proto.Ports, error) {
	var ports []config.Port
	result := r.db
	result = r.applyFilters(result, id, idLower, idUpper, filterString)
	result = result.Find(&ports)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("Got ports from db")
	protoPorts := make([]*proto.Port, len(ports))
	for i, port := range ports {
		protoPorts[i] = pkg.ConvertToProtoPortStruct(&port)
	}

	return &proto.Ports{Ports: protoPorts}, nil
}

func (r *Repo) applyFilters(view *gorm.DB, id int64, idLower int64, idUpper int64, filterString *string) *gorm.DB {
	if id != -1 {
		view = view.Where("id = ?", id)
	} else if idLower != -1 || idUpper != -1 {
		if idLower != -1 && idUpper != -1 && idLower < idUpper {
			view = view.Where("id BETWEEN ? AND ?", idLower, idUpper)
		} else if idLower != -1 {
			view = view.Where("id > ?", idLower)
		} else {
			view = view.Where("id < ?", idUpper)
		}
	}
	if *filterString != "" {
		view = view.Where("name LIKE ?", "%"+*filterString+"%")
		log.Println("Applied filters")
	}
	return view
}

func (r *Repo) UpsertPorts(protoPorts *proto.Ports) {
	for _, protoPort := range protoPorts.GetPorts() {
		port := pkg.ConvertToPortStruct(protoPort)
		result := r.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(port)
		// TODO: handle result.Error
		if result.Error != nil {
			log.Printf("Saved port \n%+v\n", *port)
		}
	}
}

func (r *Repo) CloseConnectionWithDB() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	return err
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
