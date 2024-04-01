package server

import (
	"fmt"
	"gocdn/config"
	"gocdn/lib"
	"gocdn/types"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	conf := config.GetConfig()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Get("/:fileId", func(c *fiber.Ctx) error {
		fileId := c.Params("fileId")
		filePath, err := lib.GetFilePath(fileId)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "File not found",
			})
		}

		if lib.FileExists(filePath) {
			return c.SendFile(filePath)
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "File not found",
		})
	})

	app.Post("/upload/:bigType", func(c *fiber.Ctx) error {
		fmt.Println("Headers:", c.Request().Header.String())
		fmt.Println("Body:", string(c.Body()))
		bigTypes := conf.BigTypes
		bigType := c.Params("bigType")

		if !lib.StringInSlice(bigType, bigTypes) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid bigType. Try one of these: " + lib.JoinStrings(bigTypes, ", "),
			})
		}

		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "No file uploaded",
			})
		}

		fileType := file.Header.Get("Content-Type")

		smallTypes := conf.UploadTypes
		if !lib.StringInSlice(fileType, smallTypes) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid file type. Try uploading a file with one of these types: " + lib.JoinStrings(smallTypes, ", "),
			})
		}

		key := lib.GenerateKey(16)
		lib.CreateFolder("uploads")
		lib.CreateFolder("uploads/" + bigType)
		lib.CreateFolder("uploads/" + bigType + "/" + fileType)
		lib.CreateFolder("uploads/" + bigType + "/" + fileType + "/" + key)
		lib.CreateFile("uploads/" + bigType + "/" + fileType + "/" + key + "/" + "data.yaml")

		err = c.SaveFile(file, "uploads/"+bigType+"/"+fileType+"/"+key+"/"+file.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		data := types.Upload{
			ID:   key,
			Name: file.Filename,
			Size: file.Size,
		}
		lib.WriteYAML("uploads/"+bigType+"/"+fileType+"/"+key+"/data.yaml", data)

		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
			"key":     key,
		})
	})
}
