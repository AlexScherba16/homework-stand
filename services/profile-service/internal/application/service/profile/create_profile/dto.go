package create_profile

type CreateProfileDTO struct {
	userID        int64
	name          string
	email         string
	isTaskAllowed bool
}

func NewCreateProfileDTO(userID int64, name, email string, isTaskAllowed bool) CreateProfileDTO {
	return CreateProfileDTO{
		userID:        userID,
		name:          name,
		email:         email,
		isTaskAllowed: isTaskAllowed,
	}
}
