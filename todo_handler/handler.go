package todo_handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/withoutsecondd/ToDo/internal/utils"
	"github.com/withoutsecondd/ToDo/models"
	"github.com/withoutsecondd/ToDo/service"
	"strconv"
	"strings"
)

type Handler struct {
	entityService service.EntityService
	authService   service.AuthService
}

func NewHandler(es service.EntityService, as service.AuthService) *Handler {
	return &Handler{entityService: es, authService: as}
}

func (h *Handler) SetupRoutes(a *fiber.App) {
	api := a.Group("/api")
	api.Post("/login", h.Login)

	users := api.Group("/users")
	users.Get("/", h.GetCurrentUser)
	users.Post("/", h.CreateUser)

	lists := api.Group("/lists")
	lists.Get("/", h.GetListsByCurrentUser)

	tasks := api.Group("/tasks")
	tasks.Get("/", h.GetTasksById) // Requires list_id as a query variable, if no provided, returns tasks by current user)
}

func (h *Handler) extractToken(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("no authorization token provided")
	}

	tokenSlice := strings.Split(authHeader, "Bearer ")
	if len(tokenSlice) != 2 {
		return "", errors.New("invalid token format")
	}

	return tokenSlice[1], nil
}

func (h *Handler) GetCurrentUser(c *fiber.Ctx) error {
	tokenStr, err := h.extractToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	userId, err := h.authService.AuthorizeWithToken(tokenStr)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	user, err := h.entityService.GetUserById(userId)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusNotFound, err)
	}

	return utils.FormatSuccessResponse(c, user)
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	createdUser, err := h.entityService.CreateUser(user)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
	}

	return utils.FormatSuccessResponse(c, createdUser)
}

func (h *Handler) GetListsByCurrentUser(c *fiber.Ctx) error {
	tokenStr, err := h.extractToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	userId, err := h.authService.AuthorizeWithToken(tokenStr)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	lists, err := h.entityService.GetListsByUserId(userId)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	return utils.FormatSuccessResponse(c, lists)
}

// GetTasksById returns response with tasks specified by list id they belong to.
// If no list id is provided, returns tasks by user
func (h *Handler) GetTasksById(c *fiber.Ctx) error {
	tokenStr, err := h.extractToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	userId, err := h.authService.AuthorizeWithToken(tokenStr)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	listIdStr := c.Queries()["list_id"]

	if listIdStr == "" {
		tasks, err := h.entityService.GetTasksByUserId(userId)
		if err != nil {
			return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return utils.FormatSuccessResponse(c, tasks)
	} else {
		listId, err := strconv.Atoi(listIdStr)
		if err != nil {
			return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
		}

		list, err := h.entityService.GetListById(int64(listId))

		if list.UserID != userId {
			return utils.FormatErrorResponse(c, fiber.StatusForbidden, errors.New("this list doesn't belong to current user"))
		}

		tasks, err := h.entityService.GetTasksByListId(list.ID)
		if err != nil {
			return utils.FormatErrorResponse(c, fiber.StatusInternalServerError, err)
		}

		return utils.FormatSuccessResponse(c, tasks)
	}
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var loginRequest service.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, errors.New("incorrect email or password"))
	}

	token, err := h.authService.Authenticate(&loginRequest)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	return utils.FormatSuccessResponse(c, fiber.Map{"token": token})
}
