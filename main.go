package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Emp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Person struct {
	Id   []int
	Name []string
	City []string
}

const (
	PORT = ":8080"
	HOST = "localhost"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	fmt.Println("Database connected")
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goblog")
	ErrorHandling(err)
	defer db.Close()

	// var name string
	// var city string

	// fmt.Println("Enter You Name:")
	// fmt.Scan(&name)
	// fmt.Println("Enter You City:")
	// fmt.Scan(&city)

	// insert qry
	// insert, err := db.Query("insert into employee(name,city) values(?,?)",&name,&city)

	// ErrorHandling(err)

	// defer insert.Close()

	// delete qry

	// del ,err := db.Query("delete from employee where id in (?,?,?)",7,8,9)

	// ErrorHandling(err)

	// defer del.Close()
	http.HandleFunc("/",show)
	http.HandleFunc("/index", index)
	http.HandleFunc("/process", process)
	fmt.Println("Listening on " + HOST + PORT)
	http.ListenAndServe(PORT, nil)

	// var emp Emp

	// err = db.QueryRow("select id,name,city from employee where id = ?",2).Scan(&emp.id,&emp.name,&emp.city)

	// if err != nil {
	// 	panic(err)
	// }

	// log.Println(emp.id)
	// log.Println(emp.name)
	// log.Println(emp.city)

}

func ErrorHandling(err error) {
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "process.html", r)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goblog")
	ErrorHandling(err)
	defer db.Close()

	fname := r.FormValue("fname")
	city := r.FormValue("city")
	insert, err := db.Query("insert into employee(name,city) values(?,?)", fname, city)

	ErrorHandling(err)

	defer insert.Close()
	fmt.Printf("%v", fname)
	//  data:=struct{
	// 	Fname string
	// 	City  string

	//  }{
	// 	Fname: fname,
	// 	City: city,
	//  }

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func show(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/goblog")
	ErrorHandling(err)
	defer db.Close()

	var person Person
	var emp Emp
	res, err := db.Query("select id,name,city from employee")
	ErrorHandling(err)
	fmt.Println("*************************")
	fmt.Println("ID\tNAME\tCITY")
	fmt.Println("*************************")
	for res.Next() {

		err = res.Scan(&emp.Id, &emp.Name, &emp.City)

		ErrorHandling(err)

		fmt.Printf("%v\t%v\t%v\n", emp.Id, emp.Name, emp.City)
		person.Id = append(person.Id, emp.Id)
		person.Name = append(person.Name, emp.Name)
		person.City = append(person.City, emp.City)
	}
	fmt.Println("*************************")

	

		data := Person{
			Id:   person.Id,
			Name: person.Name,
			City: person.City,
		}

		tpl.ExecuteTemplate(w,"data.html",data)

	
}
