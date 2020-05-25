package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/auth"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/util"
)

func (s *Service) sendInitial(ch channel, conn *connection, entry *member.Entry, msg Message, sprintID *uuid.UUID, sprintEntry *member.Entry) error {
	err := s.WriteMessage(conn.ID, &msg)
	if err != nil {
		return errors.Wrap(err, "error writing initial estimate message")
	}

	if sprintEntry != nil {
		err = s.sendMemberUpdate(channel{Svc: util.SvcSprint.Key, ID: *sprintID}, sprintEntry, conn.ID)
		if err != nil {
			return errors.Wrap(err, "error writing member update to sprint")
		}
	}

	err = s.sendMemberUpdate(*conn.Channel, entry, conn.ID)
	if err != nil {
		return errors.Wrap(err, "error writing member update")
	}

	err = s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, true)
	return errors.Wrap(err, "error writing online update")
}

func (s *Service) check(
	userID uuid.UUID, auths auth.Records, teamID *uuid.UUID, sprintID *uuid.UUID,
	svc util.Service, modelID uuid.UUID) (permission.Permissions, permission.Errors, error) {
	var err error

	var currTeams []uuid.UUID
	if teamID != nil {
		currTeams, err = s.teams.GetIdsByMember(userID)
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to retrieve current teams")
		}
	}

	var currSprints []uuid.UUID
	if sprintID != nil {
		currSprints, err = s.sprints.GetIdsByMember(userID)
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to retrieve current sprints")
		}
	}

	var tp *permission.Params
	if teamID != nil {
		tm, _ := s.teams.GetByID(*teamID)
		tp = &permission.Params{ID: tm.ID, Slug: tm.Slug, Title: tm.Title, Current: currTeams}
	}

	var sp *permission.Params
	if sprintID != nil {
		spr, _ := s.sprints.GetByID(*sprintID)
		sp = &permission.Params{ID: spr.ID, Slug: spr.Slug, Title: spr.Title, Current: currSprints}
	}

	var permSvc *permission.Service
	switch svc {
	case util.SvcTeam:
		permSvc = s.teams.Permissions
	case util.SvcSprint:
		permSvc = s.sprints.Permissions
	case util.SvcEstimate:
		permSvc = s.estimates.Permissions
	case util.SvcStandup:
		permSvc = s.standups.Permissions
	case util.SvcRetro:
		permSvc = s.retros.Permissions
	}

	if permSvc == nil {
		return nil, nil, errors.New("Invalid service [" + svc.Key + "]")
	}

	perms, e := permSvc.Check(s.auths.Enabled, svc, modelID, auths, tp, sp)
	return perms, e, nil
}

func (s *Service) sendPermErrors(svc util.Service, ch channel, permErrors permission.Errors) error {
	if len(permErrors) > 0 {
		msg := Message{Svc: svc.Key, Cmd: ServerCmdError, Param: "insufficient permissions"}
		return s.WriteChannel(ch, &msg)
	}
	return nil
}
