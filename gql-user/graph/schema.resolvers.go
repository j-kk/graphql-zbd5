package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/j-kk/go-graphql/graph/dtb"

	"github.com/j-kk/go-graphql/graph/generated"
	"github.com/j-kk/go-graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user model.User
	var position model.Position

	switch *input.Gender {
	case "M", "K":
		*user.Gender = model.Gender(*input.Gender)
	case "O", "Other":
		*user.Gender = model.GenderOther
	default:
		return nil, fmt.Errorf("unknown gender string format")
	}

	user.BirthYear = input.BirthYear
	user.Income = input.Income

	if input.GeoPos != nil {
		position.Heigth = input.GeoPos.Heigth
		position.Width = input.GeoPos.Width
		user.GeoPos = &position
	}

	user.Interests = input.Interests

	err := dtb.GlobalDB.AddUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
