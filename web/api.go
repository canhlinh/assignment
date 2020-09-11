package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Api struct {
	Store *Store
}

// CreateStudent creates a student
// @Summary Create a student
// @Description Create a student
// @ID create-student
// @Accept  application/json
// @Produce  json
// @Param Student body Student true "Create Student"
// @Success 200 {object} Student
// @Router /students [post]
func (a *Api) CreateStudent(ctx *gin.Context) {
	var student Student
	if err := json.NewDecoder(ctx.Request.Body).Decode(&student); err != nil {
		ctx.JSON(400, NewError(err))
		return
	}

	if err := a.Store.SaveStudent(&student); err != nil {
		ctx.JSON(400, NewError(err))
		return
	}

	ctx.JSON(201, student)
}

// GetStudent creates a student
// @Summary Get a student
// @Description Get a student
// @ID get-student
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} Student
// @Router /students/{student_id} [get]
func (a *Api) GetStudent(ctx *gin.Context) {
	// write your code
}

// ListStudent list students
// @Summary List students
// @Description List students
// @ID list-students
// @Accept  json
// @Produce  json
// @Success 200 {array} Student
// @Router /students [get]
func (a *Api) ListStudent(ctx *gin.Context) {

	students, err := a.Store.GetStudents()
	if err != nil {
		ctx.JSON(500, NewError(err))
		return
	}

	ctx.JSON(200, students)
}

// EditStudent updates a student
// @Summary Edit a student
// @Description Edit a student
// @ID edit-student
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} Student
// @Router /students/{student_id} [put]
func (a *Api) EditStudent(ctx *gin.Context) {
	// write your code
}
