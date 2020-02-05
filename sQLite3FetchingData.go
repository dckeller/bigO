package main 

import (
	"database/sql"
  "fmt"
  "time"
  _ "github.com/mattn/go-sqlite3"
)

func main() {

	// INSERT //

	// // First arg is the driver name, and the second is the database name and location //
	db, err := sql.Open("sqlite3", "/Users/david.keller/Documents/bigO/userinfo.db")
	checkErr(err)

	// Prepared statement on a connection, the stms remembers which connection was used
	// When you execute stmt, it tries to use the connection.  It will try until successful
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	// // What to add to the database
	res, err := stmt.Exec("dckeller", "marketing", "01-29-2020")
	checkErr(err)

	// // Will get the last inserted row ID 
	id, err := res.LastInsertId()
  checkErr(err)

  fmt.Println(id)


  // UPDATE //

  stmt, err = db.Prepare("update userinfo set username=? where uid=?")
  checkErr(err)

  // What to update in the databse
  res, err = stmt.Exec("dckeller", id)
  checkErr(err)

  // Will tell us how many rows were updated and deleted by statements
  affect, err := res.RowsAffected()
  checkErr(err)

  fmt.Println(affect)


  // QUERY //

  rows, err := db.Query("SELECT * FROM userinfo")
  checkErr(err)
  var uid int
  var username string
  var departname string
  var created time.Time  


 //  // Iterate over the rows with rows.Next()
  for rows.Next() {

 //  // Read the columsn in each row into variables with rows.Scan()
  	err = rows.Scan(&uid, &username, &departname, &created)
    checkErr(err)
    fmt.Println(uid)
    fmt.Println(username)
    fmt.Println(departname)
    fmt.Println(created)
	}

	rows.Close()
}

	// // DELETE //
	stmt, err := db.Prepare("DELETE FROM userinfo WHERE id=?")
	checkErr(err)

	// Which item in DB to delete
	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()

	trashSQL, err := db.Prepare("update task set is_deleted='Y',last_modified_at=datetime() where id=?")
  if err != nil {
      fmt.Println(err)
  }

  // Begin a transaction which reserves a connection to the datastore
  tx, err := db.Begin()
  if err != nil {
      fmt.Println(err)
  }


  _, err = tx.Stmt(trashSQL).Exec(id)
  if err != nil {
      fmt.Println("doing rollback")

  // If there is an error, rollback the transaction    
      tx.Rollback()
  } else {

  // Commit if there is NO error	
      tx.Commit()
  }
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}