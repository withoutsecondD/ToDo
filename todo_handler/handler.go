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
	api.Post("/login", h.login)

	users := api.Group("/users")
	users.Get("/", h.getCurrentUser)
	users.Post("/", h.createUser)

	lists := api.Group("/lists")
	lists.Get("/:listId", h.getListById)
	lists.Get("/", h.getListsByCurrentUser)

	tasks := api.Group("/tasks")
	tasks.Get("/:taskId", h.getTaskById)
	tasks.Get("/", h.getTasksById) // Requires list_id as a query variable, if no provided, returns tasks by current user
}

// extractToken extracts a token from request's Authorization header
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

// getCurrentUser returns a response with token bearer's information as a models.UserResponse
func (h *Handler) getCurrentUser(c *fiber.Ctx) error {
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

// createUser creates a user and returns a response with created user information as a models.UserResponse
func (h *Handler) createUser(c *fiber.Ctx) error {
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

// getListById returns a response with a list specified by its id
func (h *Handler) getListById(c *fiber.Ctx) error {
	tokenStr, err := h.extractToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	userId, err := h.authService.AuthorizeWithToken(tokenStr)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	listId, err := strconv.Atoi(c.Params("listId"))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, errors.New("invalid listId parameter"))
	}

	list, err := h.entityService.GetListById(int64(listId), userId)
	if err != nil {
		switch {
		case errors.As(err, &utils.DBError{}):
			return utils.FormatErrorResponse(c, fiber.StatusNotFound, err)
		case errors.As(err, &utils.ForbiddenError{}):
			return utils.FormatErrorResponse(c, fiber.StatusForbidden, err)
		}
	}

	return utils.FormatSuccessResponse(c, list)
}

// getListsByCurrentUser returns a response with lists of token bearer.
func (h *Handler) getListsByCurrentUser(c *fiber.Ctx) error {
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

// getTasksById returns a response with tasks specified by list id they belong to.
// If no list id is provided, returns tasks by token bearer.
func (h *Handler) getTasksById(c *fiber.Ctx) error {
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

		tasks, err := h.entityService.GetTasksByListId(int64(listId), userId)
		if err != nil {
			switch err.Error() {
			case "this list doesn't belong to current user":
				return utils.FormatErrorResponse(c, fiber.StatusForbidden, err)
			default:
				return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
			}
		}

		return utils.FormatSuccessResponse(c, tasks)
	}
}

// getTaskById return response with task specified by its id.
func (h *Handler) getTaskById(c *fiber.Ctx) error {
	tokenStr, err := h.extractToken(c)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	userId, err := h.authService.AuthorizeWithToken(tokenStr)
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusUnauthorized, err)
	}

	taskId, err := strconv.Atoi(c.Params("taskId"))
	if err != nil {
		return utils.FormatErrorResponse(c, fiber.StatusBadRequest, err)
	}

	task, err := h.entityService.GetTaskById(int64(taskId), userId)
	if err != nil {
		switch {
		case errors.As(err, &utils.DBError{}):
			return utils.FormatErrorResponse(c, fiber.StatusNotFound, err)
		case errors.As(err, &utils.ForbiddenError{}):
			return utils.FormatErrorResponse(c, fiber.StatusForbidden, err)
		}
	}

	return utils.FormatSuccessResponse(c, task)
}

// login is used for authentication, service.LoginRequest is used
// as credentials struct. If credentials are valid this handler will return
// response with a token.
func (h *Handler) login(c *fiber.Ctx) error {
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
