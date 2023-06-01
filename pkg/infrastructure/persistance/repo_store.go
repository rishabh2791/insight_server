package persistance

import (
	"insight/pkg/domain/entity"
	"insight/pkg/domain/repository"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepoStore struct {
	DB             *gorm.DB
	VesselRepo     repository.VesselRepository
	DeviceTypeRepo repository.DeviceTypeRepository
	DeviceRepo     repository.DeviceRepository
	DeviceDataRepo repository.DeviceDataRepository
}

func NewRepoStore() (*RepoStore, error) {
	repoStore := RepoStore{}

	mysqlURL := "administrator:canway#123@tcp(10.19.0.205:3306)/insight?parseTime=True"

	gormDB, gormErr := gorm.Open(mysql.Open(mysqlURL), &gorm.Config{
		Logger:               logger.Default.LogMode(logger.Silent),
		QueryFields:          true,
		FullSaveAssociations: true,
	})
	if gormErr != nil {
		return nil, gormErr
	}
	sqlDB, _ := gormDB.DB()
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(10000)

	repoStore.DB = gormDB
	repoStore.VesselRepo = NewVesselRepo(gormDB)
	repoStore.DeviceTypeRepo = NewDeviceTypeRepo(gormDB)
	repoStore.DeviceRepo = NewDeviceRepo(gormDB)
	repoStore.DeviceDataRepo = NewDeviceDataRepo(gormDB)

	return &repoStore, nil
}

func (repoStore *RepoStore) MigrateModels() error {
	return repoStore.DB.AutoMigrate(
		&entity.Vessel{},
		&entity.DeviceType{},
		&entity.Device{},
		&entity.DeviceData{},
	)
}
