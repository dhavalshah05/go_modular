package main

import (
	"fmt"
	"learning/utils/android"
	"learning/utils/input"
	"learning/utils/shell"
)

const projectUrl = "http://192.168.29.5:3000/dhaval/android_starter_multi_module.git"

func main() {
	fmt.Print("Enter project name: ")
	projectName := input.GetStringFromUser()

	directoryName := android.GenerateDirectoryName(projectName)
	applicationId := android.GenerateApplicationId(projectName)

	err := shell.CloneRepository(projectUrl, directoryName)
	if err != nil {
		panic(err)
	}

	err = shell.ChangeDir(directoryName)
	if err != nil {
		panic(err)
	}

	err = android.ChangeApplicationName(projectName)
	if err != nil {
		panic(err)
	}

	err = android.ChangeGradleRootProjectName(directoryName)
	if err != nil {
		panic(err)
	}

	err = android.ChangeApplicationId(applicationId)
	if err != nil {
		panic(err)
	}

	err = android.RemoveGit()
	if err != nil {
		panic(err)
	}

	err = android.InitGit()
	if err != nil {
		panic(err)
	}

	fmt.Println("Application created")
}
