#!/bin/bash
sudo systemctl stop mariadb
sudo systemctl set-environment MYSQLD_OPTS="--skip-grant-tables --skip-networking"
mysql -u root --execute="UPDATE mysql.user SET password = PASSWORD('new_password') WHERE user = 'root'; "
sudo systemctl unset-environment MYSQLD_OPTS
systemctl restart mariadb