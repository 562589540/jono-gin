package enum

//用户组 区分后台用户与前台用户

type UserGroups int

const (
	Admin UserGroups = iota + 1
)
