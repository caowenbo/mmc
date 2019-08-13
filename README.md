# mmc
mini mariadb cli tool
>usage:mmc host user password sql
*example*
```
mmc 192.168.1.1 root pwd "CREATE DATABASE IF NOT EXISTS apple DEFAULT CHARSET utf8mb4 COLLATE = utf8mb4_unicode_ci;"
mmc 192.168.1.1 root pwd "select * from mysql.user;"
mmc 192.168.1.1 root pwd "show databases;"
```
