package OSInfo

import (
	"os"
)

type EnvironmentVariable struct{}

func NewEnvironmentVariable() *EnvironmentVariable {
	evnv := new(EnvironmentVariable)

	return evnv
}

//获取所有环境变量
func (this *EnvironmentVariable) GetEnvironmentVariableList() []string {
	evnvList := os.Environ()

	return evnvList
}

//获取指定环境变量
func (this *EnvironmentVariable) GetEnvironmentVariable(evnvName string) string {
	evnvValue := os.Getenv(evnvName)

	return evnvValue
}
