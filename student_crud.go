package main

import "gorm.io/gorm"

// ทำงานกับ DB-CRUD
func createStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

//ฟังชันท์ดึงข้อมูลทั้งหมดจากตาราง
func getAllStudent(db *gorm.DB) ([]Student, error) {
	var students []Student
	result := db.Find(&students)
	return students, result.Error
}

//ฟังชันท์ดึงข้อมูลจากตาราง ที่เป็นPK คือ ID
func getStudentByID(db *gorm.DB, id int) (*Student, error) {
	var student Student

	result := db.First(&student, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return &student, nil
}

//ฟังชันท์ดึงข้อมูลจากตาราง เงือนไขไม่ใช้PK เช่น ชื่อตรงกับที่ส่งมาบางส่วน
func getStudentByName(db *gorm.DB, name string) ([]Student, error) {
	var students []Student

	result := db.Where("stu_name LIKE ?", "%"+name+"%").Find(&students)

	return students, result.Error
}

//ฟังชันท์ดึงข้อมูลจากตาราง เงื่อนไขไม่ใช่PK เช่น ชั้นปี ตรงกับที่ส่งมา
func getStudentByLevel(db *gorm.DB, StuLevel int) ([]Student, error) {
	var students []Student

	result := db.Where("stu_level = ?", StuLevel).Find(&students)

	return students, result.Error
}

//ฟังชันUpdate
func updateStudent(db *gorm.DB, id int, student *Student) error {
	return db.Model(&Student{}).Where("id = ?", id).Updates(student).Error
}

//ฟังชันลบ
func deleteStudent(db *gorm.DB, id int) error {
	return db.Delete(&Student{}, id).Error
}
