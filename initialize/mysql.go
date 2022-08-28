/*
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 16:46:45
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-22 09:43:18
 * @FilePath: /gin_micro_service/initialize/mysql.go
 * @Description: mysql init,gorm
 */

package initialize

import (
	"fmt"
	"gin_micro_service/global"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
// Author [SliverHorn](https://github.com/SliverHorn)
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
// Author [SliverHorn](https://github.com/SliverHorn)
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		logZap = global.GVA_CONFIG.Mysql.LogZap
		if logZap {
			global.GVA_LOG.Info(fmt.Sprintf(message+"\n", data...))
		} else {
			w.Writer.Printf(message, data...)
		}
	}

}

//-----------------gorm mysql init-------------------------------------
var Gorm = new(gormMysql)

type gormMysql struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *gormMysql) configs() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	_default := logger.New(NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	switch global.GVA_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}

//dont use this function
func (g *gormMysql) openMysql(dbType string) *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	if dbType == "order_center" {
		//启动order_center 数据库连接
		mysqlConfigOrder := mysql.Config{
			DSN:                       m.DsnOrder(), // DSN data source name
			DefaultStringSize:         191,          // string 类型字段的默认长度
			SkipInitializeWithVersion: false,        // 根据版本自动配置，
		}
		if db1, err1 := gorm.Open(mysql.New(mysqlConfigOrder), g.configs()); err1 != nil {
			global.GVA_LOG.Error("init order_center db faild.")
			return nil
		} else {
			sqlDB, _ := db1.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db1
		}
	} else if dbType == "customer_center" {
		//customer_center 数据库连接
		mysqlConfigOrder := mysql.Config{
			DSN:                       m.DsnCustomer(), // DSN data source name
			DefaultStringSize:         191,             // string 类型字段的默认长度
			SkipInitializeWithVersion: false,           // 根据版本自动配置，
		}
		if db1, err1 := gorm.Open(mysql.New(mysqlConfigOrder), g.configs()); err1 != nil {
			global.GVA_LOG.Error("init customer_center db faild.")
			return nil
		} else {
			sqlDB, _ := db1.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db1
		}
	} else {
		//启动 finance_center数据库连接
		mysqlConfig := mysql.Config{
			DSN:                       m.Dsn(), // DSN data source name
			DefaultStringSize:         191,     // string 类型字段的默认长度
			SkipInitializeWithVersion: false,   // 根据版本自动配置
		}
		if db, err := gorm.Open(mysql.New(mysqlConfig), g.configs()); err != nil {
			global.GVA_LOG.Error("init fiance_center db faild.")
			return nil
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db
		}
	}
}

//dont use this function.
func (g *gormMysql) InitMysql() {
	global.GVA_DB = g.openMysql("finance_center")
	if global.GVA_DB != nil {
		global.GVA_LOG.Debug("init finance_center db done.")
		//initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	//customer 库
	global.GVA_DB_CUSTOMER = g.openMysql("customer_center")
	if global.GVA_DB_CUSTOMER != nil {
		global.GVA_LOG.Info("init customer_center db done.")
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB_CUSTOMER.DB()
		defer db.Close()
	}

	global.GVA_DB_ORDER = g.openMysql("order_center") // gorm连接数据库
	if global.GVA_DB != nil {
		global.GVA_LOG.Debug("dont init order_center db done.")
		//initialize.RegisterTables(global.GVA_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

}

func OpenMysql(dbType string) *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	if dbType == "order_center" {
		//启动order_center 数据库连接
		mysqlConfigOrder := mysql.Config{
			DSN:                       m.DsnOrder(), // DSN data source name
			DefaultStringSize:         191,          // string 类型字段的默认长度
			SkipInitializeWithVersion: false,        // 根据版本自动配置，
		}
		if db1, err1 := gorm.Open(mysql.New(mysqlConfigOrder), Gorm.configs()); err1 != nil {
			global.GVA_LOG.Error("init order_center db faild.")
			return nil
		} else {
			sqlDB, _ := db1.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db1
		}
	} else if dbType == "customer_center" {
		//customer_center 数据库连接
		mysqlConfigOrder := mysql.Config{
			DSN:                       m.DsnCustomer(), // DSN data source name
			DefaultStringSize:         191,             // string 类型字段的默认长度
			SkipInitializeWithVersion: false,           // 根据版本自动配置，
		}
		if db1, err1 := gorm.Open(mysql.New(mysqlConfigOrder), Gorm.configs()); err1 != nil {
			global.GVA_LOG.Error("init customer_center db faild.")
			return nil
		} else {
			sqlDB, _ := db1.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db1
		}
	} else {
		//启动 finance_center数据库连接
		mysqlConfig := mysql.Config{
			DSN:                       m.Dsn(), // DSN data source name
			DefaultStringSize:         191,     // string 类型字段的默认长度
			SkipInitializeWithVersion: false,   // 根据版本自动配置
		}
		if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.configs()); err != nil {
			global.GVA_LOG.Error("init fiance_center db faild.")
			return nil
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(m.MaxIdleConns)
			sqlDB.SetMaxOpenConns(m.MaxOpenConns)
			return db
		}
	}
}
