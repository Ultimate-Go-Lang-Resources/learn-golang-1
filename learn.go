package main

import (
	"bytes"         //Imports the bytes package
	"encoding/json" //Imports the json package
	"errors"        //Imports the errors pasckage
	"fmt"           //Imports the formating package
	"math"          //Imports the math package
	"math/rand"     //Imports the math/rand package
	"os"            //Imports the os package
	"regexp"        //Imports the regexp package
	"runtime"       //Imports the runtime pacage
	"sort"          //Imports the sort package
	"strings"       //Imports the strings package
	"sync"          //Imports the sync package
	"sync/atomic"   //Imports the sync/atomic package
	"time"          //Imports the time package
)

var je = fmt.Println //Replacing the fmt.Println with a variable

//Global variable declaration
var g int
var g1 int = 150
var m int = 25

//Entry point of the program
func main() {
	fmt.Println("Hello world") //Prints on the screen the word Hello world
	fmt.Println(5 + 2)         //Prints the result of 5+2
	fmt.Println(5 * 2)         //Prints the result of 5/2
	var a int                  //a is a variable of type integer
	a = 50 + 25
	fmt.Println("a is:", a)              //Prints the string a is and then the value of a
	fmt.Printf("a is of type %T\n\n", a) //%T tells what type of variable is a
	var x uint64 = 50

	y := 75

	fmt.Printf("x is of type %T\n", x) //Tells what type of variable is x
	fmt.Printf("y is of type %T\n", y) //Tells what type of variable is y

	var j, k, l = 12, 13, "hello world" //Assigning variables to a specific value

	fmt.Println(j)                     //Prints the value of j
	fmt.Println(k)                     //Prints the value of k
	fmt.Println(l)                     //Prints the value of l
	fmt.Printf("j is of type %T\n", j) //Tells what type of variable is j
	fmt.Printf("k is of type %T\n", k) //Tells what type of variable is k
	fmt.Printf("l is of type %T\n", l) //Tells what type of variable is l

	fmt.Println("Hello\t\t\tWorld!")         // \t stands for tab creates space between two elements
	fmt.Println("Alarm ringing or bell\a\a") // \a stands for alarm
	/*It is a good programming practice to define constants in CAPITALS.*/
	const LENGTH int = 5 //Initialized a constant variable of type int
	const WIDTH int = 20 //Initialized a constant variable of type int
	var s int

	s = LENGTH * WIDTH    //The sum of the two constant variables
	fmt.Println("s =", s) //Prints the value of s

	if s > 100 { //If the statement is true prints the following string
		fmt.Println("s is bigger than 100")
	} else if s < 100 { //If the statement is true prints the following string
		fmt.Println("s is smaller than 100")
	} else { //If none of the statements is true prints the following string
		fmt.Println("s is equal to 100")
	}
	/* local variable definition */
	var piket int = 95
	var nota int = 5

	switch nota { //Looks the value of a variable and than prints the adequate case
	case 5:
		fmt.Println("Nota juaj eshte 5")
	case 4:
		fmt.Println("Nota juaj eshte 4")
	case 3:
		fmt.Println("Nota juaj eshte 3")
	case 2:
		fmt.Println("Nota juaj eshte 2")
	case 1:
		fmt.Println("Nota juaj eshte 1")
	}
	switch piket {
	case 95:
		fmt.Println("Shkelqyeshem")
	case 84:
		fmt.Println("Shume Mire")
	case 73:
		fmt.Println("Mire")
	case 62:
		fmt.Println("Mjaftueshem")
	case 51:
		fmt.Println("Dobet")
	default: //If none of the above options is true
		fmt.Println("Duhet te perserisni provimin")
	}
	/* local variable definition */
	points := make(chan string)
	points2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		points <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		points2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-points:
			fmt.Println("received", msg1)
		case msg2 := <-points2:
			fmt.Println("received", msg2)
		}
	}

	var i int //local variable definition
	numbers := [6]int{10, 22, 333, 45, 51, 65}

	for i = 0; i < 10; i++ { //If condition is available, then for loop executes as long as condition is true
		fmt.Println("This is the", i+1, "time the loop got executed")
	}
	for i = 0; i < 6; i++ {
		fmt.Println("The", i+1, "member is", numbers[i])
	}
	var p, q int
	//The following program uses a nested for loop to find the prime numbers from 2 to 100
	for p = 2; p < 100; p++ {
		for q = 2; q <= (p / q); q++ {
			if p%q == 0 {
				break // if factor found, not prime
			}
		}
		if q > (p / q) {
			fmt.Printf("%d is prime\n", p)
		}
	}
	var c int = 10 // local variable definition

	/* for loop execution */
	for c < 20 {
		fmt.Printf("value of c: %d\n", c)
		c++
		if c > 15 {
			break // terminate the loop using break statement
		}
		if c == 13 {
			c++
			continue
		}

	}
	var w int = 10 // local variable definition
	/* do loop execution */
LOOP:
	for w < 20 {
		if w == 15 {
			/* skip the iteration */
			w = w + 1
			goto LOOP
		}
		fmt.Printf("Value of w: %d\n", w)
		w++
	}
	/* local variable definition */
	var num1 int = 100
	var num2 int = 200
	var ret, temp1 int
	var num10 int = 100
	var num20 int = 200

	/* calling a function to get max value */
	ret = max(num1, num2)
	string1, string2 := name("Jetlum", "Ajeti")
	fmt.Printf("Max value is : %d\n", ret)
	fmt.Println(string1, string2)

	temp1 = multi(100, 200)

	fmt.Println("Value of x*y:", temp1)

	fmt.Printf("Before swap,value of num1 : %d\n", num1)
	fmt.Printf("Before swap,value of num2 : %d\n", num2)

	/* calling a function to swap the values */
	swap(num1, num2)

	fmt.Printf("After swap,value of num10 : %d\n", num1)
	fmt.Printf("After swap,value of num20 : %d\n", num2)

	fmt.Println("Before swap,value of num10:", num10)
	fmt.Println("Before swap,value of num20:", num20)
	/* calling a function to swap the values.
	 * &a indicates pointer to a ie. address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	swap1(&num10, &num20)

	fmt.Println("After swap,value of num10:", num10)
	fmt.Println("After swap,value of num20:", num20)

	//Declare a function variable
	getSquareRoot := func(g float64) float64 {
		return math.Sqrt(g)
	}

	//Use the function
	fmt.Println("Square root of number 9 is:", getSquareRoot(9))

	// nextNumber is now a function with i as 0
	nextnumber := getSequence()

	// invoke nextNumber to increase i by 1 and return the same
	fmt.Println("Value of the nextnumber for the 1st time is:", nextnumber())
	fmt.Println("Value of the nextnumber for the 2nd time is:", nextnumber())
	fmt.Println("Value of the nextnumber for the 3rd time is:", nextnumber())

	// create a new sequence and see the result, i is 0 again

	nextnumber1 := getSequence()
	/* If nextnumber1 := nextnumber then it prints number 4 and 5 instead of
	1 and 2 in the next following outputs */
	fmt.Println("Value of the nextnumber1 for the 1st time is:", nextnumber1())
	fmt.Println("Value of the nextnumber1 for the 2nd time is:", nextnumber1())

	circle := Circle{x: 0, y: 0, radius: 5}
	fmt.Printf("Circle area: %f", circle.area())

	//local variable declaration
	var a1, b1, c1 int

	/*A program can have same name for local and global variables but
	value of local variable inside a function will take preference*/
	var g1 int = 100
	// actual initialization
	a1 = 10
	b1 = 20
	c1 = a1 + b1
	g = c1 * (a1 + b1)
	//Printing the value of local variables
	fmt.Printf("\nValue of a1 = %d , b1 = %d, c1 = %d, g = %d", a1, b1, c1, g)
	fmt.Println("\nValue of g1:", g1)

	fmt.Printf("value of a1 in main() = %d\n", a1)
	c1 = sum(a1, b1)
	fmt.Printf("value of c1 in main() = %d\n", c1)

	var greeting = "Hello World!"

	fmt.Printf("normal string:")
	fmt.Printf("%s", greeting)
	fmt.Printf("\n")
	fmt.Printf("hex bytes:")
	for as := 0; as < len(greeting); as++ {
		fmt.Printf("%x", greeting[as]) //Every letter of the string is turned into hex values
	}
	fmt.Printf("\n")

	const sampleText = "\xbd\xb2\xbc\x20\xe2\x8c\x98"

	/*q flag escapes unprintable characters,
		with + flag it escapses non-ascii characters as well
	    to make output unambigous  */
	fmt.Printf("quoted string:")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")

	fmt.Println("String length is:")
	fmt.Println(len(greeting)) //Prints the length of the string

	//Concatenating strings
	greetings := []string{"Hello", "World!"}
	fmt.Println(strings.Join(greetings, ""))

	var oi, oj int

	var index [10]int // index is an array of 10 integers

	//Initialize elements of array index to 0
	for oi = 0; oi < 10; oi++ {
		index[oi] = oi + 100 //Set element at location oi to oi+100
	}

	//Output each array element's value
	for oj = 0; oj < 10; oj++ {
		fmt.Printf("Element [%d] = %d\n", oj, index[oj])
	}
	var oii, ojj int
	var twoDarrays = [3][4]int{
		{0, 1, 2, 3},   //  initializers for row indexed by 0
		{4, 5, 6, 7},   //  initializers for row indexed by 1
		{8, 9, 10, 11}, //  initializers for row indexed by 2
	}

	//Output each array element's value
	for oii = 0; oii < 3; oii++ {
		for ojj = 0; ojj < 4; ojj++ {
			fmt.Printf("2Darrays [%d][%d] = %d\n", oii, ojj, twoDarrays[oii][ojj])
		}
	}

	//An int array with 5 elements
	var balance = []int{1000, 2, 3, 17, 50}
	var avg float64

	//Pass array as an argument

	avg = getAverage(balance, 5)

	//Output the returned value
	fmt.Printf("Average value is:%f\n", avg)

	fmt.Printf("Address of a variable:%x\n", &a)
	/*
		var ip *int     //Pointer to an integer
		var fp *float32 //Pointer to a float	*/

	var ap int = 20 //Actual variable declaration
	var ip *int     //Pointer variable declaration

	ip = &ap //Store address of a in pointer variable

	fmt.Printf("Address of a variable: %x\n", &ap)

	//Address stored in pointer variable
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	//Access the value using the pointer */
	fmt.Printf("Value of *ip variable: %d\n", *ip)

	var ptr *int

	fmt.Printf("The value of ptr is:%x\n", ptr) //nil pointer

	/*To check for a nil pointer you can use an if statement as follows:

	if(ptr != nil)      //succeeds if p is not nil
	if(ptr == nil)      //succeeds if p is null 	*/

	const MAX int = 3

	aptr := []int{10, 100, 1000}
	var potr [MAX]*int
	for i = 0; i < MAX; i++ {
		potr[i] = &aptr[i] //Assign the address of integer
	}
	for i = 0; i < MAX; i++ {
		fmt.Printf("Value of a[%d] = %d\n", i, *potr[i])
	}
	var b int
	var ptr1 *int
	var pptr **int

	b = 3000

	/* take the address of var */
	ptr1 = &b

	/* take the address of ptr using address of operator & */
	pptr = &ptr1

	/* take the value using pptr */
	fmt.Printf("Value of b = %d\n", b)
	fmt.Printf("Value available at *ptr1 = %d\n", *ptr1)
	fmt.Printf("Value available at **pptr = %d\n", **pptr)

	//Local variable definition
	var aa int = 100
	var bb int = 200

	fmt.Printf("Before swap, value of aa : %d\n", aa)
	fmt.Printf("Before swap, value of bb : %d\n", bb)

	/*Calling a function to swap the values.
	 * &a6 indicates pointer to aa ie. address of variable aa and
	 * &b6 indicates pointer to bb ie. address of variable bb.
	 */
	swapptr(&aa, &bb)

	fmt.Printf("After swap, value of aa : %d\n", aa)
	fmt.Printf("After swap, value of bb : %d\n", bb)

	var Book1 Books //Declare Book1 of type Books
	var Book2 Books //Declare Book2 of type Books

	//Book specification
	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Programming Tutorial"
	Book1.book_id = 6495407

	Book2.title = "Telecom Billing"
	Book2.author = "Zara Ali"
	Book2.subject = "Telecom Billing Tutorial"
	Book2.book_id = 6495700

	//Print Book1 and Book2 info
	fmt.Println("Book1 title :", Book1.title)
	fmt.Println("Book1 author :", Book1.author)
	fmt.Println("Book1 subject :", Book1.subject)
	fmt.Println("Book1 book_id :", Book1.book_id)

	fmt.Println("Book2 title :", Book2.title)
	fmt.Println("Book2 author :", Book2.author)
	fmt.Println("Book2 subject :", Book2.subject)
	fmt.Println("Book2 book_id :", Book2.book_id)

	var Book3 Books //Declare Book3 of type Books
	var Book4 Books //Declare Book4 of type Books

	//Book specification
	Book3.title = "C++ Programming"
	Book3.author = "Bjarne Stroustrup"
	Book3.subject = "C++ Programming Tutorial"
	Book3.book_id = 6495410

	Book4.title = "C Programming"
	Book4.author = "Dennis Ritchie"
	Book4.subject = "C Programming Tutorial"
	Book4.book_id = 6495710

	printBook(Book3) //Print Book3 info

	printBook(Book4) //Print Book4 info

	var Book5 Books //Declare Book5 of type Books
	var Book6 Books //Declare Book6 of type Books

	//Book specification
	Book5.title = "C# Programming"
	Book5.author = "Anders Hejlsberg"
	Book5.subject = "C# Programming Tutorial"
	Book5.book_id = 6495411

	Book6.title = "Java Programming"
	Book6.author = "James Gosling"
	Book6.subject = "Java Programming Tutorial"
	Book6.book_id = 6495711

	structptr(&Book5) //Print Book3 info using pointer

	structptr(&Book6) //Print Book4 info using pointer

	/*	To define a slice, you can declare it as an array
		without specifying size or use make function to create the one.
			var numbers []int //a slice of unspecified size
			numbers == []int{0,0,0,0,0}
			numbers = make([]int,5,5)  //a slice of length 5 and capacity 5 */

	var numbers1 = make([]int, 5, 6) //len can not be larger than cap

	printSlice(numbers1)

	var numbers2 []int

	printSlice(numbers2)

	if numbers2 == nil {
		fmt.Println("Slice is nil")
	}

	//Create a slice

	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers3)

	//Print the original slice
	fmt.Println("Numbers == ", numbers3)

	//Print the sub slice starting from index 1(included) to index 4(exluded)
	fmt.Println("Numbers[1:4] ==", numbers3[1:4])

	//Missing lower bound implies 0
	fmt.Println("Numbers[0:3] ==", numbers3[:3])

	//Missing upper bound implies len(s)
	fmt.Println("Numbers[4:] ==", numbers3[4:])

	numbers4 := make([]int, 0, 5)
	printSlice(numbers4)

	//Print the sub slice starting from index 0(included)to index 2(exluded)
	numbers5 := numbers3[:2]
	printSlice(numbers5)

	//Print the sub slice starting form index 2(included)to index 5(exluded)
	numbers6 := numbers3[2:5]
	printSlice(numbers6)

	var numbers7 []int
	printSlice(numbers7)

	// append allows nil slice
	numbers7 = append(numbers7, 0)
	printSlice(numbers7)

	//Add one element to slice
	numbers7 = append(numbers7, 1)
	printSlice(numbers7)

	//Add more than one element at a time
	numbers7 = append(numbers7, 2, 10, 4)
	printSlice(numbers7)

	//Create a slice numbers8 with double the capacity of earlier slice
	numbers8 := make([]int, len(numbers7), (cap(numbers7) * 2))

	//Copy content of numbers7 to numbers8
	copy(numbers8, numbers7)
	printSlice(numbers8)

	/*fjala := []string{"Bitsapphire"}
	printSlice1(fjala)

	fjala1 := fjala[:5]
	printSlice1(fjala1)*/

	//Creating a slice

	nr := []int{1, 5, 3, 4, 6, 2, 8, 7}

	//Print the numbers

	for i := range nr {
		fmt.Println("Slice item", i, "is", nr[i])
	}

	//Create a map

	CapitalMap := map[string]string{"Kosovo": "Pristina", "Albania": "Tirana", "USA": "Washington D.C."}

	//Print map using keys
	for country1 := range CapitalMap {
		fmt.Println("Capital of", country1, "is", CapitalMap[country1])
	}

	//Print map using key value
	for country1, capital := range CapitalMap {
		fmt.Println("Capital of", country1, "is", capital)
	}

	var countryCapitalMap map[string]string

	//Create a map

	countryCapitalMap = make(map[string]string)

	//Insert key value pairs in the map

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Germany"] = "Berlin"
	countryCapitalMap["England"] = "London"

	//Print map using keys
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	//Test if entry is present in the map or not
	capital, ok := countryCapitalMap["Spain"]
	//If ok is true,entry is present otherwise entry is absent
	if ok {
		fmt.Println("Capital of Spain is", capital)
	} else {
		fmt.Println("Capital of Spain is not present")
	}

	//Delete an entry
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("Updated map")

	//Print map
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	var factorial1 int
	var I int = 15
	factorial1 = factorial(I)
	fmt.Printf("Factorial of %d is %d\n", I, factorial1)

	var i1 int
	fmt.Println("Fibonacci Series")
	for i1 = 0; i1 < 10; i1++ {
		fmt.Println("\t", fibonaci(i1))
	}

	var sum int = 25
	var count int = 5
	var equal float32

	equal = float32(sum) / float32(count) //Converting sum and count var from int into float32
	fmt.Println("Value of equal:", equal)

	circle1 := Circle1{v1: 5, b1: 10, radius1: 15}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println("Circle area:", getArea(circle1))
	fmt.Println("Rectangle area:", getArea(rectangle))

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	//Channel
	messages := make(chan string)

	go func() { messages <- "Go language" }()

	msg := <-messages
	fmt.Println(msg)

	//Channel Buffering
	message := make(chan string, 2)
	message <- "C++"
	message <- "Programming"

	fmt.Println(<-message)
	fmt.Println(<-message)

	//Channel Synchronization
	done := make(chan bool, 1)
	go worker1(done)

	<-done

	//Channel Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	//Timeouts

	cc1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		cc1 <- "result 1"
	}()
	select {
	case res := <-cc1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	cc2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		cc2 <- "result 2"
	}()

	select {
	case res := <-cc2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

	/*Here’s a non-blocking receive. If a value is available on messages
	then select will take the <-messages case with that value. If not it
	will immediately take the default case.
	If me remove 1from chan string it will print the default messages*/
	messages1 := make(chan string, 1)
	signals := make(chan bool)
	select {
	case msg := <-messages1:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg1 := "hi"
	select {
	case messages1 <- msg1:
		fmt.Println("sent message", msg1)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg1 := <-messages1:
		fmt.Println("received message", msg1)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	//Closing Channels

	jobs := make(chan int, 5)
	done1 := make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done1 <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done1

	//Range over channels

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	/*This range iterates over each element as it’s received from queue.
	Because we closed the channel above, the iteration terminates after receiving the 2 elements.
	If we didn’t close it we’d block on a 3rd receive in the loop.*/
	for elem := range queue {
		fmt.Println(elem)
	}

	//Timers

	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	/*<-timer2.C
	fmt.Println("Timer 2 expired")
	The first timer will expire ~2s after we start the program,
	but the second should be stopped before it has a chance to expire.
	The <-timer1.C blocks on the timer’s channel C until it sends a value
	indicating that the timer expired.*/
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	//Tickers

	/*Tickers use a similar mechanism to timers: a channel that is sent values.
	Here we’ll use the range builtin on the channel to iterate over the values as they arrive every 500ms.
	Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more values on its channel.
	We’ll stop ours after 1600ms.
	Timers are for when you want to do something once in the future - tickers are for when you want to do something
	repeatedly at regular intervals. Here’s an example of a ticker that ticks periodically until we stop it.*/

	j1 := time.Now()
	switch {
	case j1.Hour() < 14:
		fmt.Println("It's before noon exactly:", j1.Hour(), ":", j1.Minute(), ":", j1.Second())
	default:
		fmt.Println("It's after noon exactly:", j1.Hour(), ":", j1.Minute(), ":", j1.Second())
	}
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at:", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")

	//Worker pools

	job := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go worker(w, job, results)
	}
	for j := 1; j <= 9; j++ {
		job <- j
		//Rate limiting
		/*
			Rate limiting is an important mechanism for controlling resource utilization and maintaining quality of service.
			Go elegantly supports rate limiting with goroutines, channels, and tickers.*/
		requests := make(chan int, 5)
		for z := 1; z <= 5; z++ {
			requests <- z
		}
		close(requests)

		limiter := time.Tick(time.Millisecond * 200)

		for req := range requests {
			<-limiter
			fmt.Println("request", req, time.Now())
		}

		burstyLimiter := make(chan time.Time, 3)

		for z := 0; z < 3; z++ {
			burstyLimiter <- time.Now()
		}

		go func() {
			for t := range time.Tick(time.Millisecond * 200) {
				burstyLimiter <- t
			}
		}()

		burstyRequests := make(chan int, 5)
		for z := 1; z <= 5; z++ {
			burstyRequests <- z
		}
		close(burstyRequests)

		for req := range burstyRequests {
			<-burstyLimiter
			fmt.Println("request", req, time.Now())

		}
	}

	close(job)

	for a := 1; a <= 9; a++ {
		<-results
	}

	//Atomic Counters
	/*
		The primary mechanism for managing state in Go is communication over channels.
		We saw this for example with worker pools. There are a few other options for managing state though.
		 Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.*/
	var ops uint64 = 0

	for i := 50; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()

	}

	time.Sleep(time.Second)

	opsfinal := atomic.LoadUint64(&ops)

	fmt.Println("ops", opsfinal)

	//Mutexes
	/*In the previous example we saw how to manage simple counter state using atomic operations.
	For more complex state we can use a mutex to safely access data across multiple goroutines.*/

	var state = make(map[int]int) //The state will be a map.
	var mutex = &sync.Mutex{}     //This mutex will synchronize access to state.
	var opes int64 = 0            //opes will count how many operations we perform against the state.

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()        //For each read we pick a key to access, Lock() the mutex to ensure exclusive access to the state,
				total += state[key] //read the value at the chosen key,
				mutex.Unlock()      //Unlock() the mutex, and increment the ops count.
				atomic.AddInt64(&opes, 1)
				runtime.Gosched() //In order to ensure that this goroutine doesn’t starve the scheduler, we explicitly yield after each operation with runtime.Gosched().
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&opes, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opesfinal := atomic.LoadInt64(&opes)
	fmt.Println("ops", opesfinal)

	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()

	//Stateful Goroutines

	var opss int64 = 0

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	/*This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel.
	Each read requires constructing a readOp, sending it over the reads channel, and the receiving the result over the provided resp channel.*/
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&opss, 1)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&opss, 1)
			}
		}()
	}
	time.Sleep(time.Second)

	opssfinal := atomic.LoadInt64(&opss)
	fmt.Println("ops:", opssfinal)

	//Sorting

	strs := []string{"j", "e", "t", "l", "u", "m"}
	fmt.Println("Word:", strs)
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	p1 := sort.StringsAreSorted(strs)
	fmt.Println("Sorted:", p1)

	ints := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Numbers:", ints)
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	s1 := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s1)

	//Sorting by Functions

	//Sorting by Functions

	companies := []string{"Facebook", "Apple", "Microsoft", "Google"}
	fmt.Println("Unsorted", companies)
	sort.Sort(ByLength(companies))
	fmt.Println("After sorting", companies)

	/*Panic
	A panic typically means something went unexpectedly wrong.
	Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}*/

	/*Defer

	Defer is used to ensure that a function call is performed later in a program’s execution, usually for purposes of cleanup.
	defer is often used where e.g. ensure and finally would be used in other languages.*/

	f := createFile("/dekstop/defer.txt")
	defer closeFile(f)
	writeFile(f)

	//Collection Functions
	//Here we try out our various collection functions.
	var strs1 = []string{"ae", "be", "ce", "db"}
	fmt.Println(Index(strs1, "ce"))
	fmt.Println(Include(strs1, "e"))
	fmt.Println(Any(strs1, func(v string) bool {
		return strings.HasPrefix(v, "b")
	}))
	fmt.Println(All(strs1, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs1, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	//The above examples all used anonymous functions, but you can also use named functions of the correct type.
	fmt.Println(Map(strs, strings.ToUpper))

	//String Functions
	/*The standard library’s strings package provides many useful string-related functions.
	Here are some examples to give you a sense of the package.*/
	je("Contains:  ", strings.Contains("test", "es"))
	je("Count:     ", strings.Count("test", "t"))
	je("HasPrefix: ", strings.HasPrefix("test", "te"))
	je("HasSuffix: ", strings.HasSuffix("test", "st"))
	je("Index:     ", strings.Index("test", "e"))
	je("Join:      ", strings.Join([]string{"a", "b"}, " dhe "))
	je("Repeat:    ", strings.Repeat("a", 5))
	je("Replace:   ", strings.Replace("foo", "o", "0", -1))
	je("Replace:   ", strings.Replace("foo", "o", "0", 1))
	je("Split:     ", strings.Split("a-b-c-d-e", "-"))
	je("ToLower:   ", strings.ToLower("TEST"))
	je("ToUpper:   ", strings.ToUpper("test"))
	je()
	je("Len:", len("hello"))
	je("Char:", "hello"[1])

	//String Formatting
	/*Go offers excellent support for string formatting in the printf tradition.
	Here are some examples of common string formatting tasks.*/

	j5 := point{1, 2}
	fmt.Printf("%v\n", j5) //This prints an instance of our point struct.

	fmt.Printf("%+v\n", j5) //If the value is a struct, the %+v variant will include the struct’s field names.

	fmt.Printf("%#v\n", j5) //The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.

	fmt.Printf("%T\n", j5) //To print the type of a value, use %T.

	fmt.Printf("%t\n", true) //Formatting booleans is straight-forward.

	fmt.Printf("%d\n", 123) //There are many options for formatting integers. Use %d for standard, base-10 formatting.

	fmt.Println("%b\n", 14) //This prints a binary representation.

	fmt.Printf("%c\n", 33) //This prints the character corresponding to the given integer.

	fmt.Printf("%x\n", 456) //%x provides hex encoding.

	fmt.Printf("%f\n", 78.9) //There are also several formatting options for floats. For basic decimal formatting use %f.

	fmt.Printf("%e\n", 123400000.0) //%e and %E format the float in
	fmt.Printf("%E\n", 123400000.0) //(slightly different versions of) scientific notation.

	fmt.Printf("%s\n", "\"string\"") //For basic string printing use %s.

	fmt.Printf("%q\n", "\"string\"") //To double-quote strings as in Go source, use %q.

	fmt.Printf("%x\n", "hex this") //As with integers seen earlier, %x renders the string in base-16, with two output characters per byte of input.

	fmt.Printf("%p\n", &p) //To print a representation of a pointer, use %p.

	fmt.Printf("|%6d|%6d|\n", 12, 345) //When formatting numbers you will often want to control the width and precision of the resulting figure. To specify the width of an integer, use a number after the % in the verb. By default the result will be right-justified and padded with spaces.

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45) //You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //To left-justify, use the - flag.

	fmt.Printf("|%6s|%6s|\n", "foo", "b") //You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b") //To left-justify use the - flag as with numbers.

	s2 := fmt.Sprintf("a %s", "string") //So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.
	fmt.Println(s2)

	fmt.Fprintf(os.Stderr, "an %s\n", "error") //You can format+print to io.Writers other than os.Stdout using Fprintf.

	//Regular Expressions
	//Go offers built-in support for regular expressions

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach") //This tests whether a pattern matches a string.
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch") //Above we used a string pattern directly, but for other regexp tasks you’ll need to Compile an optimized Regexp struct.

	fmt.Println(r.MatchString("peach")) //Many methods are available on these structs. Here’s a match test like we saw earlier.

	fmt.Println(r.FindString("peach punch")) //This finds the match for the regexp.

	fmt.Println(r.FindStringIndex("peach punch")) //This also finds the first match but returns the start and end indexes for the match instead of the matching text.

	fmt.Println(r.FindStringSubmatch("peach punch")) //The Submatch variants include information about both the whole-pattern matches and the submatches within those matches. For example this will return information for both p([a-z]+)ch and ([a-z]+).

	fmt.Println(r.FindStringSubmatchIndex("peach punch")) //Similarly this will return information about the indexes of matches and submatches.

	fmt.Println(r.FindAllString("peach punch pinch", -1)) //The All variants of these functions apply to all matches in the input, not just the first. For example to find all matches for a regexp.

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1)) //These All variants are available for the other functions we saw above as well.

	fmt.Println(r.FindAllString("peach punch pinch", 2)) //Providing a non-negative integer as the second argument to these functions will limit the number of matches.

	fmt.Println(r.Match([]byte("peach"))) //Our examples above had string arguments and used names like MatchString. We can also provide []byte arguments and drop String from the function name.

	r = regexp.MustCompile("p([a-z]+)ch") //When creating constants with regular expressions you can use the MustCompile variation of Compile. A plain Compile won’t work for constants because it has 2 return values.
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) //The regexp package can also be used to replace subsets of strings with other values.

	in := []byte("a peach") //The Func variant allows you to transform matched text with a given function.
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	//JSON
	//Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.

	bolB, _ := json.Marshal(true) //First we’ll look at encoding basic data types to JSON strings. Here are some examples for atomic values.
	fmt.Println(string(bolB))
	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} //And here are some for slices and maps, which encode to JSON arrays and objects as you’d expect.
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{ //The JSON package can automatically encode your custom data types. It will only include exported fields in the encoded output and will by default use those names as the JSON keys.
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{ //You can use tags on struct field declarations to customize the encoded JSON key names. Check the definition of Response2 above to see an example of such tags.
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`) //Now let’s look at decoding JSON data into Go values. Here’s an example for a generic data structure.

	var dat map[string]interface{} //We need to provide a variable where the JSON package can put the decoded data. This map[string]interface{} will hold a map of strings to arbitrary data types.

	if err := json.Unmarshal(byt, &dat); err != nil { //Here’s the actual decoding, and a check for associated errors.
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) //In order to use the values in the decoded map, we’ll need to cast them to their appropriate type. For example here we cast the value in num to the expected float64 type.
	fmt.Println(num)

	strs2 := dat["strs2"].([]interface{}) //Accessing nested data requires a series of casts.
	str1 := strs2[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}` //We can also decode JSON into custom data types. This has the advantages of adding additional type-safety to our programs and eliminating the need for type assertions when accessing the decoded data.
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout) //In the examples above we always used bytes and strings as intermediates between the data and JSON representation on standard out. We can also stream JSON encodings directly to os.Writers like os.Stdout or even HTTP response bodies.
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}

/* function returning the max between two numbers */
func max(num3, num4 int) int {
	/* local variable declaration */
	var result int

	if num3 > num4 {
		result = num3
	} else {
		result = num4
	}
	return result
}
func name(x, y string) (string, string) {
	return x, y
}

func multi(x, y int) int {
	var temp int

	temp = x * y

	return temp
}
func swap(x, y int) int {
	var temp2 int

	temp2 = x /* save the value of x */
	x = y     /* put y into x */
	y = temp2 /* put temp into y */

	return temp2
}

func swap1(x *int, y *int) {
	var temp3 int

	temp3 = *x // save the value at address x
	*x = *y    // put y into x
	*y = temp3 // put temp3 to y
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/* define a circle */
type Circle struct {
	x, y, radius float64
}

/* define a method for circle */
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

/*Function parameters, formal parameters, are treated as local variables
with-in that function and they will take preference over the global variables*/
/* function to add two integers */
func sum(a1, b1 int) int {
	fmt.Printf("value of a1 in sum() = %d\n", a1)
	fmt.Printf("value of b1 in sum() = %d\n", b1)

	return a1 + b1
}

/*Initializing Local and Global Variables
When a local variable as Global variables are initialized to their corresponding 0 value. Pointer is initialized to nil.

Data Type	Initial Default Value
int	   				0
float32				0
pointer				nil 				*/

func getAverage(arr []int, size int) float64 {
	var i, sum int
	var avg float64

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float64(sum / size)

	return avg
}

func swapptr(x *int, y *int) {
	var temp6 int

	temp6 = *x // save the value at address x
	*x = *y    // put y into x
	*y = temp6 // put temp3 to y
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func printBook(book Books) {

	fmt.Println("Book title :", book.title)
	fmt.Println("Book author :", book.author)
	fmt.Println("Book subject :", book.subject)
	fmt.Println("Book book_id :", book.book_id)
}

func structptr(book *Books) {

	fmt.Println("Book title :", book.title)
	fmt.Println("Book author :", book.author)
	fmt.Println("Book subject :", book.subject)
	fmt.Println("Book book_id :", book.book_id)
}

func printSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func printSlice1(x []string) {

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func factorial(i int) int {
	if i <= 1 {
		return 1
	} else {
		return i * factorial(i-1)
	}
}

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

//Define an interface
type Shape interface {
	area() float64
}

//Define a circle
type Circle1 struct {
	v1, b1, radius1 float64
}

//Define a rectangle
type Rectangle struct {
	width, height float64
}

//Define a method for circle(implementation of Shape.area())
func (circle1 Circle1) area() float64 {
	return math.Pi * circle1.radius1 * circle1.radius1
}

//Define a method for rectangle(implemntation of Shape.area())
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

//Define a method for shape
func getArea(shape Shape) float64 {
	return shape.area()
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg + 3, nil
}

func worker1(done chan bool) {
	fmt.Printf("Working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

/* If
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
returns strings in order of the shortest to the largest*/

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

//Returns the first index of the target string t, or -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//Returns true if the target string t is in the slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//Returns true if one of the strings in the slice satisfies the predicate f.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//Returns true if all of the strings in the slice satisfy the predicate f.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//Returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

//Returns a new slice containing the results of applying the function f to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

type point struct {
	x, y int
}

func worker(id int, job <-chan int, results chan<- int) {
	for j := range job {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

type Response1 struct { //We’ll use these two structs to demonstrate encoding and decoding of custom types below.
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}
