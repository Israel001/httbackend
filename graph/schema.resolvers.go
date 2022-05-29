package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"htt/httbackend/graph/generated"
	"htt/httbackend/graph/model"
	"htt/httbackend/models"
)

func (r *mutationResolver) CreateSermon(ctx context.Context, input model.CreateSermonInput) (*model.Sermon, error) {
	s := models.Sermon{}
	s.Title = input.Title

	existingSermon := models.Sermon{}
	r.SermonRepo.DB.First(&existingSermon, &s)
	if existingSermon.ID.Valid {
		return nil, fmt.Errorf("a sermon with the same title: '%s' already exists", existingSermon.Title)
	}

	response := model.Sermon{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	tx := r.SermonRepo.DB.Create(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while creating a new sermon with title: '%s'", response.Title)
	}

	return &response, nil
}

func (r *mutationResolver) UpdateSermon(ctx context.Context, id int, input model.UpdateSermonInput) (*model.Sermon, error) {
	existingSermon := models.Sermon{}
	r.SermonRepo.DB.First(&existingSermon, id)
	if !existingSermon.ID.Valid {
		return nil, fmt.Errorf("a sermon with the ID: '%v' not found", id)
	}

	response := model.Sermon{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	if input.Title != nil {
		sermonCheck := models.Sermon{}
		r.SermonRepo.DB.First(&sermonCheck, "title=? and id<>?", response.Title, id)
		if sermonCheck.ID.Valid {
			return nil, fmt.Errorf("a sermon with the title: '%v' already exists", response.Title)
		}
	}

	response.ID = id

	tx := r.SermonRepo.DB.Save(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while updating sermon with title: '%v'", response.Title)
	}

	return &response, nil
}

func (r *mutationResolver) DeleteSermon(ctx context.Context, id int) (*model.Sermon, error) {
	response := model.Sermon{}
	result := r.SermonRepo.DB.Delete(&response, id)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, fmt.Errorf("an error has occurred while deleting sermon with ID: '%v'", id)
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("a sermon with the ID: '%v' not found", id)
	}

	return &response, nil
}

func (r *mutationResolver) CreateGallery(ctx context.Context, input model.CreateGalleryInput) (*model.Gallery, error) {
	s := models.Gallery{}
	s.Title = input.Title

	existingGallery := models.Gallery{}
	r.GalleryRepo.DB.First(&existingGallery, &s)
	if existingGallery.ID.Valid {
		return nil, fmt.Errorf("a gallery with the same title: '%s' already exists", existingGallery.Title)
	}

	response := model.Gallery{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	tx := r.GalleryRepo.DB.Create(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while creating a new gallery with title: '%s'", response.Title)
	}

	return &response, nil
}

func (r *mutationResolver) UpdateGallery(ctx context.Context, id int, input model.UpdateGalleryInput) (*model.Gallery, error) {
	existingGallery := models.Gallery{}
	r.GalleryRepo.DB.First(&existingGallery, id)
	if !existingGallery.ID.Valid {
		return nil, fmt.Errorf("a gallery with the ID: '%v' not found", id)
	}

	response := model.Gallery{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	if input.Title != nil {
		galleryCheck := models.Gallery{}
		r.GalleryRepo.DB.First(&galleryCheck, "title=? and id<>?", response.Title, id)
		if galleryCheck.ID.Valid {
			return nil, fmt.Errorf("a gallery with the title: '%v' already exists", response.Title)
		}
	}

	response.ID = id

	tx := r.GalleryRepo.DB.Save(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while updating gallery with title: '%v'", response.Title)
	}

	return &response, nil
}

func (r *mutationResolver) DeleteGallery(ctx context.Context, id int) (*model.Gallery, error) {
	response := model.Gallery{}
	result := r.GalleryRepo.DB.Delete(&response, id)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, fmt.Errorf("an error has occurred while deleting gallery with ID: '%v'", id)
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("a gallery with the ID: '%v' not found", id)
	}

	return &response, nil
}

func (r *mutationResolver) CreateContact(ctx context.Context, input model.CreateContactInput) (*model.Contact, error) {
	response := model.Contact{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	tx := r.ContactRepo.DB.Create(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while creating a new contact with email: '%s'", response.Email)
	}

	return &response, nil
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id int, input model.UpdateContactInput) (*model.Contact, error) {
	existingContact := models.Contact{}
	r.ContactRepo.DB.First(&existingContact, id)
	if !existingContact.ID.Valid {
		return nil, fmt.Errorf("a contact with the ID: '%v' not found", id)
	}

	response := model.Contact{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	response.ID = id

	tx := r.ContactRepo.DB.Save(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while updating contact with email: '%v'", response.Email)
	}

	return &response, nil
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id int) (*model.Contact, error) {
	response := model.Contact{}
	result := r.ContactRepo.DB.Delete(&response, id)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, fmt.Errorf("an error has occurred while deleting contact with ID: '%v'", id)
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("a contact with the ID: '%v' not found", id)
	}

	return &response, nil
}

func (r *mutationResolver) CreateSubscription(ctx context.Context, input model.CreateSubscriptionInput) (*model.NewsletterSubscription, error) {
	s := models.NewsletterSubscription{}
	s.Email = input.Email

	existingSubscription := models.NewsletterSubscription{}
	r.SubscriptionRepo.DB.First(&existingSubscription, &s)
	if existingSubscription.ID.Valid {
		return nil, fmt.Errorf("a subscription with the same email: '%s' already exists", existingSubscription.Email)
	}

	response := model.NewsletterSubscription{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	tx := r.SubscriptionRepo.DB.Create(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while creating a new subscription with email: '%s'", response.Email)
	}

	return &response, nil
}

func (r *mutationResolver) UpdateSubscription(ctx context.Context, id int, input model.UpdateSubscriptionInput) (*model.NewsletterSubscription, error) {
	existingSubscription := models.NewsletterSubscription{}
	r.SubscriptionRepo.DB.First(&existingSubscription, id)
	if !existingSubscription.ID.Valid {
		return nil, fmt.Errorf("a subscription with the ID: '%v' not found", id)
	}

	response := model.NewsletterSubscription{}
	res, _ := json.Marshal(input)
	json.Unmarshal(res, &response)

	if input.Email != nil {
		subscriptionCheck := models.NewsletterSubscription{}
		r.SubscriptionRepo.DB.First(&subscriptionCheck, "email=? and id<>?", response.Email, id)
		if subscriptionCheck.ID.Valid {
			return nil, fmt.Errorf("a subscription with the email: '%v' already exists", response.Email)
		}
	}

	response.ID = id

	tx := r.SubscriptionRepo.DB.Save(&response)
	if tx.Error != nil {
		fmt.Println(tx.Error.Error())
		return nil, fmt.Errorf("an error has occurred while updating subscription with email: '%v'", response.Email)
	}

	return &response, nil
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, id int) (*model.NewsletterSubscription, error) {
	response := model.NewsletterSubscription{}
	result := r.SubscriptionRepo.DB.Delete(&response, id)
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return nil, fmt.Errorf("an error has occurred while deleting subscription with ID: '%v'", id)
	} else if result.RowsAffected < 1 {
		return nil, fmt.Errorf("a subscription with the ID: '%v' not found", id)
	}

	return &response, nil
}

func (r *queryResolver) Sermons(ctx context.Context) ([]*model.Sermon, error) {
	response := []model.Sermon{}
	sermonList := []*model.Sermon{}

	r.SermonRepo.DB.Find(&response)

	for _, sermon := range response {
		sermonList = append(sermonList, &sermon)
	}

	return sermonList, nil
}

func (r *queryResolver) Sermon(ctx context.Context, id int) (*model.Sermon, error) {
	existingSermon := models.Sermon{}
	r.SermonRepo.DB.First(&existingSermon, id)
	if !existingSermon.ID.Valid {
		return nil, fmt.Errorf("a sermon with the ID: '%v' not found", id)
	}
	response := model.Sermon{}
	res, _ := json.Marshal(existingSermon)
	json.Unmarshal(res, &response)
	response.ID = int(existingSermon.ID.Int64)
	return &response, nil
}

func (r *queryResolver) Galleries(ctx context.Context) ([]*model.Gallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Gallery(ctx context.Context, id int) (*model.Gallery, error) {
	existingGallery := models.Gallery{}
	r.GalleryRepo.DB.First(&existingGallery, id)
	if !existingGallery.ID.Valid {
		return nil, fmt.Errorf("a gallery with the ID: '%v' not found", id)
	}
	response := model.Gallery{}
	res, _ := json.Marshal(existingGallery)
	json.Unmarshal(res, &response)
	response.ID = int(existingGallery.ID.Int64)
	return &response, nil
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contact(ctx context.Context, id int) (*model.Contact, error) {
	existingContact := models.Contact{}
	r.ContactRepo.DB.First(&existingContact, id)
	if !existingContact.ID.Valid {
		return nil, fmt.Errorf("a contact with the ID: '%v' not found", id)
	}
	response := model.Contact{}
	res, _ := json.Marshal(existingContact)
	json.Unmarshal(res, &response)
	response.ID = int(existingContact.ID.Int64)
	return &response, nil
}

func (r *queryResolver) Subscriptions(ctx context.Context) ([]*model.NewsletterSubscription, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Subscription(ctx context.Context, id int) (*model.NewsletterSubscription, error) {
	existingSubscription := models.NewsletterSubscription{}
	r.SubscriptionRepo.DB.First(&existingSubscription, id)
	if !existingSubscription.ID.Valid {
		return nil, fmt.Errorf("a subscription with the ID: '%v' not found", id)
	}
	response := model.NewsletterSubscription{}
	res, _ := json.Marshal(existingSubscription)
	json.Unmarshal(res, &response)
	response.ID = int(existingSubscription.ID.Int64)
	return &response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
