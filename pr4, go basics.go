package main

import (
	"fmt"
	"math"
	"strings"
)

//Задачи для практической работы на языке Go
//
//1. Задачи на линейное программирование (без условных операторов и циклов)
//
//1. Сумма цифр числа:
//Напишите программу, которая принимает целое число и вычисляет сумму его цифр.
//	Вход: `1234`
//	Выход: `10` (1 + 2 + 3 + 4)
//

func sumDigits(number int) int {
	if number == 0 {
		return 0
	}
	return number%10 + sumDigits(number/10)
}

// 2. Преобразование температуры:
// Напишите программу, которая преобразует температуру из градусов Цельсия в Фаренгейты и обратно.
// 	Вход: `25 (Celsius)`
// 	Выход: `77 (Fahrenheit)`
func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// 3. Удвоение каждого элемента массива:
// Напишите программу, которая принимает массив чисел и возвращает новый массив, где каждое число удвоено.
// 	Вход: `[1, 2, 3, 4]`
// 	Выход: `[2, 4, 6, 8]`
func doubleArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	return append([]int{arr[0] * 2}, doubleArray(arr[1:])...)
}

// 4. Объединение строк:
// Напишите программу, которая принимает несколько строк и объединяет их в одну строку через пробел.
// 	Вход: `["Hello", "world"]`
// 	Выход: `Hello world`
func joinStrings(strs []string) string {
	return strings.Join(strs, " ")
}

//5. Расчет расстояния между двумя точками:
//Напишите программу, которая вычисляет расстояние между двумя точками в 2D пространстве.
//	Вход: `(x1=1, y1=1), (x2=4, y2=5)`
//	Выход: `5.0`
//

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

//Вызов фунций

func firstBlock() {

	number := 1234
	fmt.Println("1.Сумма цифр:", sumDigits(number))
	//-----
	celsius := 25.0
	fmt.Printf("2. %.2f Celsius = %.2f Fahrenheit\n", celsius, celsiusToFahrenheit(celsius))
	fmt.Printf("2. %.2f Fahrenheit = %.2f Celsius\n", 77.0, fahrenheitToCelsius(77.0))
	arr := []int{1, 2, 3, 4}
	fmt.Println("3. Удвоенный массив:", doubleArray(arr))
	strs := []string{"Hello", "world"}
	fmt.Println("4.Объединенная строка:", joinStrings(strs))
	fmt.Printf("5.Расстояние между точками: %.2f\n", distance(1, 1, 4, 5))
}

//2. Задачи с условным оператором
//
//1. Проверка на четность/нечетность:
//Напишите программу, которая проверяет, является ли введенное число четным или нечетным.
//	Вход: `4`
//	Выход: `Четное`
//

// 2. Проверка высокосного года:
// Напишите программу, которая проверяет, является ли введенный год високосным.
// 	Вход: `2020`
// 	Выход: `Високосный`
//
// 3. Определение наибольшего из трех чисел:
// Напишите программу, которая принимает три числа и выводит наибольшее из них.
// 	Вход: `4, 9, 7`
// 	Выход: `9`
//
// 4. Категория возраста:
// Напишите программу, которая принимает возраст человека и выводит, к какой возрастной группе он относится (ребенок, подросток, взрослый, пожилой. В комментариях указать возрастные рамки).
// 	Вход: `25`
// 	Выход: `Взрослый`
//
// 5. Проверка делимости на 3 и 5:
// Напишите программу, которая проверяет, делится ли число одновременно на 3 и 5.
// 	Вход: `15`
// 	Выход: `Делится`
//
// 1. Проверка на четность/нечетность
func isEven(number int) string {
	if number%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}

// 2. Проверка високосного года
func isLeapYear(year int) string {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return "Високосный"
	}
	return "Не високосный"
}

// 3. Определение наибольшего из трех чисел
func maxOfThree(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= c {
		return b
	}
	return c
}

// 4. Категория возраста
func ageCategory(age int) string {
	// Ребенок: 0-12, Подросток: 13-17, Взрослый: 18-64, Пожилой: 65+
	if age >= 0 && age <= 12 {
		return "Ребенок"
	} else if age <= 17 {
		return "Подросток"
	} else if age <= 64 {
		return "Взрослый"
	}
	return "Пожилой"
}

// 5. Проверка делимости на 3 и 5
func isDivisibleBy3And5(number int) string {
	if number%3 == 0 && number%5 == 0 {
		return "Делится"
	}
	return "Не делится"
}

// Объединение задач в блок
func secondBlock() {
	fmt.Println("1. Четность числа 4:", isEven(4))
	fmt.Println("2. Високосный год 2020:", isLeapYear(2020))
	fmt.Println("3. Наибольшее из 4, 9, 7:", maxOfThree(4, 9, 7))
	fmt.Println("4. Возрастная категория для 25 лет:", ageCategory(25))
	fmt.Println("5. Делимость числа 15 на 3 и 5:", isDivisibleBy3And5(15))
}

// 3. Задачи на циклы
//
// 1. Факториал числа:
// Напишите программу, которая вычисляет факториал числа.
// 	Вход: `5`
// 	Выход: `120` (5! = 5 × 4 × 3 × 2 × 1)
//
// 2. Числа Фибоначчи:
// Напишите программу, которая выводит первые `n` чисел Фибоначчи.
// 	Вход: `n = 7`
// 	Выход: `0, 1, 1, 2, 3, 5, 8`
//
// 3. Реверс массива:
// Напишите программу, которая переворачивает массив чисел.
// 	Вход: `[1, 2, 3, 4, 5]`
// 	Выход: `[5, 4, 3, 2, 1]`
//
// 4. Поиск простых чисел:
// Напишите программу, которая выводит все простые числа до заданного числа.
// 	Вход: `n = 20`
// 	Выход: `2, 3, 5, 7, 11, 13, 17, 19`
//
// 5. Сумма чисел в массиве:
// Напишите программу, которая вычисляет сумму всех чисел в массиве.
// 	Вход: `[1, 2, 3, 4, 5]`
// 	Выход: `15`
//
// 1. Факториал числа
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// 2. Числа Фибоначчи
func fibonacci(n int) []int {
	if n == 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}
	seq := fibonacci(n - 1)
	return append(seq, seq[len(seq)-1]+seq[len(seq)-2])
}

// 3. Реверс массива
func reverseArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	return append([]int{arr[len(arr)-1]}, reverseArray(arr[:len(arr)-1])...)
}

// 4. Поиск простых чисел
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func primesUpTo(n int) []int {
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

// 5. Сумма чисел в массиве
func sumArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sumArray(arr[1:])
}

// Объединение задач в блок
func thirdBlock() {
	fmt.Println("1. Факториал 5:", factorial(5))
	fmt.Println("2. Первые 7 чисел Фибоначчи:", fibonacci(7))
	fmt.Println("3. Реверс массива [1, 2, 3, 4, 5]:", reverseArray([]int{1, 2, 3, 4, 5}))
	fmt.Println("4. Простые числа до 20:", primesUpTo(20))
	fmt.Println("5. Сумма чисел массива [1, 2, 3, 4, 5]:", sumArray([]int{1, 2, 3, 4, 5}))
}

func main() {
	fmt.Println("Задачи на линейное программирование ")
	firstBlock()
	fmt.Println("Задачи с условным оператором ")
	secondBlock()
	fmt.Println("Задачи на циклы ")
	thirdBlock()
}
