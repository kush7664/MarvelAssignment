package main 

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type Avengers struct{
    A_Name string `json:"a_name"`
    Max_Power int `json:"max_power"`
}

type Anti_Heroes struct{
    Anti_Heroes_Name string `json:"anti_heroes"`
    Anti_Heroes_Max_power int `json:"anti_heroes_max_power"`
}

type Mutant struct{
    M_Name string `json:"m_name"`
    M_Max_Power int `json:"m_max_power"`
}

func main(){
    fmt.Println("Starting the application...")
    response1, err := http.Get("http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b")
    response2, err := http.Get("http://www.mocky.io/v2/5ecfd630320000f1aee3d64d")
    response3, err := http.Get("http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data1, _ := ioutil.ReadAll(response1.Body)
        data2, _ := ioutil.ReadAll(response2.Body)
        data3, _ := ioutil.ReadAll(response3.Body)
        fmt.Println(string(data1))
        fmt.Println(string(data2))
        fmt.Println(string(data3))
        
    }
    fmt.Println("Terminating the application...")
}
