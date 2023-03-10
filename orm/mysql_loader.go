package orm

import (
	"time"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 调用mysql存储过程加载K线数据
type MysqlLoader struct {
	OrmDB *gorm.DB
}

func (loader *MysqlLoader) Init(dataSource string) error {
	var err error
	loader.OrmDB, err = gorm.Open(mysql.Open(dataSource),
		&gorm.Config{Logger: logger.New(log.New("-"), logger.Config{LogLevel: logger.Warn})})
	if err != nil {
		log.Errorf("Open mysql failed,dataSource:%s, err:%s\n", dataSource, err)
		return err
	}
	loader.OrmDB = loader.OrmDB.Session(&gorm.Session{CreateBatchSize: 1000})
	db, _ := loader.OrmDB.DB()
	db.SetMaxIdleConns(10) //空闲连接数
	db.SetMaxOpenConns(50) //最大连接数
	db.SetConnMaxLifetime(time.Minute * time.Duration(30))
	return nil
}
