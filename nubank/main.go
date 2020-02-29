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
	NODE_RESOURCE_PREFIX = "com.nu.production:id/"
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
	log.Printf("%#v", screen)

	entries := make([]NubankEntry, 0)
	arr := searchHierarchy(screen)

	log.Printf("Text: %#v\nEntries: %#v", arr, entries)

}

func searchHierarchy(hierarchy *Hierarchy) []string {
	var arr []string
	for _, node := range hierarchy.Node.Node {
		if node.AttrText != "" {
			arr = append(arr, node.AttrText)
		}
		for _, node2 := range node.Node.Node.Node.Node.Node {
			if node2.AttrText != "" {
				arr = append(arr, node2.AttrText)
			}
			for _, node3 := range node2.Node.Node {
				if node3.AttrText != "" {
					arr = append(arr, node3.AttrText)
				}
				for _, node4 := range node3.Node.Node {
					if node4.AttrText != "" {
						arr = append(arr, node4.AttrText)
					}
					for _, node5 := range node4.Node {
						if node5.AttrText != "" {
							arr = append(arr, node5.AttrText)
						}
						for _, node6 := range node5.Node {
							if node6.AttrText != "" {
								arr = append(arr, node6.AttrText)
							}
							for _, node7 := range node6.Node {
								if node7.AttrText != "" {
									arr = append(arr, node7.AttrText)
								}
							}
						}
					}
				}
			}
		}
	}
	return arr
}

func searchNode() *NubankEntry {
	e := &NubankEntry{}
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
