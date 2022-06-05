package gqlresolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strings"

	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/friendrequest"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/predicate"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/ent/user"
	gqlgenerated "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/generated"
	gqlmodel "github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/graph/model"
	"github.com/mayamika/2022-mai-backend-a-chakiryan/api-server/internal/model/token"
)

func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, search *string, orderBy ent.UserOrder) (*ent.UserConnection, error) {
	var where []predicate.User
	if search != nil {
		for _, s := range strings.Split(*search, " ") {
			where = append(where, user.And(
				user.Or(
					user.NameContainsFold(s),
					user.SurnameContainsFold(s),
				),
			))
		}
	}

	return ent.FromContext(ctx).User.Query().
		Where(where...).
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(&orderBy),
		)
}

func (r *userResolver) Relation(ctx context.Context, obj *ent.User) (gqlmodel.UserRelation, error) {
	t, ok := token.FromContext(ctx)
	if !ok {
		return gqlmodel.UserRelationStranger, nil
	}
	if t.UserID == obj.ID {
		return gqlmodel.UserRelationYou, nil
	}

	client := ent.FromContext(ctx)
	u, err := client.User.Get(ctx, t.UserID)
	if err != nil {
		return gqlmodel.UserRelationStranger, err
	}

	ok, err = u.QueryFriends().
		Where(user.ID(obj.ID)).
		Exist(ctx)
	if err != nil {
		return gqlmodel.UserRelationStranger, err
	}
	if ok {
		return gqlmodel.UserRelationFriend, nil
	}

	ok, err = client.FriendRequest.Query().
		Where(
			friendrequest.HasFromWith(user.ID(t.UserID)),
			friendrequest.HasToWith(user.ID(obj.ID)),
		).
		Exist(ctx)
	if err != nil {
		return gqlmodel.UserRelationStranger, err
	}
	if ok {
		return gqlmodel.UserRelationFriendRequestSent, nil
	}

	return gqlmodel.UserRelationStranger, nil
}

func (r *userResolver) Friends(ctx context.Context, obj *ent.User, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy ent.UserOrder) (*ent.UserConnection, error) {
	return ent.FromContext(ctx).User.QueryFriends(obj).
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(&orderBy),
		)
}

// User returns gqlgenerated.UserResolver implementation.
func (r *Resolver) User() gqlgenerated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
