package req

type Task struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
	Owner     string `json:"owner"`
}
