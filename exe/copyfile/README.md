# 并发Copy命令  (Same as Linux cp)

## 如何安装

```
go get github.com/relunctance/goutils
go build $GOPATH/src/github.com/relunctance/goutils/exe/copyfile/copy.go
```

## 查看帮助：
```
./copy --help
Usage of ./copy:
  -cover
    	is cover same file ,if false and filesize is equal then will not cover same file (default false)
  -input string
    	the copy input dir , the type should be dir
  -maxnum int
    	the num size  of channel , max limit is 100 (default 20)
  -output string
    	the copy input dir, the type should be dir
  -version string
    	copy version (default "1.0.0.1001")
```

## 使用示例：
```
./copy -input "you input dir" -output "you output dir" 
```
