package cmd

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/lokalise/go-lokalise-api/v5"
	"github.com/spf13/cobra"
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

// secretBundlePath mimics a pre-signed URL: both the path and the query are sensitive.
const secretBundlePath = "/objects/secret-bundle-path.zip"
const secretBundleQuery = "X-Amz-Signature=secret-signature"

// runDownload executes `file download` with a clean flag state and returns its stdout.
func runDownload(t *testing.T, client *lokalise.Api, args ...string) (string, error) {
	t.Helper()

	// cobra keeps flag values between Execute calls: reset the ones the tests touch.
	downloadAsync, downloadJsonOnly = false, false
	_ = rootCmd.PersistentFlags().Set("quiet", "false")
	t.Cleanup(func() {
		downloadAsync, downloadJsonOnly = false, false
		_ = rootCmd.PersistentFlags().Set("quiet", "false")
	})

	outputDir := t.TempDir()
	rootCmd.SetArgs(append([]string{
		"file", "download", "--dest=" + outputDir, "--unzip-to=" + outputDir, "--format=json", "--project-id=" + testProjectID,
	}, args...))
	fileDownloadCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	return captureStdout(t, rootCmd.Execute)
}

func captureStdout(t *testing.T, f func() error) (string, error) {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	original := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = original }()

	out := make(chan string)
	go func() {
		b, _ := io.ReadAll(r)
		out <- string(b)
	}()

	fErr := f()
	_ = w.Close()

	return <-out, fErr
}

func syncDownloadHandler(serverUrl string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprint(w, `{
			"project_id": "`+testProjectID+`",
			"bundle_url": "`+serverUrl+secretBundlePath+`?`+secretBundleQuery+`"
		}`)
	}
}

func assertNoBundleURL(t *testing.T, label, s string) {
	t.Helper()
	for _, secret := range []string{"secret-bundle-path", "secret-signature"} {
		if strings.Contains(s, secret) {
			t.Errorf("%s leaks the bundle URL (%q): %q", label, secret, s)
		}
	}
}

func TestFileDownload_DoesNotPrintBundleURL(t *testing.T) {
	client, mux, serverUrl, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api2/projects/%s/files/download", testProjectID), syncDownloadHandler(serverUrl))
	mux.HandleFunc(secretBundlePath, func(w http.ResponseWriter, r *http.Request) { zipHandler(w) })

	out, err := runDownload(t, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertNoBundleURL(t, "stdout", out)
	if !strings.Contains(out, "Downloading bundle (") {
		t.Errorf("expected a 'Downloading bundle (<size>)...' line, got %q", out)
	}
}

func TestAsyncFileDownload_DoesNotPrintDownloadURL(t *testing.T) {
	client, mux, serverUrl, teardown := setup()
	defer teardown()

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/files/async-download", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = fmt.Fprint(w, `{"process_id": "74738ff5-5367-5958-9aee-98fffdcd1876"}`)
		})
	mux.HandleFunc(
		"/api2/projects/"+testProjectID+"/processes/74738ff5-5367-5958-9aee-98fffdcd1876",
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = fmt.Fprint(w, `{
				"project_id": "`+testProjectID+`",
				"process": {
					"status": "finished",
					"details": {"download_url": "`+serverUrl+secretBundlePath+`?`+secretBundleQuery+`"}
				}
			}`)
		})
	mux.HandleFunc(secretBundlePath, func(w http.ResponseWriter, r *http.Request) { zipHandler(w) })

	out, err := runDownload(t, client, "--async")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assertNoBundleURL(t, "stdout", out)
	if strings.Contains(out, "Downloading file:") {
		t.Errorf("unexpected legacy 'Downloading file:' line in %q", out)
	}
}

func TestFileDownload_Quiet(t *testing.T) {
	client, mux, serverUrl, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api2/projects/%s/files/download", testProjectID), syncDownloadHandler(serverUrl))
	mux.HandleFunc(secretBundlePath, func(w http.ResponseWriter, r *http.Request) { zipHandler(w) })

	out, err := runDownload(t, client, "--quiet")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if out != "" {
		t.Errorf("expected no output with --quiet, got %q", out)
	}
}

func TestFileDownload_JsonOnlyStillReturnsBundleURL(t *testing.T) {
	client, mux, serverUrl, teardown := setup()
	defer teardown()

	mux.HandleFunc(fmt.Sprintf("/api2/projects/%s/files/download", testProjectID), syncDownloadHandler(serverUrl))

	// --quiet must not hide command results.
	out, err := runDownload(t, client, "--json-only", "--quiet")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, secretBundlePath+"?"+secretBundleQuery) {
		t.Errorf("expected --json-only to print the bundle URL, got %q", out)
	}
}

func TestFormatSize(t *testing.T) {
	for n, want := range map[int64]string{0: "0 B", 1023: "1023 B", 1024: "1.0 KiB", 1536: "1.5 KiB", 5 * 1024 * 1024: "5.0 MiB"} {
		if got := formatSize(n); got != want {
			t.Errorf("formatSize(%d) = %q, want %q", n, got, want)
		}
	}
}
