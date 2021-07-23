FROM mysql:8.0

EXPOSE 5000

COPY ./docker/mysql/config/my.cnf /etc/mysql/conf.d/my.cnf
COPY ./docker/mysql/data/init/init.sql /docker-entrypoint-initdb.d

CMD ["mysqld"]
