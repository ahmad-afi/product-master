package helper

import (
	"fmt"
	"reflect"
	"time"
)

type PaginationStruct struct {
	Limit int64 `query:"limit" db:"limit"`
	Page  int64 `query:"page" db:"page"`
}

func (p *PaginationStruct) DefaultPagination() {
	if p.Limit < 1 {
		p.Limit = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit > 1000 {
		p.Limit = 1000
	}

}

type TimeStruct struct {
	CreatedTime time.Time `bson:"createdTime,omitempty" json:"createdTime"`
	UpdatedTime time.Time `bson:"updatedTime,omitempty" json:"updatedTime"`
}

func IsEmptyStruct(structValue any, structOrigin any) bool {
	return reflect.DeepEqual(structValue, structOrigin)
}

type FilteringDate struct {
	StartDate time.Time `query:"startDate" bson:"-"`
	EndDate   time.Time `query:"endDate" bson:"-"`
}

func (f *FilteringDate) ValidDate() error {
	if f.StartDate.After(f.EndDate) {
		return fmt.Errorf("invalid date")
	}
	return nil
}

type WithCountPagination struct {
	WithCount, Pagination bool `bson:"-"`
}

type CountAggregateMongo struct {
	TotalData int64 `bson:"totalData" json:"totalData"`
}
