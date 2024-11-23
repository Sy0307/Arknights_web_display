package main

import (
    "encoding/json"
    "log"
    "net/http"
    "path/filepath"
)

// 数据结构
type Record struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
    City string `json:"city"`
}

// 全局数据存储
var records = []Record{
    {ID: 1, Name: "张三", Age: 25, City: "北京"},
    {ID: 2, Name: "李四", Age: 30, City: "上海"},
}

// API处理函数
func handleAPI(w http.ResponseWriter, r *http.Request) {
    log.Printf("Received %s request to %s", r.Method, r.URL.Path) // 添加日志

    // 设置 CORS 和 Content-Type 头
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        log.Printf("Sending records: %+v", records) // 添加日志
        if err := json.NewEncoder(w).Encode(records); err != nil {
            log.Printf("Error encoding records: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }

    case "POST":
        var newRecord Record
        if err := json.NewDecoder(r.Body).Decode(&newRecord); err != nil {
            log.Printf("Error decoding request body: %v", err)
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        records = append(records, newRecord)
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"status": "success"})
        log.Printf("Added new record: %+v", newRecord)

    case "OPTIONS":
        w.WriteHeader(http.StatusOK)
    }
}

func main() {
    mux := http.NewServeMux()

    // API 路由
    mux.HandleFunc("/api/data", handleAPI)

    // 处理 submit 页面
    mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join("front-end", "submit.html"))
    })

    // 处理主页
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path == "/" {
            http.ServeFile(w, r, filepath.Join("front-end", "index.html"))
            return
        }
        // 处理其他静态文件
        http.FileServer(http.Dir("front-end")).ServeHTTP(w, r)
    })

    log.Printf("Server starting on http://localhost:8831")
    if err := http.ListenAndServe(":8831", mux); err != nil {
        log.Fatal(err)
    }
}