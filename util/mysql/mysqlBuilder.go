package mysql

import (
	"database/sql"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"strconv"
	"strings"
)

/**
	自己封装的mysql构造器
 */

type SqlBuilder struct {
	field string
	where []WhereBuilder
	value []interface{}
	order string
	table string
	start int
	limit int
	lastSql string
	alias string
	join []joinBuilder
	lock bool //mysql自带悲观锁
	tx *sql.Tx // 事务管理
}

func NewSqlBuilder() *SqlBuilder{
	//初始化sql构造器
	return &SqlBuilder{
		field:"",
		where:make([]WhereBuilder,0),
		order:"",
		table:"",
		start:-1,
		limit:-1,
		value:make([]interface{},0),
		alias:"",
		lock:false,
		tx:nil,
	}
}

func NewSqlBuilderWithTable(table string) *SqlBuilder{
	//初始化sql构造器
	return &SqlBuilder{
		field:"",
		where:make([]WhereBuilder,0),
		order:"",
		table:table,
		start:-1,
		limit:-1,
		value:make([]interface{},0),
		alias:"",
		lock:false,
		tx:nil,
	}
}

/**
	使用string的切片设置field字段
*/
func (this *SqlBuilder) Select(fields ...string) *SqlBuilder {
	for _,field := range fields {
		if field != "*" {
			this.field += "`" + field + "`" + ","
		} else {
			this.field += field + ","
		}
	}
	this.field = strings.TrimRight(this.field,",")
	return this
}

/**
	设置表别名
 */
func  (this *SqlBuilder) Alias(alias string) *SqlBuilder {
	this.alias = alias
	return this
}

/**
	使用string的切片设置field字段
*/
func (this *SqlBuilder) In(field string,scope []int64) *SqlBuilder {
	length := len(scope)
	conditions := "(" + strings.TrimRight(strings.Repeat("?,",length),",") + ")"
	this.where = append(this.where,WhereBuilder{Field:field,Operation:"in",Expression:conditions})

	for _,item := range scope {
		this.value = append(this.value,item)
	}
	return this
}

/**
	使用string的切片设置field字段
*/
func (this *SqlBuilder) Between(field string,start,end int64) *SqlBuilder {
	this.where = append(this.where,WhereBuilder{Field:field,Operation:"between",Expression:"? AND ?"})
	this.value = append(this.value,start,end)
	return this
}

/**
	设置where条件
*/
func (this *SqlBuilder) Where(field string,operation string,value interface{}) *SqlBuilder{
	this.where = append(this.where,WhereBuilder{field,operation,"?",value})
	return this
}

/**
设置查询表
*/
func (this *SqlBuilder) From(table string) *SqlBuilder {
	this.table = "`" + table + "`"
	return this
}

/**
设置查询条数
1.一个参数 limit
2.两个参数 start,limit
*/
func (this *SqlBuilder) Limit(params ...int) *SqlBuilder {
	//只处理前两个参数
	if len(params) > 2 {
		params = params[0:2]
	}
	//只有一个参数设置limit
	if len(params) == 1 {
		this.start = 0
		this.limit = params[0]
	}
	//如果有两个参数的情况下同时设置start和limit
	if len(params) == 2 {
		this.start = params[0]
		this.limit = params[1]
	}
	return this
}


/**
	设置order排序字段
*/
func (this *SqlBuilder) OrderBy(order string) *SqlBuilder {
	this.order = order
	return this
}

/**
	获取第一条数据
*/
func (this *SqlBuilder) First() (*sql.Row,error) {
	stmt, err := this.Limit(1).query()
	if err != nil {
		return nil,err
	}
	defer stmt.Close()
	defer this.flush()
	return stmt.QueryRow(this.value...), nil
}

/**
	获取所有数据
*/
func (this *SqlBuilder) All() (*sql.Rows,error){
	stmt, err := this.query()
	if err != nil {
		return nil,err
	}
	rows, err := stmt.Query(this.value...)
	if err != nil {
		return nil,err
	}
	defer stmt.Close()
	defer this.flush()
	return rows,nil
}

/**
	新增插入语句
*/
func (this *SqlBuilder) Insert(insertData map[string]interface{}) (int64,error) {
	sql := "INSERT INTO " + this.table
	fieldCollection := "("
	valueCollection := "("
	for field,value := range insertData{
		fieldCollection += "`" + field +"`" + ","
		valueCollection += "?,"
		this.value = append(this.value,value)
	}
	//访问完毕清除条件
	defer func() {
		this.value = nil
	}()

	field := strings.TrimRight(fieldCollection, ",") + ")"
	value := strings.TrimRight(valueCollection, ",") + ")"
	sql += field + "VALUES" + value

	//更新最后一条执行的sql
	this.lastSql = sql
	//是否开启事务
	if this.tx == nil {
		result, err := pool.Exec(sql, this.value...)
		if err != nil {
			return 0,err
		}
		lastId,err := result.LastInsertId()
		if err != nil {
			return 0,err
		}
		return lastId,nil
	} else {
		result, err := this.tx.Exec(sql, this.value...)
		if err != nil {
			return 0,err
		}
		lastId,err := result.LastInsertId()
		if err != nil {
			return 0,err
		}
		return lastId,nil
	}
}

/**
更新语句
*/
func (this *SqlBuilder) Update(updateData map[string]interface{}) (int64,error){

	sql := "UPDATE " + this.table + " SET "
	fieldCollection := ""
	for field,value := range updateData{
		fieldCollection += "`" + field +"`" + " = " + "?" + "AND"
		this.value = append(this.value,value)
	}
	fieldCollection = strings.TrimRight(fieldCollection, "AND") + " "
	sql += fieldCollection
	//如果设置了where条件
	if len(this.where) != 0 {
		condition := " WHERE " + this.parseWhere()
		sql += condition
	}
	//更新最后一条执行的sql
	this.lastSql = sql

	defer this.flush()
	if this.tx == nil {
		result, err := pool.Exec(sql, this.value...)
		if err != nil {
			return 0,err
		}
		lastId,err := result.RowsAffected()
		if err != nil {
			return 0,err
		}
		return lastId,nil
	} else {
		result, err := this.tx.Exec(sql, this.value...)
		if err != nil {
			return 0,err
		}
		lastId,err := result.RowsAffected()
		if err != nil {
			return 0,err
		}
		return lastId,nil
	}
}

/**
	删除语句
*/
func (this *SqlBuilder) Delete() (int64,error){
	if len(this.where) == 0 {
		return 0,errors.New("you can not execute delete without condition")
	}

	sql := "DELETE FROM " + this.table + " WHERE " + this.parseWhere()
	this.lastSql = sql

	defer this.flush()
	if this.tx == nil {
		result, err := pool.Exec(sql, this.value...)
		if err != nil {
			return 0,err
		}
		affectRows,err := result.RowsAffected()
		if err != nil {
			return 0,err
		}
		return affectRows,nil
	} else {
		result, err := this.tx.Exec(sql, this.value...)
		if err != nil {
			return 0,err
		}
		affectRows,err := result.RowsAffected()
		if err != nil {
			return 0,err
		}
		return affectRows,nil
	}
}

func (this *SqlBuilder) query() (*sql.Stmt,error) {

	//判断是否设置了字段,默认全查
	if this.field == "" {
		this.field = "*"
	}

	//如果设置了表别名
	sql := ""
	if this.alias != "" {
		sql = fmt.Sprintf("SELECT %s FROM %s as %s",this.field,this.table,this.alias)
	} else {
		sql = fmt.Sprintf("SELECT %s FROM %s",this.field,this.table)
	}

	//存在关联条件
	if len(this.join) != 0 {
		sql += this.parseJoinCondition()
	}

	//如果设置了where条件
	if len(this.where) != 0 {
		condition := " WHERE " + this.parseWhere()
		sql += condition
	}
	//如果设置了排序字段
	if this.order != "" {
		sql += "ORDER BY " + this.order
	}
	//如果设置了limit
	if this.limit != -1 {
		limit := ""
		if this.start != -1 {
			limit = strconv.FormatInt(int64(this.limit), 10)
		} else {
			limit = strconv.FormatInt(int64(this.start),10) + "," + strconv.FormatInt(int64(this.limit), 10) //构造limit条件
		}
		sql += " limit " + limit
	}
	//如果加锁了
	if this.lock {
		sql += " FOR UPDATE"
	}
	this.lastSql = sql //初始化最后一条执行的sql

	//判断是否开启事务
	if this.tx == nil {
		stmt, err := pool.Prepare(sql)
		if err != nil {
			return nil,errors.New(fmt.Sprintf("error when prepare sql #%v,error is :#%v",sql,err))
		}
		return stmt,nil
	} else {
		stmt, err := this.tx.Prepare(sql)
		if err != nil {
			return nil,errors.New(fmt.Sprintf("error when prepare sql #%v,error is :#%v",sql,err))
		}
		return stmt,nil
	}
}

/**
	统计
 */
func (this *SqlBuilder) Count() (int64,error){
	var counter int64
	sql := "SELECT COUNT(*) FROM " + this.table + this.parseWhere()
	err := pool.QueryRow(sql).Scan(&counter)
	if err != nil {
		return 0,err
	}
	return counter,nil
}

/**
	直接执行查询的sql语句
*/
func (this *SqlBuilder) Query(sql string) *sql.Rows{
	rows, _ := pool.Query(sql)
	return rows
}

/**
	直接执行增删改的sql语句
*/
func (this *SqlBuilder) Exec(sql string) sql.Result{
	result, _ := pool.Exec(sql)
	return result
}

/**
	解析where条件
 */
func (this *SqlBuilder)parseWhere() string {
	condition := ""
	for _,item := range this.where {
		condition += "`" + item.Field + "` " + transferExpresstionIntoOperation(item.Operation) + " " + item.Expression + "AND"
		this.value = append(this.value, item.Value)
	}
	return strings.TrimRight(condition,"AND")
}

/**
	表达式转义
 */
func transferExpresstionIntoOperation(expression string) string {
	expression = strings.ToLower(expression)
	if expression == "eq" {
		return "="
	} else if expression == "lt" {
		return ">"
	} else if expression == "gt" {
		return "<"
	} else if expression == "elt" {
		return ">="
	} else if expression == "egt" {
		return "<="
	} else if expression == "neq" {
		return "<>"
	} else {
		return expression
	}
}

//获取最后一条执行的sql
func (this *SqlBuilder) GetLastSql() string{
	return this.lastSql
}

/**
	左关联
 */
func (this *SqlBuilder) LeftJoin(accessoryTable,subjectTableField,expression,accessoryTableField string) *SqlBuilder{
	return this._join("LEFT JOIN",subjectTableField,expression,accessoryTable,accessoryTableField)
}

/**
	innerJoin
 */
func (this *SqlBuilder) InnerJoin(accessoryTable,subjectTableField,expression,accessoryTableField string) *SqlBuilder{
	return this._join("INNER JOIN",subjectTableField,expression,accessoryTable,accessoryTableField)
}

/**
	右关联
 */
func (this *SqlBuilder) RightJoin(accessoryTable,subjectTableField,expression,accessoryTableField string) *SqlBuilder{
	return this._join("RIGHT JOIN",subjectTableField,expression,accessoryTable,accessoryTableField)
}

/**
	实现关联逻辑
	joinType 限制关联的方向
 */
func (this *SqlBuilder) _join(joinType,subjectTableField,expression,accessoryTable,accessoryTableField string) *SqlBuilder {
	//初始化主表
	this.join = append(this.join,joinBuilder{joinType,subjectTableField,expression,accessoryTable,accessoryTableField})
	return this
}

/**
	解析join语句
 */
func (this *SqlBuilder) parseJoinCondition() string {
	condition := ""
	for _,item := range this.join {
		condition += fmt.Sprintf(" %s `%s` ON %s %s %s",item.direction,item.accessoryTable,item.subjectField,transferExpresstionIntoOperation(item.expression),item.accessoryField)
	}
	return condition
}

func (this *SqlBuilder) Lock() *SqlBuilder{
	this.lock = true
	return this
}

/**
	开启事务
 */
func (this *SqlBuilder) StartTransaction() *SqlBuilder{
	tx, _ := pool.Begin()
	this.tx = tx
	return this
}

/**
	事务提交
 */
func (this *SqlBuilder) Commit() error {
	var err error
	if this.tx != nil {
		err = this.tx.Commit()
		this.flush()
	} else {
		err = errors.New("transaction not start")
	}
	return err
}

/**
	事务回滚
 */
func (this *SqlBuilder) Rollback() error {
	var err error
	if this.tx != nil {
		err = this.tx.Rollback()
		this.flush()
	} else {
		err = errors.New("transaction not start")
	}
	return err
}

/**
	清空查询缓存
 */
func (this *SqlBuilder) flush() {
	this.value = nil
	this.where = nil
	this.lock = false
	this.start = -1
	this.limit = -1
	this.alias = ""
	this.order = ""
	this.join = nil
	this.tx = nil
}