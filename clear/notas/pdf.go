package notas

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

var lines []string
var results Results

// ParsePDF reads and parses orders from broker pdf
func ParsePDF() {
	// err := unlockPdf("pdf/nota-de-corretagem.pdf", "pdf/nota-de-corretagem-UNLOCKED.pdf", "485")
	// if err != nil {
	// 	panic(err.Error())
	// }
	err := parse("pdf/nota-de-corretagem-UNLOCKED.pdf")
	if err != nil {
		panic(err.Error())
	}
	return
}

func parse(inputPath string) error {
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
	positions := make(map[string]Position)

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