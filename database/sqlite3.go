package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

// 计数器模型
type Counter struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"` // 标签名
	Count uint   // 访问计数
}

// 初始化数据库连接
func InitDB(dbFile string, debug bool) error {
	var err error
	config := &gorm.Config{}
	if debug {
		// 调试模式下显示详细日志
		config.Logger = gormlogger.Default.LogMode(gormlogger.Info)
	} else {
		// 非调试模式禁用日志
		config.Logger = gormlogger.Default.LogMode(gormlogger.Silent)
	}
	DB, err = gorm.Open(sqlite.Open(dbFile), config)
	if err != nil {
		return err
	}

	// 自动迁移表结构
	return DB.AutoMigrate(&Counter{})
}

// 增加计数并返回当前值
func IncrementCounter(name string) (uint, error) {
	var counter Counter
	result := DB.Where("name = ?", name).First(&counter)

	if result.Error == gorm.ErrRecordNotFound {
		// 新创建计数器
		counter = Counter{Name: name, Count: 1}
		if err := DB.Create(&counter).Error; err != nil {
			return 0, err
		}
		return 1, nil
	} else if result.Error != nil {
		return 0, result.Error
	}

	// 更新计数器
	counter.Count++
	if err := DB.Save(&counter).Error; err != nil {
		return 0, err
	}
	return counter.Count, nil
}

// 查询计数器的当前值（不增加计数）
func GetCount(name string) (uint, error) {
	var counter Counter
	result := DB.Where("name = ?", name).First(&counter)

	if result.Error == gorm.ErrRecordNotFound {
		return 0, nil
	} else if result.Error != nil {
		return 0, result.Error
	}
	return counter.Count, nil
}
