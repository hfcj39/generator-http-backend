package initialize

import (
	"<%= displayName %>/global"
	"<%= displayName %>/service"
	"<%= displayName %>/source"
	"reflect"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() error {
	existRule := []gormadapter.CasbinRule{}
	var e = service.Casbin()
	var updateMap = make(map[string]string)
	global.DB.Find(&existRule)
	for _, v := range existRule {
		for i, r := range source.Casbin_rules {
			if reflect.DeepEqual(v.V1, r[1]) {
				source.Casbin_rules = append(source.Casbin_rules[:i], source.Casbin_rules[i+1:]...)
				if len(r) > 3 && !reflect.DeepEqual(v.V3, r[3]) {
					updateMap[v.V1] = r[3]
				}
			}
		}
	}

	// 添加原先没有的规则
	_, err := e.AddPolicies(source.Casbin_rules)
	if err != nil {
		global.LOG.Error(err.Error())
	}
	err = e.SavePolicy()
	if err != nil {
		global.LOG.Error(err.Error())
		return err
	}

	// 手动更新button,自带adaptor不知道为什么更新失败,先手动更新代替
	for key := range updateMap {
		err := global.DB.Table("casbin_rule").Where("v1 = ?", key).Update("v3", updateMap[key]).Error
		if err != nil {
			global.LOG.Error(err.Error())
		}
	}
	err = global.DB.Table("casbin_rule").Where("v3 = ?", "").Update("v3", "-").Error
	if err != nil {
		global.LOG.Error(err.Error())
	}
	global.LOG.Info("初始化casbin数据成功")
	return nil
}
