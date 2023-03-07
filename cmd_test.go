package cmdrsc

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCmd(t *testing.T) {
	log := logrus.New()
	entry := logrus.NewEntry(log)

	n := NewExecutor(entry, 1)
	parameter := ""
	cmdTmpl := fmt.Sprintf(" ls")
	stdout, _, err := n.RunCmd(n.RefreshCmd(cmdTmpl), parameter)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout)
}
