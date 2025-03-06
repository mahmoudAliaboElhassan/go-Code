package main

import (
	"errors"
	"fmt"
)

func main() {

    // // function
    // fmt.Println("Hello, Go on Windows!")
    // fmt.Println(printName(("Mahmoud")))
    // fmt.Println(printAge((21)))

    // // scan input
    // // var age int
    // // fmt.Println("Enter your age :")
    // // fmt.Scan(&age)
    // // fmt.Println("Your Age is ",age)
    
    // // variables
    // num:=5
    // fmt.Println(num)

    // var fac string="Computer and Information Science"
    // fmt.Println("My Faculty is ",fac)
    // // panic("error happend unexpected")
    // var x,y int =5,10
    // fmt.Println(x)
    // fmt.Println(y)
    // var message="Hello Mahmoud"
    // fmt.Println(message)

    // const PI float32=3.14
    // fmt.Println(PI)

    // // loops
    // for i:=0;i<5;i++{
    //     fmt.Println(i)
    // }
    // // while
    // j:=0
    // for j<5{
    //     fmt.Println(j)
    //     j++
    // }
    // // infinte loop
    // // for{
    // //     fmt.Println("Mahmoud")
    // // }

    // // switch
    // day:="monday"
    // if(day=="monday"){
    //     fmt.Println("you suggested right")
    // }else {
    //     fmt.Println("you suggested false")
    // }
    // switch day{
    //     // switch has break inside it
    // case "saturday":
    //     fmt.Println("day is saturday") 
    //     fallthrough // Fall through to the next case

    // case "monday":
    //     fmt.Println("day is monday")
    //     fallthrough // make break does not work

    // case "Friday":
    //     fmt.Println("Weekend is coming!")
    //     fallthrough // Fall through to the next case
    // default:
    //     fmt.Println("Not a day")
    // }

    // // array    
    // var arr [3]string
    // arr[0]="Mahmoud"
    // arr[1]="Ali"
    // arr[2]="Hassan"

    // for i:=0;i<len(arr);i++{
    //     fmt.Println(arr[i])
    // }

    // myarr:=[3] int{1,3,5}
    // fmt.Println(myarr)
    // Cities:=[...]string{"m","a","h","m","o","u","d"} // dynamic length
    // fmt.Println(Cities)
    // theArr:=[2][2] int{
    //     {1,2},
    //     {3,4},
    // }
    // fmt.Println(theArr)
    // countries:=[]string {"Egypt","Sua"}
    // countries=append(countries,"palestine")
    // fmt.Println(countries)
    // sum,sub,multiply,division:=arith(3,2)
    // fmt.Println(sum)
    // fmt.Println(sub)
    // fmt.Println(multiply)
    // fmt.Println(division)
    // multy(1,2,3)
    // multy("a","b","c")


    // go justloop()
    // // go here make it run in a seperated thread but main
    // // will not wait it to return so main run synchronous and leave end main
    // fmt.Println("End")
    // time.Sleep(time.Second)



    // Call a function that might panic
        // fmt.Println("Starting the program")
        // triggerPanic()
        // fmt.Println("Program continues after panic") // This line won't be reached if panic isn't recovered

        fmt.Println("Starting the program")
        safeTriggerPanic()
        fmt.Println("Program continues after panic")
        res,err:=division(10,2)
        if(err!=nil){
            fmt.Println(err)
        }else{
            fmt.Println(res)
        }
        res1,err1:=division(10,0)
        if(err1!=nil){
            fmt.Println(err1)
        }else{
            fmt.Println(res1)
        }

}
func triggerPanic() {
    
    defer fmt.Println("Exiting triggerPanic")
    // Defer function to run before the function exits

    // Panic here
    panic("Something went wrong!")

    // This code won't run because of the panic
    fmt.Println("This won't be printed")
    
}

func safeTriggerPanic() {
    // Defer function to run before the function exits
    defer fmt.Println("Exiting safeTriggerPanic")

    // Recover from panic to prevent the program from crashing
    // The recover function in Go is used to recover from a panic and resume normal execution
    defer func() {
    recover()
    }()

    // Panic here
    panic("Something went wrong!")


}

func arith(x, y int) (int, int, int, int) {
    return x + y, x - y, x * y, x / y
}

func multy(param ...interface{}){
    for i:=0;i<len(param);i++{
    fmt.Println(param[i])
    }
    fmt.Println("================")
}

func printName(name string)(string){
    return name
}
func printAge(age int)(ag int){
    ag=age+20
    return
}

func justloop(){
    for i:=0;i<5;i++{
        fmt.Println("line")
        
    }
}
func division(x int ,y int)(int ,error){
    if(y==0){
        return 0, errors.New("Divison by zero is crime")
    } 
    return x/y,nil
     
}