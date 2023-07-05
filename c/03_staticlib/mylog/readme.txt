# 编译不链接
gcc -c mylog.c -o mylog.o

# 打包 必须lib开头
ar rcs libmylog.a mylog.o

# 复制libmylog.a 去lib
cp libmylog.a ../lib/

# 复制head.h 去 include
cp head.h ../include/
# 程序运行需要静态库和头文件