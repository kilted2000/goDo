package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World!")
}

//treat each CRUD element as a seperate DSA problem using PREP

//GET-retrive items on the task list
//POST-send new item to add to talk list
//DELETE-delete task from task list
//PATCH/PUT-edit task list
//print list to terminal using fmt.Println()

// package main

// import (
// 	"net/http"
//   "github.com/gin-gonic/gin"
// )

// type TaskItem struct{
//   ID string `json:"id"`
//   Content string `json:"content"`
//   Checked bool `json:"checked"`
// }

// var tasks = []TaskItem{
//   {ID: "1", Content: "walk the dog", Checked: false},
//   {ID: "2", Content: "feed the dog", Checked: false},
//   {ID: "3", Content: "water the dog", Checked: false},
//   {ID: "4", Content: "clip dogs toenails", Checked: false},
// }
// //main function needed in all Go apps
// func main(){
//  router := gin.Default()
//     router.GET("/tasks", getAllTasks)
//     //router.GET("/tasks/:id", getTasksByID)
//     router.POST("/tasks", postTasks)
//     router.PUT("/tasks/:id", editTasksById)
//     router.DELETE("/tasks/:id", deleteTaskById)

//     router.Run("localhost:8080")

// }

// func getAllTasks(c *gin.Context) {
//     c.IndentedJSON(http.StatusOK, tasks)
// }

// func postTasks(c *gin.Context) {
//     var newTask TaskItem

//     if err := c.BindJSON(&newTask); err != nil {
//         return
//     }

//     tasks = append(tasks, newTask)
//     c.IndentedJSON(http.StatusCreated, newTask)
// }

// func editTasksById(c *gin.Context){
// id := c.Param("id")
//  for _, a := range tasks {
//         if a.ID == id {
//             c.IndentedJSON(http.StatusOK, a)
//             return
//         }
// }
// }


// func deleteTaskById(c *gin.Context){
// id := c.Param("id")
//  for _, a := range tasks {
//         if a.ID == id {
//             c.IndentedJSON(http.StatusOK, a)
//             return
//         }
// }
//  }