package notas

import (
	"fmt"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

var lines []string
var results Results

// ParsePDF reads and parses orders from broker pdf
func ParsePDF() {
	var err error
	pdfs := []string{
		"pdf/NotaCorretagem12.pdf",
		"pdf/NotaCorretagem11.pdf",
	}
	for _, p := range pdfs {
		if err = unlockPdf(p, p, "485") ; err != nil {
			fmt.Printf("ERROR: Error unlocking pdf %q: %#v", p, err)
			panic(err.Error())
		}
		if err = parse(p); err != nil {
			fmt.Printf("ERROR: Error parsing pdf %q: %#v", p, err)
			panic(err.Error())
		}
	}
}

func parse(inputPath string) error {
	fmt.Printf("Parsing PDF %q\n", inputPath)
	
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}
	positions := make(DayTradePositions)
	swings := make(SwingTradePositions)

	// swings := make(SwingTradePositions)
	results = make(Results)

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err
		}
		fmt.Printf("Extracting page %d...\n", pageNum)

		extract(text, positions)

		if !strings.Contains(text, "CONTINUA...") {
			updatePositions(positions, swings)
			calculateResults(results, positions, swings)
		}
	}
	
	printResults(results)

	// err = json.WriteJSON("results.json", results)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// err = json.WriteJSON("positions.json", positions)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	return nil
}

func unlockPdf(inputPath string, outputPath string, password string) error {
	pdfWriter := pdf.NewPdfWriter()

	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	isEncrypted, err := pdfReader.IsEncrypted()
	if err != nil {
		return err
	}

	// Try decrypting both with given password and an empty one if that fails.
	if isEncrypted {
		auth, err := pdfReader.Decrypt([]byte(password))
		if err != nil {
			return err
		}
		if !auth {
			return fmt.Errorf("Wrong password")
		}
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		err = pdfWriter.AddPage(page)
		if err != nil {
			return err
		}
	}

	fWrite, err := os.Create(outputPath)
	if err != nil {
		return err
	}

	defer fWrite.Close()

	err = pdfWriter.Write(fWrite)
	if err != nil {
		return err
	}

	return nil
}