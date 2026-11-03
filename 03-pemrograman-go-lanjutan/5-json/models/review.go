package models

type Review struct {
	Username       string `json:"username"`
	Score          int    `json:"score"`
	RestaurantName string `json:"restaurant_name"`
	Comment        string `json:"comment"`
}
