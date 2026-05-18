package helper

import "fmt"

func WaitingApprovalEmail(documentName, url string) string {

	return fmt.Sprintf(`
		<h2>Document "%s" is waiting for your approval</h2>
		<p>Please click the link below to review and approve/reject the document:</p>
		<a href="%s">%s</a>
	`, documentName, url, url)

}

func ApprovedEmail(documentName, url string) string {

	return fmt.Sprintf(`
	<h2>Your document "%s" has been approved</h2>
	<p>You can view the document details at the link below:</p>
	<a href="%s">%s</a>
`, documentName, url, url)

}

func RejectedEmail(documentName, url, reason string) string {

	return fmt.Sprintf(`
	<h2>Your document "%s" has been rejected</h2>
	<p>Reason for rejection: %s</p>
	<p>You can view the document details at the link below:</p>
	<a href="%s">%s</a>
`, documentName, reason, url, url)

}

func ExpirationReminderEmail(documentName, url, expiredDate string) string {

	return fmt.Sprintf(`
	<h2>Your document "%s" is expiring soon</h2>
	<p>The document is set to expire on %s</p>
	<p>You can view the document details at the link below:</p>
	<a href="%s">%s</a>
`, documentName, expiredDate, url, url)

}
