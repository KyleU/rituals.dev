package retro

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"github.com/kyleu/rituals.dev/app/comment"

	"github.com/kyleu/npn/npnservice/user"
	"github.com/kyleu/rituals.dev/app/history"
	"github.com/kyleu/rituals.dev/app/session"

	"github.com/kyleu/rituals.dev/app/permission"

	"github.com/kyleu/rituals.dev/app/action"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	Data   *session.DataServices
	db     *npndatabase.Service
	logger logur.Logger
	svc    util.Service
}

func NewService(actions *action.Service, users *user.Service, comments *comment.Service, db *npndatabase.Service, logger logur.Logger) *Service {
	svc := util.SvcRetro
	logger = logur.WithFields(logger, map[string]interface{}{npncore.KeyService: svc.Key})

	data := session.DataServices{
		Svc:         svc,
		Members:     member.NewService(actions, users, db, logger, svc),
		Comments:    comments,
		Permissions: permission.NewService(actions, db, logger, svc),
		History:     history.NewService(actions, db, logger, svc),
		Actions:     actions,
	}

	return &Service{Data: &data, db: db, logger: logger, svc: svc}
}

func (s *Service) New(title string, userID uuid.UUID, memberName string, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID) (*Session, error) {
	slug, err := s.Data.History.NewSlugFor(nil, title)
	if err != nil {
		return nil, errors.Wrap(err, "error creating retro slug")
	}

	model := NewSession(title, slug, userID, categories, teamID, sprintID)

	q := npndatabase.SQLInsert(s.svc.Key, []string{npncore.KeyID, npncore.KeySlug, npncore.KeyTitle, npncore.WithDBID(util.SvcTeam.Key), npncore.WithDBID(util.SvcSprint.Key), npncore.KeyOwner, npncore.KeyStatus, npncore.Plural(npncore.KeyCategory)}, 1)
	categoriesString := npndatabase.ArrayToString(model.Categories)
	err = s.db.Insert(q, nil, model.ID, slug, model.Title, model.TeamID, model.SprintID, model.Owner, model.Status.String(), categoriesString)
	if err != nil {
		return nil, errors.Wrap(err, "error saving new retro session")
	}

	s.Data.Members.Register(model.ID, userID, memberName, member.RoleOwner)

	s.Data.Actions.Post(s.svc, model.ID, userID, action.ActCreate, nil)
	s.Data.Actions.PostRef(util.SvcSprint, model.SprintID, s.svc, model.ID, userID, action.ActContentAdd)
	s.Data.Actions.PostRef(util.SvcTeam, model.TeamID, s.svc, model.ID, userID, action.ActContentAdd)

	return &model, nil
}

func (s *Service) List(params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retro sessions: %+v", err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByID(id uuid.UUID) *Session {
	dto := &sessionDTO{}
	q := npndatabase.SQLSelectSimple("*", s.svc.Key, npncore.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting retro by id [%v]: %+v", id, err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetBySlug(slug string) *Session {
	var dto = &sessionDTO{}
	q := npndatabase.SQLSelectSimple("*", s.svc.Key, "slug = $1")
	err := s.db.Get(dto, q, nil, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		s.logger.Error(fmt.Sprintf("error getting retro by slug [%v]: %+v", slug, err))
		return nil
	}
	return dto.toSession()
}

func (s *Service) GetByMember(userID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	t := "retro join retro_member m on id = m." + npncore.WithDBID(s.svc.Key)
	q := npndatabase.SQLSelect("retro.*", t, "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, userID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros for user [%v]: %+v", userID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByTeamID(teamID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "team_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, teamID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros for team [%v]: %+v", teamID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetBySprintID(sprintID uuid.UUID, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, sprintID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros for sprint [%v]: %+v", sprintID, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) GetByCreated(d *time.Time, params *npncore.Params) Sessions {
	params = npncore.ParamsWithDefaultOrdering(s.svc.Key, params, npncore.DefaultCreatedOrdering...)
	var dtos []sessionDTO
	q := npndatabase.SQLSelect("*", s.svc.Key, "created between $1 and $2", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, d, d.Add(npncore.HoursInDay*time.Hour))
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving retros created on [%v]: %+v", d, err))
		return nil
	}
	return toSessions(dtos)
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, categories []string, teamID *uuid.UUID, sprintID *uuid.UUID, userID uuid.UUID) error {
	cols := []string{"title", npncore.Plural(npncore.KeyCategory), npncore.WithDBID(util.SvcTeam.Key), npncore.WithDBID(util.SvcSprint.Key)}
	q := npndatabase.SQLUpdate(s.svc.Key, cols, fmt.Sprintf("%v = $%v", npncore.KeyID, len(cols)+1))
	categoriesString := npndatabase.ArrayToString(categories)
	err := s.db.UpdateOne(q, nil, title, categoriesString, teamID, sprintID, sessionID)
	s.Data.Actions.Post(s.svc, sessionID, userID, action.ActUpdate, nil)
	return errors.Wrap(err, "error updating retro session")
}

func toSessions(dtos []sessionDTO) Sessions {
	ret := make(Sessions, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toSession())
	}
	return ret
}
