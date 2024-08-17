package mysql

import (
	"time"

	"github.com/me2seeks/cola/config"
	"github.com/me2seeks/cola/internal/models"
	"github.com/me2seeks/cola/internal/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() *gorm.DB {
	logger := logger.Logger

	Client, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.Cfg.MySQL.Dsn(), // DSN data source name
		DefaultStringSize:         256,                    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                  // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// 	Logger: logger,
	})
	if err != nil {
		logger.Fatalf("failed to connect to mysql: %v", err)
		return nil
	}

	// 连接池
	sqlDB, err := Client.DB()
	if err != nil {
		logger.Fatalf("failed to get db: %v", err)
		return nil
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Migrate the schema
	err = Client.AutoMigrate(models.User{}, models.UserAuth{})
	if err != nil {
		logger.Fatalf("failed to migrate schema: %v", err)
		return nil
	}

	return Client
}
