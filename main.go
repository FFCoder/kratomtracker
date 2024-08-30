package main

import (
	"context"
	"database/sql"
	"embed"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"kratomTracker/doses"
	"kratomTracker/notificationmanager"
	"kratomTracker/notificationmanager/services"
	"kratomTracker/remindersManager"
	"net/http"
	"os"
	"strconv"
	"time"
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

// printLocalTime prints the local time to the console.
func printLocalTime() {
	fmt.Println("Local Time: ", time.Now().Local().Format("2006-01-02 15:04:05"))
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
	ctx, cancel := context.WithCancel(context.Background())

	// Print the local time to the console.
	printLocalTime()

	// Get Flags
	portFlag := flag.Int("port", 8080, "Port to run the server on")
	flag.Parse()

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

	// Setup Reminder Manager
	reminderManager := remindersManager.NewReminderManager()
	go reminderManager.Start(ctx)
	nextDoseTime, err := doseRepo.GetNextDoseTime()
	if err != nil {
		fmt.Println("Error getting next dose time: ", err)
	} else {
		reminderManager.SetReminder("Take Kratom", nextDoseTime)
	}

	// Setup API Server
	devMode := os.Getenv("DEV_MODE") == "true"
	if devMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.GET("/assets/*filepath", gin.WrapH(fsHandler()))
	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString(indexHTML)
	})

	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		context.Header("Access-Control-Allow-Headers", "Content-Type")
	})

	g := router.Group("/api")

	g.GET("/doses", doses.GetAllDoses(doseRepo))
	g.GET("/doses/today", doses.GetAllDosesToday(doseRepo))
	g.GET("/doses/next", doses.GetNextDoseTime(doseRepo))
	g.POST("/doses", doses.AddDose(doseRepo))
	g.POST("/doses/now", doses.AddDoseNow(doseRepo))

	portNumber := *portFlag
	portNumberStr := strconv.Itoa(portNumber)
	fmt.Println("Starting server on port " + portNumberStr)
	apiErr := router.Run(":" + portNumberStr)
	if apiErr != nil {
		panic(apiErr)
	}
	cancel()

}
