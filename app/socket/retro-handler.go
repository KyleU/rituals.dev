package socket

import (
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/util"
)

func onRetroSessionSave(s *Service, ch Channel, userID uuid.UUID, param retroSessionSaveParams) error {
	dataSvc := s.retros
	title := util.ServiceTitle(util.SvcRetro, param.Title)

	categories := npndatabase.StringToArray(param.Categories)
	if len(categories) == 0 {
		categories = retro.DefaultCategories
	}

	sprintID := npncore.GetUUIDFromString(param.SprintID)
	teamID := npncore.GetUUIDFromString(param.TeamID)

	curr := dataSvc.GetByID(ch.ID)
	if curr == nil {
		return errors.New("no retro available with id [" + ch.ID.String() + "]")
	}

	sr := s.checkPerms(userID, curr.TeamID, curr.SprintID, ch.Svc, ch.ID)
	if sr != nil {
		return sr
	}

	teamChanged := differentPointerValues(curr.TeamID, teamID)
	sprintChanged := differentPointerValues(curr.SprintID, sprintID)

	msg := "saving retro session [%s] with categories [%s], sprint [%s] and team [%s]"
	s.Logger.Debug(fmt.Sprintf(msg, title, npncore.OxfordComma(categories, "and"), sprintID, teamID))

	err := dataSvc.UpdateSession(ch.ID, title, categories, teamID, sprintID, userID)
	if err != nil {
		return errors.Wrap(err, "error updating retro session")
	}

	if title != curr.Title {
		slug, err := dataSvc.Data.History.UpdateSlug(curr.ID, curr.Slug, curr.Title, title, userID)
		if err != nil {
			return errors.Wrap(err, "error updating retro slug from ["+curr.Slug+"] to ["+slug+"]")
		}
	}

	err = sendRetroSessionUpdate(s, ch)
	if err != nil {
		return errors.Wrap(err, "error sending retro session")
	}

	if teamChanged {
		tm := s.teams.GetByIDPointer(teamID)
		err = sendTeamUpdate(s, ch, curr.TeamID, tm)
		if err != nil {
			return errors.Wrap(err, "error sending team for updated retro session")
		}
	}

	if sprintChanged {
		spr := s.sprints.GetByIDPointer(sprintID)
		err = sendSprintUpdate(s, ch, curr.SprintID, spr)
		if err != nil {
			return errors.Wrap(err, "error sending sprint for updated retro session")
		}
	}

	err = s.updatePerms(ch, userID, teamID, sprintID, dataSvc.Data.Permissions, param.Permissions)
	if err != nil {
		return errors.Wrap(err, "error updating permissions")
	}

	return nil
}

func sendRetroSessionUpdate(s *Service, ch Channel) error {
	sess := s.retros.GetByID(ch.ID)
	if sess == nil {
		return errors.New("cannot load retro session [" + ch.ID.String() + "]")
	}

	err := s.WriteChannel(ch, NewMessage(util.SvcRetro, ServerCmdSessionUpdate, sess))
	return errors.Wrap(err, "error sending retro session")
}
