package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/personal-finance-crawlers/nubank/types"
)

func main() {
	screen := &types.Hierarchy{}
	b, err := ioutil.ReadFile("screens/nubank_credit_entries.xml")
	if err != nil {
		log.Fatalf("ioutil.ReadFile failed with %s\n", err)
	}
	if err := xml.Unmarshal(b, &screen); err != nil {
		log.Fatalf("xml.Unmarshal failed with %s\n", err)
	}
	log.Printf("Found screen with %d bytes", len(b))

	entries, values := searchHierarchy(screen)

	log.Printf("Entries found: %d", len(entries))

	for i, entry := range entries {
		log.Printf("Entry[%d]: %#v", i+1, entry)
	}
	for i, entry := range values {
		log.Printf("Value[%d]: %#v", i+1, entry)
	}

	encodeToCsvFile(values)
	encodeToJSONFile(entries)
}

func searchHierarchy(hierarchy *types.Hierarchy) ([]*types.NubankEntry, [][]string) {
	var entries []*types.NubankEntry
	values := [][]string{}
	for _, parentNode := range hierarchy.Node.Node {
		for _, parentNode2 := range parentNode.Node.Node.Node.Node.Node {
			for _, parentNode3 := range parentNode2.Node.Node {
				for _, parentNode4 := range parentNode3.Node.Node {
					for _, parentNode5 := range parentNode4.Node {
						for _, parentNode6 := range parentNode5.Node {
							for _, parentNode7 := range parentNode6.Node {
								if strings.Contains(parentNode7.ResourceID, "TV") {
									if entry, value := searchNodes(parentNode6.Node); entry != nil {
										entries = append(entries, entry)
										values = append(values, value)
										log.Printf("Found entry, breaking parentNode6 loop...")
										break
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return entries, values
}

func searchNodes(nodes []types.LeafNode) (*types.NubankEntry, []string) {
	e := &types.NubankEntry{}
	s := make([]string, 0)
	for _, node := range nodes {
		if strings.Contains(node.ResourceID, "title") {
			e.Category = node.AttrText
			s = append(s, node.AttrText)
		}
		if strings.Contains(node.ResourceID, "description") {
			e.Place = node.AttrText
			s = append(s, node.AttrText)
		}
		if strings.Contains(node.ResourceID, "amount") {
			text := strings.Replace(node.AttrText, "R$ ", "", -1)
			text = strings.Replace(text, ",", ".", -1)
			float, err := strconv.ParseFloat(text, 64)
			if err != nil {
				panic(err)
			}
			e.Value = float
			s = append(s, text)
		}
		if strings.Contains(node.ResourceID, "date") {
			e.Time = node.AttrText
			s = append(s, node.AttrText)
		}
	}
	return e, s
}

func nubank() {
	cmd("adb shell input tap 525 1495")
	time.Sleep(2 * time.Second)
	cmd("adb shell uiautomator dump")
	time.Sleep(2 * time.Second)
	b := cmd("adb shell cat /sdcard/window_dump.xml")
	var screen *types.Hierarchy

	if err := xml.Unmarshal(b, screen); err != nil {
		log.Fatalf("xml.Unmarshal failed with %s\n", err)
	}
	log.Printf("%#v", screen)
}

func cmd(cmd string) []byte {
	args := strings.Split(cmd, " ")
	if len(args) <= 1 {
		log.Fatalln("cmd parameter has to be <COMMAND + ' ' + ARGS...>")
	}
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.CombinedOutput() failed with %s\n", err)
	}
	return out
}

func encodeToCsvFile(data [][]string) {
	file, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			panic(err)
		}
	}
}

func encodeToJSONFile(data []*types.NubankEntry) {
	file, err := os.Create("result.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(data)
}
