package main

import (
    "encoding/json"
    "fmt"
	"html"
    "net/http"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var localhost = "127.0.0.1"
var itb = "167.205.67.251"

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Book-oo! You are here %q. This is made by Balya Ibnu Sulistiyono (18215005)", html.EscapeString(r.URL.Path))
}

func ConnectToDB() (*sql.DB) {
	//connect to db
	db, err := sql.Open("mysql","root:@tcp("+itb+":3306)/book-oo")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DoesKeyExists(db *sql.DB, table string, keyname string, id string) (bool) {
	var isAvailable bool
	queryRow := "SELECT IF(COUNT(*), 'true', 'false') FROM " + table + " WHERE " + keyname + "=\"" + id + "\""
	err := db.QueryRow(queryRow).Scan(&isAvailable)
	if err!= nil {
		log.Fatal(err)
	}
	return isAvailable
}

func GetAllMember(w http.ResponseWriter, r *http.Request) {
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//variable identification
	member := Member {}
	var members Members
	
	//execute query
	rows, err := db.Query("SELECT students_id, name, social_media_contact FROM member")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//making arrays of member
	for rows.Next() {
		err := rows.Scan(&member.Students_id, &member.Name, &member.Social_media_contact)
		if err != nil {
			log.Fatal(err)
		} else {
			members = append(members,member)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	//json encoding
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
	enc := json.NewEncoder(w)
	enc.SetIndent("","\t")
	enc.Encode(members); 
	if err != nil {
        panic(err)
    }
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//variable identification
	book := Book {}
	var books Books
	
	//execute query
	rows, err := db.Query("SELECT isbn, book_title FROM book")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//making arrays of book
	for rows.Next() {
		err := rows.Scan(&book.Isbn, &book.Book_title)
		if err != nil {
			log.Fatal(err)
		} else {
			books = append(books,book)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	//json encoding
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    enc := json.NewEncoder(w)
	enc.SetIndent("","\t")
	enc.Encode(books); 
	if err != nil {
        panic(err)
    }
}

//access variable value: mux.Vars(http.Request)["<variable-name>"]
//<variable-name> == {<variable-name>}

func GetMember(w http.ResponseWriter, r *http.Request) {
	//getting member id
	students_id := mux.Vars(r)["member_id"]
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//variable identification
	member := Member {}
	
	//execute query
	rows, err := db.Query("SELECT students_id, name, social_media_contact FROM member WHERE students_id=" + students_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//making member
	for rows.Next() {
		err := rows.Scan(&member.Students_id, &member.Name, &member.Social_media_contact)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	//json encoding
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    enc := json.NewEncoder(w)
	enc.SetIndent("","\t")
	enc.Encode(member);
	if err != nil {
        panic(err)
    }
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	//getting book id
	isbn := mux.Vars(r)["isbn"]
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//variable identification
	book := Book {}
	
	//execute query
	rows, err := db.Query("SELECT isbn, book_title FROM book WHERE isbn=\"" + isbn + "\"")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//making book
	for rows.Next() {
		err := rows.Scan(&book.Isbn, &book.Book_title)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	//json encoding
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    enc := json.NewEncoder(w)
	enc.SetIndent("","\t")
	enc.Encode(book);
	if err != nil {
        panic(err)
    }
}

func GetMemberShelf(w http.ResponseWriter, r *http.Request) {
	//getting member id
	students_id := mux.Vars(r)["member_id"]
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//variable identification
	book := BookOwned {}
	var shelf BookOwneds
	
	//execute query
	rows, err := db.Query("SELECT book.isbn, book.book_title, owns.number_of_copy FROM owns INNER JOIN member ON member.students_id = owns.students_id INNER JOIN book ON book.isbn = owns.isbn WHERE owns.students_id =" + students_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//making book
	for rows.Next() {
		err := rows.Scan(&book.Isbn, &book.Book_title, &book.Number_of_copy)
		if err != nil {
			log.Fatal(err)
		} else {
			//append shelf
			shelf = append(shelf, book)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	//json encoding
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    enc := json.NewEncoder(w)
	enc.SetIndent("","\t")
	enc.Encode(shelf);
	if err != nil {
        panic(err)
    }
}


func GetBookOwnedBy(w http.ResponseWriter, r *http.Request) {
	//getting book id
	isbn := mux.Vars(r)["isbn"]
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//variable identification
	member := Member {}
	var members Members
	
	//execute query
	rows, err := db.Query("SELECT member.students_id, member.name, member.social_media_contact FROM owns INNER JOIN member ON member.students_id = owns.students_id INNER JOIN book ON book.isbn = owns.isbn WHERE owns.isbn =\"" + isbn +"\"")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	
	//making arrays of member
	for rows.Next() {
		err := rows.Scan(&member.Students_id, &member.Name, &member.Social_media_contact)
		if err != nil {
			log.Fatal(err)
		} else {
			members = append(members,member)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	
	//json encoding
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    enc := json.NewEncoder(w)
	enc.SetIndent("","\t")
	enc.Encode(members);
	if err != nil {
        panic(err)
    }
}

func AddMember(w http.ResponseWriter, r *http.Request) {
	//getting json from client
	decoder := json.NewDecoder(r.Body)
	
	var member Member
	if err := decoder.Decode(&member); err != nil {
		panic(err)
	}
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//execute query
	if (!DoesKeyExists(db, "member","students_id",member.Students_id)) {
		rows, err := db.Query("CALL add_member(\""+member.Students_id+"\",\""+member.Name+"\",\""+member.Social_media_contact+"\")")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	//getting json from client
	decoder := json.NewDecoder(r.Body)
	
	var book Book
	if err := decoder.Decode(&book); err != nil {
		panic(err)
	}
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//execute query
	if (!DoesKeyExists(db, "book","isbn",book.Isbn)) {
		rows, err := db.Query("CALL add_book(\""+book.Isbn+"\",\""+book.Book_title+"\")")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	}
}

func AddToShelf(w http.ResponseWriter, r *http.Request) {
	//getting json from client
	decoder := json.NewDecoder(r.Body)
	
	var owns Owns
	if err := decoder.Decode(&owns); err != nil {
		panic(err)
	}
	
	//connect to db
	db := ConnectToDB()
	defer db.Close()
	
	//execute query
	rows, err := db.Query("CALL add_to_shelf(\""+owns.Students_id+"\",\""+owns.Isbn+"\",\""+owns.Number_of_copy+"\")")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}