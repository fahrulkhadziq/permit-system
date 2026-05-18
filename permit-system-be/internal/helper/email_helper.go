package helper

import "permit-license/internal/model"

func ExtractEmails(users []model.User) []string {
	var emails []string

	for _, user := range users {
		emails = append(emails, user.Email)
	}
	return emails
}

func MergeEmails(emailGroups ...[]string) []string {
	emailMap := map[string]bool{}

	var result []string

	for _, group := range emailGroups {
		for _, email := range group {
			if !emailMap[email] {
				emailMap[email] = true

				result = append(result, email)
			}
		}
	}
	return result
}

func UserEmail(user *model.User) []string {
	return []string{user.Email}
}
