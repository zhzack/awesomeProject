package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"github.com/robotn/gohook"
	"github.com/tarm/serial"
	"log"
	"time"
)

var screenX, screenY, mouseX, mouseY, curMouPosX, curMouPosY int
var data []string
var isWrite = false

func low() {
	evChan := hook.Start()
	defer hook.End()
	for ev := range evChan {
		//fmt.Println(ev)
		if ev.Kind == 9 {
			curMouPosX = int(ev.X) * 240 / 2560
			curMouPosY = int(ev.Y) * 240 / 1440
		} else if ev.Kind >= 6 && ev.Kind <= 7 {
			sprintData := fmt.Sprintf("%v%v", ev.Kind, ev.Button)
			sentData(sprintData)
		}

	}
}

func init() {
	screenX, screenY = robotgo.GetScreenSize()
	println(screenX, screenY)
	mouseX, mouseY = robotgo.GetMousePos()
	println(mouseX, mouseY)
}

func sentDataByte(s []byte) {
	_, err := conn.Write(s)
	if err != nil {
		log.Println(err)
		return
	}
}

func sentData(s string) {
	for {
		if !isWrite {
			isWrite = true
			sentDataByte([]byte(s))
			isWrite = false
			return
		}
		time.Sleep(time.Millisecond * 100)
	}

}

var err error
var conn *serial.Port

func mainfff() {

	//设置串口编号
	ser := &serial.Config{Name: "COM47", Baud: 1500000}
	//打开串口
	conn, err = serial.OpenPort(ser)
	if err != nil {
		log.Fatal(err)
	}
	//sentData(Conn, "ssss")
	go low()
	go func() {
		for {
			if mouseX != curMouPosX || mouseY != curMouPosY {
				mouseX = curMouPosX
				mouseY = curMouPosY
				sentData(fmt.Sprintf("%v,%v,%v", 9, mouseX, mouseY))
			}
			time.Sleep(time.Millisecond * 100)

		}
	}()

	//保持数据持续接收
	for {
		buf := make([]byte, 1024)
		lens, err := conn.Read(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		revData := buf[:lens]
		fmt.Println(string(revData))
	}

}
