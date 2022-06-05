// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodel

import (
	"fmt"
	"io"
	"strconv"
)

type UserRelation string

const (
	UserRelationStranger          UserRelation = "STRANGER"
	UserRelationYou               UserRelation = "YOU"
	UserRelationFriendRequestSent UserRelation = "FRIEND_REQUEST_SENT"
	UserRelationFriend            UserRelation = "FRIEND"
)

var AllUserRelation = []UserRelation{
	UserRelationStranger,
	UserRelationYou,
	UserRelationFriendRequestSent,
	UserRelationFriend,
}

func (e UserRelation) IsValid() bool {
	switch e {
	case UserRelationStranger, UserRelationYou, UserRelationFriendRequestSent, UserRelationFriend:
		return true
	}
	return false
}

func (e UserRelation) String() string {
	return string(e)
}

func (e *UserRelation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserRelation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserRelation", str)
	}
	return nil
}

func (e UserRelation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}