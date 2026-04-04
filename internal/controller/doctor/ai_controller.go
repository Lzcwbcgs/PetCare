package doctor

type AIController struct{}

func NewAI() *AIController {
	return &AIController{}
}

func derefPage(value *int, fallback int) int {
	if value == nil {
		return fallback
	}
	return *value
}
