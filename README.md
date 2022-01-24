# 防止手滑小软件

### 使用方法

```text
Usage: main.exe --exec-path=STRING

Flags:
  -h, --help                Show context-sensitive help.
      --exec-path=STRING    要防止手滑的可执行文件路径
      --times=1             启动之前的确认次数，默认为 1
      --args=ARGS,...       要传递给可执行文件的参数
```


### 注意事项

- 暂时没有在 macOS、Linux上测试过
- 可执行文件参数传递可能不正常，没有完整测试
- 不要用防止手滑来启动防止手滑...
