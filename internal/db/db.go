package db

import (
	"fmt"
	"log"
	"time"

	"vericred/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// Use the provided managed PostgreSQL connection string directly
	dsn := "postgresql://ikvvhzvdwblfanaplwtv:skogjnwlomkhojmbhptqsbwdgjhofd@9qasp5v56q8ckkf5dc.leapcellpool.com:6438/bieoohmpqpdwdkumqegq?sslmode=require"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("connection to db failed:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get db from GORM: ", err)
	}
	sqlDB.SetConnMaxLifetime(time.Hour)
	fmt.Println("(SUCCESS): connected to database successfully ")

	models.InitDB(DB)

	// Drop any stale/incorrect FK created previously by older model tags
	DB.Exec("ALTER TABLE credentials DROP CONSTRAINT IF EXISTS fk_organizations_credentials;")

	// AutoMigrate required tables
	if err = DB.AutoMigrate(&models.Accounts{}); err != nil {
		log.Fatal("AutoMigration failed for Accounts: ", err)
	}
	if err = DB.AutoMigrate(&models.Organization{}); err != nil {
		log.Fatal("AutoMigration failed for Organization: ", err)
	}
	if err = DB.AutoMigrate(&models.Users{}); err != nil {
		log.Fatal("AutoMigration failed for Users: ", err)
	}
	if err = DB.AutoMigrate(&models.PendingRequest{}); err != nil {
		log.Fatal("AutoMigration failed for PendingRequest: ", err)
	}
	if err = DB.AutoMigrate(&models.Credential{}); err != nil {
		log.Fatal("AutoMigration failed for Credential: ", err)
	}
	if err = DB.AutoMigrate(&models.Transaction{}); err != nil {
		log.Fatal("AutoMigration failed for Transaction: ", err)
	}

	// Ensure correct FKs exist for Credential relations
	_ = DB.Migrator().CreateConstraint(&models.Credential{}, "User")
	_ = DB.Migrator().CreateConstraint(&models.Credential{}, "Organization")
}
