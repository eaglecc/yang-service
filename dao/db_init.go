package dao

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"yang-service/commons/constparam"
	"yang-service/config"
	"yang-service/entity"
)

var _db *gorm.DB

func ConnectMySQLDB() {
	// 1. 准备数据库的参数
	username := config.Conf.Db.UserName
	password := config.Conf.Db.PassWord
	host := config.Conf.Db.Host
	port := config.Conf.Db.Port
	dbname := config.Conf.Db.DBName
	myConf := config.Conf.Db.Myconfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", username, password, host, port, dbname, myConf)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("Asia/Shanghai")
			return time.Now().In(loc)
		},
	})

	if err != nil {
		zap.S().Error("数据库连接失败")
		panic(err)
	}
	// 2. 迁移结构体schema
	err = db.AutoMigrate(
		&entity.AngryHistoryRecord{})
	if err != nil {
		zap.S().Error("迁移结构体至数据库失败")
	}
	_db = db
	constparam.YANG_DB = db
	zap.S().Info("Connect MySQL Success...")
	return
}

//func NewDBClient(ctx context.Context) *gorm.DB {
//	db := _db
//	return db.WithContext(ctx)
//}
