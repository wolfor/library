package FileHelper

import (
	"github.com/widuu/goini"
)

type INIHelper struct {
	conf *goini.Config
}

func NewINIHelper(iniFilePath string) *INIHelper {
	helper := new(INIHelper)
	helper.conf = goini.SetConfig(iniFilePath)

	return helper
}

//获取指定section下key的值
func (this *INIHelper) GetValue(section, keyName string) string {
	return this.conf.GetValue(section, keyName)
}

//设置指定section下的key
func (this *INIHelper) SetValue(section, keyName, val string) bool {
	return this.conf.SetValue(section, keyName, val)
}

//删除指定section下的key
func (this *INIHelper) DeleteValue(section, keyName string) bool {
	return this.conf.DeleteValue(section, keyName)
}

//读取ini列表
func (this *INIHelper) ReadList() []map[string]map[string]string {
	return this.conf.ReadList()
}
