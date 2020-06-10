package json

import (
	"encoding/json"
	"os"
	"time"

	jap "github.com/ferkze/jsonappender"
)

// AppendArrayJSON escreve dados marshaveis no arquivo YYYY-MM-DD-data.json
func AppendArrayJSON(fileName string, data interface{}) (err error) {
	outDir := "out/"
	if fileName == "" {
		fileName = time.Now().Format("2006-01-02-15-04-05")+"-data.json"
	}
	
	a, err := jap.JSONAppender(outDir + fileName)
	if err != nil {
		return err
	}

	var b []byte
	b, err = json.Marshal(&data)
  if err != nil {
    return
  }

  if _, err = a.Write(b); err != nil {
    return
  }

  if err = a.Close(); err != nil {
    return
  }
	return
}

// WriteJSON escreve dados marshaveis no arquivo .json
func WriteJSON(fileName string, data interface{}) (err error) {
	outDir := "out/"
	if fileName == "" {
		fileName = time.Now().Format("2006-01-02-15-04-05")+"-data.json"
	}

	file := outDir+fileName
	err = checkFile(file)
	if err != nil {
		return
	}
		
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()  

	return json.NewEncoder(f).Encode(data)
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}