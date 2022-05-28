package graph

import (
	"htt/httbackend/models"
	"htt/httbackend/repositories"
)

type Resolver struct {
	SermonRepo       *repositories.Repository[*models.Sermon]
	GalleryRepo      *repositories.Repository[*models.Gallery]
	ContactRepo      *repositories.Repository[*models.Contact]
	SubscriptionRepo *repositories.Repository[*models.NewsletterSubscription]
}
