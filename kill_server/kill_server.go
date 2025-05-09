package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

func main(){
    err := KillServer("server.pid")
    if err != nil {
        fmt.Println("ERROR:", err)
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("not found")
        }
        for e := err; e != nil ; e = errors.Unwrap(e) {
            fmt.Printf("> %s\n", e)
        }
    }   
}
func KillServer(pidFile string) error {
    file, err := os.Open(pidFile)
    if err != nil {
        return err
    }
    /*
    - defer happens when function exits, no matter what
    - defer works at the function level
    - defer are executed in reverse order (stack, LIFO)
    Idiom: try to acquire a resource, check for error, defer release
    */
    defer file.Close()
    var pid int
    if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
        return fmt.Errorf("%q - bad pid: %w", pidFile, err)
    }
    slog.Info("killing", "pid", pid)

    if err := os.Remove(pidFile); err != nil {
        // Not failing, only warning
        slog.Warn("delete", "file", pidFile, "error", err)
    }
    return nil
}