package database

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id       int64     `orm:"pk;auto"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
	Name     string    `orm:"size(128)"`
	Email    string    `orm:"index;unique;size(64)"`
	Password string    `orm:"index;size(60)"`
	Team     *Team     `orm:"rel(fk);null"`
}

// try and authenticate the user, err is nil if auth succeeds
func (u *User) Authenticate(p *string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(*p))
	if err == nil {
		log.Printf("Authentication succeeded for user [%s]\n", u.Email)
	} else {
		log.Printf("Authentication failed for user [%s]\n", u.Email)
	}
	return err
}

// FIXME: move this to models/team.go
type Team struct {
	Id      int64     `orm:"pk;auto"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Name    string    `orm:"size(128)"`
	Email   string    `orm:"index;size(64)"`
	Url     string    `orm:"size(1024)"`
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "teresa.sqlite") // FIXME: make it configurable and per-env
	orm.RegisterModel(new(User), new(Team))

	// create the database -- FIXME
	if true {
		// Database alias.
		name := "default"

		// Drop table and re-create.
		force := false

		// Print log.
		verbose := true

		// Error.
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			fmt.Println(err)
		}
	}
}
