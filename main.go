package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func loadEnvFromFile(envPath string) error {
	file, err := os.Open(envPath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		os.Setenv(key, val)
	}
	return scanner.Err()
}

func main() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Ошибка получения пути к .exe:", err)
		return
	}
	baseDir := filepath.Dir(exePath)
	envPath := filepath.Join(baseDir, "assets", "config.txt")

	for {
		err = loadEnvFromFile(envPath)
		if err != nil {
			fmt.Println("[ОШИБКА] Не удалось загрузить файл конфигурации:", envPath)
			fmt.Println("Причина:", err)
		} else {
			fmt.Println("[OK] Конфигурация успешно загружена из:", envPath)
			fmt.Println("[Список переменных из config.txt]:")
			file, err := os.Open(envPath)
			if err != nil {
				fmt.Println("[ОШИБКА] Не удалось повторно открыть config.txt для вывода переменных:", err)
			} else {
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := strings.TrimSpace(scanner.Text())
					if line == "" || strings.HasPrefix(line, "#") {
						continue
					}
					fmt.Println("  ", line)
				}
				if err := scanner.Err(); err != nil {
					fmt.Println("[ОШИБКА] Ошибка при чтении config.txt:", err)
				}
				file.Close()
			}
		}
		fmt.Println("-----------------------------")
		time.Sleep(5 * time.Second)
	}
}
