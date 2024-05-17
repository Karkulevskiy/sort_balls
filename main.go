package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Идея:
// Нужно смотреть на то, что сможет ли i - й контейнер
// поместить в себя m шариков определенного цвета
// Для этого достаточно посчитать сумму шариков в i - м контейнере
// И соотнести ее с суммой шариков определенного цвета (отсортировав оба массива)

// main Входная точка для запуска алгоритма
func main() {
	containers := GetData() // Получаем контейнеры с шариками

	couldBeSorted := Orginize(containers) // Можно ли отсортировать

	if couldBeSorted {
		fmt.Println("yes") // Отсортировать можно
	} else {
		fmt.Println("no") // Отсортировать нельзя
	}
}

// Orginize проверяет можно ли отсортировать контейнеры
func Orginize(containers [][]int) bool {
	// totalBallsInContainer содержит сумму шариков в каждом контейнере
	totalBallsInContainer := make([]int64, len(containers))

	// Считаем сумму всех шариков в контейнере
	for i := 0; i < len(containers); i++ {
		totalBallsInContainer[i] = Sum(containers[i])
	}

	// totalCertainColors содержит общее количество возможных цветов
	totalCertainColors := make([]int64, len(containers))

	// Считаем цвет шарика в каждом контейнере
	for i := 0; i < len(containers); i++ {
		totalCertainColor := int64(0)

		// Для подсчета бежим по колонке матрицы
		for j := 0; j < len(containers); j++ {
			totalCertainColor += int64(containers[j][i])
		}

		// Сохраняем результат
		totalCertainColors[i] = totalCertainColor
	}

	// Сортируем слайсы в порядке возрастания
	slices.Sort(totalBallsInContainer)
	slices.Sort(totalCertainColors)

	// Смотрим, если наши слайсы равны, значит мы можем выполнить перестановки
	return slices.Equal(totalBallsInContainer, totalCertainColors)
}

// GetData получает данные о кол-ве контейнеров и цветах
func GetData() [][]int {
	scanner := bufio.NewScanner(os.Stdin) // Заводим scanner для считывания с консоли

	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text()) // Получаем общее кол-во контейнеров

	containers := make([][]int, n) // Контейнеры
	for i := 0; i < n; i++ {
		container := make([]int, n) // Текущий контейнер

		scanner.Scan()

		// Преобразуем строку в слайс цветов
		colorsArr := strings.Split(scanner.Text(), " ")

		for j, color := range colorsArr {
			// Преобразуем string в int
			colorInt, _ := strconv.Atoi(color)

			// Сохраняем цвет в слайс
			container[j] = colorInt
		}

		// Добавляем конейнер
		containers[i] = container
	}

	// Возвращаем конейнеры
	return containers
}

// Sum считает сумму всех шариков в контейнере
func Sum(container []int) int64 {
	totalSum := int64(0) // Общая шариков в контейнере
	for i := 0; i < len(container); i++ {
		totalSum += int64(container[i])
	}

	return totalSum
}
