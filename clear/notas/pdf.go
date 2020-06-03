package notas

import (
	"fmt"
	"os"
	"strings"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

var lines []string

func Test() {
	err := unlockPdf("../nota-de-corretagem.pdf", "../nota-de-corretagem-2.pdf", "485")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	err = outputPdfText("../nota-de-corretagem-2.pdf")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	return
}

func outputPdfText(inputPath string) error {
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

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
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

		lines = strings.Split(text, "\n")

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		for _, line := range lines {
			texts := strings.Split(line, " ")
			if strings.HasPrefix(texts[0], "C") || strings.HasPrefix(texts[0], "V") {
				fmt.Printf("\"%s\"\n", line)
			}
		}
		fmt.Println("------------------------------")
	}

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