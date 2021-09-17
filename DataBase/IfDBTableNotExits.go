package DataBase

import "log"

func TableNotExits() {
	UserTable := "CREATE TABLE IF NOT EXISTS `gokufa`.`users` ( `id` INT UNSIGNED AUTO_INCREMENT NOT NULL ,`email` VARCHAR(255) NOT NULL , `password` VARCHAR(255) NOT NULL ,`status` int(11) NOT NULL DEFAULT 1,`role` int(11) NOT NULL DEFAULT 1 COMMENT '1=user 2=employee 3= admin',`image` varchar(255) NOT NULL DEFAULT 'default.png', UNIQUE KEY `email` (`email`), PRIMARY KEY (`id`)) ENGINE = InnoDB;"
	_, err := Db.Query(UserTable)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
