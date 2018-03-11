FROM mysql:5.7
RUN usermod -u 1000 mysql
RUN chmod -R 777 /var/run/mysqld
COPY init/my.cnf /etc/mysql/conf.d/