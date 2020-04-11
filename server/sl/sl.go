// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package sl

import (
	"github.com/mattermost/mattermost-plugin-solar-lottery/server/config"
	"github.com/mattermost/mattermost-plugin-solar-lottery/server/utils/bot"
	"github.com/mattermost/mattermost-plugin-solar-lottery/server/utils/md"
	"github.com/mattermost/mattermost-plugin-solar-lottery/server/utils/types"
)

const (
	ctxActingUserID   = "ActingUserID"
	ctxActingUsername = "ActingUsername"
	ctxAPI            = "API"
	ctxInput          = "Input"
	ctxFill           = "Fill"
	ctxForce          = "Force"
	ctxInterval       = "Interval"
	ctxRotationID     = "RotationID"
	ctxSkill          = "Skill"
	ctxSkillLevel     = "SkillLevel"
	ctxSourceName     = "SourceName"
	ctxStarting       = "Starting"
	ctxTaskID         = "TaskID"
	ctxUnavailable    = "Unavailable"
	ctxUserIDs        = "UserIDs"
	ctxUsernames      = "Usernames"
	ctxUsers          = "Users"
)

type SL interface {
	RotationService
	SkillService
	UserService
	TaskService

	PluginAPI
	bot.Logger

	ActingUser() (*User, error)
	Config() *config.Config

	LoadUsers(mattermostUserIDs *types.IDSet) (*Users, error)
	LoadMattermostUserByUsername(username string) (*User, error)
}

type sl struct {
	*Service
	bot.Logger

	conf *config.Config

	// set by Service.ActingAs.
	actingMattermostUserID types.ID

	// set by withActingUser or withActingUserExpanded.
	actingUser *User

	// Stack of loggers
	loggers []bot.Logger
}

func (sl *sl) Config() *config.Config {
	return sl.conf
}

func (sl *sl) LogAPI(msg md.Markdowner) {
	sl.Infof("%s: %s", sl.actingUser.Markdown(), msg.Markdown())
}