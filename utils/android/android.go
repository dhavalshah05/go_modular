package android

import (
	"fmt"
	"learning/utils/io"
	"learning/utils/shell"
	"strings"
)

func GenerateDirectoryName(projectName string) string {
	projectName = strings.TrimSpace(projectName)
	projectName = strings.ReplaceAll(projectName, " ", "_")
	projectName = strings.ReplaceAll(projectName, "-", "_")
	projectName = strings.ToLower(projectName)
	return projectName
}

func GenerateApplicationId(projectName string) string {
	projectName = strings.TrimSpace(projectName)
	projectName = strings.ReplaceAll(projectName, " ", "")
	projectName = strings.ReplaceAll(projectName, "-", "")
	projectName = strings.ReplaceAll(projectName, "_", "")
	projectName = strings.ToLower(projectName)
	return fmt.Sprintf("com.%s.app", projectName)
}

func ChangeApplicationName(projectName string) error {
	path := "./app/src/main/res/values/strings.xml"

	fileContent, err := io.ReadFile(path)
	if err != nil {
		return err
	}

	fileContent = strings.ReplaceAll(fileContent, "AndroidStarter", projectName)
	return io.WriteFile(path, fileContent)
}

func ChangeGradleRootProjectName(dirName string) error {
	path := "./settings.gradle"

	fileContent, err := io.ReadFile(path)
	if err != nil {
		return err
	}

	fileContent = strings.ReplaceAll(fileContent, "AndroidStarter", dirName)
	return io.WriteFile(path, fileContent)
}

func ChangeApplicationId(applicationId string) error {
	path := "./buildSrc/src/main/java/ProjectConfig.kt"

	fileContent, err := io.ReadFile(path)
	if err != nil {
		return err
	}

	fileContent = strings.ReplaceAll(fileContent, "com.androidstarter.app", applicationId)
	return io.WriteFile(path, fileContent)
}

func RemoveGit() error {
	_, _, err := shell.ExecuteShellCommand("rm -rf .git")
	if err != nil {
		return err
	}
	return nil
}

func InitGit() error {
	_, _, err := shell.ExecuteShellCommand("git init")
	if err != nil {
		return err
	}
	return nil
}
