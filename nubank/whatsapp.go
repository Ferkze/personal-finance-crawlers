package main

import (
	"time"
)

func whatsapp() {
	_ = cmd("adb shell input tap 980 2222")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input tap 910 170")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input text amo")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input tap 518 333")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input tap 330 1524")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input text oi")
	_ = cmd("adb shell input keyevent 62")
	_ = cmd("adb shell input text amo")
	_ = cmd("adb shell input keyevent 62")
	_ = cmd("adb shell input text da")
	_ = cmd("adb shell input keyevent 62")
	_ = cmd("adb shell input text minha")
	_ = cmd("adb shell input keyevent 62")
	_ = cmd("adb shell input text vida")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input tap 992 1524")
	time.Sleep(1 * time.Second)
	_ = cmd("adb shell input swipe 992 1524 992 1524 5000")

	// input swipe 100 500 100 500 250
}
