package service

import (
	"<%= displayName %>/global"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(global.DB)
	e, _ := casbin.NewEnforcer(global.CONFIG.Casbin.ModelPath, a)
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = e.LoadPolicy()
	return e
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	key1 := strings.Split(name1, "?")[0]
	return util.KeyMatch2(key1, name2), nil
}

func GetCasbinRuleList() []gormadapter.CasbinRule {
	// e := Casbin()
	// list := e.GetPolicy()
	var casbin_rule []gormadapter.CasbinRule
	global.DB.Find(&casbin_rule)
	return casbin_rule
}

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

func UpdateCasbin(oldPolicy []string, newPolicy []string) (bool, error) {
	e := Casbin()
	rst, err := e.UpdatePolicy(oldPolicy, newPolicy)
	return rst, err
}

func UpdateCasbinById(id uint, value0 string, value2 string, value3 string) error {
	return global.DB.Model(&gormadapter.CasbinRule{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{"V0": value0, "V2": value2, "V3": value3}).
		Error
}
