package service

import (
	"auth/auth/dto"
	"auth/auth/model"
	"auth/global"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"strings"
)


func UpdateAuth(authorityId string, authInfo []dto.AuthInfo) error {
	ClearCasbin(0, authorityId)
	var rules [][]string
	for _, v := range authInfo {
		cm := model.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	enforcer := GetEnforcer()
	success, _ := enforcer.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []dto.AuthInfo) {
	enforcer := GetEnforcer()
	list := enforcer.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, dto.AuthInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

func ClearCasbin(v int, p ...string) bool {
	enforcer := GetEnforcer()
	success, _ := enforcer.RemoveFilteredPolicy(v, p...)
	return success

}

func GetEnforcer() *casbin.Enforcer {
	admin := global.CONFIG.Mysql
	a, err := gormAdapter.NewAdapter(global.CONFIG.System.DbType, admin.Username+":"+admin.Password+"@("+admin.Path+")/"+admin.Dbname, true)
	if err != nil {
		global.LOG.Error("cas链接数据库错误",zap.Any("cas :",err))
		return nil
	}
	enforcer, _ := casbin.NewEnforcer(global.CONFIG.Auth.ModelPath, a)
	enforcer.AddFunction("ParamsMatch", ParamsMatchFunc)
	_ = enforcer.LoadPolicy()
	return enforcer
}

func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}

func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}