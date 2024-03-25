package model

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    Name   string `json:"name"`
    Status string `json:"status"`
    Description string `json:"description"`
    Title string `json:"title"`
}
