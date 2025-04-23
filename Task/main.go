// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/joho/godotenv"
// )

// var sqlDB *sql.DB

// // Load .env file
// func loadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("❌ Error loading .env file")
// 	}
// }

// func connect() {
// 	loadEnv()
// 	user := os.Getenv("DB_USER")
// 	pass := os.Getenv("DB_PASS")
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	name := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, name)

// 	var err error
// 	sqlDB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatal("❌ Connection error:", err)
// 	}
// 	if err := sqlDB.Ping(); err != nil {
// 		log.Fatal("❌ Ping error:", err)
// 	}
// 	fmt.Println("✅ Connected to MySQL!")
// }

// func insertEmployee(name, city string, age int) {
// 	_, err := sqlDB.Exec(`INSERT INTO employee (name, city, age) VALUES (?, ?, ?)`, name, city, age)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("✅ Employee inserted")
// }

// func updateEmployee(id int, name, city string, age int) {
// 	_, err := sqlDB.Exec(`UPDATE employee SET name = ?, city = ?, age = ? WHERE id = ?`, name, city, age, id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("✅ Employee updated")
// }

// func deleteEmployee(id int) {
// 	_, err := sqlDB.Exec(`DELETE FROM employee WHERE id = ?`, id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("✅ Employee deleted")
// }

// func getAllEmployees() {
// 	rows, err := sqlDB.Query(`SELECT id, name, city, age FROM employee`)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("🧾 All Employees:")
// 	for rows.Next() {
// 		var id, age int
// 		var name, city string
// 		rows.Scan(&id, &name, &city, &age)
// 		fmt.Printf("- ID: %d | %s | %s | Age: %d\n", id, name, city, age)
// 	}
// }

// func getEmployeeByID(id int) {
// 	row := sqlDB.QueryRow(`SELECT id, name, city, age FROM employee WHERE id = ?`, id)
// 	var eid, age int
// 	var name, city string
// 	err := row.Scan(&eid, &name, &city, &age)
// 	if err != nil {
// 		fmt.Println("❌ Employee not found")
// 		return
// 	}
// 	fmt.Printf("🧍 Found Employee: ID: %d | %s | %s | Age: %d\n", eid, name, city, age)
// }

// func countEmployees() {
// 	row := sqlDB.QueryRow(`SELECT COUNT(*) FROM employee`)
// 	var count int
// 	row.Scan(&count)
// 	fmt.Printf("📌 Total Employees: %d\n", count)
// }

// func oldestAndYoungestEmployee() {
// 	row := sqlDB.QueryRow(`SELECT MIN(age), MAX(age) FROM employee`)
// 	var minAge, maxAge int
// 	row.Scan(&minAge, &maxAge)
// 	fmt.Printf("👶 Youngest Age: %d\n", minAge)
// 	fmt.Printf("👴 Oldest Age: %d\n", maxAge)
// }

// func orderEmployeesByAge(order string) {
// 	if order != "ASC" && order != "DESC" {
// 		fmt.Println("❌ Invalid order. Use 'ASC' or 'DESC'.")
// 		return
// 	}

// 	query := fmt.Sprintf(`SELECT name, city, age FROM employee ORDER BY age %s`, order)
// 	rows, err := sqlDB.Query(query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	fmt.Println("📊 Employees Ordered by Age (" + order + "):")
// 	for rows.Next() {
// 		var name, city string
// 		var age int
// 		rows.Scan(&name, &city, &age)
// 		fmt.Printf("- %s | %s | Age: %d\n", name, city, age)
// 	}
// }

// // --- Helper Input Function ---
// func input(prompt string) string {
// 	fmt.Print(prompt)
// 	var response string
// 	fmt.Scanln(&response)
// 	return response
// }

// // --- Menu System ---
// func main() {

// 	connect()
// createEmployeeTableIfNotExists()

// 	for {
// 		fmt.Println("\n📋 EMPLOYEE MANAGEMENT SYSTEM")
// 		fmt.Println("1. Add Employee")
// 		fmt.Println("2. Update Employee")
// 		fmt.Println("3. Delete Employee")
// 		fmt.Println("4. View All Employees")
// 		fmt.Println("5. View Employee by ID")
// 		fmt.Println("6. Count Employees")
// 		fmt.Println("7. Youngest & Oldest Age")
// 		fmt.Println("8. Order Employees by Age")
// 		fmt.Println("0. Exit")
// 		fmt.Print("➡ Choose an option: ")

// 		var choice int
// 		fmt.Scanln(&choice)

// 		switch choice {
// 		case 1:
// 			name := input("Enter name: ")
// 			city := input("Enter city: ")
// 			ageStr := input("Enter age: ")
// 			age, _ := strconv.Atoi(ageStr)
// 			insertEmployee(name, city, age)

// 		case 2:
// 			idStr := input("Enter employee ID to update: ")
// 			id, _ := strconv.Atoi(idStr)
// 			name := input("Enter new name: ")
// 			city := input("Enter new city: ")
// 			ageStr := input("Enter new age: ")
// 			age, _ := strconv.Atoi(ageStr)
// 			updateEmployee(id, name, city, age)

// 		case 3:
// 			idStr := input("Enter employee ID to delete: ")
// 			id, _ := strconv.Atoi(idStr)
// 			deleteEmployee(id)

// 		case 4:
// 			getAllEmployees()

// 		case 5:
// 			idStr := input("Enter employee ID to search: ")
// 			id, _ := strconv.Atoi(idStr)
// 			getEmployeeByID(id)

// 		case 6:
// 			countEmployees()

// 		case 7:
// 			oldestAndYoungestEmployee()

// 		case 8:
// 			order := input("Enter order (ASC/DESC): ")
// 			orderEmployeesByAge(order)

// 		case 0:
// 			fmt.Println("👋 Exiting Bye...")
// 			os.Exit(0)

// 		default:
// 			fmt.Println("❌ Invalid choice. Try again.")
// 		}
// 	}
// }

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var sqlDB *sql.DB
var tmpl *template.Template

type Employee struct {
	ID   int
	Name string
	City string
	Age  int
}






func createEmployeeTableIfNotExists() {
	query := `
	CREATE TABLE IF NOT EXISTS employee (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		city VARCHAR(100) NOT NULL,
		age INT NOT NULL
	);
	`

	_, err := sqlDB.Exec(query)
	if err != nil {
		log.Fatal("❌ Failed to create employee table:", err)
	} else {
		fmt.Println("🛠️ Table checked/created: employee")
	}
}



func main() {
	connect()
	createEmployeeTableIfNotExists()
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", showEmployees)
	http.HandleFunc("/add", addEmployee)
	http.HandleFunc("/delete", deleteEmployee)
	http.HandleFunc("/edit", editEmployee) // Register the new edit route

	fmt.Println("🌐 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("❌ Connection error:", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("❌ Ping error:", err)
	}
	fmt.Println("✅ Connected to MySQL!")
}

func showEmployees(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")
	order := r.URL.Query().Get("order") 
	if order == "" {
		order = "id"
	}

	query := "SELECT id, name, city, age FROM employee"
	countQuery := "SELECT COUNT(*) FROM employee"

	var args []interface{}

	// Add search filtering if provided
	if search != "" {
		query += " WHERE name LIKE ? OR city LIKE ?"
		countQuery += " WHERE name LIKE ? OR city LIKE ?"
		likeSearch := "%" + search + "%"
		args = append(args, likeSearch, likeSearch)
	}

	query += " ORDER BY " + order

	rows, err := sqlDB.Query(query, args...)
	if err != nil {
		http.Error(w, "DB query error", 500)
		return
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var emp Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.City, &emp.Age)
		employees = append(employees, emp)
	}

 	var totalCount int
	err = sqlDB.QueryRow(countQuery, args...).Scan(&totalCount)
	if err != nil {
		http.Error(w, "DB count error", 500)
		return
	}

	data := struct {
		Employees   []Employee
		TotalCount  int
		SearchValue string
		OrderBy     string
	}{
		Employees:   employees,
		TotalCount:  totalCount,
		SearchValue: search,
		OrderBy:     order,
	}

	tmpl.ExecuteTemplate(w, "index.html", data)
}


func addEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		city := r.FormValue("city")
		ageStr := r.FormValue("age")
		age, _ := strconv.Atoi(ageStr)

		_, err := sqlDB.Exec("INSERT INTO employee (name, city, age) VALUES (?, ?, ?)", name, city, age)
		if err != nil {
			http.Error(w, "Insert error", 500)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "add.html", nil)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id != "" {
		_, err := sqlDB.Exec("DELETE FROM employee WHERE id = ?", id)
		if err != nil {
			http.Error(w, "Delete error", 500)
			return
		}
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Employee ID is missing", http.StatusBadRequest)
		return
	}

	// Fetch the employee details
	var emp Employee
	err := sqlDB.QueryRow("SELECT id, name, city, age FROM employee WHERE id = ?", id).Scan(&emp.ID, &emp.Name, &emp.City, &emp.Age)
	if err != nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// If the request is POST, update the employee details
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		city := r.FormValue("city")
		ageStr := r.FormValue("age")
		age, _ := strconv.Atoi(ageStr)

		// Update the employee in the database
		_, err := sqlDB.Exec("UPDATE employee SET name = ?, city = ?, age = ? WHERE id = ?", name, city, age, id)
		if err != nil {
			http.Error(w, "Update error", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// If GET request, display the current employee data in the form
	tmpl.ExecuteTemplate(w, "edit.html", emp)
}

