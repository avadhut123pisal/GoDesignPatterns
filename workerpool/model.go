package workerpool

// Every task is of job type
type Job struct {
	JobId    int
	Resource interface{}
}

// result processor function should return Result type value 
type Result struct {
	JobId  int
	Result interface{}
	IsErr  bool
}
// Worker pool 
type Pool struct {
	NumOfRoutines             int           // no. of concurrent running goroutines  
	JobChannel                chan Job      // channel for jobs (user tasks)
	ResultChannel             chan Result   // channel to keep track of processed results
	Done                      chan bool     // channel to track completion of tasks
	IsResultProcessorRequired bool          // to decide if user has result processing function
	IsCompleted               bool          // future scope for all tasks completion status
}

// type for processor function (dependency injection)
type JobProcessFunction func(resource Job) (Result, error)

// type for result process function (dependency injection)
type ResultProcessFunction func(result Result) error
