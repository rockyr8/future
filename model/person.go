package model

import (
	"fmt"
	"time"
	
	db "future/database"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	// rs, err := db.SqlDB.Exec("call test1(?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	// db.RedisSet(p.FirstName, p.LastName, 60)
	id, err = rs.LastInsertId()
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func GetRedisV(key string) (rat string) {
	// rat = db.RedisGet(key)
	ms := int64(db.RedisClient.TTL(key).Val() / time.Second)
	rat = fmt.Sprintf("%v",ms)
	return
}

func SetRedisV(key string, val string, sec time.Duration) bool {	
	return db.RedisSet(key, val, sec)==nil
}
