package handlers

import (
	"github.com/theunhackable/gator/internal/models"
)

func RegisterHandlers(cmds *models.Commands) {
	cmds.Register("login", HandlerLogin)
	cmds.Register("register", HandlerRegister)

}
