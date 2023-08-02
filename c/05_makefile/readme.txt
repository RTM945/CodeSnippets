### 变量

a=b

$(AR)=ar
$(CC)=cc
$(CXX)=g++
$@=目标全称
$<=第一个依赖
$^=所有依赖

### 模式匹配

% 通配符匹配一个字符串
%.o:%.c 意为name相同 如 a.o:a.c

%.o:%.c
    $(CC) -c $< -o $@

意为
目标 xx.o 依赖 xx.c
操作为 cc -c xx.c -o xx.o
可以自动匹配所有文件

### 函数

$(wildcard PATTERN...)
获取指定目录下符合PATTERN的文件列表 
多个PATTERN用空格隔开
$(wildcard *.c ./sub/*.c)

$(patsubst <pattern>,<replacement>,<text>)
在text中找到匹配pattern的字符用replacement替换
返回替换后的值
多个text用空格 tab 换行分割

$(patsubst %.c,%.o,a.c b.c)
返回 a.o b.o

### tips

Makefile:5: *** missing separator.  Stop.
缩进问题 必须用tab

如果target和已经存在的文件名称冲突
如存在clean文件
执行make clean不会执行clean的命令

.PHONY: clean
表示clean是一个伪目标，而不是一个真正的文件
无论clean文件是否存在 都会执行clean的命令