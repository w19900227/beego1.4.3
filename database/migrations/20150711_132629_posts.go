package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Posts_20150711_132629 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Posts_20150711_132629{}
	m.Created = "20150711_132629"
	migration.Register("Posts_20150711_132629", m)
}

// Run the migrations
func (m *Posts_20150711_132629) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE `posts` (
		`id` INT(11) UNSIGNED AUTO_INCREMENT NOT NULL,
		`title` varchar(255) NOT NULL,
		`content` text NOT NULL,
		`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		`updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
		PRIMARY KEY (`id`)
		) engine=InnoDB DEFAULT CHARSET=utf8;")
	/*
	INSERT INTO `posts` (`id`, `user_id`, `category_id`, `status`, `title`, `content`, `created_at`, `updated_at`) VALUES
	(1, 4, 4, 1, '21', '2', '2015-07-13 02:41:30', '0000-00-00 00:00:00'),
	(2, 4, 3, 2, '2', '2', '2015-07-13 02:41:30', '0000-00-00 00:00:00');
	*/

}

// Reverse the migrations
func (m *Posts_20150711_132629) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	use m.Sql("DROP TABLE posts") to reverse schema update

}
//bee generate scaffold post -fields="title:string,body:text"