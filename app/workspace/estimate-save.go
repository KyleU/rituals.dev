package workspace

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/estimate/ehistory"
	"github.com/kyleu/rituals/app/estimate/emember"
	"github.com/kyleu/rituals/app/util"
)

var DefaultEstimateChoices = []string{"0", "1", "2", "3", "5", "8", "13", "100"}

func (s *Service) CreateEstimate(
	ctx context.Context, id uuid.UUID, title string, user uuid.UUID, name string, picture string, teamID *uuid.UUID, sprintID *uuid.UUID, logger util.Logger,
) (*estimate.Estimate, *emember.EstimateMember, error) {
	slug := s.e.Slugify(ctx, id, title, "", s.eh, nil, logger)
	model := &estimate.Estimate{
		ID: id, Slug: slug, Title: title, Status: enum.SessionStatusNew, TeamID: teamID, SprintID: sprintID, Choices: DefaultEstimateChoices, Created: time.Now(),
	}
	err := s.e.Create(ctx, nil, logger, model)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save estimate")
	}
	member, err := s.em.Register(ctx, model.ID, user, name, picture, enum.MemberStatusOwner, nil, s.a, s.send, s.us, logger)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to save estimate owner")
	}
	if teamID != nil {
		msg := map[string]any{"type": enum.ModelServiceEstimate, "id": model.ID, "title": model.Title, "path": model.PublicWebPath(), "icon": model.IconSafe()}
		err = s.send(enum.ModelServiceTeam, *teamID, action.ActChildAdd, msg, &user, logger)
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to notify team")
		}
	}
	if sprintID != nil {
		msg := map[string]any{"type": enum.ModelServiceEstimate, "id": model.ID, "title": model.Title, "path": model.PublicWebPath(), "icon": model.IconSafe()}
		err = s.send(enum.ModelServiceSprint, *sprintID, action.ActChildAdd, msg, &user, logger)
		if err != nil {
			return nil, nil, errors.Wrap(err, "unable to notify sprint")
		}
	}
	return model, member, nil
}

func (s *Service) SaveEstimate(ctx context.Context, e *estimate.Estimate, user uuid.UUID, tx *sqlx.Tx, logger util.Logger) (*estimate.Estimate, error) {
	curr, err := s.e.Get(ctx, tx, e.ID, logger)
	if err != nil {
		return nil, err
	}
	if curr == nil {
		return nil, errors.Errorf("no existing estimate found with id [%s]", e.ID.String())
	}

	if curr.Slug != e.Slug {
		_ = s.eh.DeleteWhere(ctx, tx, "estimate_id = $1 and slug = $2", -1, logger, e.ID, curr.Slug)
		hist := &ehistory.EstimateHistory{Slug: curr.Slug, EstimateID: e.ID, EstimateName: e.TitleString(), Created: time.Now()}
		err = s.eh.Create(ctx, tx, logger, hist)
		if err != nil {
			return nil, err
		}
	}

	if len(e.Choices) == 0 {
		e.Choices = DefaultEstimateChoices
	}

	err = s.e.Save(ctx, tx, logger, e)
	if err != nil {
		return nil, err
	}

	act := action.New(enum.ModelServiceEstimate, e.ID, user, "update", util.ValueMap{}, "")
	err = s.a.Save(ctx, tx, logger, act)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) DeleteEstimate(ctx context.Context, fe *FullEstimate, logger util.Logger) error {
	tx, err := s.db.StartTransaction(logger)
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	for _, v := range fe.Votes {
		err = s.v.Delete(ctx, tx, v.StoryID, v.UserID, logger)
		if err != nil {
			return err
		}
	}

	for _, st := range fe.Stories {
		err = s.st.Delete(ctx, tx, st.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, h := range fe.Histories {
		err = s.eh.Delete(ctx, tx, h.Slug, logger)
		if err != nil {
			return err
		}
	}

	for _, c := range fe.Comments {
		err = s.c.Delete(ctx, tx, c.ID, logger)
		if err != nil {
			return err
		}
	}

	for _, p := range fe.Permissions {
		err = s.ep.Delete(ctx, tx, p.EstimateID, p.Key, p.Value, logger)
		if err != nil {
			return err
		}
	}

	for _, m := range fe.Members {
		err = s.em.Delete(ctx, tx, m.EstimateID, m.UserID, logger)
		if err != nil {
			return err
		}
	}

	err = s.e.Delete(ctx, tx, fe.Estimate.ID, logger)
	if err != nil {
		return err
	}

	if fe.Estimate.TeamID != nil {
		msg := util.ValueMap{"type": enum.ModelServiceEstimate, "id": fe.Estimate.ID}
		_ = s.send(enum.ModelServiceTeam, *fe.Estimate.TeamID, action.ActChildRemove, msg, &fe.Self.UserID, logger)
	}
	if fe.Estimate.SprintID != nil {
		msg := util.ValueMap{"type": enum.ModelServiceEstimate, "id": fe.Estimate.ID}
		_ = s.send(enum.ModelServiceSprint, *fe.Estimate.SprintID, action.ActChildRemove, msg, &fe.Self.UserID, logger)
	}
	err = s.send(enum.ModelServiceEstimate, fe.Estimate.ID, action.ActReset, nil, nil, logger)
	if err != nil {
		return err
	}

	return tx.Commit()
}
