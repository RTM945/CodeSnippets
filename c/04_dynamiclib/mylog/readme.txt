# 加编译选项 -fpic 
cd lib
gcc -c -fPIC mylog.c 

gcc -shared mylog.o -o libmylog.so

# 复制libmylog.so 去lib
cp libmylog.so ../lib/

# 复制head.h 去 include
cp head.h ../include/

# 回根目录尝试执行