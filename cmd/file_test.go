package cmd

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func zipHandler(w http.ResponseWriter) {
	// Create a buffer to store the ZIP file
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	// Add files to the ZIP archive
	files := map[string]string{
		"en.json": "{\"key1\": \"value1\"}",
		"lv.json": "{\"key1\": \"value1\"}",
	}

	for name, content := range files {
		f, err := zipWriter.Create(name)
		if err != nil {
			http.Error(w, "Failed to create ZIP entry", http.StatusInternalServerError)
			return
		}
		_, err = f.Write([]byte(content))
		if err != nil {
			http.Error(w, "Failed to write to ZIP entry", http.StatusInternalServerError)
			return
		}
	}

	// Close the ZIP writer
	if err := zipWriter.Close(); err != nil {
		http.Error(w, "Failed to close ZIP archive", http.StatusInternalServerError)
		return
	}

	// Set the response headers
	w.Header().Set("Content-Disposition", "attachment; filename=files.zip")
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))

	// Write the ZIP file to the response
	_, err := w.Write(buf.Bytes())
	if err != nil {
		http.Error(w, "Failed to send ZIP file", http.StatusInternalServerError)
	}
}

func TestFileDownload_WithoutWarning(t *testing.T) {
	client, mux, serverUrl, teardown := setup()
	defer teardown()

	outputDir, err := os.MkdirTemp("", "lokalise-cli-2-go-test")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/files/download", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "POST")
			testHeader(t, r, "X-Api-Token", testApiToken)
			data := `{
				"format": "json",
				"original_filenames": true,
				"include_description": true,
				"replace_breaks": true
			}`

			req := new(bytes.Buffer)
			_ = json.Compact(req, []byte(data))

			testBody(t, r, req.String())

			_, _ = fmt.Fprint(w, `{
				"project_id": "`+testProjectID+`",
				"bundle_url": "`+serverUrl+`/objects/export-file.zip"
			}`)
		})

	mux.HandleFunc(
		"/objects/export-file.zip",
		func(w http.ResponseWriter, r *http.Request) {
			zipHandler(w)
		})

	var buf bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&buf)

	args := []string{"file", "download", "--dest=" + outputDir, "--unzip-to=" + outputDir, "--format=json", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	fileDownloadCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	_ = rootCmd.Execute()

	if !bytes.Equal(buf.Bytes(), []byte("")) {
		t.Errorf("Expected log output to be empty")
	}

	log.SetOutput(originalOutput)
	defer os.RemoveAll(outputDir)
}

func TestFileDownload_WithWarning(t *testing.T) {
	client, mux, serverUrl, teardown := setup()
	defer teardown()

	outputDir, err := os.MkdirTemp("", "lokalise-cli-2-go-test")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/files/download", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "POST")
			testHeader(t, r, "X-Api-Token", testApiToken)
			data := `{
				"format": "json",
				"original_filenames": true,
				"include_description": true,
				"replace_breaks": true
			}`

			req := new(bytes.Buffer)
			_ = json.Compact(req, []byte(data))

			testBody(t, r, req.String())

			w.Header().Set("X-Response-Too-Big", "Project too big for sync export. Please use our async export endpoint instead. (/files/async-download)")

			_, _ = fmt.Fprint(w, `{
				"project_id": "`+testProjectID+`",
				"bundle_url": "`+serverUrl+`/objects/export-file.zip"
			}`)
		})

	mux.HandleFunc(
		"/objects/export-file.zip",
		func(w http.ResponseWriter, r *http.Request) {
			zipHandler(w)
		})

	var buf bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&buf)

	args := []string{"file", "download", "--dest=" + outputDir, "--unzip-to=" + outputDir, "--format=json", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	fileDownloadCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	_ = rootCmd.Execute()

	expected := "Warning: Project too big for sync export. Please use our async export endpoint instead. (/files/async-download)\n"
	if !bytes.Contains(buf.Bytes(), []byte(expected)) {
		t.Errorf("Expected log output to contain %q, but got %q", expected, buf.String())
	}

	log.SetOutput(originalOutput)
	defer os.RemoveAll(outputDir)
}

func TestAsyncFileDownload_Error(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	outputDir, err := os.MkdirTemp("", "lokalise-cli-2-go-test")
	if err != nil {
		t.Fatal("Failed to create temporary directory:", err)
	}

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/files/async-download", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "POST")
			testHeader(t, r, "X-Api-Token", testApiToken)
			data := `{
				"format": "json",
				"original_filenames": true,
				"include_description": true,
				"replace_breaks": true
			}`

			req := new(bytes.Buffer)
			_ = json.Compact(req, []byte(data))

			testBody(t, r, req.String())

			_, _ = fmt.Fprint(w, `{
				"process_id": "74738ff5-5367-5958-9aee-98fffdcd1876"
			}`)
		})

	mux.HandleFunc(
		"/api2/projects/"+testProjectID+"/processes/74738ff5-5367-5958-9aee-98fffdcd1876",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "GET")
			testHeader(t, r, "X-Api-Token", testApiToken)

			_, _ = fmt.Fprint(w, `{
                "project_id": "`+testProjectID+`",
                "process": {
                    "status": "failed",
					"message": "Error by test"
				}
			}`)
		})

	var buf bytes.Buffer
	originalOutput := log.Writer()
	log.SetOutput(&buf)

	args := []string{"file", "download", "--async", "--dest=" + outputDir, "--unzip-to=" + outputDir, "--format=json", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	fileDownloadCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	cmdErr := rootCmd.Execute()

	if cmdErr == nil || cmdErr.Error() != "Download failed: Error by test" {
		t.Errorf("Expected error message to be 'Download failed: Error by test', but got %q", cmdErr)
	}

	log.SetOutput(originalOutput)
	defer os.RemoveAll(outputDir)
}
