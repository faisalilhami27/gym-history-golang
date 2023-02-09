package category

type CategoryInput struct {
	Name string `bson:"name"  json:"name" binding:"required"`
}
