// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package sl

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost-plugin-solar-lottery/server/utils/types"
)

type UserMap map[string]*User

func (m UserMap) Clone(deep bool) UserMap {
	users := UserMap{}
	for id, user := range m {
		if deep {
			users[id] = user.Clone()
		} else {
			users[id] = user
		}
	}
	return users
}

func (sl *sl) ExpandUserMap(users UserMap) error {
	for _, user := range users {
		err := sl.ExpandUser(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m UserMap) MarkdownWithSkills() string {
	out := []string{}
	for _, user := range m {
		out = append(out, fmt.Sprintf("%s %s", user.Markdown(), user.MarkdownSkills()))
	}
	return strings.Join(out, ", ")
}

func (m UserMap) Markdown() string {
	out := []string{}
	for _, user := range m {
		out = append(out, user.Markdown())
	}
	return strings.Join(out, ", ")
}

func (m UserMap) String() string {
	out := []string{}
	for _, user := range m {
		out = append(out, user.String())
	}
	return strings.Join(out, ", ")
}

func (m UserMap) IDs() *types.Set {
	set := types.NewSet()
	for id := range m {
		set.Add(id)
	}
	return set
}

// Sorted returns all users, sorted by ID. It is used in testing, so it returns
// []User rather than a []*User to make it easier to compare with expected
// results.
func (m UserMap) Sorted() []User {
	out := []User{}
	for _, id := range m.IDs().Sorted() {
		out = append(out, *m[id])
	}
	return out
}

func (users UserMap) Qualified(need *Need) UserMap {
	qualified := UserMap{}
	for id, user := range users {
		if user.IsQualified(need) {
			qualified[id] = user
		}
	}
	return qualified
}