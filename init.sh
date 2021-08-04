#!/bin/bash
/etc/init.d/cron start
echo "export SESSDATA=$SESSDATA" > cookie/cookie
echo "export BILI_JCT=$BILI_JCT" >> cookie/cookie
echo "export DEDEUSERID=$DEDEUSERID" >> cookie/cookie
bash
