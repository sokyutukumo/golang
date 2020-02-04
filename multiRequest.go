package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    )


func GetContent (key string, url string, c chan map[string]string) {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    
    s := string(body[:])
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    result := make(map[string]string)
    result[key] = s
    c <- result
}

func GetContent2(urlList map[string]string, c chan map[string]string) {
    defer close(c)
    for key, url := range urlList {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Println(err)
            continue
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        
        s := string(body[:])
        
        if err != nil {
            fmt.Println(err)
            continue
        }
        
        result := make(map[string]string)
        result[key] = s
        c <- result        
    }
}


func main() {
    // Your code here!
    
    urlList := make(map[string]string)
    urlList["y"] = "https://www.yahoo.co.jp/"
    urlList["g"] = "https://www.google.co.jp/"
    
    c := make(chan map[string]string, len(urlList))
    go GetContent2(urlList, c)
    
    for result := range c {
        for key, content := range result {
            fmt.Println(key)
            fmt.Println(content)
        }
    }
    
    
    // for key, url := range urlList {
    //     go GetContent(key, url, c)
    // }

    // for i := 0; i < len(urlList); i++ {
    //     fmt.Println(<-c)
    // }
    // close(c)
    // fmt.Println(result01)
    // fmt.Println("********-----------")
    // fmt.Println(result02)
}
