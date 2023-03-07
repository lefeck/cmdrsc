# cmdrsc

this is a simple command Framework lib.

## useage

```go
go get  github.com/lefeck/cmdrsc
```


## for example 
```go
import (
	"fmt"
	"github.com/lefeck/cmdrsc"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	entry := logrus.NewEntry(log)
	e := cdmrsc.NewExecutor(entry, 1)

	parameter := ""
	cmdTmpl := fmt.Sprintf(" ls -l")
	stdout, _, err := e.RunCmd(e.RefreshCmd(cmdTmpl), parameter)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stdout)
}

```