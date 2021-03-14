package handlers

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sinanganizz/password-generator/database"
)

func Homepage(c *fiber.Ctx) error {

	passTitle := c.Query("passTitle")
	rPassword := c.Query("rpassword")
	var status bool = false
	fmt.Println("PassTitle: " + passTitle)
	fmt.Println("rPassword: " + rPassword)
	fmt.Println("")

	if passTitle != "" && rPassword != "" {
		rslt := database.SavePassword(rPassword, passTitle)
		status = rslt
	}
	charList := [73]int{33, 38, 40, 41, 43, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 61, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122}

	source := rand.NewSource(time.Now().UnixNano())
	randss := rand.New(source)
	pass_length := randss.Intn(10) + 9
	var password string
	password = ""

	for i := 0; i < pass_length; i++ {
		char := string(charList[randss.Intn(73)])
		password += char
	}
	fmt.Println(password)
	return c.Render("index", fiber.Map{
		"Password": password,
		"Status":   status,
	})
}

func MyPasswords(c *fiber.Ctx) error {

	passwords := database.GetPasswords()
	fmt.Println(passwords)
	fmt.Println()
	status := true
	return c.Render("mypasswords", fiber.Map{
		"Passwords": passwords,
		"Status":    status,
	})
}

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "404 - Page Not Found!",
	})
}
