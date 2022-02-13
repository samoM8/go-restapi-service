package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/go-restapi-service/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.Model(&models.User{}).CreateTable(opts)
	if createError != nil {
		log.Printf("Error while creating user table, Reason: %v\n", createError)
		return createError
	}

	log.Printf("User table created")
	return nil
}

// GetAllUsers Get all users from database
func GetAllUsers(c *gin.Context) {
	var users []models.User

	//err := db.Model(&users).Select() //Get users with no groups
	err := db.Model(&users).
		Column("user.*").
		Relation("Group").
		Select()

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All users",
		"data":    users,
	})
	return
}

// GetSingleUser Get single user
func GetSingleUser(c *gin.Context) {
	userId := c.Param("userId")

	//user := &models.User{ID: userId}
	//err := db.Model(user).WherePK().Select() //get only user without group
	/*	i get some syntax error in where
		user := new(models.User)
		err := db.Model(user).
			Column("user.*").
			Relation("Group").
			Where("user.id = ?", userId).
			Select()
	*/

	var users []models.User
	err := db.Model(&users).
		Column("user.*").
		Relation("Group").
		Select()

	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	//todo: fix if find a better solution
	var user models.User
	found := false
	for _, u := range users {
		if userId == u.ID {
			user = u
			found = true
			break
		}
	}
	if !found {
		log.Printf("User with given id does not exist")
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single user",
		"data":    user,
	})
	return
}

// CreateUser Creates a new user
func CreateUser(c *gin.Context) {
	var user models.User
	bindErr := c.BindJSON(&user)
	if bindErr != nil {
		log.Printf("Error while binding data, Reason: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Incorrect parameters",
		})
		return
	}

	id := uuid.New().String()
	password, _ := HashPassword(user.Password)
	newUser := &models.User{
		ID:       id,
		Email:    user.Email,
		Password: password,
		Name:     user.Name,
		GroupId:  user.GroupId,
	}

	_, insertError := db.Model(newUser).Insert()
	if insertError != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created successfully",
	})
	return
}

// EditUser Edits a user
func EditUser(c *gin.Context) {
	userId := c.Param("userId")

	oldUser := &models.User{ID: userId}
	errOldUser := db.Model(oldUser).WherePK().Select() //get only user without group
	if errOldUser != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", errOldUser)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	user := &models.User{ID: userId}
	bindErr := c.BindJSON(&user)

	if bindErr != nil {
		log.Printf("Error while binding data, Reason: %v\n", bindErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Incorrect parameters",
		})
		return
	}

	// Use data from old user if not in request body
	if user.Email == "" {
		user.Email = oldUser.Email
	}
	if user.Password == "" {
		user.Password = oldUser.Password
	} else {
		user.Password, _ = HashPassword(user.Password)
	}
	if user.Name == "" {
		user.Name = oldUser.Name
	}
	if user.GroupId == "" {
		user.GroupId = oldUser.GroupId
	}

	_, err := db.Model(user).WherePK().Update()
	if err != nil {
		log.Printf("Error while editing user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User edited successfully",
	})
	return
}

// DeleteUser Deletes a user by id
func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &models.User{ID: userId}
	_, err := db.Model(user).WherePK().Delete()

	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
