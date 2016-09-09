package main

import (
    "fmt"

    "github.com/ready-steady/hdf5"
)

func main() {
    put("data.h5")
    get("data.h5")
}

func put(path string) {
    file, _ := hdf5.Create(path)
    defer file.Close()

    A := 42
    file.Put("A", A)

    B := []float64{1, 2, 3}
    file.Put("B", B)

    C := struct {
        D int
        E []float64
    }{
        D: 42,
        E: []float64{1, 2, 3},
    }
    file.Put("C", C)
}

func get(path string) {
    file, _ := hdf5.Open(path)
    defer file.Close()

    A := 0
    file.Get("A", &A)
    fmt.Println(A)

    B := []float64{}
    file.Get("B", &B)
    fmt.Println(B)

    C := struct {
        D int
        E []float64
    }{}
    file.Get("C", &C)
    fmt.Println(C)
}
