# Go Language Exercise Book

## 1. Basics: Variables, Data Types, and Constants

1. Write a program that declares variables of different types (int, string, bool, float64), assigns values to them, and prints them to the screen using fmt.Println.

2. Create a constant named PI with the value 3.14159 and use it to calculate the area of a circle based on the radius (int) entered from the keyboard. Print the result.

3. Implement a program that takes two ints from the user, performs arithmetic operations (+, -, *, /), and prints the results using short variable declaration (:=).

## 2. Control Structures: if, switch, for

1. Write a program that checks if the entered number (int) is even or odd using if-else, and prints the corresponding message.

2. Use switch to handle the day of the week (entered as int from 1 to 7), printing "Weekday" for 1-5 and "Weekend" for 6-7. Add default for invalid input.

3. Implement a for loop that sums all numbers from 1 to N (entered by the user) and prints the sum.

## 3. Functions

1. Write a function add(a, b int) int that returns the sum of two ints, and call it in main with example values, printing the result.

2. Create a function factorial(n int) int that calculates the factorial of a number using recursion, and test it for n=5 in main.

3. Implement a function swap(a, b string) (string, string) that swaps two strings and returns them. Call it in main and print.

## 4. Arrays, Slices, and Maps

1. Create an array [5]int, fill it with numbers from 1 to 5, and print all elements using for-range.

2. Write a program that creates a string slice, adds 3 elements to it using append, removes the second element (using slicing), and prints the updated slice.

3. Implement a map map[string]int for storing names and ages (add 3 pairs), then print all keys and values using for-range.

## 5. Structures and Methods

1. Define a structure Person with fields Name (string) and Age (int), create an instance, and print its fields in main.

2. Add a method greet(p *Person) string to the Person structure that returns the string "Hello, I'm [Name]". Call the method and print the result.

3. Create a structure Rectangle with fields Width and Height (float64), add a method area() float64 to calculate the area, and test it in main.

## 6. Interfaces

1. Define an interface Shape with a method area() float64. Create a structure Circle implementing this interface (area = πr²), and call area() in main.

2. Extend the Shape interface with a method perimeter() float64. Add a structure Rectangle implementing the interface, and in main call both methods for different shapes.

3. Write a function printArea(s Shape) that takes a Shape interface and prints its area. Test with Circle and Rectangle.

## 7. Goroutines and Channels

1. Launch a simple goroutine using go that prints "Hello from goroutine" after a 1-second delay (time.Sleep), and in main print "Hello from main".

2. Create a channel ch chan int, send the number 42 to it from a goroutine, receive it in main, and print.

3. Implement two goroutines: one sends numbers 1-5 to the channel, the other receives and prints them. Use close(ch) after sending.

## 8. Error Handling

1. Write a function divide(a, b float64) (float64, error) that returns a/b or an error if b=0. Handle it in main.

2. Use errors.New to create an error in a function readFile(filename string) error, simulating file reading. Check the error in main with if err != nil.

3. Implement a function sqrt(x float64) (float64, error) with a check for x < 0, using math.Sqrt. Handle the error and print the value or message.

## 9. Packages and Modules

1. Create a new module with go mod init example.com/mymod, add a function hello() in a separate file, import and call it in main.

2. Write a package mathutils with a function max(a, b int) int, import it in main, and use it for two numbers.

3. Add an exported constant PI to the mathutils package, import it as . (blank import) for side effects, and use it in main.

## 10. Testing

1. Write a function sum(a, b int) int and create a test file _test.go with TestSum, checking sum(2, 3) == 5 with t.Errorf.

2. Add a test table for sum with multiple cases (including negative numbers) using t.Run.

3. Write a function isEven(n int) bool and tests, including TestIsEven with subtests for even/odd and zero.