// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package command

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"

	"github.com/mattermost/mattermost-plugin-solar-lottery/server/api"
)

func (c *Command) qualifyUsers(parameters []string) (string, error) {
	var usernames, skillName string
	var level api.Level
	fs := pflag.NewFlagSet("", pflag.ContinueOnError)
	withSkillFlags(fs, &skillName, &level)
	fs.StringVarP(&usernames, flagUsers, flagPUsers, "", "users to qualify")
	err := fs.Parse(parameters)
	if err != nil {
		return c.flagUsage(fs), err
	}

	if skillName == "" || level == 0 {
		return c.flagUsage(fs), errors.New("must provide --level and --skill values")
	}

	err = c.API.Qualify(usernames, skillName, level)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Qualified %s as %s", usernames, api.MarkdownSkillLevel(skillName, level)), nil
}
