package handlerUser

import (
	"belajar-go/utils"
	"net/http"
)

func (s *UserController) FindByIdHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok := ctx.Value("user_id").(int)
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	user, err := s.service.GetProfile(ctx, userId)

	if err != nil {
		utils.Error(w, http.StatusNotFound, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "user profile retrieved successfully", user)
}
