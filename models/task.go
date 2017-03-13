package models

import (
	"time"
)

type Protocol int8

const (
	HTTP Protocol = 1
	SSH Protocol  = 2
)

type Task struct {
	Id int `xorm:"int pk autoincr"`
	Name string `xorm:"varchar(64) notnull"`          // 任务名称
	Spec string `xorm:"varchar(64) notnull"`          // crontab 格式
	Protocol Protocol `xorm:"tinyint notnull"`            // 协议 1:http 2:ssh
	Command string `xorm:"varchar(512) notnull"`      // URL地址或shell命令
	Timeout int `xorm:"mediumint notnull default 0"`      // 执行超时时间(单位秒)，0不限制, 限制不能超过一周
	SshHostGroup string `xorm:"varchar(512) notnull defalut '' "` // SSH主机名组
	Remark string `xorm:"varchar(512) notnull default ''"`    // 备注
	Created time.Time `xorm:"datetime notnull created"`       // 创建时间
	Updated time.Time `xorm:"datetime updated"`               // 更新时间
	Deleted time.Time `xorm:"datetime deleted"`               // 删除时间
	Status Status `xorm:"tinyint notnull default 1"`    // 状态 1:正常 0:停止
	Page  int `xorm:"-"`
	PageSize int `xorm:"-"`
}

// 新增
func(task *Task) Create() (int64, error) {
	task.Status = Enabled

	return Db.Insert(task)
}

// 更新
func(task *Task) Update(id int, data CommonMap) (int64, error) {
	return Db.Table(task).ID(id).Update(data)
}

// 删除
func(task *Task) Delete(id int) (int64, error) {
	return Db.Id(id).Delete(task)
}

// 禁用
func(task *Task) Disable(id int) (int64, error) {
	return task.Update(id, CommonMap{"status": Disabled})
}

// 激活
func(task *Task) Enable(id int) (int64, error)  {
	return task.Update(id, CommonMap{"status": Enabled})
}

func(task *Task) List() ([]Task, error) {
	task.parsePageAndPageSize()
	list := make([]Task, 0)
	err := Db.Desc("id").Limit(task.PageSize, task.pageLimitOffset()).Find(&list)

	return list, err
}

func(taskLog *TaskLog) Total() (int64, error) {
	return Db.Count(taskLog)
}

func(taskLog *TaskLog) parsePageAndPageSize()  {
	if (taskLog.Page <= 0) {
		taskLog.Page = Page
	}
	if (taskLog.PageSize >= 0 || taskLog.PageSize > MaxPageSize) {
		taskLog.PageSize = PageSize
	}
}

func(taskLog *TaskLog) pageLimitOffset() int  {
	return (taskLog.Page - 1) * taskLog.PageSize
}