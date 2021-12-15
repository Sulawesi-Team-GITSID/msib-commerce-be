package main

import (
	"backend-service/configuration/config"
	"backend-service/entity"
	"fmt"
	"os"
	"sort"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config, err := config.NewConfig(".env")
	checkError(err)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		//dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	checkError(err)

	sqlDB, err := db.DB()
	if err != nil {
		defer sqlDB.Close()
	}

	executePendingMigrations(db)

	// Migrate rest of the models
	log.Info().Msg("AutoMigrate Model [table_name]")
	db.AutoMigrate(&entity.Credential{})
	log.Info().Msg("  TableModel [" + (&entity.Credential{}).TableName() + "]")
	db.AutoMigrate(&entity.Profile{})
	log.Info().Msg("  TableModel [" + (&entity.Profile{}).TableName() + "]")
	db.AutoMigrate(&entity.Game{})
	log.Info().Msg("  TableModel [" + (&entity.Game{}).TableName() + "]")
	db.AutoMigrate(&entity.Voucher{})
	log.Info().Msg("  TableModel [" + (&entity.Voucher{}).TableName() + "]")
	db.AutoMigrate(&entity.Verification{})
	log.Info().Msg("  TableModel [" + (&entity.Verification{}).TableName() + "]")
	db.AutoMigrate(&entity.GiftCard{})
	log.Info().Msg("  TableModel [" + (&entity.GiftCard{}).TableName() + "]")
	db.AutoMigrate(&entity.Review{})
	log.Info().Msg("  TableModel [" + (&entity.Review{}).TableName() + "]")
	db.AutoMigrate(&entity.SuperAdmin{})
	log.Info().Msg("  TableModel [" + (&entity.SuperAdmin{}).TableName() + "]")
	db.AutoMigrate(&entity.Shop{})
	log.Info().Msg("  TableModel [" + (&entity.Shop{}).TableName() + "]")
	db.AutoMigrate(&entity.Genre{})
	log.Info().Msg("  TableModel [" + (&entity.Genre{}).TableName() + "]")
	db.AutoMigrate(&entity.Tags{})
	log.Info().Msg("  TableModel [" + (&entity.Tags{}).TableName() + "]")
	db.AutoMigrate(&entity.Tags_detail{})
	log.Info().Msg("  TableModel [" + (&entity.Tags_detail{}).TableName() + "]")
	db.AutoMigrate(&entity.File{})
	log.Info().Msg("  TableModel [" + (&entity.File{}).TableName() + "]")
	db.AutoMigrate(&entity.Wishlist{})
	log.Info().Msg("  TableModel [" + (&entity.Wishlist{}).TableName() + "]")

	// db.AutoMigrate(&entity.Users{})
	// log.Info().Msg("  TableModel [" + (&entity.Users{}).TableName() + "]")

}

func executePendingMigrations(db *gorm.DB) {
	db.AutoMigrate(&MigrationHistoryModel{})
	lastMigration := MigrationHistoryModel{}
	skipMigration := db.Order("migration_id desc").Limit(1).Find(&lastMigration).RowsAffected > 0

	// skip to last migration
	keys := make([]string, 0, len(migrations))
	for k := range migrations {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// run all migrations in one transaction
	if len(migrations) == 0 {
		log.Info().Msg("No pending migrations")
	} else {
		db.Transaction(func(tx *gorm.DB) error {
			for _, k := range keys {
				if skipMigration {
					if k == lastMigration.MigrationID {
						skipMigration = false
					}
				} else {
					log.Info().Msg("  " + k)
					tx.Transaction(func(subTx *gorm.DB) error {
						// run migration update
						checkError(migrations[k](subTx))
						// insert migration id into history
						checkError(subTx.Create(MigrationHistoryModel{MigrationID: k}).Error)
						return nil
					})
				}
			}
			return nil
		})
	}
}

type mFunc func(tx *gorm.DB) error

var migrations = make(map[string]mFunc)

// MigrationHistoryModel model migration
type MigrationHistoryModel struct {
	MigrationID string `gorm:"type:text;primaryKey"`
}

// TableName name of migration table
func (model *MigrationHistoryModel) TableName() string {
	return "migration_history"
}

func checkError(err error) {
	if err != nil {
		log.Fatal().Err(err)
		panic(err)
	}
}

// func registerMigration(id string, fm mFunc) {
// 	migrations[id] = fm
// }
