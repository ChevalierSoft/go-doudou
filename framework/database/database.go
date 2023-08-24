package database

import (
	"github.com/unionj-cloud/go-doudou/v2/framework/config"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/cast"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/errorx"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/stringutils"
	"github.com/unionj-cloud/go-doudou/v2/toolkit/zlogger"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

const (
	DriverMysql      = "mysql"
	DriverPostgres   = "postgres"
	DriverSqlite     = "sqlite"
	DriverSqlserver  = "sqlserver"
	DriverTidb       = "tidb"
	DriverClickhouse = "clickhouse"
)

var Db *gorm.DB

func init() {
	if cast.ToBoolOrDefault(config.GddDBDisableAutoConfigure.Load(), config.DefaultGddDBDisableAutoConfigure) {
		return
	}
	slowThreshold, err := time.ParseDuration(config.GddDBLogSlowThreshold.Load())
	if err != nil {
		zlogger.Debug().Msgf("Parse %s %s as time.Duration failed: %s, use default %s instead.\n", string(config.GddDBLogSlowThreshold), config.GddDBLogSlowThreshold.Load(), err.Error(), config.DefaultGddDBLogSlowThreshold)
		slowThreshold, _ = time.ParseDuration(config.DefaultGddDBLogSlowThreshold)
	}
	logLevel := config.DefaultGddDBLogLevel
	if stringutils.IsNotEmpty(config.GddDBLogLevel.Load()) {
		switch strings.ToLower(config.GddDBLogLevel.Load()) {
		case "silent":
			logLevel = logger.Silent
		case "error":
			logLevel = logger.Error
		case "warn":
			logLevel = logger.Warn
		case "info":
			logLevel = logger.Info
		}
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             slowThreshold,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: cast.ToBoolOrDefault(config.GddDBLogIgnoreRecordNotFoundError.Load(), config.DefaultGddDBLogIgnoreRecordNotFoundError),
			ParameterizedQueries:      cast.ToBoolOrDefault(config.GddDBLogParameterizedQueries.Load(), config.DefaultGddDBLogParameterizedQueries),
			Colorful:                  false,
		},
	)
	gormConf := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	tablePrefix := strings.TrimSuffix(config.GddDBTablePrefix.LoadOrDefault(config.DefaultGddDBTablePrefix), ".")
	if stringutils.IsNotEmpty(tablePrefix) {
		gormConf.NamingStrategy = schema.NamingStrategy{
			TablePrefix: tablePrefix + ".",
		}
	}
	dsn := config.GddDBDsn.Load()
	if stringutils.IsEmpty(dsn) {
		return
	}
	driver := config.GddDBDriver.Load()
	if stringutils.IsEmpty(driver) {
		errorx.Panic("Database driver is missing")
	}
	switch driver {
	case DriverMysql, DriverTidb:
		conf := mysql.Config{
			DSN:                           dsn, // data source name
			SkipInitializeWithVersion:     cast.ToBoolOrDefault(config.GddDBMysqlSkipInitializeWithVersion.Load(), config.DefaultGddDBMysqlSkipInitializeWithVersion),
			DefaultStringSize:             uint(cast.ToIntOrDefault(config.GddDBMysqlDefaultStringSize.Load(), config.DefaultGddDBMysqlDefaultStringSize)),
			DisableWithReturning:          cast.ToBoolOrDefault(config.GddDBMysqlDisableWithReturning.Load(), config.DefaultGddDBMysqlDisableWithReturning),
			DisableDatetimePrecision:      cast.ToBoolOrDefault(config.GddDBMysqlDisableDatetimePrecision.Load(), config.DefaultGddDBMysqlDisableDatetimePrecision),
			DontSupportRenameIndex:        cast.ToBoolOrDefault(config.GddDBMysqlDontSupportRenameIndex.Load(), config.DefaultGddDBMysqlDontSupportRenameIndex),
			DontSupportRenameColumn:       cast.ToBoolOrDefault(config.GddDBMysqlDontSupportRenameColumn.Load(), config.DefaultGddDBMysqlDontSupportRenameColumn),
			DontSupportForShareClause:     cast.ToBoolOrDefault(config.GddDBMysqlDontSupportForShareClause.Load(), config.DefaultGddDBMysqlDontSupportForShareClause),
			DontSupportNullAsDefaultValue: cast.ToBoolOrDefault(config.GddDBMysqlDontSupportNullAsDefaultValue.Load(), config.DefaultGddDBMysqlDontSupportNullAsDefaultValue),
			DontSupportRenameColumnUnique: cast.ToBoolOrDefault(config.GddDBMysqlDontSupportRenameColumnUnique.Load(), config.DefaultGddDBMysqlDontSupportRenameColumnUnique),
		}
		Db, err = gorm.Open(mysql.New(conf), gormConf)
	case DriverPostgres:
		conf := postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: cast.ToBoolOrDefault(config.GddDBPostgresPreferSimpleProtocol.Load(), config.DefaultGddDBPostgresPreferSimpleProtocol),
			WithoutReturning:     cast.ToBoolOrDefault(config.GddDBPostgresWithoutReturning.Load(), config.DefaultGddDBPostgresWithoutReturning),
		}
		Db, err = gorm.Open(postgres.New(conf), gormConf)
	case DriverSqlite:
		Db, err = gorm.Open(sqlite.Open(dsn), gormConf)
	case DriverSqlserver:
		Db, err = gorm.Open(sqlserver.Open(dsn), gormConf)
	case DriverClickhouse:
		Db, err = gorm.Open(clickhouse.Open(dsn), gormConf)
	default:
		errorx.Panic("Not support driver")
	}
	if err != nil {
		errorx.Panic(err.Error())
	}
	sqlDB, err := Db.DB()
	if err != nil {
		errorx.Panic(err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(cast.ToIntOrDefault(config.GddDBMaxIdleConns.Load(), config.DefaultGddDBMaxIdleConns))

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(cast.ToIntOrDefault(config.GddDBMaxOpenConns.Load(), config.DefaultGddDBMaxOpenConns))

	maxLifetime, err := time.ParseDuration(config.GddDBConnMaxLifetime.Load())
	if err != nil {
		zlogger.Debug().Msgf("Parse %s %s as time.Duration failed: %s, use default %d instead.\n", string(config.GddDBConnMaxLifetime), config.GddDBConnMaxLifetime.Load(), err.Error(), config.DefaultGddDBConnMaxLifetime)
		maxLifetime = config.DefaultGddDBConnMaxLifetime
	}
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(maxLifetime)

	maxIdleTime, err := time.ParseDuration(config.GddDBConnMaxIdleTime.Load())
	if err != nil {
		zlogger.Debug().Msgf("Parse %s %s as time.Duration failed: %s, use default %d instead.\n", string(config.GddDBConnMaxIdleTime), config.GddDBConnMaxIdleTime.Load(), err.Error(), config.DefaultGddDBConnMaxIdleTime)
		maxIdleTime = config.DefaultGddDBConnMaxIdleTime
	}
	sqlDB.SetConnMaxIdleTime(maxIdleTime)
}

func NewDb(conf config.Config) (db *gorm.DB) {
	var err error
	slowThreshold, _ := time.ParseDuration("200ms")
	if stringutils.IsNotEmpty(conf.Db.Log.SlowThreshold) {
		if value, err := time.ParseDuration(conf.Db.Log.SlowThreshold); err == nil {
			slowThreshold = value
		} else {
			zlogger.Error().Err(err).Msg(err.Error())
		}
	}
	logLevel := logger.Warn
	if stringutils.IsNotEmpty(conf.Db.Log.Level) {
		switch strings.ToLower(conf.Db.Log.Level) {
		case "silent":
			logLevel = logger.Silent
		case "error":
			logLevel = logger.Error
		case "warn":
			logLevel = logger.Warn
		case "info":
			logLevel = logger.Info
		}
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             slowThreshold,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: conf.Db.Log.IgnoreRecordNotFoundError,
			ParameterizedQueries:      conf.Db.Log.ParameterizedQueries,
			Colorful:                  false,
		},
	)
	gormConf := &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	tablePrefix := strings.TrimSuffix(conf.Db.Table.Prefix, ".")
	if stringutils.IsNotEmpty(tablePrefix) {
		gormConf.NamingStrategy = schema.NamingStrategy{
			TablePrefix: tablePrefix + ".",
		}
	}
	dsn := conf.Db.Dsn
	if stringutils.IsEmpty(dsn) {
		return
	}
	driver := conf.Db.Driver
	if stringutils.IsEmpty(driver) {
		errorx.Panic("Database driver is missing")
	}
	switch driver {
	case DriverMysql, DriverTidb:
		conf := mysql.Config{
			DSN:                           dsn, // data source name
			SkipInitializeWithVersion:     conf.Db.Mysql.SkipInitializeWithVersion,
			DefaultStringSize:             uint(conf.Db.Mysql.DefaultStringSize),
			DisableWithReturning:          conf.Db.Mysql.DisableWithReturning,
			DisableDatetimePrecision:      conf.Db.Mysql.DisableDatetimePrecision,
			DontSupportRenameIndex:        conf.Db.Mysql.DontSupportRenameIndex,
			DontSupportRenameColumn:       conf.Db.Mysql.DontSupportRenameColumn,
			DontSupportForShareClause:     conf.Db.Mysql.DontSupportForShareClause,
			DontSupportNullAsDefaultValue: conf.Db.Mysql.DontSupportNullAsDefaultValue,
			DontSupportRenameColumnUnique: conf.Db.Mysql.DontSupportRenameColumnUnique,
		}
		db, err = gorm.Open(mysql.New(conf), gormConf)
	case DriverPostgres:
		conf := postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: conf.Db.Postgres.PreferSimpleProtocol,
			WithoutReturning:     conf.Db.Postgres.WithoutReturning,
		}
		db, err = gorm.Open(postgres.New(conf), gormConf)
	case DriverSqlite:
		db, err = gorm.Open(sqlite.Open(dsn), gormConf)
	case DriverSqlserver:
		db, err = gorm.Open(sqlserver.Open(dsn), gormConf)
	case DriverClickhouse:
		db, err = gorm.Open(clickhouse.Open(dsn), gormConf)
	default:
		errorx.Panic("Not support driver")
	}
	if err != nil {
		errorx.Panic(err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		errorx.Panic(err.Error())
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(conf.Db.Pool.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(conf.Db.Pool.MaxOpenConns)

	maxLifetime := time.Duration(-1)
	if stringutils.IsNotEmpty(conf.Db.Pool.ConnMaxLifetime) {
		if value, err := time.ParseDuration(conf.Db.Pool.ConnMaxLifetime); err == nil {
			maxLifetime = value
		} else {
			zlogger.Error().Err(err).Msg(err.Error())
		}
	}
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(maxLifetime)

	maxIdleTime := time.Duration(-1)
	if stringutils.IsNotEmpty(conf.Db.Pool.ConnMaxIdleTime) {
		if value, err := time.ParseDuration(conf.Db.Pool.ConnMaxIdleTime); err == nil {
			maxIdleTime = value
		} else {
			zlogger.Error().Err(err).Msg(err.Error())
		}
	}
	sqlDB.SetConnMaxIdleTime(maxIdleTime)
	return
}
