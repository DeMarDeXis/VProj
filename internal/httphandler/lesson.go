package httphandler

import (
	"courses/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type GetAllDoneLesson struct {
	Data []domain.Lesson `json:"data"`
}

func (h *Handler) createLesson(w http.ResponseWriter, r *http.Request) {
	var input domain.Lesson

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.Lesson.CreateLesson(&input)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{"id": id}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) getLessonByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid name")
		return
	}

	lesson, err := h.service.Lesson.GetLessonByName(name)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get lesson by name: %s, %e", name, err))
		return
	}
	if lesson == nil {
		newErrorResponse(w, h.logg, http.StatusNotFound, "lesson not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//response := map[string]interface{}{"lesson": lesson}
	json.NewEncoder(w).Encode(lesson)
}

func (h *Handler) getLessonByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}
	lesson, err := h.service.Lesson.GetLessonByID(idInt)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get lesson by id: %s, %e", id, err))
		return
	}
	if lesson == nil {
		newErrorResponse(w, h.logg, http.StatusNotFound, "lesson not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//response := map[string]interface{}{"lesson": lesson}
	json.NewEncoder(w).Encode(lesson)
}

func (h *Handler) getAllDoneLessons(w http.ResponseWriter, r *http.Request) {
	lists, err := h.service.Lesson.GetAllDoneLesson()
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all done lessons: %e", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//response := GetAllDoneLesson{
	//	Data: lists,
	//}
	json.NewEncoder(w).Encode(lists)
}

func (h *Handler) getAllDoneLessonsByCourse(w http.ResponseWriter, r *http.Request) {
	course := chi.URLParam(r, "course")
	if course == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid course")
		return
	}

	courseInt, err := strconv.Atoi(course)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid course")
		return
	}

	lists, err := h.service.Lesson.GetAllDoneLessonByCourse(courseInt)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to get all done lessons: %e", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lists)
}

func (h *Handler) sendLessonForMarking(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.service.Lesson.SendLessonForMarking(idInt); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to send lesson for marking: %s, %e", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateLessonStatus(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, "invalid id")
		return
	}

	var input domain.UpdateLessonStatus
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}
	if err := input.Validate(); err != nil {
		newErrorResponse(w, h.logg, http.StatusBadRequest, err.Error())
		return
	}

	if input.LessonID != nil {
		idInt = *input.LessonID
	}

	if err := h.service.Lesson.UpdateLessonStatus(idInt, *input.Status); err != nil {
		newErrorResponse(w, h.logg, http.StatusInternalServerError, fmt.Sprintf("failed to update lesson status: %s, %e", id, err))
		return
	}

	w.WriteHeader(http.StatusOK)
}
