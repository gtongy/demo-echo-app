# データベース作成
mysql -u root -proot -e "\
    CREATE DATABASE development_database CHARACTER SET utf8;
    CREATE DATABASE test_database CHARACTER SET utf8;
    CREATE DATABASE production_database CHARACTER SET utf8;
"

TABLE_SQL=$(cat<<"EOS"

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`password` char(64) NOT NULL,
`email` varchar(64) NOT NULL DEFAULT '',
`access_token` varchar(64) NOT NULL DEFAULT '',
`created_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tasks`;

CREATE TABLE `tasks` (
`id` int(11) NOT NULL AUTO_INCREMENT,
`user_id` int(11) NOT NULL,
`title` varchar(64) NOT NULL DEFAULT '',
`completed` bool NOT NULL,
`completed_at` datetime DEFAULT CURRENT_TIMESTAMP,
`updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

EOS
)

mysql -u root -proot -e "\
    USE development_database;
    $TABLE_SQL

    USE test_database;
    $TABLE_SQL

    USE production_database;
    $TABLE_SQL
"

echo 'complete create table!'