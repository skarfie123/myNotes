package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func add(x int, y int) int { // can also return multiple results with eg. (int,int)
	return x + y
}

func multiply(x int, y int) (z int) {
	z = x * y
	return // naked return
}

func deferDemo() {
	t := 1
	defer fmt.Println("Defer this", t) // evaluates parameters now but calls function at return, Last-In-First-Out
	t += 1
	defer fmt.Println("Defer this", t)
	t += 1
	fmt.Println("Print this", t)
	return
}

var a int
var b = 2

type Place struct {
	X, Y float32
	Name string
}

var m map[int]Place // for a map that links ints to Places

type Vertex struct {
	X, Y float64
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0 // each adder has its own sum
	return func(x int) int { //returns the adder function
		sum += x
		return sum
	}
}

func (v Vertex) Abs() float64 { // specify Vertex as the reciver to make it a method of the Vertex type
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 { // can only make methods for types defined in the same package
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
func (v *Vertex) Scale(f float64) { // use pointer reciever to modify original, use value receiver to work on a copy
	v.X = v.X * f
	v.Y = v.Y * f
}

type I interface {
	// an interface states requirements to be of that type
	M() // any other type with an M function can be put in an I type interface object
}
type T struct {
	S string
}

func (t *T) M() {
	if t == nil { // handle before t.S to avoid nil pointer error
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

type MyError struct {
	//my type of error
	When time.Time
	What string
}

func (e *MyError) Error() string { //how the error is to be displayed
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}
func run() error { //example function that returns an error, normally only return one if something went wrong, otherwise nil
	return &MyError{ // return an error with the current time and a message
		time.Now(),
		"it didn't work",
	}
}

func mainOriginal() { // program starts in main functipn in main package. ///correct?

	//Functions
	fmt.Println(add(3, 4))
	fmt.Println(multiply(3, 4))
	fmt.Println(rand.Intn(4)) /// why always same number?

	//Variables
	a = 1
	var c int = 3
	d := 4 // shorthand only in functions
	fmt.Println(a, b, c, d)

	//For/While
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// for ; i<10 ; {} = for i<10 {} => while
	// for {} => while true

	//If/Else
	if a > 10 {
		fmt.Println("a>10")
	}
	if v := 2 * a; v > 10 { // short statement, v can only be used in the if and else set
		fmt.Println("v>10", v)
	} else if v > 5 {
		fmt.Println("10>v>5", v)
	} else {
		fmt.Println("v<5", v)
	}

	//Defer
	deferDemo()

	//Pointers
	var p *int = &a
	fmt.Println(*p)

	//Structs
	home := Place{1, 2, "Home"}
	sch := Place{Name: "School"}
	home.X = 5
	fmt.Println(home, sch)
	// if p is pointer to home, (*p).X gives "Home" but shorthand p.x

	//Arrays
	var aa [3]int // currently [000]
	fmt.Println(aa)
	aa[0] = 1
	bb := [3]int{1, 2, 3}
	fmt.Println(aa, bb)

	//Slices
	var cc []int = bb[0:2] // other egs. [:2],[:],[1:]
	var dd []int = bb[1:3]
	fmt.Println(bb, cc, dd) // cc, dd are references to bb
	bb[1] = 4
	fmt.Println(bb, cc, dd) // all changed not just bb
	ee := []int{4, 5, 6}    // literal, hidden array
	// array xx, with slice yy
	// xx => [000000000000000000000]
	// yy =>        [000000000]
	//            len<------->
	//            cap<------------>
	ff := make([]int, 3, 5)
	fmt.Println(ff, len(ff), cap(ff))
	var gg []int // => currently nil
	fmt.Println(ee, gg)
	aaa := [][]int{ //slice of an array of slices
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	fmt.Println(aaa)
	ee = append(ee, 2, 3, 4)
	fmt.Println(ee)

	//Range
	for i, v := range ee { // range gives index and value i,v
		fmt.Println(i, v)
	}
	// i,v => both
	// _,v => just value
	// i = i,_ => just index

	//Maps
	m = make(map[int]Place)
	m[1] = Place{
		40.68433, -74.39967, "Bell Labs",
	}
	fmt.Println(m[1])
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	elem, ok := m["Bell Labs"]
	fmt.Println(elem, ok)
	delete(m, "Bell Labs")
	elem, ok = m["Bell Labs"]
	fmt.Println(elem, ok)

	//Passing Functions
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	//Function Closures
	pos, neg := adder(), adder() //returns a function thats bound to a specific sum
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i), // each one adds to its own sum
			neg(-2*i),
		)
	}

	//Methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
	mf := MyFloat(-math.Sqrt2)
	fmt.Println(mf.Abs())

	//Interfaces
	describe := func(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
	var i I
	i = F(math.Pi)
	describe(i)
	i.M()
	i = &T{"Hello"}
	describe(i)
	i.M()
	var t *T
	i = t
	describe(i)
	i.M()
	var j interface{} // empty interface accepts anything (since zero requirements)
	j = 42
	s, ok := j.(int) // or just s
	fmt.Println(s, ok)
	s2, ok2 := j.(string) // just s2 would call a panic
	fmt.Println(s2, ok2)

	//Type Switch using empty interface
	switch v := j.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
	//Stringer
	//write a String method to make your type printable via a Stringer (such as the Stringer in Sprintf)

	//Errors
	if err := run(); err != nil { //if an error ocurred
		fmt.Println(err)
	}

	//Readers
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8) //slice of 8 bytes
	for {
		n, err := r.Read(b) //read the next 8 bytes, n<=8, err = io.EOF if end of file reached
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n]) //prints as much as we got that round
		if err == io.EOF {
			break
		}
	}
}

func goRoutines() {

	//Go Routines
	say := func(s string) {
		for i := 0; i < 3; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}
	go say("I'm from the other thread") //starts another thread to execute this task
	say("I'm from the main thread")     //executes int he same thread as before
	// the other thread is cut when the main thread finishes unless you sync/wait

	//Channels
	// use channels to send data back from go routines
	sum := func(s []int, c chan int) {
		x := 0
		for _, v := range s {
			x += v
		}
		c <- x // send sum to c
	}
	s := []int{7, 2, 8, -9, 4, 0}
	ch := make(chan int)
	go sum(s[:len(s)/2], ch) // split summation work between two threads
	go sum(s[len(s)/2:], ch)
	x, y := <-ch, <-ch // receive from ch, x or y depends on which order each half finishes, it waits until both results are back
	fmt.Println(x, y, x+y) // then add to get the final sum

	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c) // close channel to indicate no more will be sent, can check with v, ok := <-ch, ok will be false if closed
		// only sender should close, never the reciever
	}
	c := make(chan int, 10) // buffered channel limits the number it can receive (error if exceeded)
	go fibonacci(cap(c), c) // cap gives the buffer length of channel
	for i := range c { //iterates until channel is closed
		fmt.Println(i)
	}

}

func selectBlock() {
	//Select
	// waits until one of the cases can run, if multiple ready, chooses at random
	fibonacci := func(c, quit chan int) {
		x, y := 0, 1
		for { //infinite loop
			select {
			case c <- x: // always ready so keeps feeding into channel
				x, y = y, x+y
			case <-quit: // ready when asked to quit, doesn't have to be 0
				fmt.Println("quit")
				return
			}
		}
	}
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) //receives from channel
		}
		quit <- 0 //asks to quit after 10 recieved
	}()
	fibonacci(c, quit) //starts the fibonacci generator

	// use default when other cases are not ready
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}
// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}
// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}
func syncMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

func main() {
	//mainOriginal()

	//goRoutines()
	//selectBlock()
	syncMutex()
	//see learningGO/channels for mutex and channel use
	//see learningGO/http/client for flag use
}
