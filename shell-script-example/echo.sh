#!/usr/bin/env bash

echo "\"\""

#  读取输入
echo -e "请输入name: \c"
read name
echo "your name:"         ${name}

# 显示换行 -e \n
echo  "不加-e \n"
echo -e "加-e \n"

# 显示不换行 -e \c
echo -e "abc\c"
echo "def"


# 覆盖文件
echo "新内容1" > myfile
echo "新内容2 覆盖" > myfile
echo "新内容3 追加" >> myfile

# 显示命令执行结果
echo "当前时间为:"`date`