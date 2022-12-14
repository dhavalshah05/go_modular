package main

import (
	_ "embed"
	"fmt"
	"learning/utils/android"
	"learning/utils/input"
	"learning/utils/shell"
)

//go:embed seed/project.zip
var projectBytes []byte

func main() {
	fmt.Print("Enter project name: ")
	projectName := input.GetStringFromUser()
	directoryName := android.GenerateDirectoryName(projectName)
	applicationId := android.GenerateApplicationId(projectName)

	err := android.CreateAndroidProject(directoryName, projectBytes)
	handleError(err)

	err = shell.ChangeDir(directoryName)
	handleError(err)

	err = android.ChangeApplicationName(projectName)
	handleError(err)

	err = android.ChangeGradleRootProjectName(directoryName)
	handleError(err)

	err = android.ChangeApplicationId(applicationId)
	handleError(err)

	err = android.RemoveGit()
	handleError(err)

	err = android.InitGit()
	handleError(err)

	err = android.MakeGradleExecutable()
	handleError(err)

	fmt.Println("Application created")
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
