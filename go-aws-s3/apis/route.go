package apis

import (
	"fmt"
	"mime/multipart"

	"github.com/adwinugroho/go-aws-s3/config"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/labstack/echo/v4"
)

type (
	awsService interface {
		UploadFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
		ListObject() (*s3.ListObjectsOutput, error)
		DeleteObjects() (string, error)
		DeleteObject(key string) (string, error)
	}

	MyAWSService struct {
		service awsService
	}
)

func Init(e *echo.Echo, service awsService) {
	impService := &MyAWSService{service}
	myAwsRoute := e.Group("/aws-service")
	myAwsRoute.POST("/uploads3", impService.upload)
	myAwsRoute.POST("/list", impService.list)
	myAwsRoute.POST("/delete-list", impService.deleteList)
	myAwsRoute.POST("/delete", impService.delete)
}

func (impService *MyAWSService) upload(c echo.Context) error {
	file, header, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":       false,
			"errorMessage": "Bad Request",
		})
	}
	filename, err := impService.service.UploadFile(file, header)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       false,
			"errorMessage": "Internal Server Error",
		})
	}
	filepath := "https://" + config.BUCKET_NAME + "." + "s3-" + config.AWS_REGION + ".amazonaws.com/" + filename
	return c.JSON(200, map[string]interface{}{
		"status":  true,
		"message": fmt.Sprintf("Upload successfully %s", filepath),
	})
}

func (impService *MyAWSService) list(c echo.Context) error {
	list, err := impService.service.ListObject()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       false,
			"errorMessage": "Internal Server Error",
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status": true,
		"data":   list,
	})
}

func (impService *MyAWSService) deleteList(c echo.Context) error {
	message, err := impService.service.DeleteObjects()
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       false,
			"errorMessage": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  true,
		"message": message,
	})
}

func (impService *MyAWSService) delete(c echo.Context) error {
	body := map[string]interface{}{}
	message, err := impService.service.DeleteObject(body["key"].(string))
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":       false,
			"errorMessage": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"status":  true,
		"message": message,
	})
}
