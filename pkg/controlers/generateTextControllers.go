package controlers

import (
	"encoding/json"
	"github.com/unidoc/unipdf/v3/model"
	"io"
	"net/http"
	"os"
	"textQuillBackend/pkg/utils"
)

type PDFTextResponse struct {
	Text string `json:"text"`
}

func HandleGenerateText(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form to get the file part.
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data.
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a temporary file to store the uploaded file.
	tempFile, err := os.CreateTemp("", "uploaded-*.pdf") //directory the file
	if err != nil {
		http.Error(w, "Failed to create temporary file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// Copy the file contents to the temporary file.
	_, err = io.Copy(tempFile, file)
	if err != nil {
		http.Error(w, "Failed to copy file", http.StatusInternalServerError)
		return
	}

	// Open the uploaded PDF file.
	pdfFile, err := os.Open(tempFile.Name())
	if err != nil {
		http.Error(w, "Failed to open PDF file", http.StatusInternalServerError)
		return
	}
	defer pdfFile.Close()

	// Extract text from the PDF.
	pdfReader, err := model.NewPdfReader(pdfFile)
	if err != nil {
		http.Error(w, "Failed to create PDF reader", http.StatusInternalServerError)
		return
	}

	// Extract text from all pages.
	var extractedText string
	numPages, _ := pdfReader.GetNumPages()
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			http.Error(w, "Failed to get PDF page", http.StatusInternalServerError)
			return
		}

		extractedText += utils.ExtractTextFromPage(page)
	}

	// Return the extracted text as JSON.
	response := PDFTextResponse{Text: extractedText}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
