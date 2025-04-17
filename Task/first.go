package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var sqlDB *sql.DB

func connect() {
	var err error
	sqlDB, err = sql.Open("mysql", "root:mahmoud123@tcp(127.0.0.1:3306)/company")
	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to MySQL!")
}

func insertEmployee(name, city string, age int) {
	_, err := sqlDB.Exec(`INSERT INTO employee (name, city, age) VALUES (?, ?, ?)`, name, city, age)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Employee inserted")
}

func updateEmployee(id int, name, city string, age int) {
	_, err := sqlDB.Exec(`UPDATE employee SET name = ?, city = ?, age = ? WHERE id = ?`, name, city, age, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Employee updated")
}

func deleteEmployee(id int) {
	_, err := sqlDB.Exec(`DELETE FROM employee WHERE id = ?`, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("✅ Employee deleted")
}

func getAllEmployees() {
	rows, err := sqlDB.Query(`SELECT id, name, city, age FROM employee`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("🧾 All Employees:")
	for rows.Next() {
		var id, age int
		var name, city string
		rows.Scan(&id, &name, &city, &age)
		fmt.Printf("- ID: %d | %s | %s | Age: %d\n", id, name, city, age)
	}
}

func getEmployeeByID(id int) {
	row := sqlDB.QueryRow(`SELECT id, name, city, age FROM employee WHERE id = ?`, id)
	var eid, age int
	var name, city string
	err := row.Scan(&eid, &name, &city, &age)
	if err != nil {
		fmt.Println("❌ Employee not found")
		return
	}
	fmt.Printf("🧍 Found Employee: ID: %d | %s | %s | Age: %d\n", eid, name, city, age)
}

func countEmployees() {
	row := sqlDB.QueryRow(`SELECT COUNT(*) FROM employee`)
	var count int
	row.Scan(&count)
	fmt.Printf("📌 Total Employees: %d\n", count)
}

func oldestAndYoungestEmployee() {
	row := sqlDB.QueryRow(`SELECT MIN(age), MAX(age) FROM employee`)
	var minAge, maxAge int
	row.Scan(&minAge, &maxAge)
	fmt.Printf("👶 Youngest Age: %d\n", minAge)
	fmt.Printf("👴 Oldest Age: %d\n", maxAge)
}

func orderEmployeesByAge(order string) {
	if order != "ASC" && order != "DESC" {
		fmt.Println("❌ Invalid order. Use 'ASC' or 'DESC'.")
		return
	}

	query := fmt.Sprintf(`SELECT name, city, age FROM employee ORDER BY age %s`, order)
	rows, err := sqlDB.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	fmt.Println("📊 Employees Ordered by Age (" + order + "):")
	for rows.Next() {
		var name, city string
		var age int
		rows.Scan(&name, &city, &age)
		fmt.Printf("- %s | %s | Age: %d\n", name, city, age)
	}
}

// --- Helper Input Function ---
func input(prompt string) string {
	fmt.Print(prompt)
	var response string
	fmt.Scanln(&response)
	return response
}

// --- Menu System ---
func main() {
	connect()

	for {
		fmt.Println("\n📋 EMPLOYEE MANAGEMENT SYSTEM")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Update Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. View All Employees")
		fmt.Println("5. View Employee by ID")
		fmt.Println("6. Count Employees")
		fmt.Println("7. Youngest & Oldest Age")
		fmt.Println("8. Order Employees by Age")
		fmt.Println("0. Exit")
		fmt.Print("➡ Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			name := input("Enter name: ")
			city := input("Enter city: ")
			ageStr := input("Enter age: ")
			age, _ := strconv.Atoi(ageStr)
			insertEmployee(name, city, age)

		case 2:
			idStr := input("Enter employee ID to update: ")
			id, _ := strconv.Atoi(idStr)
			name := input("Enter new name: ")
			city := input("Enter new city: ")
			ageStr := input("Enter new age: ")
			age, _ := strconv.Atoi(ageStr)
			updateEmployee(id, name, city, age)

		case 3:
			idStr := input("Enter employee ID to delete: ")
			id, _ := strconv.Atoi(idStr)
			deleteEmployee(id)

		case 4:
			getAllEmployees()

		case 5:
			idStr := input("Enter employee ID to search: ")
			id, _ := strconv.Atoi(idStr)
			getEmployeeByID(id)

		case 6:
			countEmployees()

		case 7:
			oldestAndYoungestEmployee()

		case 8:
			order := input("Enter order (ASC/DESC): ")
			orderEmployeesByAge(order)

		case 0:
			fmt.Println("👋 Exiting...")
			os.Exit(0)

		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/dialog"
// 	"fyne.io/fyne/v2/layout"
// 	"fyne.io/fyne/v2/widget"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var sqlDB *sql.DB

// func connect() {
// 	var err error
// 	sqlDB, err = sql.Open("mysql", "root:mahmoud123@tcp(127.0.0.1:3306)/company")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err = sqlDB.Ping(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("✅ Connected to MySQL!")
// }

// func insertEmployee(name, city string, age int) {
// 	_, err := sqlDB.Exec(`INSERT INTO employee (name, city, age) VALUES (?, ?, ?)`, name, city, age)
// 	if err != nil {
// 		dialog.ShowError(err, nil)
// 		return
// 	}
// 	fmt.Println("✅ Employee inserted")
// }

// func getAllEmployees() []string {
// 	rows, err := sqlDB.Query(`SELECT id, name, city, age FROM employee`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var data []string
// 	for rows.Next() {
// 		var id, age int
// 		var name, city string
// 		err := rows.Scan(&id, &name, &city, &age)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		data = append(data, fmt.Sprintf("ID: %d | %s | %s | Age: %d", id, name, city, age))
// 	}
// 	return data
// }

// func main() {
// 	connect()
// 	a := app.New()
// 	w := a.NewWindow("Employee Manager")
// 	w.Resize(fyne.NewSize(500, 500))

// 	employeeList := widget.NewList(
// 		func() int { return len(getAllEmployees()) },
// 		func() fyne.CanvasObject { return widget.NewLabel("") },
// 		func(i widget.ListItemID, o fyne.CanvasObject) {
// 			employees := getAllEmployees()
// 			if i < len(employees) {
// 				o.(*widget.Label).SetText(employees[i])
// 			}
// 		},
// 	)

// 	nameEntry := widget.NewEntry()
// 	nameEntry.SetPlaceHolder("Name")

// 	cityEntry := widget.NewEntry()
// 	cityEntry.SetPlaceHolder("City")

// 	ageEntry := widget.NewEntry()
// 	ageEntry.SetPlaceHolder("Age")

// 	addButton := widget.NewButton("Add Employee", func() {
// 		name := nameEntry.Text
// 		city := cityEntry.Text
// 		age, err := strconv.Atoi(ageEntry.Text)
// 		if err != nil {
// 			dialog.ShowError(fmt.Errorf("Invalid age input"), w)
// 			return
// 		}
// 		insertEmployee(name, city, age)
// 		employeeList.Refresh()
// 		nameEntry.SetText("")
// 		cityEntry.SetText("")
// 		ageEntry.SetText("")
// 	})

// 	// Layout
// 	form := container.NewVBox(
// 		widget.NewLabel("Add New Employee"),
// 		nameEntry,
// 		cityEntry,
// 		ageEntry,
// 		addButton,
// 	)

// 	mainContainer := container.NewBorder(form, nil, nil, nil, container.NewVScroll(employeeList))
// 	w.SetContent(container.New(layout.NewVBoxLayout(), mainContainer))
// 	w.ShowAndRun()
// }
