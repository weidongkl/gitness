// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package events

import (
	"context"

	"github.com/harness/gitness/events"

	"github.com/rs/zerolog/log"
)

const TagCreatedEvent events.EventType = "tagcreated"

type TagCreatedPayload struct {
	RepoID      int64  `json:"repo_id"`
	PrincipalID int64  `json:"principal_id"`
	Ref         string `json:"ref"`
	SHA         string `json:"sha"`
}

func (r *Reporter) TagCreated(ctx context.Context, payload *TagCreatedPayload) {
	eventID, err := events.ReporterSendEvent(r.innerReporter, ctx, TagCreatedEvent, payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("failed to send tag created event")
		return
	}

	log.Ctx(ctx).Debug().Msgf("reported tag created event with id '%s'", eventID)
}

func (r *Reader) RegisterTagCreated(fn func(context.Context, *events.Event[*TagCreatedPayload]) error) error {
	return events.ReaderRegisterEvent(r.innerReader, TagCreatedEvent, fn)
}

const TagUpdatedEvent events.EventType = "tagupdated"

type TagUpdatedPayload struct {
	RepoID      int64  `json:"repo_id"`
	PrincipalID int64  `json:"principal_id"`
	Ref         string `json:"ref"`
	OldSHA      string `json:"old_sha"`
	NewSHA      string `json:"new_sha"`
	Forced      bool   `json:"forced"`
}

func (r *Reporter) TagUpdated(ctx context.Context, payload *TagUpdatedPayload) {
	eventID, err := events.ReporterSendEvent(r.innerReporter, ctx, TagUpdatedEvent, payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("failed to send tag updated event")
		return
	}

	log.Ctx(ctx).Debug().Msgf("reported tag updated event with id '%s'", eventID)
}

func (r *Reader) RegisterTagUpdated(fn func(context.Context, *events.Event[*TagUpdatedPayload]) error) error {
	return events.ReaderRegisterEvent(r.innerReader, TagUpdatedEvent, fn)
}

const TagDeletedEvent events.EventType = "tagdeleted"

type TagDeletedPayload struct {
	RepoID      int64  `json:"repo_id"`
	PrincipalID int64  `json:"principal_id"`
	Ref         string `json:"ref"`
	SHA         string `json:"sha"`
}

func (r *Reporter) TagDeleted(ctx context.Context, payload *TagDeletedPayload) {
	eventID, err := events.ReporterSendEvent(r.innerReporter, ctx, TagDeletedEvent, payload)
	if err != nil {
		log.Ctx(ctx).Err(err).Msgf("failed to send tag deleted event")
		return
	}

	log.Ctx(ctx).Debug().Msgf("reported tag deleted event with id '%s'", eventID)
}

func (r *Reader) RegisterTagDeleted(fn func(context.Context, *events.Event[*TagDeletedPayload]) error) error {
	return events.ReaderRegisterEvent(r.innerReader, TagDeletedEvent, fn)
}