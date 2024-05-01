package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	reader := NewLogHolder()
	entryScriptPath := "/opt/tiger/data/log/toutiao/log/projects/p_5349367315982062959/.main_02171077407456500000000000000000000ffff0afe6265fa35c5"
	projectDir := "/opt/tiger/data/log/toutiao/log/projects/p_5349367315982062959"
	var envVariableMap = make(map[string]string)
	envVars := os.Environ()
	//envVariableMap["GOPATH"] = os.Getenv("GOPATH")
	//envVariableMap["GOROOT"] = os.Getenv("GOROOT")
	//envVariableMap["GOROOT"] = os.Getenv("GOROOT")
	for key, value := range envVariableMap {
		envVars = append(envVars, key+"="+value)
	}
	cmd := PrepareCmd(entryScriptPath, projectDir, envVars)
	fmt.Printf("dir=%s, %v\n", cmd.Dir, cmd)
	fmt.Printf("dir=%v\n", envVars)
	exitCode, err := RunWithCmd(cmd, reader, time.Duration(3600)*time.Second)
	fmt.Printf("%d %v\n", exitCode, err)
	fmt.Println(reader.String())
}
