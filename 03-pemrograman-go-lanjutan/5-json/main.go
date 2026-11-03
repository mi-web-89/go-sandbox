package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"review-json/models"
)

func main() {
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("cannot get current file info")
	}

	projectDir, err := filepath.Abs(filepath.Dir(sourceFile))
	if err != nil {
		log.Fatal(err)
	}

	fileName := filepath.Join(projectDir, "reviews.json")

	byteValue, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var reviews []models.Review
	// decode JSON data into the reviews slice
	if err := json.Unmarshal(byteValue, &reviews); err != nil {
		log.Fatal(err)
	}

	for i, review := range reviews {
		reviews[i].RestaurantName = strings.TrimSpace(review.RestaurantName)
		if reviews[i].RestaurantName == "" {
			reviews[i].RestaurantName = "N/A"
		}

		reviews[i].Username = strings.TrimSpace(review.Username)
		if reviews[i].Username == "" {
			reviews[i].Username = "N/A"
		}

		if review.Score < 0 {
			reviews[i].Score = 0
		}

		reviews[i].Comment = strings.TrimSpace(review.Comment)
		if reviews[i].Comment == "" {
			reviews[i].Comment = "N/A"
		}

		fmt.Printf("Restaurant: %s\n", reviews[i].RestaurantName)
		fmt.Printf("User: %s\n", reviews[i].Username)
		fmt.Printf("Score: %d\n", review.Score)
		fmt.Printf("Comment: %s\n\n", reviews[i].Comment)
	}

	updatedReviews, err := json.MarshalIndent(reviews, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	newFIleName := filepath.Join(projectDir, "updated_reviews.json")

	if err := os.WriteFile(newFIleName, updatedReviews, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Println("reviews updated successfully")
}
