package initialize

import (
	"<%= displayName %>/global"
	"<%= displayName %>/model"
	"<%= displayName %>/source"
	"<%= displayName %>/utils"
	internal "<%= displayName %>/utils"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB
//@return: *gorm.DB

func Gorm() *gorm.DB {
	//switch global.CONFIG.System.DbType {
	//default:
	return GormPostgres()
	//}
}

//@function: Tables
//@description: 注册数据库表专用
//@param: db *gorm.DB

func PostgresTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.Role{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
		model.ServerConfig{},
		model.Button{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
	// init data
	err = utils.InitDB(
		source.Role,
		source.Admin,
		source.BaseMenu,
		source.RolesMenus,
		source.Config,
		source.Button,
	)
	if err != nil {
		global.LOG.Error("init data failed", zap.Any("err", err))
		os.Exit(0)
	}
}

//
//@function: GormPostgres
//@description: 初始化pg数据库
//@return: *gorm.DB

func GormPostgres() *gorm.DB {
	m := global.CONFIG.Postgres
	dsn := m.Dsn()
	if db, err := gorm.Open(postgres.Open(dsn), gormConfig(m.LogMode)); err != nil {
		global.LOG.Error("Postgres启动异常", zap.Any("err", err))
		os.Exit(0)
		return nil
	} else {
		global.LOG.Info("数据库连接成功！")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.CONFIG.Postgres.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}
