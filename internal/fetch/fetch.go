package fetch

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/joshmenden/aocplz/internal/printit"
)

func FetchDayInput(day, year *int) (err error) {
	printit.Info(fmt.Sprintf("Fetching input and creating test file for AOC%v day %v", *year, *day))
	solutionDir, err := createDir(*day)
	if err != nil {
		return
	}
	printit.Info(fmt.Sprintf("Created new directory: %s", solutionDir))

	err = fetchInput(*day, *year, solutionDir)
	if err != nil {
		return
	}
	printit.Info("Created new input.txt file with data")

	filePath, err := createSolutionFiles(solutionDir)
	if err != nil {
		return
	}
	printit.Info(fmt.Sprintf("Created new test files: %s", filePath))

	ok := openPuzzle(*day, *year)
	if !ok {
		return fmt.Errorf("could not open browser to puzzle")
	}
	printit.Info("Opening browser to relevant puzzle...")

	return
}

func createDir(day int) (dir string, err error) {
	dir = fmt.Sprintf("%s/day-%v", os.Getenv("AOCPLZ_ROOT_DIR"), day)
	err = os.Mkdir(dir, 0755)

	return
}

func getRawDataFromURL(url string, aocReq bool) (*[]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if aocReq {
		req.AddCookie(&http.Cookie{Name: "session", Value: os.Getenv("AOCPLZ_SESSION_TOKEN")})
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &bytes, nil
}

func fetchInput(day, year int, dir string) (err error) {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)
	raw, err := getRawDataFromURL(url, true)
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprintf("%s/input.txt", dir), *raw, 0644)
	if err != nil {
		return
	}

	err = os.WriteFile(fmt.Sprintf("%s/sample.txt", dir), nil, 0644)
	if err != nil {
		return
	}

	return
}

func copyFile(srcPath *string, destPath string, srcBytes *[]byte) (err error) {
	var bytesToCopy []byte
	if srcBytes == nil && srcPath != nil {
		bytesToCopy, err = ioutil.ReadFile(*srcPath)
		if err != nil {
			return err
		}
	} else if srcBytes != nil {
		bytesToCopy = *srcBytes
	} else {
		return fmt.Errorf("given neither a file path or bytes to copy")
	}

	err = ioutil.WriteFile(destPath, bytesToCopy, 0644)
	if err != nil {
		return
	}

	err = os.Chmod(destPath, 0700)
	if err != nil {
		return
	}

	return
}

func copyLocalFiles(templateDir, fileNames, solutionDir string) (files []string, err error) {
	tmplsStrArr := strings.Split(fileNames, ",")
	for _, tmplName := range tmplsStrArr {
		destName := strings.ReplaceAll(tmplName, ".tmpl", "")
		srcPath := fmt.Sprintf("%s/%s", templateDir, tmplName)
		destPath := fmt.Sprintf("%s/%s", solutionDir, destName)
		err = copyFile(&srcPath, destPath, nil)
		if err != nil {
			return
		} else {
			files = append(files, destName)
		}
	}

	return
}

func copyGlobalFilesFromGithub(solutionDir string) (files []string, err error) {
	solutionFileData, err := getRawDataFromURL("https://raw.githubusercontent.com/joshmenden/aocplz/main/templates/a.rb.tmpl", false)
	if err != nil {
		return
	}

	err = copyFile(nil, fmt.Sprintf("%s/%s", solutionDir, "a.rb"), solutionFileData)
	if err != nil {
		return
	} else {
		files = append(files, "github.com/.../a.rb")
	}

	gemfileData, err := getRawDataFromURL("https://raw.githubusercontent.com/joshmenden/aocplz/main/templates/Gemfile.tmpl", false)
	if err != nil {
		return
	}

	err = copyFile(nil, fmt.Sprintf("%s/%s", solutionDir, "Gemfile"), gemfileData)
	if err != nil {
		return
	} else {
		files = append(files, "github.com/.../Gemfile")
	}

	aocSolutionData, err := getRawDataFromURL("https://raw.githubusercontent.com/joshmenden/aocplz/main/templates/aoc_solution.rb.tmpl", false)
	if err != nil {
		return
	}

	err = copyFile(nil, fmt.Sprintf("%s/%s", os.Getenv("AOCPLZ_ROOT_DIR"), "aoc_solution.rb"), aocSolutionData)
	if err != nil {
		return
	} else {
		files = append(files, "github.com/.../aoc_solution.rb")
	}

	return
}

func createSolutionFiles(solutionDir string) (files []string, err error) {
	tmplsDir := os.Getenv("AOCPLZ_TEMPLATES_DIR")
	tmplsStr := os.Getenv("AOCPLZ_TEMPLATE_FILES")

	if tmplsDir != "" && tmplsStr != "" {
		return copyLocalFiles(tmplsDir, tmplsStr, solutionDir)
	} else {
		return copyGlobalFilesFromGithub(solutionDir)
	}
}

func openPuzzle(day, year int) bool {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v", year, day)

	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}

	cmd := exec.Command(args[0], append(args[1:], url)...)

	return cmd.Start() == nil
}
