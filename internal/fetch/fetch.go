package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joshmenden/aocplz/internal/printit"
)

func FetchDayInput(day, year *int, dontOpen *bool) (err error) {
	printit.Info(fmt.Sprintf("Fetching input and creating test file for AOC%v day %v", *year, *day))
	dailySolutionDir, err := createDailySolutionDir(*day, *year)
	if err != nil {
		return
	}
	printit.Info(fmt.Sprintf("Created new directory: %s", dailySolutionDir))

	err = fetchInput(*day, *year, dailySolutionDir)
	if err != nil {
		return
	}
	printit.Info("Created new .txt files with the sample and the custom test input")

	filePath, err := createSolutionFile(dailySolutionDir)
	if err != nil {
		return
	}
	printit.Info(fmt.Sprintf("Created new solution file: %s", filePath))

	if !*dontOpen {
		ok := openPuzzle(*day, *year)
		if !ok {
			return fmt.Errorf("could not open browser to puzzle")
		}
		printit.Info("Opening browser to relevant puzzle...")
	}

	return
}

func createDailySolutionDir(day int, year int) (dir string, err error) {
	dir = fmt.Sprintf("%s/%v/day-%v", os.Getenv("AOCPLZ_ROOT_DIR"), year, day)
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

	bytes, err := io.ReadAll(resp.Body)
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

func copyFile(srcPath *string, destPath string, srcBytes *[]byte, executable *bool) (err error) {
	var bytesToCopy []byte
	if srcBytes == nil && srcPath != nil {
		bytesToCopy, err = os.ReadFile(*srcPath)
		if err != nil {
			return err
		}
	} else if srcBytes != nil {
		bytesToCopy = *srcBytes
	} else {
		return fmt.Errorf("given neither a file path or bytes to copy")
	}

	err = os.WriteFile(destPath, bytesToCopy, 0644)
	if err != nil {
		return
	}

	if executable != nil && *executable {
		err = os.Chmod(destPath, 0700)
		if err != nil {
			return
		}
	}

	return
}

func createSolutionFile(dailySolutionDir string) (solutionPath string, err error) {
	parentDir := filepath.Dir(dailySolutionDir)
	files, err := os.ReadDir(parentDir)
	if err != nil {
		return "", err
	}

	var templatePath string
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".tmpl") {
			templatePath = filepath.Join(parentDir, file.Name())
			break
		}
	}

	if templatePath == "" {
		return "", fmt.Errorf("no template file found in %s", parentDir)
	}

	destFilename := strings.ReplaceAll(filepath.Base(templatePath), ".tmpl", "")
	destPath := filepath.Join(dailySolutionDir, destFilename)
	boolValue := true
	err = copyFile(&templatePath, destPath, nil, &boolValue)
	if err != nil {
		return "", err
	}

	return destPath, nil
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
