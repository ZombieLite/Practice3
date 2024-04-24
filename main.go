package main

import (
	"fmt"
)

func isPrime(n int) bool {

	// Если число меньше или равно 1, оно не является простым.
	if n <= 1 {
		return false
	}

	// Если число меньше или равно 3, оно является простым.
	if n <= 3 {
		return true
	}

	// Если число делится на 2 или 3, оно не является простым.
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Проверяем деление на нечетные числа, начиная с 5 и шагая на 6. Т.к до 5 мы проверили.
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	// Если число не делится ни на одно из проверенных чисел, оно является простым.
	return true
}

// Функция getPrimesInRange возвращает слайс простых чисел в заданном диапазоне.
// принимает два аргумента: минимальное и максимальное значения диапазона.
func getPrimesInRange(min, max int) []int {
	// Создаем пустой слайс для хранения простых чисел.
	primes := make([]int, 0)
	// Проверяем каждое число в диапазоне и добавляем простые числа в слайс.
	for i := min; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	// Возвращаем слайс простых чисел.
	return primes
}

func main() {
	primes := getPrimesInRange(11, 20)
	fmt.Println(primes)

	primes = getPrimesInRange(51, 120)
	fmt.Println(primes)

	primes = getPrimesInRange(24, 32)
	fmt.Println(primes)
}
