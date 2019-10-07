#!/usr/bin/env bash
fistFunc(){
echo `date`
}
#运行
fistFunc

#  函数返回值
returnFunc(){

    return 5
}
returnFunc
echo "返回值为$?"


# args
argsFunc(){
echo $2
echo ${1}
echo ${10}
}
argsFunc 1 2 1 1 1 1 1 11 1 11  1 1 11