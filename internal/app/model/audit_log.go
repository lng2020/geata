package model

import (
	"time"

	"xorm.io/xorm"
)

type AuditLog struct {
	ID         int64     `xorm:"pk autoincr 'id'" json:"id"`
	Message    string    `xorm:"'message'" json:"message"`
	CreateTime time.Time `xorm:"'create_time' created" json:"createTime"`
}

func (a *AuditLog) TableName() string {
	return "audit_log"
}

func CreateAuditLog(engine *xorm.Engine, log *AuditLog) error {
	_, err := engine.Insert(log)
	return err
}

func GetAuditLogByID(engine *xorm.Engine, id int64) (*AuditLog, error) {
	log := new(AuditLog)
	has, err := engine.ID(id).Get(log)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return log, nil
}

func GetAllAuditLogs(engine *xorm.Engine) ([]*AuditLog, error) {
	var logs []*AuditLog
	err := engine.Find(&logs)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
