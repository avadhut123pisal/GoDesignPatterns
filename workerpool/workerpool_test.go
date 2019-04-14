package workerpool

import (
	"log"
	"testing"
)

// unit test cases
func TestWorkerpoolWithResultProcessorFunction(t *testing.T) {
	resources := []int{3, 4, 65, 656, 2, 76, 7}
	var resourcesList []interface{}
	
	for _, resource := range resources {
		resourcesList = append(resourcesList, resource)
	}
	poolObj := GenerateNewPool(6, true)
	poolObj.Start(resourcesList, demoJobProcessorFunction, demoResultProcessFunction)
}

func TestWorkerpoolWithoutResultProcessorFunction(t *testing.T) {
	resources := []int{3, 4, 65, 656, 2, 76, 7}
	var resourcesList []interface{}
	
	for _, resource := range resources {
		resourcesList = append(resourcesList, resource)
	}
	poolObj := GenerateNewPool(6, false)
	poolObj.Start(resourcesList, demoJobProcessorFunction, demoResultProcessFunction)
}

func demoJobProcessorFunction(job Job) (Result, error) {
	log.Printf("Job %v proceesed", job)
	result := Result{}
	return result, nil
}

func demoResultProcessFunction(result Result) error {
	log.Printf("Result is %v", result)
	return nil
}
