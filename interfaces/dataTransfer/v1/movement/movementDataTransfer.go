package category

type MovementInput struct {
	Name string `bson:"name"  json:"name" binding:"required"`
}
