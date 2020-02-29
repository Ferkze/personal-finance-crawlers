package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
// NODE_RESOURCE_PREFIX = "com.nu.production:id/"
)

func main() {
	screen := &Hierarchy{}
	b, err := ioutil.ReadFile("screens/nubank_credit_entries.xml")
	if err != nil {
		log.Fatalf("ioutil.ReadFile failed with %s\n", err)
	}
	if err := xml.Unmarshal(b, &screen); err != nil {
		log.Fatalf("xml.Unmarshal failed with %s\n", err)
	}
	log.Printf("Found screen with %d bytes", len(b))

	entries := searchHierarchy(screen)

	log.Printf("Entries found: %d", len(entries))

	for i, entry := range entries {
		log.Printf("Entry[%d]: %#v", i+1, entry)
	}

}

func searchHierarchy(hierarchy *Hierarchy) []*NubankEntry {
	var entries []*NubankEntry
	for _, parentNode := range hierarchy.Node.Node {
		for _, parentNode2 := range parentNode.Node.Node.Node.Node.Node {
			for _, parentNode3 := range parentNode2.Node.Node {
				for _, parentNode4 := range parentNode3.Node.Node {
					for _, parentNode5 := range parentNode4.Node {
						for _, parentNode6 := range parentNode5.Node {
							for _, parentNode7 := range parentNode6.Node {
								if strings.Contains(parentNode7.ResourceID, "TV") {
									if entry := searchNodes(parentNode6.Node); entry != nil {
										entries = append(entries, entry)
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
	return entries
}

func searchNodes(nodes []LeafNode) *NubankEntry {
	e := &NubankEntry{}
	for _, node := range nodes {
		if strings.Contains(node.ResourceID, "title") {
			e.Category = node.AttrText
		}
		if strings.Contains(node.ResourceID, "description") {
			e.Place = node.AttrText
		}
		if strings.Contains(node.ResourceID, "amount") {
			e.Value = node.AttrText
		}
		if strings.Contains(node.ResourceID, "date") {
			e.Time = node.AttrText
		}
	}
	return e
}

func nubank() {
	cmd("adb shell input tap 525 1495")
	time.Sleep(2 * time.Second)
	cmd("adb shell uiautomator dump")
	time.Sleep(2 * time.Second)
	b := cmd("adb shell cat /sdcard/window_dump.xml")
	var screen *Hierarchy

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
