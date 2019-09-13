package mysql
type joinBuilder struct {
	direction string //关联方向
	subjectField string //主表关联字段
	expression string //关联表达式
	accessoryTable string //关联的从表
	accessoryField string //关联的从表的字段
}