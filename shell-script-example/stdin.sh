#!/usr/bin/env bash

wc -l < a > b

#将 stdout 和 stderr 合并后重定向到 file
# command > file 2>&1
aa > err 2>&1
aa >> err 2>&1



wc -l <<EOF

11
1

1
1
1
1
EOF