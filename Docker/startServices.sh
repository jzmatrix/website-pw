#!/bin/bash

/usr/sbin/postfix -c /etc/postfix start

/usr/sbin/cron

/run-httpd.sh

mkdir /run/sshd
mkdir /var/run/sshd

/usr/sbin/sshd -D
