package beanstalk

import (
	"testing"
	"time"
)

type TestConsumer struct{}

func (tc *TestConsumer) FinishJob(job *Job, method JobMethod, priority uint32, delay time.Duration) error {
	return nil
}

func NewTestJob() *Job {
	return &Job{
		ID:     12345,
		Body:   []byte("Hello World"),
		TTR:    time.Duration(1),
		Finish: &TestConsumer{}}
}

func TestBuryJob(t *testing.T) {
	job := NewTestJob()
	if err := job.Bury(1024); err != nil {
		t.Fatalf("Unexpected error from Bury: %s", err)
	}
}

func TestDeleteJob(t *testing.T) {
	job := NewTestJob()
	if err := job.Delete(); err != nil {
		t.Fatalf("Unexpected error from Delete: %s", err)
	}
}

func TestReleaseJob(t *testing.T) {
	job := NewTestJob()
	if err := job.Release(1024, time.Duration(time.Second)); err != nil {
		t.Fatalf("Unexpected error from Release: %s", err)
	}
}

func TestDoubleFinalizeJob(t *testing.T) {
	job := NewTestJob()
	if err := job.Delete(); err != nil {
		t.Fatalf("Unexpected error from Delete: %s", err)
	}
	if err := job.Delete(); err != ErrJobFinished {
		t.Fatalf("Expected ErrJobFinished, but got: %s", err)
	}

}
