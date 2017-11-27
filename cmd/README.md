## installation

go get github.com/relunctance/goutils/cmd

## 示例1:
```go
cmd.RunCommand("ls /home/")
```
## 示例2:
```go
cmd.Debug = true
RunCommand("ps auxwww | grep init | grep -v grep")
```
## 示例3:
```go

//获取当前执行脚本的pid
cmdCode := `ps aux  | grep 'cmd/_test' | grep -v "grep" | grep -v 'go run' | awk '{print $2}'`
cmdResult , err := RunCommand(cmdCode)
if err != nil {
    panic(err)
}
fmt.Println(cmdResult);
```
