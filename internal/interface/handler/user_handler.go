package handler

import (
"net/http"

"github.com/gin-gonic/gin"
"github.com/google/uuid"
"web-server/internal/application/usecase"
"web-server/internal/domain/entity"
"web-server/internal/infrastructure/middleware"
)

type UserHandler struct {
userUseCase *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
return &UserHandler{
userUseCase: uc,
}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
var user entity.User
if err := c.ShouldBindJSON(&user); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

// Generate UUID
user.ID = uuid.New().String()

// Set default role if not provided
if user.Role == "" {
user.Role = "user"
}

// Hash password
hashedPassword, err := entity.HashPassword(user.Password)
if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
return
}
user.Password = hashedPassword

if err := h.userUseCase.CreateUser(&user); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

// Don't return the password in the response
user.Password = ""
c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
id := c.Param("id")
user, err := h.userUseCase.GetUser(id)
if err != nil {
c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
return
}

c.JSON(http.StatusOK, user)
}

func (h *UserHandler) LoginUser(c *gin.Context) {
var loginRequest struct {
Email    string `json:"email" binding:"required,email"`
Password string `json:"password" binding:"required"`
}

if err := c.ShouldBindJSON(&loginRequest); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

user, err := h.userUseCase.GetUserByEmail(loginRequest.Email)
if err != nil {
c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
return
}

// Verify password
if !entity.CheckPassword(loginRequest.Password, user.Password) {
c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
return
}

// Generate JWT token
token, err := middleware.GenerateToken(user.ID, user.Role)
if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
return
}

c.JSON(http.StatusOK, gin.H{
"token": token,
"user": gin.H{
"id":       user.ID,
"username": user.Username,
"email":    user.Email,
"role":     user.Role,
},
})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
id := c.Param("id")
var user entity.User
if err := c.ShouldBindJSON(&user); err != nil {
c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
return
}

user.ID = id // Ensure the ID matches the URL parameter
if err := h.userUseCase.UpdateUser(&user); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
id := c.Param("id")
if err := h.userUseCase.DeleteUser(id); err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
users, err := h.userUseCase.ListUsers()
if err != nil {
c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
return
}

c.JSON(http.StatusOK, users)
}