package main

import (
	//"context"
	"testing"
	//"time"
	"strings"
	"github.com/google/uuid"
	//"sync"

	"github.com/testcontainers/testcontainers-go"
	//"github.com/testcontainers/testcontainers-go/wait"
)

/*
// TODO: Set up the Docker Compose test env as a singleton. That way you 
// can call GetTestEnv() at the start of each test case, and DestroyTestEnv()
// at the end of some test cases - only if the test env has been compromised by
// executing the test case
type TestEnvDetail struct {
	compose testcontainers
}
var testEnv *TestEnvDetail

func GetTestEnv() *TestEnvDetail {
    sync.Once.Do(func() {
        testEnv = &TestEnvDetail{state: "off"}
    })
    return testEnv
}

func DestroyTestEnv() *TestEnvDetail {
	execError := compose.Down()
	return execError, compose
}
*/

func TestWithCompose(t *testing.T) {
	//ctx := context.Background()
	composeFilePaths := []string {"./docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(map[string]string {
			"key1": "value1",
			"key2": "value2",
		}).
		Invoke()
	err := execError.Error
	if err != nil {
		t.Error(err)
		//return fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	}
	//return nil
	//composeFilePaths = []string{"testresources/docker-compose.yml"}

	//compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError = compose.Down()
	err = execError.Error
	if err != nil {
		t.Error(err)
		//return fmt.Errorf("Could not run compose file: %v - %v", composeFilePaths, err)
	}
	//return nil

}