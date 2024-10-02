// Assignment
// In the Fantasy Quest game menu, we show users the email addresses of their friends. We need that email to be a clickable hyperlink. When they click the hyperlink their default email client should open with a new message ready to send to the address they clicked on.

// Complete the getMailtoLinkForEmail function. It should return a "mailto" hyperlink for the given email.

package main

func getMailtoLinkForEmail(email string) string {
	return "mailto:" + email
}
