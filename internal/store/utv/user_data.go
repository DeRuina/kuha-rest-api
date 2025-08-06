package utv

import (
	"context"
	"database/sql"
	"encoding/json"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type UserDataStore struct {
	db *sql.DB
}

type DeviceInfo struct {
	Connected  bool `json:"connected"`
	DataExists bool `json:"data_exists"`
}

type DeviceStatus struct {
	Garmin DeviceInfo `json:"garmin"`
	Oura   DeviceInfo `json:"oura"`
	Polar  DeviceInfo `json:"polar"`
	Suunto DeviceInfo `json:"suunto"`
}

func (s *UserDataStore) GetUserData(ctx context.Context, userID uuid.UUID) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetUserData(ctx, userID)
}

func (s *UserDataStore) UpsertUserData(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.UpsertUserData(ctx, utvsqlc.UpsertUserDataParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *UserDataStore) DeleteUserData(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.DeleteUserData(ctx, userID)
}

func (s *UserDataStore) GetUserIDBySportID(ctx context.Context, sportID string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetUserIDBySportID(ctx, sportID)
}

func (s *UserDataStore) GetUserDeviceStatus(ctx context.Context, userID uuid.UUID) (DeviceStatus, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := utvsqlc.New(s.db)
	var status DeviceStatus

	if garmin, err := q.GetGarminStatus(ctx, userID); err == nil {
		status.Garmin = DeviceInfo{Connected: garmin.Connected, DataExists: garmin.DataExists}
	} else {
		return status, err
	}

	if oura, err := q.GetOuraStatus(ctx, userID); err == nil {
		status.Oura = DeviceInfo{Connected: oura.Connected, DataExists: oura.DataExists}
	} else {
		return status, err
	}

	if polar, err := q.GetPolarStatus(ctx, userID); err == nil {
		status.Polar = DeviceInfo{Connected: polar.Connected, DataExists: polar.DataExists}
	} else {
		return status, err
	}

	if suunto, err := q.GetSuuntoStatus(ctx, userID); err == nil {
		status.Suunto = DeviceInfo{Connected: suunto.Connected, DataExists: suunto.DataExists}
	} else {
		return status, err
	}

	return status, nil
}
