package main

import (
	"log"
	"net/http"

	"github.com/7t1cker/bp/api/handlers"
	"github.com/7t1cker/bp/api/handlers/handlers_learn"
	"github.com/7t1cker/bp/api/handlers/handlers_qests"
	"github.com/7t1cker/bp/api/handlers/handlers_skills"
	"github.com/7t1cker/bp/api/handlers/handlers_user"
	"github.com/7t1cker/bp/api/handlers/hot_qests"
	"github.com/7t1cker/bp/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
    db, err := db.Connect()
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }
    defer db.Close()

    r := gin.Default()

    // User routes
    r.POST("/api/v2/users", handlers_user.CreateUser(db))//+
	r.POST("/api/v2/add-skill-from-learning", handlers_user.AddSkillFromLearning(db))//
    r.POST("/api/v2/login", handlers_user.Login(db))//+
    r.POST("/api/v2/logout", handlers_user.Logout(db))//+
	r.GET("/api/v2/users/:user_id/assigned_quests", handlers_user.GetAssignedQuests(db))//+
	r.GET("/api/v2/user", handlers_user.KL(db))//+
    // Division, Group  routes
    r.POST("/api/v2/divisions", handlers.CreateDivisions(db))//+
    r.POST("/api/v2/groups", handlers.CreateGroups(db))//+
    

    // Quest routes
    r.POST("/api/v2/quest", handlers_qests.CreateQuest(db))//+
    r.POST("/api/v2/quest/hot", hot_qests.CreateHotTask(db))//+
    r.POST("/api/v2/quest-complite", handlers_qests.MarkQuestAsDone(db))//+
    r.PUT("/api/v2/purpose-qests", handlers_qests.UpdateAssignee(db))//+
    r.GET("/api/v2/quest", handlers_qests.GetQests(db))//+

    // Learn routes
    r.POST("/api/v2/learn", handlers_learn.CreateLearn(db))//+
    r.PUT("/api/v2/learn/:id", handlers_learn.UpdateLearnTitle(db))//+
	r.GET("/api/v2/learn", handlers_learn.GetAllLearnings(db))//+
	r.GET("/api/v2/learn/:id", handlers_learn.GetLearningByID(db))//+
	r.DELETE("/api/v2/learn/:id", handlers_learn.DeleteLearning(db))//+
	

    // Skill routes
	r.POST("/api/v2/skills", handlers_skills.CreateSkills(db))//+
    r.GET("/api/v2/skills", handlers_skills.GetAllSkills(db))//+
	

    // Start server
    log.Println("Server is running on port 8000")
    http.ListenAndServe(":8000", r)
}
