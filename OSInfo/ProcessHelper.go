package OSInfo

import (
	"strings"

	ps "github.com/keybase/go-ps"
)

type ProcessHelper struct{}

func NewProcessHelper() *ProcessHelper {
	helper := new(ProcessHelper)

	return helper
}

func (this *ProcessHelper) ProcessList() []ps.Process {
	processArrary, err := ps.Processes()

	if err != nil {
		return nil
	}

	return processArrary
}

func (this *ProcessHelper) IsExist(processName string) bool {
	processArrary := this.ProcessList()

	if processArrary == nil || len(processArrary) <= 0 {
		return false
	}

	var isExist bool = false

	for _, p := range processArrary {
		descr := p.Executable()
		if strings.Contains(descr, processName) {
			isExist = true
			break
		}

	}

	return isExist
}
