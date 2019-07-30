package main

import (
	"fmt"
	"os/exec"
	"os"
	"strings"
)
func exec_command(program string, args ...string) {
    cmd := exec.Command(program, args...)
    cmd.Stdin = os.Stdin;
    cmd.Stdout = os.Stdout;
    cmd.Stderr = os.Stderr;
    err := cmd.Run() 
    if err != nil {
        fmt.Printf("%v\n", err)
    }
}

func main() {
    search_string := os.Args[1]
    search_path := os.Args[2]
	cmd := exec.Command("grep", "-irn",search_string,search_path )

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	for _, line := range strings.Split(strings.TrimSuffix(string(out),"\n"),"\n") {
		fmt.Println(line)
		line_splice := strings.Split(line,":")
		// Only for Debug
		fmt.Printf("file name: %s\t line no: %s\t\n",line_splice[0],line_splice[1])
		exec_command("vim" , line_splice[0],"+" + line_splice[1])
	}
}
