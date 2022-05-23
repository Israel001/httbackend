package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"htt/httbackend/graph/generated"
	"htt/httbackend/graph/model"
)

func (r *mutationResolver) CreateSermon(ctx context.Context, input model.CreateSermonInput) (*model.Sermon, error) {
	return nil, nil
}

func (r *mutationResolver) UpdateSermon(ctx context.Context, id string, input model.UpdateSermonInput) (*model.Sermon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSermon(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateGallery(ctx context.Context, input model.CreateGalleryInput) (*model.Gallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateGallery(ctx context.Context, id string, input model.UpdateGalleryInput) (*model.Gallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteGallery(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateContact(ctx context.Context, input model.CreateContactInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateContact(ctx context.Context, id string, input model.UpdateContactInput) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateSubscription(ctx context.Context, input model.CreateSubscriptionInput) (*model.NewsletterSubscription, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSubscription(ctx context.Context, id string, input model.UpdateSubscriptionInput) (*model.NewsletterSubscription, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteSubscription(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sermons(ctx context.Context) ([]*model.Sermon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Sermon(ctx context.Context, id string) (*model.Sermon, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Galleries(ctx context.Context) ([]*model.Gallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Gallery(ctx context.Context, id string) (*model.Gallery, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Contact(ctx context.Context, id string) (*model.Contact, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Subscriptions(ctx context.Context) ([]*model.NewsletterSubscription, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Subscription(ctx context.Context, id string) (*model.NewsletterSubscription, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
