package main

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Job struct {
	Progress int16
	ID       uuid.UUID
}

type SuccessResposnse struct {
	JobID uuid.UUID
}

type JobRequest struct {
	ID string `param:"jobId"`
}

var jobs []*Job

func NewJob() *Job {
	return &Job{
		Progress: 0,
		ID:       uuid.New(),
	}
}

func main() {

	e := echo.New()

	e.POST("/api/setJob", func(c echo.Context) error {

		newJob := NewJob()
		jobs = append(jobs, newJob)
		UpdateJobProgress(newJob)

		return c.JSON(http.StatusOK, SuccessResposnse{JobID: newJob.ID})
	})

	e.GET("/api/getJob/:jobId", func(c echo.Context) error {
		jobReq := JobRequest{
			ID: c.Param("jobId"),
		}
		var job *Job
		for _, element := range jobs {
			if element.ID.String() == jobReq.ID {
				job = element
			}
		}

		return c.JSON(http.StatusOK, job.Progress)
	})

	e.Logger.Fatal(e.Start(":8000"))
}

func UpdateJobProgress(job *Job) {
	if job.Progress == 100 {
		return
	}
	time.AfterFunc(time.Duration(int64(time.Second)), func() {
		job.Progress += 1
		UpdateJobProgress(job)
	})
}
