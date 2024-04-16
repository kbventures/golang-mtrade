// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/wpcodevo/golang-gorm-postgres/initializers"
// )

// var (
// 	server *gin.Engine
// )

// func init() {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("? Could not load environment variables", err)
// 	}

// 	initializers.ConnectDB(&config)

// 	server = gin.Default()
// }

// func main() {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("? Could not load environment variables", err)
// 	}

// 	router := server.Group("/api")
// 	router.GET("/healthchecker", func(ctx *gin.Context) {
// 		message := "Welcome to Golang with Gorm and Postgres"
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
// 	})

// 	log.Fatal(server.Run(":" + config.ServerPort))
// }

// package main

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// 	"github.com/wpcodevo/golang-gorm-postgres/initializers"
// )

// var (
// 	server *gin.Engine
// 	db     *sql.DB
// )

// func init() {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("Could not load environment variables", err)
// 	}

// 	db, err = sql.Open("postgres", config.DBConnectionString)
// 	if err != nil {
// 		log.Fatal("Failed to connect to database", err)
// 	}

// 	server = gin.Default()
// }

// func main() {
// 	defer db.Close()

// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("Could not load environment variables", err)
// 	}

// 	router := server.Group("/api")
// 	router.GET("/healthchecker", func(ctx *gin.Context) {
// 		message := "Welcome to Golang with Gorm and Postgres"
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
// 	})

// 	router.GET("/users", func(ctx *gin.Context) {
// 		rows, err := db.Query("SELECT * FROM \"User\"") // Use double quotes to escape reserved keyword "User"
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
// 			return
// 		}
// 		defer rows.Close()

// 		users := []map[string]interface{}{}
// 		for rows.Next() {
// 			user := make(map[string]interface{})
// 			err := rows.Scan(&user["id"], &user["name"], &user["email"]) // Adjust field names as per your table schema
// 			if err != nil {
// 				log.Println("Error scanning row:", err)
// 				continue
// 			}
// 			users = append(users, user)
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{"users": users})
// 	})

// 	log.Fatal(server.Run(":" + config.ServerPort))
// }

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/wpcodevo/golang-gorm-postgres/initializers"
// )

// var (
// 	server *gin.Engine
// )

// func init() {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("Could not load environment variables", err)
// 	}

// 	initializers.ConnectDB(&config)

// 	server = gin.Default()
// }

// func main() {
// 	config, err := initializers.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("Could not load environment variables", err)
// 	}

// 	router := server.Group("/api")
// 	router.GET("/healthchecker", func(ctx *gin.Context) {
// 		message := "Welcome to Golang with Gorm and Postgres"
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
// 	})

// 	router.GET("/user", func(ctx *gin.Context) {
// 		var user []User
// 		result := initializers.DB.Find(&user)
// 		if result.Error != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
// 			return
// 		}

// 		ctx.JSON(http.StatusOK, gin.H{"users": user})
// 	})

// 	log.Fatal(server.Run(":" + config.ServerPort))
// }

// // Define your User model struct here
// type User struct {
// 	ID    uint   `gorm:"primaryKey"`
// 	Name  string `gorm:"not null"`
// 	Email string `gorm:"unique;not null"`
// }

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/user", func(ctx *gin.Context) {
		// Get the underlying *sql.DB instance from GORM's DB
		db, err := initializers.DB.DB()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get database connection"})
			return
		}
		defer db.Close()

		rows, err := db.Query("SELECT id, username, name FROM \"User\"")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to execute query", "details": err.Error()})
			return
		}
		defer rows.Close()

		// Process query results
		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Username, &user.Name)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row", "details": err.Error()})
				return
			}
			users = append(users, user)
		}

		// Return users as JSON response
		ctx.JSON(http.StatusOK, gin.H{"users": users})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}

// Define your User model struct here
// User represents a user in the database
type User struct {
	ID            string     `json:"id"`
	Username      string     `json:"username"`
	Name          *string    `json:"name"`
	Password      *string    `json:"password"`
	EmailVerified *time.Time `json:"emailVerified,omitempty"` // Optional field
	Image         *string    `json:"image"`
}
