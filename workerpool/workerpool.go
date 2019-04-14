package workerpool

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// This function returns a workerpool object with necessary configuration
func GenerateNewPool(numOfRoutines int, isResultProcessorRequired bool) *Pool {
	poolObj := &Pool{NumOfRoutines: numOfRoutines}
	poolObj.JobChannel = make(chan Job, numOfRoutines)
	poolObj.ResultChannel = make(chan Result, numOfRoutines)
	if isResultProcessorRequired {
		poolObj.IsResultProcessorRequired = true
	}
	poolObj.Done = make(chan bool)
	return poolObj
}

// allocate functions creates job list from user provided data objects
func (p *Pool) allocate(resources []interface{}) {
	log.Printf("Allocating %d jobs", len(resources))
	defer close(p.JobChannel)
	for i, v := range resources {
		job := Job{JobId: i, Resource: v}
		p.JobChannel <- job
	}
	log.Printf("Job Allocation completed")
}

func (p *Pool) Start(resources []interface{}, jobProcessFunction JobProcessFunction, resultProcessFunction ResultProcessFunction) error {
	log.Printf("Worker pool started!")
	startTime := time.Now()
	go p.allocate(resources)
	// Check if IsResultProcessorRequired then call collect otherise call taskCompletionStatusChecker
	if p.IsResultProcessorRequired {
		go p.collect(resultProcessFunction)
	} else {
		go p.taskCompletionStatusChecker()
	}
	go p.workerpool(jobProcessFunction)
	// blocks the execution until all tasks are done
	<-p.Done
	processTime := time.Now().Sub(startTime)
	fmt.Printf("Process took %d time for completion", processTime)
	return nil
}

func (p *Pool) work(wg *sync.WaitGroup, jobProcessFunction JobProcessFunction) {
	log.Println("Worker started!")
	defer wg.Done()
	for job := range p.JobChannel {
		result, err := jobProcessFunction(job)
		if err != nil {
			log.Printf("Error: Error occured in jobProcessFunction for job %v", job)
			return
		}
		p.ResultChannel <- Result{JobId: job.JobId, Result: result}
	}
	fmt.Println("Worked completed")
}

func (p *Pool) workerpool(jobProcessFunction JobProcessFunction) {
	defer close(p.ResultChannel)
	wg := sync.WaitGroup{}
	for worker := 0; worker < p.NumOfRoutines; worker++ {
		wg.Add(1)
		go p.work(&wg, jobProcessFunction)
	}
	wg.Wait()
	log.Println("All workers done.")
}

func (p *Pool) collect(resultProcessFunction ResultProcessFunction) {
	for res := range p.ResultChannel {
		err := resultProcessFunction(res)
		if err != nil {
			log.Printf("Error: Error occured in resultProcessFunction for job %v", res.JobId)
			return
		}
		log.Printf("Job with id: [%d] completed", res.JobId)
	}
	p.Done <- true
	p.IsCompleted = true
	fmt.Println("All results are collected!")
}

// This function is for checking tasks completion status
func (p *Pool) taskCompletionStatusChecker() {
	for {
		select {
		case _, channelStatus := <-p.ResultChannel:
			if !channelStatus {
				p.Done <- true
				break
			} else {
				log.Println("Tasks are not complted yet")
			}
		}
	}
}
