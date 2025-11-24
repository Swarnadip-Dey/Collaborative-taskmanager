package services

import (
	"log"
	"time"

	"gorm.io/gorm"
)

// StartHealthMonitor launches a background goroutine that periodically pings the
// database to ensure the connection remains healthy. It logs any errors but does
// not terminate the application – the API can continue serving requests while
// the monitor runs.
func StartHealthMonitor(database *gorm.DB, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			<-ticker.C
			sqlDB, err := database.DB()
			if err != nil {
				log.Printf("Health monitor: failed to obtain underlying DB: %v", err)
				continue
			}
			if err := sqlDB.Ping(); err != nil {
				log.Printf("Health monitor: DB ping failed: %v", err)
			} else {
				log.Println("Health monitor: DB connection healthy")
			}
		}
	}()
}
