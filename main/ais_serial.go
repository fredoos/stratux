package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func initAISSerial() {
    go readAISSerial("/dev/ttyAMA0")
}

func readAISSerial(port string) {
    f, err := os.Open(port)
    if err != nil {
        log.Printf("Erreur lors de l'ouverture du port AIS série: %s\n", err)
        return
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if strings.HasPrefix(line, "!AIVDM") {
            handleAISLine(line)
        }
    }

    if err := scanner.Err(); err != nil {
        log.Printf("Erreur lors de la lecture du port AIS série: %s\n", err)
    }
}

func handleAISLine(line string) {
    processNMEALine(line)
}
