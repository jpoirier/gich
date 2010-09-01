package main

import (
	"exec"
	"io/ioutil"
	"fmt"
	"syscall"
)

//	cmd, err := exec.Run( "./test", argv, "", nil, exec.DevNull, exec.Pipe, exec.DevNull )

var exeExt string = func () string {
			if syscall.OS == "windows" {
				return ".exe"
			}
			return ""
		}()

func TestFlagH() {
	//             stdin, stdout, stderr
	cmd, err := exec.Run("./giche"+exeExt, []string{"-h"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)
	if err != nil {
		panic("-h: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		panic("-h read:" + err.String())
	}
	if string(buf) != helpMsg {
		panic("-h helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-h close: " + err.String())
	}
}

func TestFlagHelp() {
	//             stdin, stdout, stderr
	cmd, err := exec.Run("./giche"+exeExt, []string{"-help"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)
	if err != nil {
		panic("-help: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		panic("-help read:" + err.String())
	}
	if string(buf) != helpMsg {
		panic("-help helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-help close: " + err.String())
	}
}

func TestNoArgs() {
	//             stdin, stdout, stderr
	cmd, err := exec.Run("./giche"+exeExt, []string{""}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)
	if err != nil {
		panic("NoArgs: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		panic("NoArgs read:" + err.String())
	}
	if string(buf) != helpMsg {
		panic("NoArgs helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("NoArgs close: " + err.String())
	}
}

func TestInvalidArg() {
	//             stdin, stdout, stderr
	cmd, err := exec.Run("./giche"+exeExt, []string{"ZZZ"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)
	if err != nil {
		panic("InvalidArg: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		panic("InvalidArg read:" + err.String())
	}
	if string(buf) != helpMsg {
		panic("InvalidArg helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("InvalidArg close: " + err.String())
	}
}

/*
func TestFlagA() {
	//             stdin, stdout, stderr
	cmd, err := exec.Run("./giche"+exeExt, []string{"-a"}, nil, "", exec.DevNull, exec.Pipe, exec.Pipe)
	if err != nil {
		panic("-a: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		panic("-a read:" + err.String())
	}
	if string(buf) != helpMsg {
		panic("-a helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-a close: " + err.String())
	}
}

func TestFlagS() {
	//             stdin, stdout, stderr
	cmd, err := exec.Run("./giche"+exeExt, []string{"-s"}, nil, "", exec.DevNull, exec.DevNull, exec.Pipe)
	if err != nil {
		panic("-s: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stderr)
	if err != nil {
		panic("-s read:" + err.String())
	}
	if string(buf) != helpMsg {
		panic("-s helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-s close: " + err.String())
	}
}
*/

func main() {
	TestFlagH()
	TestFlagHelp()
	TestNoArgs()
	TestInvalidArg()
//	TestFlagA()
//	TestFlagS()
	fmt.Println("PASS")
}

//func BenchmarkCrc(b *testing.B) {
//	b.StopTimer()

//	// data creation
//	data := make([]uint8, 765)

//	for i := 0; i < len(data); i += 4 {
//		// 0x12345678
//		data[i + 0] = 0x12
//		data[i + 1] = 0x34
//		data[i + 2] = 0x56
//		data[i + 3] = 0x78
//	}

//	b.StartTimer()

//	for i := 0; i < len(tables); i++ {
//		_, _ = Crc(data, tables[i], (6114), crcSzs[i])
//	}
//}
