package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/j-kk/go-graphql/graph/dtb"
	"github.com/j-kk/go-graphql/graph/generated"
	"github.com/j-kk/go-graphql/graph/model"
)

func (r *mutationResolver) RegisterView(ctx context.Context, input model.NewView) (*model.View, error) {
	var adView model.View

	adView.UserID = input.UserID
	adView.AdID = input.AdID

	err := dtb.GlobalDB.RegisterView(&adView)
	if err != nil {
		return nil, err
	}

	return &adView, nil
}

func (r *userResolver) GeoPos(ctx context.Context, obj *model.User) (*model.Position, error) {
	posID := obj.GeoPosID

	pos, err := dtb.GlobalDB.GetPosition(*posID)
	if err != nil {
		return nil, err
	}

	return pos, err
}

func (r *viewResolver) Ad(ctx context.Context, obj *model.View) (*model.Ad, error) {
	adID := obj.AdID

	ad, err := dtb.GlobalDB.GetAd(adID)
	if err != nil {
		return nil, err
	}

	return ad, err
}

func (r *viewResolver) User(ctx context.Context, obj *model.View) (*model.User, error) {
	userID := obj.UserID

	user, err := dtb.GlobalDB.GetUser(userID)
	if err != nil {
		return nil, err
	}

	return user, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

// View returns generated.ViewResolver implementation.
func (r *Resolver) View() generated.ViewResolver { return &viewResolver{r} }

type mutationResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
type viewResolver struct{ *Resolver }
