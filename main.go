package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/handiism/crypto/algorithm"
)

func main() {
	app := fiber.New()
	app.Post("/encipher", Enchiper())
	app.Post("/decipher", Dechiper())
	app.Static("/", "./view")
	log.Fatal(app.Listen(":8000"))
}

type Request struct {
	Algorithm  string `json:"algorithm"`
	Plaintext  string `json:"plaintext,omitempty"`
	Chipertext string `json:"cipHertext,omitempty"`
	Key        Key    `json:"key"`
}

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}
type Key struct {
	Numbers []int  `json:"numbers,omitempty"`
	Text    string `json:"text,omitempty"`
}

func Enchiper() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var request Request
		if err := json.Unmarshal(body, &request); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
		}

		if request.Algorithm == "affine" {
			affine, err := algorithm.NewAffine(request.Key.Numbers[0], request.Key.Numbers[1])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}
			chiper := affine.Encipher(request.Plaintext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": chiper}})
		} else if request.Algorithm == "caesar" {
			caesar := algorithm.NewCaesar(request.Key.Numbers[0])
			chiper := caesar.Encipher(request.Plaintext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": chiper}})
		} else if request.Algorithm == "railfence" {
			railfence, err := algorithm.NewRailfence(request.Key.Numbers[0])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}

			chiper, err := railfence.Encipher(request.Plaintext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}

			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": chiper}})
		} else if request.Algorithm == "super" {
			super, err := algorithm.NewSuper(request.Key.Numbers[0], request.Key.Numbers[1])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}
			chiper, err := super.Encipher(request.Plaintext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}
			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": chiper}})
		}

		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": "unknown algorithm"}})
	}
}

func Dechiper() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var request Request
		if err := json.Unmarshal(body, &request); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
		}

		if request.Algorithm == "affine" {
			affine, err := algorithm.NewAffine(request.Key.Numbers[0], request.Key.Numbers[1])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}
			plain := affine.Decipher(request.Chipertext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		} else if request.Algorithm == "caesar" {
			caesar := algorithm.NewCaesar(request.Key.Numbers[0])
			plain := caesar.Decipher(request.Chipertext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		} else if request.Algorithm == "railfence" {
			railfence, err := algorithm.NewRailfence(request.Key.Numbers[0])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}

			plain, err := railfence.Decipher(request.Chipertext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}

			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		} else if request.Algorithm == "super" {
			super, err := algorithm.NewSuper(request.Key.Numbers[0], request.Key.Numbers[1])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}
			plain, err := super.Decipher(request.Chipertext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": err.Error()}})
			}
			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		}

		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(Response{Status: "failed", Data: fiber.Map{"message": "unknown algorithm"}})
	}
}
