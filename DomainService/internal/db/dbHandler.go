package db

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

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

func ConnectWithGorm() *Repo {
	dsn := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	return NewRepo(db)
}

func (r *Repo) ListPorts() map[int32]int32 {
	var ports []Ports
	res := r.db.Find(&ports)
	if res.Error != nil {
		log.Fatal("Failed to find table that fits given structure")
	}
	result := make(map[int32]int32)
	for _, port := range ports {
		result[port.id] = port.port
	}
	return result
}

func (r *Repo) UpsertPort(port int32) {
	toSave := &Ports{port: port}
	//r.db.Save(toSave)
	r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"port"}), // column needed to be updated
	}).Create(toSave)
	log.Println("Saved port ", port)
}

// Update columns to new value on `id` conflict
// INSERT INTO "users" *** ON CONFLICT ("id") DO UPDATE SET "name"="excluded"."name", "age"="excluded"."age"; PostgreSQL

func (r *Repo) FilterBetweenIndex(indexLower, indexUpper int32) []Ports {
	var ports []Ports
	res := r.db.Where("id BETWEEN ? AND ?", indexLower, indexUpper).Find(&ports)
	if res.Error != nil {
		log.Fatal("Failed to find table that fits given structure")
	}
	return ports
}

func (r *Repo) FilterBetweenPort(portLower, portUpper int) []Ports {
	var ports []Ports
	res := r.db.Where("port BETWEEN ? AND ?", portLower, portUpper).Find(&ports)
	if res.Error != nil {
		log.Fatal("Failed to find table that fits given structure")
	}
	return ports
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
