package controllers

import "secondmemory/models"

type (
	RecordResource struct {
		Data models.Record `json:"data"`
	}
	RecordsResource struct {
		Data []models.Record `json:"data"`
	}
)
