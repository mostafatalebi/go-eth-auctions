package shared

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Handler = func(workerIndex int, job interface{})

// Worker is a job manager based on go routines. Define the number of internal
// workers, and start publishing jobs using Worker Publish() API. It will distribute the job
// among its internal pool. The worker never exit and are always listening.
// To exit the worker, simply call Exit(), it will register a sigterm BUT the associated
// job channel will NOT be closed, because channel is externally passed to this instance,
// and other processes might be using the channel.
type Worker struct {
	bufferSize     int
	jobChannel     chan interface{}
	numberOfWorker int
	sigTerm        chan os.Signal
	do             Handler

	waiter *sync.WaitGroup
}

func NewWorkerManager(bufferSize, numberOfWorkers int, jobChannel chan any) *Worker {
	if jobChannel == nil {
		jobChannel = make(chan any, bufferSize)
	}
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	return &Worker{
		bufferSize:     bufferSize,
		numberOfWorker: numberOfWorkers,
		jobChannel:     jobChannel,
		sigTerm:        sigChan,
		waiter:         &sync.WaitGroup{},
	}
}

func (w *Worker) GetUnreadCount() int64 {
	if w.jobChannel == nil {
		return 0
	}
	return int64(len(w.jobChannel))
}

func (w *Worker) JobEvents() chan interface{} {
	return w.jobChannel
}

func (w *Worker) SetWorker(worker Handler) {
	w.do = worker
}

// Enqueue if no external channel is passed, an internal one is created then you can
// use this function to publish to it
func (w *Worker) Enqueue(val interface{}) {
	w.jobChannel <- val
}

// doSafe is a wrapper around the worker function, it will catch any error and
// call the error handler if defined
func (w *Worker) doSafe(workerIndex int, job interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("worker recovered, error: " + r.(error).Error())
		}
	}()
	if w.do == nil {
		return
	}
	w.do(workerIndex, job)
}

// Start starts the process of workers
func (w *Worker) Start() error {
	w.waiter.Add(w.numberOfWorker)
	for i := 0; i < w.numberOfWorker; i++ {
		go func(index int) {
			for {
				select {
				case job := <-w.jobChannel:
					w.doSafe(index, job)
				case <-w.sigTerm:
					w.waiter.Done()
					return
				}
			}
		}(i)
	}
	w.waiter.Wait()

	return errors.New("workers terminated")
}

// Exit calling this function would immediately exit from all sub-processes
func (w *Worker) Exit() {
	fmt.Println("exiting workers...")
	for i := 0; i < w.numberOfWorker; i++ {
		w.sigTerm <- syscall.SIGSTOP
	}
}
