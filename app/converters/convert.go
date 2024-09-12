package converters

import (
	"Restful-Perpustakaan-API/app/models"
)

// ConvertMemberToUser converts a member.Member to a user.User.
func ConvertMemberToUser(m *models.Member) *models.User {
	return &models.User{
		ID:       m.ID,
		Username: m.Name,
		// Map other fields as necessary
	}
}

// ConvertUserToMember converts a user.User to a member.Member.
func ConvertUserToMember(u *models.User) *models.Member {
	return &models.Member{
		ID:   u.ID,
		Name: u.Username,
		// Map other fields as necessary
	}
}
