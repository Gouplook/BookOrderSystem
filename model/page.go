/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2020/12/9 上午10:01
@Description:

*********************************************/
package model

//Page 结构
type Page struct {
	Books       []*Book // 每页查询出来的图书存放的切片
	PageNo      int64   // 当前页
	PageSize    int64   // 每页显示的条数
	TotalPageNo int64   // 总页数，通过计算得到
	TotalRecord int64   // 总记录数，通过查询数据库得到
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	UserName    string
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (p *Page) GetPrePageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	// 没有上一页，表明已经到了首页
	return 1
}

//GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo
}
