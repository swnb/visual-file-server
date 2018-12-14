package config

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// resolveEnvConfigLine define the resolve ruler of env config
func resolveEnvConfigLine(envLine string, envMap *map[string]string) {
	commonIndex := strings.Index(envLine, ";;")
	if commonIndex != -1 {
		envLine = envLine[:commonIndex] // remove comment
	}

	kv := strings.Split(envLine, "=")
	if len(kv) != 2 {
		return
	}

	key := strings.TrimSpace(kv[0])
	value := strings.TrimSpace(kv[1])

	(*envMap)[key] = value
}

func resolveEnv() map[string]string {
	envMap := make(map[string]string)

	file, err := os.Open("./env")
	if err != nil {
		return envMap
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)

	for {
		envLine, err := fileReader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		} else if err != io.EOF {
			resolveEnvConfigLine(envLine, &envMap)
		} else {
			// when comes to the end of line,
			// ReadString('\n') will throw err io.EOF,
			// but we want it to continue resolve config Line
			resolveEnvConfigLine(envLine, &envMap)
			break
		}
	}

	return envMap
}
