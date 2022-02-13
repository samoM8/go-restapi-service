package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-restapi-service/models"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func CreateGroupTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&models.Group{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating group table, Reason: %v\n", createError)
		return createError
	}

	log.Printf("Group table created")
	return nil
}

var db *pg.DB

func InitializeDB(dbConn *pg.DB) {
	db = dbConn
}

// GetAllGroups Get all groups from database
func GetAllGroups(c *gin.Context) {
	var groups []models.Group

	err := db.Model(&groups).Select()

	if err != nil {
		log.Printf("Error while getting all groups, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All groups",
		"data":    groups,
	})
	return
}

// GetSingleGroup Get group from id
func GetSingleGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	group := &models.Group{ID: groupId}
	err := db.Model(group).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting a single group, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Group not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single group",
		"data":    group,
	})
	return
}

// CreateGroup Creates a new group
func CreateGroup(c *gin.Context) {
	var group models.Group
	bindErr := c.BindJSON(&group)
	if bindErr != nil {
		log.Printf("Error while binding data, Reason: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Incorrect parameters",
		})
		return
	}
	id := uuid.New().String()

	newGroup := &models.Group{
		ID:   id,
		Name: group.Name,
	}

	_, insertError := db.Model(newGroup).Insert()
	if insertError != nil {
		log.Printf("Error while inserting new group into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Group created successfully",
	})
	return
}

// EditGroup Edits a group
func EditGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	var group models.Group
	bindErr := c.BindJSON(&group)
	if bindErr != nil {
		log.Printf("Error while binding data, Reason: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Incorrect parameters",
		})
		return
	}
	_, err := db.Model(&models.Group{}).Set("name = ?",
		group.Name).Where("id = ?", groupId).Update()

	if err != nil {
		log.Printf("Error while editing group, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Group edited successfully",
	})
	return
}

// DeleteGroup Deletes a group by id
func DeleteGroup(c *gin.Context) {
	groupId := c.Param("groupId")
	group := &models.Group{ID: groupId}
	_, err := db.Model(group).WherePK().Delete()

	if err != nil {
		log.Printf("Error while deleting a single group, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Group deleted successfully",
	})
	return
}
