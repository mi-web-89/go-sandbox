package lib

import (
	"book/models"
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func readBookInput(reader *bufio.Reader) (models.BookInput, error) {
	title, err := readLine(reader, "Enter Title: ")
	if err != nil {
		return models.BookInput{}, fmt.Errorf("title: %w", err)
	}

	author, err := readLine(reader, "Enter Author: ")
	if err != nil {
		return models.BookInput{}, fmt.Errorf("author: %w", err)
	}

	category, err := readLine(reader, "Enter Category: ")
	if err != nil {
		return models.BookInput{}, fmt.Errorf("category: %w", err)
	}

	publisher, err := readLine(reader, "Enter Publisher: ")
	if err != nil {
		return models.BookInput{}, fmt.Errorf("publisher: %w", err)
	}

	yearText, err := readLine(reader, "Enter Year of Publication: ")
	if err != nil {
		return models.BookInput{}, fmt.Errorf("year of publication: %w", err)
	}

	year, err := strconv.Atoi(strings.TrimSpace(yearText))
	if err != nil {
		return models.BookInput{}, fmt.Errorf("year of publication must be a whole number: %w", err)
	}

	return models.BookInput{
		Title:             title,
		Author:            author,
		Category:          category,
		Publisher:         publisher,
		YearOfPublication: year,
	}, nil
}

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)
	line, err := reader.ReadString('\n')
	if err != nil && !(errors.Is(err, io.EOF) && line != "") {
		return "", err
	}

	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	return line, nil
}
