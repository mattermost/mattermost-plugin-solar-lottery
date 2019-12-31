// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package command

import (
	"fmt"

	"github.com/mattermost/mattermost-plugin-solar-lottery/server/api"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
)

func (c *Command) join(parameters []string) (string, error) {
	var rotationID, rotationName string
	users := ""
	graceShifts := 0
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.StringVarP(&users, flagUsers, flagPUsers, "", "add nother users to rotation.")
	fs.IntVar(&graceShifts, flagGrace, 0, "start with N grace shifts.")
	withRotationFlags(fs, &rotationID, &rotationName)
	err := fs.Parse(parameters)
	if err != nil {
		return c.subUsage(fs), err
	}

	rotationID, err = c.parseRotationFlags(rotationID, rotationName)
	if err != nil {
		return "", err
	}
	rotation, err := c.API.LoadRotation(rotationID)
	if err != nil {
		return "", err
	}

	added, err := c.API.AddRotationUsers(rotation, users, graceShifts)
	if err != nil {
		return "", errors.WithMessagef(err, "failed, %s might have been updated", api.MarkdownUserMap(added))
	}

	return fmt.Sprintf("%s joined rotation %s", api.MarkdownUserMap(added), rotation.Name), nil
}

func (c *Command) leave(parameters []string) (string, error) {
	var rotationID, rotationName string
	users := ""
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.StringVar(&users, flagUsers, "", "remove other users from rotation.")
	withRotationFlags(fs, &rotationID, &rotationName)
	err := fs.Parse(parameters)
	if err != nil {
		return c.subUsage(fs), err
	}

	rotationID, err = c.parseRotationFlags(rotationID, rotationName)
	if err != nil {
		return "", err
	}
	rotation, err := c.API.LoadRotation(rotationID)
	if err != nil {
		return "", err
	}

	deleted, err := c.API.DeleteRotationUsers(rotation, users)
	if err != nil {
		return "", errors.WithMessagef(err, "failed, %s might have been updated", api.MarkdownUserMap(deleted))
	}

	return fmt.Sprintf("%s left rotation %s", api.MarkdownUserMap(deleted), rotation.Name), nil
}
