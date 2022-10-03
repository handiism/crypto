package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/handiism/crypto/algorithm"
	"github.com/pkg/browser"
)

//go:embed view/*
var view embed.FS

func main() {
	app := fiber.New()
	app.Post("/encipher", Encipher())
	app.Post("/decipher", Decipher())
	app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(view),
		PathPrefix: "view",
		Browse:     true,
	}))
	go browser.OpenURL("http://127.0.0.1:8000/")
	go log.Fatal(app.Listen(":8000"))
}

type Request struct {
	Algorithm  string `json:"algorithm"`
	Plaintext  string `json:"plaintext,omitempty"`
	Ciphertext string `json:"ciphertext,omitempty"`
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

func Encipher() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var request Request
		if err := json.Unmarshal(body, &request); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
		}

		if request.Algorithm == "affine" {
			affine, err := algorithm.NewAffine(request.Key.Numbers[0], request.Key.Numbers[1])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}
			cipher := affine.Encipher(request.Plaintext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": cipher}})
		} else if request.Algorithm == "caesar" {
			caesar := algorithm.NewCaesar(request.Key.Numbers[0])
			cipher := caesar.Encipher(request.Plaintext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": cipher}})
		} else if request.Algorithm == "otp" {
			otp := algorithm.NewOTP(request.Key.Text)
			cipher, err := otp.Encrypt(request.Plaintext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}

			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": cipher}})
		} else if request.Algorithm == "super" {
			super, err := algorithm.NewSuper(request.Key.Numbers[0], request.Key.Numbers[1], request.Key.Text)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}
			cipher, err := super.Encipher(request.Plaintext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}
			return c.JSON(Response{Status: "success", Data: fiber.Map{"ciphertext": cipher}})
		}

		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": "unknown algorithm"}})
	}
}

func Decipher() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var request Request
		if err := json.Unmarshal(body, &request); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
		}

		if request.Algorithm == "affine" {
			affine, err := algorithm.NewAffine(request.Key.Numbers[0], request.Key.Numbers[1])
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}
			plain := affine.Decipher(request.Ciphertext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		} else if request.Algorithm == "caesar" {
			caesar := algorithm.NewCaesar(request.Key.Numbers[0])
			plain := caesar.Decipher(request.Ciphertext)
			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		} else if request.Algorithm == "otp" {
			otp := algorithm.NewOTP(request.Key.Text)
			plain, err := otp.Decrypt(request.Ciphertext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}

			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		} else if request.Algorithm == "super" {
			super, err := algorithm.NewSuper(request.Key.Numbers[0], request.Key.Numbers[1], request.Key.Text)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}
			plain, err := super.Decipher(request.Ciphertext)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": err.Error()}})
			}
			return c.JSON(Response{Status: "success", Data: fiber.Map{"plaintext": plain}})
		}

		c.Status(http.StatusUnprocessableEntity)
		return c.JSON(Response{Status: "fail", Data: fiber.Map{"message": "unknown algorithm"}})
	}
}
