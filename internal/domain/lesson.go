package domain

import (
	"errors"
	"fmt"
)

type Lesson struct {
	LessonID    int     `json:"lesson_id" db:"lesson_id"`
	CourseID    int     `json:"course_id" db:"course_id"`
	Name        string  `json:"name" db:"lesson_name"`
	Description string  `json:"description" db:"lesson_description"`
	Status      *string `json:"status" db:"status"`
}

type UpdateLessonStatus struct {
	LessonID *int    `json:"lesson_id" db:"lesson_id"`
	Status   *string `json:"status" db:"lesson_status"`
}

func (l Lesson) Validate() error {
	if l.CourseID == 0 || l.Name == "" || l.Description == "" {
		return errors.New(fmt.Sprintf("invalid lesson id: %d, name: %s, description: %s", l.CourseID, l.Name, l.Description))
	}
	return nil
}

func (l UpdateLessonStatus) Validate() error {
	if l.LessonID == nil && l.Status == nil {
		return errors.New(fmt.Sprintf("invalid lesson id: %d or status: %d", l.LessonID, l.Status))
	}
	return nil
}

//{
//	"course_id": 1
//	"name": "check",
//	"description": "check description"
//	"status": "done"
//}
