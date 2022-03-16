#!/bin/bash/env bash
echo "初始化中....."
git pull origin main;
git add -A;
ss=`date` >/dev/null
read -p "输入日志,按Enter键跳过 :" log
if  [ ! -n "$log" ] ;then
    git commit -m "${ss}";
else git commit -m "${log} ${ss}";
fi;
git push origin main;
echo "远程推送完成"
