package server

import (
	"encoding/json"
	"fmt"
)

type Siswa struct {
	Nim   string `json:"nim"`
	Nama  string `json:"nama"`
	Nilai int    `json:"nilai"`
}

func GetSiswaObject() {
	var jsonString = `{"nim": "234234", "nama": "Irawan Kilmer", "nilai": 98}`
	var jsonData = []byte(jsonString)

	var data Siswa

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nim:", data.Nim)
	fmt.Println("Nama:", data.Nama)
	fmt.Println("Nilai:", data.Nilai)
}

func GetMapInterface() {
	var jsonString = `{"nim": "23423", "nama": "Irawan Kilmer", "nilai": 90}`
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Nim", data["nim"])
	fmt.Println("Nama", data["nama"])
	fmt.Println("Nilai", data["nilai"])
}

func GetArrayObject() {
	var jsonString = `[
		{"nim": "2342", "nama": "Jajang Nurjaman", "nilai": 87},
		{"nim": "9867", "nama": "Irawan Kilmer", "nilai": 98},
		{"nim": "1237", "nama": "Ghania Jilbi", "nilai": 94}
	]`

	var data []Siswa
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i, d := range data {
		fmt.Printf("No: %d | Nim: %s | Nama: %s | Nilai: %d \n", i+1, d.Nim, d.Nama, d.Nilai)
	}
}

func GetObjectJson() {
	var object = []Siswa{
		{"23423", "Ghania Jilbi", 90},
		{"91863", "Hikamul Albi", 90},
		{"91283", "Rhania Syaranopa", 90},
	}

	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)

	fmt.Println(jsonString)
}
