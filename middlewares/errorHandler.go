package middlewares

import (
	"github.com/gin-gonic/gin"
)

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type CustomizedError struct {
	Message string
}

func (e *CustomizedError) Error() string {
	return e.Message
}

type ServerInternalError struct {
	Message string
}

func (e *ServerInternalError) Error() string {
	return e.Message
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case *ValidationError:
					c.JSON(400, gin.H{"error": "Validation Error"})
					c.Abort()
				case *CustomizedError:
					c.JSON(400, gin.H{"error": e.Error()})
					c.Abort()
				case *ServerInternalError:
					c.JSON(500, gin.H{"error": "Server Internal Error"})
					c.Abort()
				case error:
					c.JSON(500, gin.H{"error": "Unexpected Error"})
					c.Abort()
				default:
					c.JSON(500, gin.H{"error": "Unknown Error"})
					c.Abort()
				}
			}
		}()
		c.Next()
	}
}
