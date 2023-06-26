# 可执行文件
gcc helloworld.c -o helloworld

# 预处理不编译
gcc -E helloworld.c -o helloworld.i

# 编译不汇编
gcc -S helloworld.i -o helloworld.s

# 汇编不链接
gcc -c helloworld.s -o helloworld.o

