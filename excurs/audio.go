package main

import (
    "fmt"
    "log"

    "github.com/go-audio/audio"
    "github.com/go-audio/audio/wav"
    "github.com/go-audio/wav"
)

func main() {
    // Öffnen Sie das Audiogerät und beginnen Sie mit der Aufnahme
    mic, err := audio.Open(audio.BufferSize(44100*4), audio.DevInput)
    if err != nil {
        log.Fatal(err)
    }
    defer mic.Close()

    // Öffnen Sie eine WAV-Datei zum Speichern der aufgenommenen Sprache
    wavFile, err := wav.NewFile("recorded.wav", mic.Format)
    if err != nil {
        log.Fatal(err)
    }
    defer wavFile.Close()

    fmt.Println("Sprechen Sie etwas. Drücken Sie Enter, um die Aufnahme zu beenden.")

    // Starten Sie die Aufnahme
    err = mic.Start()
    if err != nil {
        log.Fatal(err)
    }

    // Warten auf Enter-Taste, um die Aufnahme zu beenden
    var input string
    fmt.Scanln(&input)

    // Stoppen Sie die Aufnahme
    err = mic.Stop()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Aufnahme beendet.")

    // Lesen Sie die aufgenommenen Daten und schreiben Sie sie in die WAV-Datei
    buffer := make([]int, mic.Size())
    _, err = mic.Read(buffer)
    if err != nil {
        log.Fatal(err)
    }

    _, err = wavFile.WriteInts(buffer)
    if err != nil {
        log.Fatal(err)
    }
}
