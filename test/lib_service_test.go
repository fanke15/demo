package test

import (
	"cron/internal/lib/service"
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestInitConfig(t *testing.T) {

	conf := service.GetConfig()

	fmt.Println(1111, conf.Mod, conf.Port, conf.Version)

	fmt.Println(2222, conf.Mysqls["default"])
}

func TestInitLog(t *testing.T) {
	type Data struct {
		ID   int             `json:"id"`
		Type int             `json:"type"`
		Name string          `json:"name"`
		Val  decimal.Decimal `json:"val"`
	}
	d := Data{1, 8, "dataInfo", decimal.NewFromFloat(6.666666)}

	service.GetLog().Info("操作成功！", "name", "zkf")
	service.GetLog().Error("操作失败！", "name", "zkf", "data", d, "other", map[string]interface{}{
		"my":   "zkf",
		"test": 8888,
	})
}

type Acc struct {
	ID                uint   `gorm:"column:id;primary_key" json:"id"` // 自增Id主键
	UserAddress       string `gorm:"column:user_address;" json:"user_address"`
	UserAccount       string `gorm:"column:user_account;" json:"user_account"`
	Strategy          string `gorm:"column:strategy;" json:"strategy"` // email
	Email             string `gorm:"column:email;" json:"email"`       // email
	EffectiveLast     int64  `gorm:"column:effective_last;" json:"effective_last"`
	EffectiveStart    int64  `gorm:"column:effective_start;" json:"effective_start"`
	EffectiveDuration int64  `gorm:"column:effective_duration;" json:"effective_duration"`
}

func (m *Acc) TableName() string {
	return "onepiece_account"
}

// 测试mysql
func TestConnMysql(t *testing.T) {

	service.GetMysql().Model(&Acc{}).Where("id > ?", 10000).Update("email", "hello@qq.com")

}
