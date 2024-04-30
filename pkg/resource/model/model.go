package model

import (
	pkgConfig "base/pkg/config"
	"base/pkg/constant"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Database struct {
	Template *gorm.DB
	Oracle   *gorm.DB
}

type HealthCheckResponse struct {
	Name   string
	Type   string
	Host   string
	Status string
	Error  error             `json:"error,omitempty"`
	Detail map[string]string `json:"detail,omitempty"`
}

// Ping performs a health check on the database connection and returns an array of HealthCheckResponse.
// Each time a new database connection is added, make sure to include a call to Ping to check its health.
// It utilizes the provided Database configuration to establish the connection.
func (d *Database) Ping(ctx context.Context, cfg pkgConfig.Config) []HealthCheckResponse {
	var res []HealthCheckResponse

	// template
	// Establish connection with the database template and check its health.
	dbTemplate, errTemplate := d.Template.DB()
	if errTemplate == nil {
		errTemplate = dbTemplate.PingContext(ctx)
	}

	// Generate HealthCheckResponse based on the status of the database template connection.
	res = append(res, generateHealthCheckResponse(cfg.Database.Template, errTemplate))

	// oracle
	dbOracle, errOracle := d.Oracle.DB()
	if errOracle == nil {
		errOracle = dbOracle.PingContext(ctx)
	}

	res = append(res, generateHealthCheckResponse(cfg.Database.Oracle, errOracle))

	return res
}

// generateHealthCheckResponse generates a HealthCheckResponse based on the provided database configuration
// and the error occurred during the health check.
// It determines the status of the connection (Up or Down) and includes any error information.
func generateHealthCheckResponse(db pkgConfig.DatabaseCfg, err error) HealthCheckResponse {
	status := constant.Up
	if err != nil {
		status = constant.Down
	}

	return HealthCheckResponse{
		Name:   db.Name,
		Type:   fmt.Sprintf("%s/%s", constant.ServiceDatabase, db.Driver),
		Host:   db.Host,
		Status: status,
		Error:  err,
	}
}
