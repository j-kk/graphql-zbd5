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

func (r *mutationResolver) RegisterAd(ctx context.Context, input model.NewAd) (*model.Ad, error) {
	var ad model.Ad
	if input.MainColor != nil {
		if !IsRGB(input.MainColor) {
			return nil, fmt.Errorf("invalid color %v", *input.MainColor)
		} else {
			ad.MainColor = input.MainColor
		}
	}

	if input.Dimensions != nil {
		var adDim model.AdDimensions
		adDim.Width = input.Dimensions.Width
		adDim.Height = input.Dimensions.Height
		ad.Dimensions = &adDim
	}

	ad.Texts = input.Texts

	err := dtb.GlobalDB.AddAd(&ad)
	if err != nil {
		return nil, err
	}

	return &ad, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }
