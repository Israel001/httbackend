// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Contact struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Message   string `json:"message"`
}

type CreateContactInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Message   string `json:"message"`
}

type CreateGalleryInput struct {
	Title string `json:"title"`
	Image string `json:"image"`
	Date  string `json:"date"`
}

type CreateSermonInput struct {
	Title    string `json:"title"`
	Video    string `json:"video"`
	Message  string `json:"message"`
	Date     string `json:"date"`
	Image    string `json:"image"`
	Preacher string `json:"preacher"`
}

type CreateSubscriptionInput struct {
	Email string `json:"email"`
}

type Gallery struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Date  string `json:"date"`
}

type NewsletterSubscription struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Sermon struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Video    string `json:"video"`
	Message  string `json:"message"`
	Date     string `json:"date"`
	Image    string `json:"image"`
	Preacher string `json:"preacher"`
}

type UpdateContactInput struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
	Message   *string `json:"message"`
}

type UpdateGalleryInput struct {
	Title *string `json:"title"`
	Image *string `json:"image"`
	Date  *string `json:"date"`
}

type UpdateSermonInput struct {
	Title    *string `json:"title"`
	Video    *string `json:"video"`
	Message  *string `json:"message"`
	Date     *string `json:"date"`
	Image    *string `json:"image"`
	Preacher *string `json:"preacher"`
}

type UpdateSubscriptionInput struct {
	Email *string `json:"email"`
}
