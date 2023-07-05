# 先进mylog制作动态库

# 和静态库使用方法类似 可以编译
gcc main.c -o main -I ./include/ -l mylog -L ./lib/

# 但执行会报错
# ./main: error while loading shared libraries: 
# libmylog.so: cannot open shared object file: No such file or directory

# 动态库通过ld-linux.so来获取绝对路径
# 按 DT_RPATH -> PATH::LD_LIBRARY_PATH -> /etc/ld.so.cache -> /lib/ -> /usr/lib/ 顺序查找

# ldd main 查看动态库依赖关系
# linux-vdso.so.1 (0x00007ffda75d7000)
# libmylog.so => not found
# libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f077b11f000)
# /lib64/ld-linux-x86-64.so.2 (0x00007f077b353000)

# 当前shell下有效
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:<绝对路径>/c/04_dynamiclib/lib

# 编辑 ~/.bashrc 在最后添加上面的命令
source ~/.bashrc

# 编辑 /etc/profile 在最后添加上面的命令
source /etc/profile

# 修改 /etc/ld.so.conf 加上绝对路径
ldconfig
