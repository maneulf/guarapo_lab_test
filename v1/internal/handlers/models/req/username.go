package req

type Username struct {
	Username string `json:"username" binding:"required"`
}
