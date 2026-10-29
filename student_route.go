package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

//ทำAPI

func studentRoute(app *fiber.App, db *gorm.DB) {
	//api post for create (Endpoint /api/students)

	app.Post("/api/students", func(c fiber.Ctx) error {
		student := new(Student)

		// นำข้อมูลที่ส่งมาเก็บในตัวแปร พร้อมทั้งตรวจสอบ
		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		if error := createStudent(db, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		return c.Status(201).JSON(fiber.Map{
			"message": "Student created successfully",
			"data":    student,
		})
	})
	//API Get for Read (Endpoint:/api/students)
	app.Get("/api/students", func(c fiber.Ctx) error {

		var students []Student

		//------------------------------------------
		//สร้างตัวแปรเก็บข้อมูลที่เป็นเงื่อนไขที่ส่งมาแบบquery string
		StuName := c.Query("StuName")
		if StuName == "" {
			StuName = c.Query("name")
		}
		if StuName == "" {
			StuName = c.Query("stu_name")
		}

		StuLevel := c.Query("StuLevel")
		if StuLevel == "" {
			StuLevel = c.Query("level")
		}
		if StuLevel == "" {
			StuLevel = c.Query("stu_level")
		}

		var error error

		if StuName != "" {
			//ดึงข้อมูลเงื่อนไข StuName
			students, error = getStudentByName(db, StuName)
		} else if StuLevel != "" {
			//ดึงข้อมูลเงื่อนไข StuLevel
			StuLevelInt, err := strconv.Atoi(StuLevel)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "Invalid request parameter",
				})
			}
			students, error = getStudentByLevel(db, StuLevelInt)
		} else {
			//ดึงข้อมูลปกติ
			students, error = getAllStudent(db)
		}

		//------------------------------------------

		if errors.Is(error, gorm.ErrRecordNotFound) || ((StuName != "" || StuLevel != "") && len(students) == 0 && error == nil) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		//ตรวจสอบว่ามีการทำงานผิดพลาดหรือไม่
		if error != nil {
			log.Println(error)
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Sucess",
			"data":    students,
		})
	})
	//API Get for Read (Endpoint:/api/students/:id)
	app.Get("api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		student, error := getStudentByID(db, idInt)

		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		if error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Sucess",
			"data":    student,
		})

	})
	//API Put for Update (Endpoint:/api/students)
	app.Put("/api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		student, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		students := new(Student)

		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}
		if error := updateStudent(db, idInt, students); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Student updated successfully",
			"data":    students,
		})
	})
	//API Delete for Delete (Endpoint:/api/students/:id)
	app.Delete("api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}
		if error := deleteStudent(db, idInt); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Student deleted successfully",
		})
	})
}
