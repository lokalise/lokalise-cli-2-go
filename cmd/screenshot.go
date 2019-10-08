package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/lokalise/go-lokalise-api"
	"github.com/spf13/cobra"
)

var (
	screenshotId        int64
	newScreenshot       lokalise.NewScreenshot
	newScreenshotFile   string
	newScreenshotOcr    bool
	newScreenshotKeyIds []uint
)

// screenshotCmd represents the screenshot command
var screenshotCmd = &cobra.Command{
	Use: "screenshot",
}

var screenshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists project screenshots",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Screenshots().List(projectId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Adds a screenshot to the project",
	RunE: func(*cobra.Command, []string) error {
		// preparing screenshot
		data, err := screenshotToBase64(newScreenshotFile)
		if err != nil {
			return err
		}
		newScreenshot.Body = data
		newScreenshot.Ocr = &newScreenshotOcr
		for _, id := range newScreenshotKeyIds {
			newScreenshot.KeyIDs = append(newScreenshot.KeyIDs, int64(id))
		}

		s := Api.Screenshots()
		resp, err := s.Create(projectId, []lokalise.NewScreenshot{newScreenshot})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves a screenshot ",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Screenshots().Retrieve(projectId, screenshotId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the properties of a screenshot",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Screenshots().Update(projectId, screenshotId, lokalise.UpdateScreenshot{})
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

var screenshotDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a screenshot from the project.",
	RunE: func(*cobra.Command, []string) error {

		resp, err := Api.Screenshots().Delete(projectId, screenshotId)
		if err != nil {
			return err
		}
		return printJson(resp)
	},
}

func init() {
	screenshotCmd.AddCommand(screenshotListCmd, screenshotCreateCmd, screenshotRetrieveCmd,
		screenshotUpdateCmd, screenshotDeleteCmd)
	rootCmd.AddCommand(screenshotCmd)

	// general flags
	flagProjectId(screenshotCmd, true)

	// Create
	fs := screenshotCreateCmd.Flags()
	fs.StringVar(&newScreenshotFile, "file", "", "Path to file")
	_ = screenshotCreateCmd.MarkFlagRequired("file")
	fs.StringVar(&newScreenshot.Title, "title", "", "Screenshot title")
	fs.StringVar(&newScreenshot.Description, "description", "", "Screenshot description")
	fs.BoolVar(&newScreenshotOcr, "ocr", false, "Try to recognize translations on the image")
	fs.UintSliceVar(&newScreenshotKeyIds, "key-ids", []uint{}, "Attach the screenshot to key IDs specified")
	fs.StringSliceVar(&newScreenshot.Tags, "tags", []string{}, "List of tags to add to the uploaded screenshot")

	// Other
	flagScreenshotId(screenshotUpdateCmd)
	flagScreenshotId(screenshotRetrieveCmd)
	flagScreenshotId(screenshotDeleteCmd)
}

func flagScreenshotId(cmd *cobra.Command) {
	cmd.Flags().Int64Var(&screenshotId, "screenshot-id", 0, "A unique identifier of screenshot (required)")
	_ = cmd.MarkFlagRequired("screenshot-id")
}

func screenshotToBase64(path string) (string, error) {

	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	//noinspection GoUnhandledErrorResult
	defer f.Close()

	cType, err := getFileContentType(f)
	if err != nil {
		return "", err
	}

	if cType != "image/png" && cType != "image/jpeg" {
		return "", fmt.Errorf("unsupported image type: expected image/png or image/jpeg, got %s", cType)
	}

	buf, err := ioutil.ReadFile(path)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return fmt.Sprintf("data:%s;base64,%s", cType, imgBase64Str), nil
}

func getFileContentType(out *os.File) (string, error) {
	// Only the first 512 bytes are used to sniff the content type.
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
