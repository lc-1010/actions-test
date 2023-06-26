package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	// id
	ID uint32 `gorm:"primary_key" json:"id"`
 
	CratedOn uint32 `json:"crated_on"`
	 
	CratedBy string `json:"crated_by"`
 
	ModifiedOn uint32 `json:"modified_on"`
	 
	ModifiedBy string `json:"modified_by"`
	 
	DeletedOn uint32 `json:"deleted_on"`
	 
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(dbsetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbsetting.UserName,
		dbsetting.Password,
		dbsetting.Host,
		dbsetting.DBName,
		dbsetting.Charset,
		dbsetting.ParseTime)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.Logger.LogMode(logger.Info)
	}
	
	DB, err := db.DB()
	if err != nil {
		return nil, err
	}

	DB.SetMaxIdleConns(100)
	DB.SetMaxOpenConns(100)
	return db, nil
}
