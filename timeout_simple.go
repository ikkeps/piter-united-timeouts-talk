package main

import (
    "fmt"
    "time"
    "math/rand"
)

type result struct{
    Data string
}

func doRequest() result{
    
    time.Sleep(time.Duration(rand.Float64()*1.5*float64(time.Second)))
    return result{
        Data: "some data",
    }
}

//START OMIT
func requestWithTimeout(timeout time.Duration){
    deadline := time.After(timeout)
    c := make(chan result)
    
    go func(){
        c <- doRequest() // HL
    }()
    
    select{
    case res := <- c:
        fmt.Printf("%+v\n", res)
    case <-deadline:
        fmt.Println("Deadline")
    }
}
//END OMIT

func main(){
    for n:=0;n<10;n++{
    requestWithTimeout(time.Second)
    }
}


