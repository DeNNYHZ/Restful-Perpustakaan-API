package converters

import (
	"Restful-Perpustakaan-API/app/user"
	"Restful-Perpustakaan-API/member"
)

// ConvertMemberToUser converts a member.Member to a user.User.
func ConvertMemberToUser(m *member.Member) *user.User {
	return &user.User{
		ID:       m.ID,
		Username: m.Name,
		// Map other fields as necessary
	}
}

// ConvertUserToMember converts a user.User to a member.Member.
func ConvertUserToMember(u *user.User) *member.Member {
	return &member.Member{
		ID:   u.ID,
		Name: u.Username,
		// Map other fields as necessary
	}
}
