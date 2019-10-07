#!/usr/bin/env bash

a=100
b=100
if test ${a} -lt ${b}; then
    echo "="
    else
    echo "!="
fi
# 数值
# if test ${a} -eq ${b} 等于
# if test ${a} -ne ${b} 不等于
# if test ${a} -gt ${b} 大于
# if test ${a} -ge ${b} 大于等于
# if test ${a} -lt ${b} 小于
# if test ${a} -le ${b} 小于等于
#
if test $[a+0] = ${b}; then
    echo "="
    else
    echo "!="
fi
## $[] 可以在里面进行数值计算
a=$[1+2+3]
echo $[1+2+3] ${a}


#字符串 test
# = 等于
# != 不等于
# -z 比较长度
# -n 长度不为00
stringA="helloworld!"
stringB="hellowold!"
if test ${stringA} != ${stringB}; then
    echo "="
else
    echo "!="
fi

#文件测试
# -e xxx 存在文件或目录
fileName="test.sh"
if test -e ${fileName}; then
   echo "${fileName} 文件存在"
else
    echo "${fileName} 文件不存在"
fi
# -r xxx 存在文件或目录并且可读
fileName="test.sh"
if test -r ${fileName}; then
   echo "${fileName} 存在文件或目录并且可读"
else
    echo "${fileName} 不存在文件或目录或者不可读"
fi
# -w xxx 存在文件或目录并且可写
fileName="test.sh"
if test -w ${fileName}; then
   echo "${fileName} 存在文件或目录并且可写"
else
    echo "${fileName} 不存在存在文件或目录或者可读"
fi
# -x xxx 存在文件或目录并且可执行
fileName="myfile"
if test -x ${fileName}; then
   echo "${fileName} 存在文件或目录并且可执行"
else
    echo "${fileName} 不存在存在文件或目录或者不可以执行"
fi

# -s xxx 存在文件或目录并且不为空
fileName="empty"
if test -f ${fileName}; then
    rm ${fileName}
    else
    touch ${fileName}
fi

if test -s ${fileName}; then
   echo "${fileName} 存在文件或目录并且不为空"
else
    echo "${fileName} 不存在存在文件或目录或者为空"
fi
# -d xxx 存在目录
fileName="test.sh"
if test -d ${fileName}; then
   echo "${fileName} 目录存在"
else
    echo "${fileName} 目录不存在"
fi

# -f xxx 存在文件
fileName="test.sh"
if test -f ${fileName}; then
   echo "${fileName} 普通文件存在"
else
    echo "${fileName} 普通文件不存在"
fi
# -c xxx 存在文件或目录并且字符型特殊文件
fileName="test.sh"
if test -c ${fileName}; then
   echo "${fileName} 普通文件存在"
else
    echo "${fileName} 普通文件不存在"
fi
# -b xxx 存在文件或目录并且块特殊文件










