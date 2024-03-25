package upload

import (
	"fmt"
	"github.com/go-chi/render"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

type ObjectStorage interface {
	Upload(path string, object int64) (int64, error)
}

func New(log *slog.Logger) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("File Upload Endpoint Hit")

		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, handler, err := r.FormFile("file")
		if err != nil {
			log.Error("Error Retrieving the File", "err", err)
			return
		}
		defer file.Close()
		log.Info("Uploaded File:", "file", handler.Filename, "size", handler.Size, "mine", handler.Header)

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		tempFile, err := os.CreateTemp("./storage", fmt.Sprintf("upload-*%s", filepath.Ext(handler.Filename)))
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)
		// return that we have successfully uploaded our file!
		log.Info("Successfully Uploaded File")
		render.JSON(w, r, "Successfully Uploaded File")
	}

}
