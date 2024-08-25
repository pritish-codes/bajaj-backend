package routes

import (
	// Replace 'myapp' with your module name

	"sort"
	"strconv"
	"unicode"

	"github.com/gofiber/fiber/v2"
)

type req struct {
	Data []string `json:"data"`
}

type Response struct {
	IsSuccess                bool     `json:"is_success"`
	UserID                   string   `json:"user_id"`
	Email                    string   `json:"email"`
	RollNumber               string   `json:"roll_number"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet string   `json:"highest_lowercase_alphabet"`
}

func GetUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"operation_code": 1,
	})
}

func PostUser(c *fiber.Ctx) error {
r := new(req)
	if err := c.BodyParser(r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"is_success": false,
			"message":    "Invalid request",
		})
	}

	var numbers, alphabets []string
	var lowercaseAlphabets []string

	for _, val := range r.Data {
		if _, err := strconv.Atoi(val); err == nil {
			numbers = append(numbers, val)
		} else if unicode.IsLetter([]rune(val)[0]) {
			alphabets = append(alphabets, val)
			if unicode.IsLower([]rune(val)[0]) {
				lowercaseAlphabets = append(lowercaseAlphabets, val)
			}
		}
	}

	// Find the highest lowercase alphabet
	highestLowercaseAlphabet := ""
	if len(lowercaseAlphabets) > 0 {
		sort.Strings(lowercaseAlphabets)
		highestLowercaseAlphabet = lowercaseAlphabets[len(lowercaseAlphabets)-1]
	}

	// Respond with the processed data
	response := Response{
		IsSuccess:                true,
		UserID:                   "john_doe_17091999",
		Email:                    "john@xyz.com",
		RollNumber:               "ABCD123",
		Numbers:                  numbers,
		Alphabets:                alphabets,
		HighestLowercaseAlphabet: highestLowercaseAlphabet,
	}

	return c.JSON(response)
}
