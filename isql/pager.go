package isql

import (
	"context"
	"fmt"
	"reflect"

	sq "github.com/Masterminds/squirrel"
	"github.com/lann/builder"
)

// Page Requests
/*
名称	             类型	       分页方式	                说明
count	            Integer	    普通分页/标记分页	         本页记录数
total_count	        Integer	       普通分页	           满足条件的记录总数
page_size	        Integer	    普通分页/标记分页	       实际每页最多记录数
page_number	        Integer	       普通分页	              当前页号
page_count	        Integer	       普通分页	              总分页数量
page_marker_field	String	       标记分页	             分页标记字段
page_marker	        Any	           标记分页	             下一页标记值
*/
// type Page struct {
// 	Count           int         `json:"count"`             // 本页记录数
// 	TotalCount      int64       `json:"total_count"`       // 满足条件的记录总数（一般需扫描全表，不常用）
// 	PageSize        int64       `json:"page_size"`         // 实际每页最多记录数
// 	PageNumber      int64       `json:"page_number"`       // 当前页号（一般需扫描全表，不常用）
// 	PageCount       int64       `json:"page_count"`        // 总分页数量（一般需扫描全表，不常用）
// 	OrderMethod     string      `json:"order_method"`      // 排序方式：asc| desc
// 	PageMarkerField string      `json:"page_marker_field"` // 分页标记字段： id 或者其他可排序字段名称(建议字段约束：唯一&有序)
// 	PageMarker      interface{} `json:"page_marker"`       // 下一页标记值： id 或者其他可排序字段值(建议字段约束：唯一&有序)
// }
type Page struct {
	Count           uint64      `json:"count"`             // 本页记录数
	TotalCount      uint64      `json:"total_count"`       // 满足条件的记录总数（一般需扫描全表，不常用）
	PageSize        uint64      `json:"page_size"`         // 实际每页最多记录数
	PageNumber      uint64      `json:"page_number"`       // 当前页号（一般需扫描全表，不常用）
	PageCount       uint64      `json:"page_count"`        // 总分页数量（一般需扫描全表，不常用）
	OrderMethod     string      `json:"order_method"`      // 排序方式：asc| desc
	PageMarkerField string      `json:"page_marker_field"` // 分页标记字段： id 或者其他可排序字段名称(建议字段约束：唯一&有序)
	PageMarker      interface{} `json:"page_marker"`       // 下一页标记值： id 或者其他可排序字段值(建议字段约束：唯一&有序)

	pageMarkerAsc bool
}

// NewPager New Pager
func NewPager(pageSize, pageNumber uint64, pageMarker interface{}, pageMarkerField string, pageMarkerAsc bool) (p *Page) {
	p = &Page{
		Count:           0,
		TotalCount:      0,
		PageSize:        pageSize,
		PageNumber:      pageNumber,
		PageCount:       0,
		OrderMethod:     "",
		PageMarkerField: pageMarkerField,
		PageMarker:      pageMarker,
		pageMarkerAsc:   pageMarkerAsc,
	}
	if p.PageSize == 0 {
		p.PageSize = 20
	}

	return
}

func (p *Page) start() uint64 {
	if p.PageNumber == 0 {
		return 0
	}
	return (p.PageNumber - 1) * p.PageSize
}

// Result result
type Result struct {
	Data     interface{}
	PageInfo *Page
}

// Count generates a select count statement.
func Count(b sq.SelectBuilder) (int64, error) {
	sb := builder.Delete(b, "Columns").(sq.SelectBuilder)
	sb = sb.Column(sq.Alias(sq.Select("count(*) as count"), "count"))
	var r struct {
		Count int64 `db:"count"`
	}
	err := sb.Scan(&r)
	return r.Count, err
}

// DoQuery Do Query
func (p *Page) DoQuery(ctx context.Context, q sq.SelectBuilder, v interface{}) (ret *Result, err error) {
	query := q

	// totalCount
	var totalCount int64
	if p.PageMarker != nil || p.PageNumber == 0 {
		if p.pageMarkerAsc {
			query = query.Where(sq.Expr(fmt.Sprintf("%s > ?", p.PageMarkerField), p.PageMarker))
		} else {
			query = query.Where(sq.Expr(fmt.Sprintf("%s < ?", p.PageMarkerField), p.PageMarker))
		}
		totalCount = 0
	} else {
		totalCount, err = Count(query)
		if err != nil {
			return
		}
	}

	// result
	err = query.Offset(p.start()).Limit(p.PageSize).Scan(&v)
	if err != nil {
		return
	}

	// count
	var count uint64
	bh, ok := v.(reflect.SliceHeader)
	if ok {
		count = uint64(bh.Len)
	}

	ret = &Result{
		Data: v,
		PageInfo: &Page{
			Count:           count,
			TotalCount:      uint64(totalCount),
			PageSize:        p.PageSize,
			PageNumber:      p.PageNumber,
			PageCount:       0,
			OrderMethod:     "",
			PageMarkerField: "",
			PageMarker:      p.PageMarker,
		},
	}
	ret.PageInfo.PageCount = ret.PageInfo.TotalCount / ret.PageInfo.PageSize

	return
}
