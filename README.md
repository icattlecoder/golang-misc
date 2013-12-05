golang-misc
===========

## progressBar
golang命令行进度提示条.
    
``` go
import "progressbar"
pbar:=progressbar.New("label")
pbar.SetVal(10)
```

显示效果:

```
10%[===========                                             ]label
```
