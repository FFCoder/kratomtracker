package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"kratomTracker/doses"
	"kratomTracker/notificationmanager"
	"kratomTracker/notificationmanager/services"
	"net/http"
	"os"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

//go:embed frontend/dist/assets/*
var publicAssets embed.FS

//go:embed frontend/dist/index.html
var indexHTML string

func handleRepoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fsHandler() http.Handler {
	sub, err := fs.Sub(publicAssets, "frontend/dist")
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(sub))
}

func main() {
	RESEND_API_KEY, resendAPIKeyExists := os.LookupEnv("RESEND_API_KEY")
	RESEND_FROM_EMAIL, resendFromExists := os.LookupEnv("RESEND_FROM_EMAIL")

	// Setup Database
	db, err := sql.Open("sqlite3", "kratom_tracker_app.db?_time_format=sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup Repositories
	notifRepo, err := notificationmanager.NewSqliteNotificationRepository(db)
	handleRepoErr(err)
	defer notifRepo.Close()

	// Setup Notification Manager
	notifManager := notificationmanager.NewNotificationManager(notifRepo)
	notifManager.AddService(&services.ConsoleNotificationServce{})
	var emailService notificationmanager.NotificationService

	// Setup Email Notification Service
	if !resendAPIKeyExists || !resendFromExists {
		fmt.Println("Email notification service not configured. Please set RESEND_API_KEY and RESEND_FROM_EMAIL environment variables.")
	} else {
		fmt.Println("Email notification service configured.")
		emailService = services.NewEmailNotificationService(RESEND_API_KEY, RESEND_FROM_EMAIL)

		//emailService.AddEmail("jonathon@jonathonchambers.com")
		err = notifManager.AddService(emailService)
		if err != nil {
			fmt.Printf("Error adding email: %s due to an error: %s", emailService, err)
		}
	}

	// Setup Dose Repository
	doseRepo, err := doses.NewSqliteDoseRepository(db, notifManager)
	handleRepoErr(err)
	defer doseRepo.Close()

	// Setup API Server
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/assets/*filepath", gin.WrapH(fsHandler()))
	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString(indexHTML)
	})

	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	})

	g := router.Group("/api")

	g.GET("/doses", doses.GetAllDoses(doseRepo))
	g.GET("/doses/today", doses.GetAllDosesToday(doseRepo))
	g.POST("/doses", doses.AddDose(doseRepo))
	g.POST("/doses/now", doses.AddDoseNow(doseRepo))

	fmt.Println("Starting server on port 8080")
	apiErr := router.Run(":8080")
	if apiErr != nil {
		panic(apiErr)
	}

}
