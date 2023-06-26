# 编译不链接
gcc -c mylog.c -o mylog.o

# 打包 必须lib开头
ar rcs libmylog.a mylog.o

