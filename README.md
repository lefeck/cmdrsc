# cmdrsc

this is a simple command Framework lib.

## useage method

```go
go get  github.com/lefeck/cmdrsc
```


## for example 
```go
func main() {
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
```