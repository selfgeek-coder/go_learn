package main

import "fmt"

type Book struct {
	Author string
	Title  string
}

func main() {
	// инициализация слайса (массива)
	
	/* ints := []int{
		1, 2, 3,
	}

	fmt.Println(ints) */

	books := []Book{
		{Author: "Pushkin", Title: "Roman N"},
		{Author: "Andrzej Sapkowski", Title: "The Witcher"},
		{Author: "J.K. Rowling", Title: "Harry Potter"},
		{Author: "Robert Kiyosaki", Title: "Rich Dad"},
	}

	// можно пройтись по каждой книге
	// for _, value := range books {
	// 	fmt.Println(value)
	// }

	// можно обратится через индекс
	// fmt.Println(books[0])

	// добавление нового элемента в слайс
	new_book := Book{
		Author: "Sun Tzu", Title: "The Art of War",
	}

	books = append(books, new_book)

	for index, book := range books {
		fmt.Printf("%d. %s, %s\n\n", index, book.Title, book.Author)
	}
}