// gorm example - Go ORM (Object Relational Mapper) see docs: https://github.com/jinzhu/gorm
//				  mapping to SQL server 2017 (locally) using UI SSMS 18.2,
//                and installing sample db: GORMExampleDB, and maybe AdventureWorks2017 mapping an existing view (vEmployee)
//  1) install by, cmd line:  C:\>  go get -u github.com/jinzhu/gorm,  installs in GOPATH=C:\Users\MartyMazurik\go      (gorm.a)
//  2) make sure SQL Server installed, create a user: gorm pwd: gorm, give db_owner, db_datareader, db_datawriter to user: gorm
//  3) through SQL Server Config mgr, make sure to enable protocol for TCP/IP
//  4) restart SQL Server service (we're using 1433 port)

package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql" // mssql not 'officially' supported so ad dialects to allow it, from blogpost: https://stackoverflow.com/questions/40110670/how-to-connect-to-microsoft-sql-server-using-gorm-in-golang
)

func main() {
	db, err := gorm.Open("mssql", "sqlserver//gorm:gorm@localhost:1433?database=gorm")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// connection to db verification
	err = db.DB().Ping()
	if err != nil {
		panic(err.Error())
	} else {
		println("Connection to db was established.")
	}

	db.Set("gorm:table_options", "ENGINE=GORMExampleDB CHARSET=utf8")
	db.Debug().CreateTable(&User{}) // bug!   creates it in the Master DB, done with the ORM for now
	println("done.")

}

type User struct {
	gorm.Model
	Username  string
	FirstName string
	LastName  string
}
