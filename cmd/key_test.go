package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/spf13/cobra"
)

func TestKeyBulkDelete(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/keys", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "DELETE")
			testHeader(t, r, "X-Api-Token", testApiToken)

			data := `{"keys":[123,456]}`
			req := new(bytes.Buffer)
			_ = json.Compact(req, []byte(data))
			testBody(t, r, req.String())

			_, _ = fmt.Fprint(w, `{
				"project_id": "`+testProjectID+`",
				"keys_removed": true,
				"keys_locked": 0
			}`)
		})

	args := []string{"key", "bulk-delete", "--key-ids=123,456", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	keyBulkDeleteCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	if err := rootCmd.Execute(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestKeyBulkUpdate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/keys", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "PUT")
			testHeader(t, r, "X-Api-Token", testApiToken)

			_, _ = fmt.Fprint(w, `{
				"project_id": "`+testProjectID+`",
				"keys": [
					{
						"key_id": 123,
						"key_name": {
							"ios": "updated_key",
							"android": "updated_key",
							"web": "updated_key",
							"other": "updated_key"
						},
						"platforms": ["web"],
						"filenames": {},
						"description": "Updated description",
						"tags": ["updated"],
						"comments": [],
						"screenshots": [],
						"translations": []
					}
				]
			}`)
		})

	keysJSON := `[{"key_id":123,"description":"Updated description","tags":["updated"]}]`
	args := []string{"key", "bulk-update", "--keys=" + keysJSON, "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	keyBulkUpdateCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	if err := rootCmd.Execute(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestKeyBulkCreate(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc(
		fmt.Sprintf("/api2/projects/%s/keys", testProjectID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			testMethod(t, r, "POST")
			testHeader(t, r, "X-Api-Token", testApiToken)

			_, _ = fmt.Fprint(w, `{
				"project_id": "`+testProjectID+`",
				"keys": [
					{
						"key_id": 789,
						"key_name": {
							"ios": "welcome",
							"android": "welcome",
							"web": "welcome",
							"other": "welcome"
						},
						"platforms": ["web"],
						"filenames": {},
						"description": "",
						"tags": [],
						"comments": [],
						"screenshots": [],
						"translations": []
					},
					{
						"key_id": 790,
						"key_name": {
							"ios": "goodbye",
							"android": "goodbye",
							"web": "goodbye",
							"other": "goodbye"
						},
						"platforms": ["ios", "android"],
						"filenames": {},
						"description": "",
						"tags": ["v2"],
						"comments": [],
						"screenshots": [],
						"translations": []
					}
				]
			}`)
		})

	keysJSON := `[{"key_name":"welcome","platforms":["web"],"tags":[]},{"key_name":"goodbye","platforms":["ios","android"],"tags":["v2"]}]`
	args := []string{"key", "bulk-create", "--keys=" + keysJSON, "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	keyBulkCreateCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	if err := rootCmd.Execute(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestKeyBulkUpdate_InvalidJSON(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	args := []string{"key", "bulk-update", "--keys=not-json", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	keyBulkUpdateCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	err := rootCmd.Execute()
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}

func TestKeyBulkCreate_InvalidJSON(t *testing.T) {
	client, _, _, teardown := setup()
	defer teardown()

	args := []string{"key", "bulk-create", "--keys=not-json", "--project-id=" + testProjectID}
	rootCmd.SetArgs(args)
	keyBulkCreateCmd.PreRun = func(cmd *cobra.Command, args []string) {
		Api = client
	}

	err := rootCmd.Execute()
	if err == nil {
		t.Error("Expected error for invalid JSON, got nil")
	}
}
