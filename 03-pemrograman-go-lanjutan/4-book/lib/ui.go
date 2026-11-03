package lib

import (
	"book/services"
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func DisplayMenu() {
	reader := bufio.NewReader(os.Stdin)
	bookService := services.BookService{}

	for {
		fmt.Println("==== Book Menu ====")
		fmt.Println("1. List Book")
		fmt.Println("2. Add Book")
		fmt.Println("3. Update Book")
		fmt.Println("4. Delete Book")
		fmt.Println("5. keluar")
		fmt.Println("==============")

		line, err := readLine(reader, "Select Menu: ")
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println()
				return
			}
			fmt.Printf("An error occurred while reading the menu choice: %v\n", err)
			return
		}

		choice, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil || choice < 1 || choice > 5 {
			fmt.Println("Invalid menu choice; enter a number from 1 to 5.")
			continue
		}

		switch choice {
		case 1:
			books := bookService.GetAll()
			if len(books) == 0 {
				fmt.Println("Book list is empty")
				continue
			}

			for _, book := range books {
				fmt.Println(strings.Repeat("=", 20))
				fmt.Printf("ID: %s\n", book.ID)
				fmt.Printf("Title: %s\n", book.Title)
				fmt.Printf("Author: %s\n", book.Author)
				fmt.Printf("Category: %s\n", book.Category)
				fmt.Printf("Publisher: %s\n", book.Publisher)
				fmt.Printf("Year of Publication: %d\n", book.YearOfPublication)
				fmt.Println(strings.Repeat("=", 20))
			}
		case 2:
			input, err := readBookInput(reader)
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\nInput ended; add operation cancelled.")
					return
				}
				fmt.Printf("Unable to read book details: %v\n", err)
				continue
			}

			bookService.Add(input)
		case 3:
			id, err := readLine(reader, "Enter ID: ")
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\nInput ended; update operation cancelled.")
					return
				}
				fmt.Printf("Unable to read book ID: %v\n", err)
				continue
			}

			input, err := readBookInput(reader)
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\nInput ended; update operation cancelled.")
					return
				}
				fmt.Printf("Unable to read book details: %v\n", err)
				continue
			}

			if err := bookService.Update(input, id); err != nil {
				fmt.Printf("Update failed: %v\n", err)
			}
		case 4:
			id, err := readLine(reader, "Enter ID: ")
			if err != nil {
				if errors.Is(err, io.EOF) {
					fmt.Println("\nInput ended; delete operation cancelled.")
					return
				}
				fmt.Printf("Unable to read book ID: %v\n", err)
				continue
			}

			if err := bookService.Delete(id); err != nil {
				fmt.Printf("Delete failed: %v\n", err)
			}
		case 5:
			fmt.Println("Goodbye...")
			return
		}
	}
}
