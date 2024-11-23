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
    // 设置 CORS 头
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Set("Content-Type", "application/json")

    switch r.Method {
    case "GET":
        if err := json.NewEncoder(w).Encode(records); err != nil {
            log.Printf("Error encoding GET response: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        log.Println("GET request handled successfully")
        
    case "POST":
        // 读取请求体
        var newRecord Record
        err := json.NewDecoder(r.Body).Decode(&newRecord)
        if err != nil {
            log.Printf("Error decoding request body: %v", err)
            http.Error(w, "Invalid request body", http.StatusBadRequest)
            return
        }

        // 打印接收到的数据
        log.Printf("Received new record: %+v", newRecord)

        // 添加到记录中
        records = append(records, newRecord)
        
        // 返回成功响应
        response := map[string]interface{}{
            "success": true,
            "message": "Data added successfully",
            "data": newRecord,
        }
        
        if err := json.NewEncoder(w).Encode(response); err != nil {
            log.Printf("Error encoding POST response: %v", err)
            http.Error(w, "Internal server error", http.StatusInternalServerError)
            return
        }
        log.Println("POST request handled successfully")

    case "OPTIONS":
        response := map[string]bool{"ok": true}
        json.NewEncoder(w).Encode(response)
        log.Println("OPTIONS request handled")

    default:
        response := map[string]string{"error": "Method not allowed"}
        w.WriteHeader(http.StatusMethodNotAllowed)
        json.NewEncoder(w).Encode(response)
        log.Printf("Unsupported method: %s", r.Method)
    }
}

func main() {
    // 创建多路复用器
    mux := http.NewServeMux()

    // API路由
    mux.HandleFunc("/api/data", handleAPI)

    // 表格页面
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join("front-end", "index.html"))
    })

    // 提交页面
    mux.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, filepath.Join("front-end", "submit.html"))
    })

    // 静态文件服务
    mux.Handle("/src/", http.StripPrefix("/src/", http.FileServer(http.Dir("front-end/src"))))
    mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("front-end/public"))))

    // 启动服务器
    log.Println("Server starting on http://localhost:8831")
    if err := http.ListenAndServe(":8831", mux); err != nil {
        log.Fatal(err)
    }
}