package main

import (
	"exec"
	"io/ioutil"
	"fmt"
	"syscall"
//	"testing"
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
	c := "giche" + exeExt
	cmd, err := exec.Run(c, []string{c, "-h"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("-h: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
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
	c := "./giche" + exeExt
	cmd, err := exec.Run(c, []string{c, "-help"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("-help: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
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
	c := "giche" + exeExt
	cmd, err := exec.Run(c, []string{c}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("NoArgs: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
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
	c := "giche" + exeExt
	cmd, err := exec.Run(c, []string{c, "ZZZ"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("InvalidArg: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		panic("InvalidArg read:" + err.String())
	}
	if string(buf) != "" {
		panic("InvalidArg helpMsg: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("InvalidArg close: " + err.String())
	}
}

/*
func TestFlagA() {
	//             stdin, stdout, stderr
	c := "giche" + exeExt
	cmd, err := exec.Run("./giche"+exeExt, []string{"-a"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("-a: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
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
*/

func TestFlagS1() {
	//             stdin, stdout, stderr
	c := "giche" + exeExt
	cmd, err := exec.Run(c, []string{c, "-s", "invalid_arg"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("-s: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		panic("-s read:" + err.String())
	}
	if string(buf) != "None" {
		panic("-s : got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-s close: " + err.String())
	}
}

func TestFlagS2() {
	//             stdin, stdout, stderr
	c := "giche" + exeExt
	cmd, err := exec.Run(c, []string{c, "-s", "cat"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("-s: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		panic("-s read:" + err.String())
	}
	if string(buf) != "Found" {
		panic("-s : got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-s close: " + err.String())
	}
}


func TestMisc() {
	//             stdin, stdout, stderr
	c := "giche" + exeExt
	cmd, err := exec.Run(c, []string{c, "cat"}, nil, "", exec.DevNull, exec.Pipe, exec.DevNull)
	if err != nil {
		panic("TestMisc: " + err.String())
	}
	buf, err := ioutil.ReadAll(cmd.Stdout)
	if err != nil {
		panic("TestMisc read:" + err.String())
	}
	if string(buf) != "/bin/cat\n" {
		panic("TestMisc cat: got " + string(buf))
	}
	if err = cmd.Close(); err != nil {
		panic("-s close: " + err.String())
	}
}

  //func Benchmark() {
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

func main() {
	TestFlagH()
	TestFlagHelp()
	TestNoArgs()
	TestInvalidArg()
	TestMisc()
//	if syscall.OS == "windows" {
//	} else {

//	}
//	TestFlagA()
	TestFlagS1()
	TestFlagS2()
	fmt.Println("PASS")
}


