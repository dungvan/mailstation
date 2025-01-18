package otpservice

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"

	"github.com/dungvan/mailstation/common/memcache"
	v "github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	mc        memcache.Client
	validator = v.New()
)

func Init(dbConn *gorm.DB, cacheClient memcache.Client) {
	db = dbConn
	mc = cacheClient
}

type TemplateModel struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	ServiceName   string         `json:"serviceName" validate:"required"`
	Sender        string         `json:"sender" validate:"required"`
	SubjectRegexp *regexp.Regexp `json:"subjectRegexp" validate:"required"`
	BodyRegexp    *regexp.Regexp `json:"bodyRegexp" validate:"required"`
	ParamReplaced string         `json:"paramReplaced" validate:"required"`
}

func (t *TemplateModel) UnmarshalJSON(data []byte) error {
	type Alias TemplateModel
	aux := &struct {
		SubjectRegexp string `json:"subjectRegexp"`
		BodyRegexp    string `json:"bodyRegexp"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	subjectRegexp, err := regexp.Compile(aux.SubjectRegexp)
	if err != nil {
		return err
	}
	t.SubjectRegexp = subjectRegexp

	bodyRegexp, err := regexp.Compile(aux.BodyRegexp)
	if err != nil {
		return err
	}
	t.BodyRegexp = bodyRegexp

	return nil
}

func (t *TemplateModel) MarshalJSON() ([]byte, error) {
	type Alias TemplateModel
	return json.Marshal(&struct {
		SubjectRegexp string `json:"subjectRegexp"`
		BodyRegexp    string `json:"bodyRegexp"`
		*Alias
	}{
		SubjectRegexp: t.SubjectRegexp.String(),
		BodyRegexp:    t.BodyRegexp.String(),
		Alias:         (*Alias)(t),
	})
}

func (t *TemplateModel) Validate() error {
	return validator.Struct(t)
}

func (t *TemplateModel) ExtractOTP(content string) string {
	return t.BodyRegexp.ReplaceAllString(content, t.ParamReplaced)
}

func GetTemplateBySender(sender string) ([]*TemplateModel, error) {
	var templates []*TemplateModel

	if err := mc.Get(context.Background(), sender, &templates); err == nil && len(templates) > 0 {
		// cache hit
		return templates, nil
	}
	// cache miss
	if err := db.Where("sender = ?", sender).Find(&templates).Error; err != nil {
		return nil, err
	}

	var errs = make([]error, 0)

	for idx := len(templates) - 1; idx >= 0; idx-- {
		if err := templates[idx].Validate(); err != nil {
			errs = append(errs, err)
			if idx == len(templates)-1 {
				templates = templates[:idx]
			} else {
				templates = append(templates[:idx], templates[idx+1:]...)
			}
		}
	}

	if len(templates) > 0 {
		// cache update
		if err := mc.Set(context.Background(), sender, templates); err != nil {
			errs = append(errs, err)
		}
	}

	return templates, errors.Join(errs...)
}
