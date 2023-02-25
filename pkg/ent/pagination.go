// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/vwenkk/load/pkg/ent/group"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type GroupPager struct {
	Order  OrderFunc
	Filter func(*GroupQuery) (*GroupQuery, error)
}

// GroupPaginateOption enables pagination customization.
type GroupPaginateOption func(*GroupPager)

// DefaultGroupOrder is the default ordering of Group.
var DefaultGroupOrder = Desc(group.FieldID)

func newGroupPager(opts []GroupPaginateOption) (*GroupPager, error) {
	pager := &GroupPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultGroupOrder
	}
	return pager, nil
}

func (p *GroupPager) ApplyFilter(query *GroupQuery) (*GroupQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// GroupPageList is Group PageList result.
type GroupPageList struct {
	List        []*Group     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (gr *GroupQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...GroupPaginateOption,
) (*GroupPageList, error) {

	pager, err := newGroupPager(opts)
	if err != nil {
		return nil, err
	}

	if gr, err = pager.ApplyFilter(gr); err != nil {
		return nil, err
	}

	ret := &GroupPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := gr.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		gr = gr.Order(pager.Order)
	} else {
		gr = gr.Order(DefaultGroupOrder)
	}

	gr = gr.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := gr.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
