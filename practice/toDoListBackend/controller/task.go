package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/easy-to-study/go_study/practice/toDoListBackend/model"
	"github.com/easy-to-study/go_study/practice/toDoListBackend/view"
)

// TaskController require *sql.Db to initialize
// This controller hove CRUD methods
type TaskController struct {
	Db *sql.DB
}

// GetTasks return All Tasks
func (tc *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tasks, err := model.GetTasks(ctx, tc.Db)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get tasks error: %v", err))
		return
	}
	view.RenderTasks(w, tasks)
}

// GetTask は path に含まれる uuid に一致する tasks テーブルの レコードを返す
func (tc *TaskController) GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskUUID := params["uuid"]
	ctx := context.Background()
	exist, err := model.CheckTaskExist(ctx, tc.Db, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("check task exist error: %v", err))
		return
	}
	if !exist {
		view.RenderNotFound(w, "tasks", taskUUID)
		return
	}

	task, err := model.GetTask(ctx, tc.Db, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get tasks error: %v", err))
		return
	}
	view.RenderTask(w, task, http.StatusOK)
}

// CreateTask create new Task, and return that Task
func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		return
	}
	var task model.Task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("json parse error: %v", err))
		return
	}

	ctx := context.Background()
	insertID, err := model.CreateTask(ctx, tc.Db, &task)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create task error: %v", err))
		return
	}
	createdTask, err := model.GetTaskByID(ctx, tc.Db, insertID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get task error: %v", err))
		return
	}
	view.RenderTask(w, createdTask, http.StatusCreated)
}

// PutTask replace specified Task, and return that Task
func (tc *TaskController) PutTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskUUID := params["uuid"]
	ctx := context.Background()
	exist, err := model.CheckTaskExist(ctx, tc.Db, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("check task exist error: %v", err))
		return
	}
	if !exist {
		view.RenderNotFound(w, "tasks", taskUUID)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		view.RenderBadRequest(w, []string{fmt.Sprintf("read post body error: %v", err)})
		return
	}

	var task model.Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("json parse error: %v", err))
		return
	}

	err = model.UpdateTask(ctx, tc.Db, &task, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create task error: %v", err))
		return
	}
	updatedTask, err := model.GetTask(ctx, tc.Db, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("get task error: %v", err))
		return
	}
	view.RenderTask(w, updatedTask, http.StatusOK)
}

// DeleteTask delete specified Task, and return only status code
func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	params := mux.Vars(r)
	taskUUID := params["uuid"]
	ctx := context.Background()
	exist, err := model.CheckTaskExist(ctx, tc.Db, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("check task exist error: %v", err))
		return
	}
	if !exist {
		view.RenderNotFound(w, "tasks", taskUUID)
		return
	}

	err = model.DeleteTask(ctx, tc.Db, taskUUID)
	if err != nil {
		view.RenderInternalServerError(w, fmt.Sprintf("create task error: %v", err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
