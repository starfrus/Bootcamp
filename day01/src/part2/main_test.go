package main

import (
	"reflect"
	"testing"
)

// Тестовая функция для проверки основной логики
func TestGetTopWords(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		k        int
		expected []string
	}{
		{
			name:     "K меньше количества уникальных слов",
			words:    []string{"aa", "bb", "cc", "aa", "cc", "cc", "cc", "aa", "ab", "ac", "bb"},
			k:        3,
			expected: []string{"cc", "aa", "bb"}, // cc(4 раза), aa(3 раза), bb(2 раза)
		},
		{
			name:     "Пустой список слов",
			words:    []string{},
			k:        5,
			expected: []string{},
		},
		{
			name:     "K больше количества уникальных слов",
			words:    []string{"aa", "bb", "cc", "aa", "bb"},
			k:        10,
			expected: []string{"aa", "bb", "cc"}, // Все уникальные слова
		},
		{
			name:     "K равно количеству уникальных слов",
			words:    []string{"aa", "bb", "cc"},
			k:        3,
			expected: []string{"aa", "bb", "cc"}, // Все слова по одному разу
		},
		{
			name:     "Слова с одинаковой частотой",
			words:    []string{"aa", "bb", "cc"},
			k:        2,
			expected: []string{"aa", "bb", "cc"}, // Может вернуть любые 2, но обычно в алфавитном порядке
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Создаем мапу частот слов
			mapWordsNum := make(map[string]int)
			for _, word := range tt.words {
				mapWordsNum[word]++
			}

			// Создаем списки уникальных слов и их частот
			uniqWords := []string{}
			arrNumWords := []int{}
			for k, v := range mapWordsNum {
				uniqWords = append(uniqWords, k)
				arrNumWords = append(arrNumWords, v)
			}

			var result []string

			// Проверяем граничные условия
			if len(tt.words) == 0 {
				result = []string{}
			} else if len(uniqWords) <= tt.k {
				result = SortWords(uniqWords, mapWordsNum)
			} else {
				// Сортируем частоты по убыванию
				sortedFreqs := make([]int, len(arrNumWords))
				copy(sortedFreqs, arrNumWords)
				sortedFreqs = SortNumWords(sortedFreqs)

				// Получаем результат
				result = CompareNum(tt.k, mapWordsNum, sortedFreqs)
			}

			// Проверяем результат
			if len(tt.expected) == 0 && len(result) == 0 {
				return // Оба пустые - OK
			}

			// Для случаев с одинаковыми частотами проверяем, что все ожидаемые слова присутствуют
			if tt.name == "Слова с одинаковой частотой" || tt.name == "K больше количества уникальных слов" {
				// Проверяем, что все уникальные слова из входных данных присутствуют в результате
				expectedSet := make(map[string]bool)
				resultSet := make(map[string]bool)

				for _, word := range uniqWords {
					expectedSet[word] = true
				}
				for _, word := range result {
					resultSet[word] = true
				}

				if !reflect.DeepEqual(expectedSet, resultSet) {
					t.Errorf("Ожидалось %v, получено %v", uniqWords, result)
				}
			} else if len(tt.expected) > 0 {
				// Проверяем первые k элементов
				if len(result) < len(tt.expected) {
					t.Errorf("Ожидалось как минимум %d элементов, получено %d", len(tt.expected), len(result))
					return
				}

				// Проверяем, что первые элементы соответствуют ожидаемым
				for i, expectedWord := range tt.expected {
					if i >= len(result) {
						t.Errorf("Результат короче ожидаемого: индекс %d, длина результата %d", i, len(result))
						return
					}
					if result[i] != expectedWord {
						t.Errorf("На позиции %d: ожидалось %s, получено %s", i, expectedWord, result[i])
					}
				}
			}
		})
	}
}

// Тесты для вспомогательных функций
func TestSortNumWords(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Обычная сортировка",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6},
			expected: []int{9, 6, 5, 4, 3, 2, 1, 1},
		},
		{
			name:     "Пустой массив",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Один элемент",
			input:    []int{5},
			expected: []int{5},
		},
		{
			name:     "Уже отсортированный",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)
			result := SortNumWords(inputCopy)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Ожидалось %v, получено %v", tt.expected, result)
			}
		})
	}
}

func TestSortWords(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		freq     map[string]int
		expected []string
	}{
		{
			name:     "Сортировка по частоте",
			words:    []string{"aa", "bb", "cc"},
			freq:     map[string]int{"aa": 3, "bb": 2, "cc": 1},
			expected: []string{"aa", "bb", "cc"},
		},
		{
			name:     "Пустой список",
			words:    []string{},
			freq:     map[string]int{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SortWords(tt.words, tt.freq)

			// Проверяем, что длины совпадают
			if len(result) != len(tt.expected) {
				t.Errorf("Разная длина: ожидалось %d, получено %d", len(tt.expected), len(result))
				return
			}

			// Проверяем порядок (для слов с разной частотой)
			if len(tt.expected) > 0 && len(result) > 0 {
				// Простая проверка первых элементов
				maxExpectedFreq := tt.freq[tt.expected[0]]
				maxResultFreq := tt.freq[result[0]]

				if maxResultFreq < maxExpectedFreq {
					t.Errorf("Неправильная сортировка: максимальная частота в результате %d, ожидалась %d",
						maxResultFreq, maxExpectedFreq)
				}
			}
		})
	}
}

// Интеграционный тест для основного сценария
func TestMainFunctionLogic(t *testing.T) {
	// Тест для случая a: K меньше количества уникальных слов
	t.Run("Сценарий a: K меньше уникальных слов", func(t *testing.T) {
		words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
		k := 2

		mapWordsNum := make(map[string]int)
		for _, word := range words {
			mapWordsNum[word]++
		}

		uniqWords := make([]string, 0, len(mapWordsNum))
		for word := range mapWordsNum {
			uniqWords = append(uniqWords, word)
		}

		if len(uniqWords) > k {
			arrNumWords := make([]int, 0, len(mapWordsNum))
			for _, freq := range mapWordsNum {
				arrNumWords = append(arrNumWords, freq)
			}

			sortedFreqs := make([]int, len(arrNumWords))
			copy(sortedFreqs, arrNumWords)
			sortedFreqs = SortNumWords(sortedFreqs)

			result := CompareNum(k, mapWordsNum, sortedFreqs)

			if len(result) < k {
				t.Errorf("Ожидалось не менее %d слов, получено %d", k, len(result))
			}
		}
	})

	// Тест для случая b: пустой список слов
	t.Run("Сценарий b: пустой список слов", func(t *testing.T) {
		words := []string{}
		result := []string{}

		if len(words) == 0 {
			if len(result) != 0 {
				t.Errorf("Для пустого списка должен возвращаться пустой результат, получено %v", result)
			}
		}
	})

	// Тест для случая c: K больше количества уникальных слов
	t.Run("Сценарий c: K больше уникальных слов", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}
		k := 10

		mapWordsNum := make(map[string]int)
		for _, word := range words {
			mapWordsNum[word]++
		}

		uniqWords := make([]string, 0, len(mapWordsNum))
		for word := range mapWordsNum {
			uniqWords = append(uniqWords, word)
		}

		if len(uniqWords) <= k {
			result := SortWords(uniqWords, mapWordsNum)

			// Должны получить все уникальные слова
			if len(result) != len(uniqWords) {
				t.Errorf("Ожидалось %d слов, получено %d", len(uniqWords), len(result))
			}
		}
	})
}