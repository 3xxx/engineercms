// (c) Copyright 2015-2017 JONNALAGADDA Srinivas
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

// CREATE TABLE users_master (
//     id INT NOT NULL AUTO_INCREMENT,
//     first_name VARCHAR(30) NOT NULL,
//     last_name VARCHAR(30) NOT NULL,
//     email VARCHAR(100) NOT NULL,
//     active TINYINT(1) NOT NULL,
//     PRIMARY KEY (id),
//     UNIQUE (email)
// );
type users_master struct {
	Id         int64
	First_name string `orm:"size(30)"`
	Last_name  string `orm:"size(30)"`
	Email      string `orm:"unique;size(100)"`
	Active     bool
}

// CREATE TABLE wf_ac_group_hierarchy (
//     id INT NOT NULL AUTO_INCREMENT,
//     ac_id INT NOT NULL,
//     group_id INT NOT NULL,
//     reports_to INT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (ac_id) REFERENCES wf_access_contexts(id),
//     FOREIGN KEY (group_id) REFERENCES wf_groups_master(id),
//     FOREIGN KEY (reports_to) REFERENCES wf_groups_master(id),
//     UNIQUE (ac_id, group_id)
// );
type wf_ac_group_hierarchy struct {
	Id         int64
	Ac         *wf_access_contexts `orm:"rel(fk);unique"`
	Group      *wf_groups_master   `orm:"rel(fk);unique"`
	Reports_to *wf_groups_master   `orm:"rel(fk);column(reports_to)"`
}

// CREATE TABLE wf_ac_group_roles (
//     id INT NOT NULL AUTO_INCREMENT,
//     ac_id INT NOT NULL,
//     group_id INT NOT NULL,
//     role_id INT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (ac_id) REFERENCES wf_access_contexts(id),
//     FOREIGN KEY (group_id) REFERENCES wf_groups_master(id),
//     FOREIGN KEY (role_id) REFERENCES wf_roles_master(id)
// );
type wf_ac_group_roles struct {
	Id    int64
	Ac    *wf_access_contexts `orm:"rel(fk)"`
	Group *wf_groups_master   `orm:"rel(fk)"`
	Role  *wf_roles_master    `orm:"rel(fk)"`
}

// CREATE OR REPLACE VIEW wf_ac_perms_v AS
// SELECT ac_grs.ac_id, ac_grs.group_id, gu.user_id, ac_grs.role_id, rdas.doctype_id, rdas.docaction_id
// FROM wf_ac_group_roles ac_grs
// JOIN wf_group_users gu ON ac_grs.group_id = gu.group_id
// JOIN wf_role_docactions rdas ON ac_grs.role_id = rdas.role_id;

// CREATE TABLE wf_access_contexts (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(100) NOT NULL,
//     active TINYINT(1) NOT NULL,
//     PRIMARY KEY (id),
//     UNIQUE (name)
// );
type wf_access_contexts struct {
	Id     int64
	Name   string `orm:"unique;size(100)"`
	Active bool
}

// CREATE TABLE wf_docactions_master (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(100) NOT NULL,
//     reconfirm TINYINT(1) NOT NULL,
//     PRIMARY KEY (id),
//     UNIQUE (name)
// );
type wf_docactions_master struct {
	Id        int64  `orm:"pk"`
	Name      string `orm:"unique"`
	Reconfirm bool
}

// CREATE TABLE wf_docevent_application (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     doc_id INT NOT NULL,
//     from_state_id INT NOT NULL,
//     docevent_id INT NOT NULL,
//     to_state_id INT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (from_state_id) REFERENCES wf_docstates_master(id),
//     FOREIGN KEY (docevent_id) REFERENCES wf_docevents(id),
//     FOREIGN KEY (to_state_id) REFERENCES wf_docstates_master(id)
// );
type wf_docevent_application struct {
	Id         int64
	Doctype    *wf_doctypes_master `orm:"rel(fk)"`
	Doc_id     int64
	From_state *wf_docstates_master `orm:"rel(fk);column(from_state_id)"`
	Docevent   *wf_docevents        `orm:"rel(fk)"`
	To_state   *wf_docstates_master `orm:"rel(fk);column(to_state_id)"`
}

// CREATE TABLE wf_docevents (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     doc_id INT NOT NULL,
//     docstate_id INT NOT NULL,
//     docaction_id INT NOT NULL,
//     group_id INT NOT NULL,
//     data TEXT,
//     ctime TIMESTAMP NOT NULL,
//     status ENUM('A', 'P') NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (docstate_id) REFERENCES wf_docstates_master(id),
//     FOREIGN KEY (docaction_id) REFERENCES wf_docactions_master(id),
//     FOREIGN KEY (group_id) REFERENCES wf_groups_master(id)
// );
type wf_docevents struct {
	Id        int64
	Doctype   *wf_doctypes_master `orm:"rel(fk)"`
	DocId     int64
	Docstate  *wf_docstates_master  `orm:"rel(fk)"`
	Docaction *wf_docactions_master `orm:"rel(fk)"`
	GroupId   int64
	Data      string
	Ctime     time.Time
	Status    string `orm:"size(1)"`
}

// CREATE TABLE wf_docstate_transitions (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     from_state_id INT NOT NULL,
//     docaction_id INT NOT NULL,
//     to_state_id INT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (from_state_id) REFERENCES wf_docstates_master(id),
//     FOREIGN KEY (docaction_id) REFERENCES wf_docactions_master(id),
//     FOREIGN KEY (to_state_id) REFERENCES wf_docstates_master(id),
//     UNIQUE (doctype_id, from_state_id, docaction_id, to_state_id)
// );
type wf_docstate_transitions struct {
	Id            int64
	Doctype       *wf_doctypes_master   `orm:"rel(fk);unique"`
	From_state_id *wf_docstates_master  `orm:"rel(fk);unique;column(from_state_id)"`
	Docaction     *wf_docactions_master `orm:"rel(fk);unique"`
	To_state      *wf_docstates_master  `orm:"rel(fk);unique;column(to_state_id)"`
}

// CREATE TABLE wf_docstates_master (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(100) NOT NULL,
//     PRIMARY KEY (id),
//     UNIQUE (name)
// );
// INSERT INTO wf_docstates_master(name)
// VALUES('__RESERVED_CHILD_STATE__');
type wf_docstates_master struct {
	Id   int64  `orm:"pk"`
	Name string `orm:"unique;size(100)"`
}

// CREATE TABLE wf_doctypes_master (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(100) NOT NULL,
//     PRIMARY KEY (id),
//     UNIQUE (name)
// );
type wf_doctypes_master struct {
	Id   int64  `orm:"pk"`
	Name string `orm:"unique;size(100)"`
}

// CREATE TABLE wf_document_blobs (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     doc_id INT NOT NULL,
//     sha1sum CHAR(40) NOT NULL,
//     name TEXT NOT NULL,
//     path TEXT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     UNIQUE (doctype_id, doc_id, sha1sum)
// );
type wf_document_blobs struct {
	Id      int64
	Doctype *wf_doctypes_master `orm:"rel(fk);unique"`
	Doc_id  int64               `orm:"unique"`
	Sha1sum string              `orm:"type(char);size(40);unique"`
	Name    string
	Path    string
}

// CREATE TABLE wf_document_tags (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     doc_id INT NOT NULL,
//     tag VARCHAR(50) NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     UNIQUE (doctype_id, doc_id, tag)
// );
type wf_document_tags struct {
	Id      int64
	Doctype *wf_doctypes_master `orm:"rel(fk);unique"`
	Doc_id  int64               `orm:"unique"`
	Tag     string              `orm:"size(50);unique"`
}

// CREATE TABLE wf_group_users (
//     id INT NOT NULL AUTO_INCREMENT,
//     group_id INT NOT NULL,
//     user_id INT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (group_id) REFERENCES wf_groups_master(id),
//     UNIQUE (group_id, user_id)
// );
type wf_group_users struct {
	Id      int64
	Group   *wf_groups_master `orm:"rel(fk);unique"`
	User_id int64             `orm:"unique"`
}

// CREATE TABLE wf_groups_master (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(100) NOT NULL,
//     group_type ENUM('G', 'S'),
//     PRIMARY KEY (id),
//     UNIQUE (name)
// );
type wf_groups_master struct {
	Id         int64
	Name       string `orm:"size(100);unique"`
	Group_type string `orm:"size(1)"`
}

// CREATE TABLE wf_mailboxes (
//     id INT NOT NULL AUTO_INCREMENT,
//     group_id INT NOT NULL,
//     message_id INT NOT NULL,
//     unread TINYINT(1) NOT NULL,
//     ctime TIMESTAMP NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (group_id) REFERENCES wf_groups_master(id),
//     FOREIGN KEY (message_id) REFERENCES wf_messages(id),
//     UNIQUE (group_id, message_id)
// );
type wf_mailboxes struct {
	Id      int64
	Group   *wf_groups_master `orm:"rel(fk);unique"`
	Message *wf_messages      `orm:"rel(fk);unique"`
	Unread  bool
	Ctime   time.Time
}

// CREATE TABLE wf_messages (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     doc_id INT NOT NULL,
//     docevent_id INT NOT NULL,
//     title VARCHAR(250) NOT NULL,
//     data TEXT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (docevent_id) REFERENCES wf_docevents(id),
//     UNIQUE (doctype_id, doc_id, docevent_id)
// );
type wf_messages struct {
	Id       int64
	Doctype  *wf_doctypes_master `orm:"rel(fk);unique"`
	Doc_id   int64               `orm:"unique"`
	Docevent *wf_docevents       `orm:"rel(fk);unique"`
	Title    string              `orm:"size(250)"`
	Data     string
}

// CREATE TABLE wf_role_docactions (
//     id INT NOT NULL AUTO_INCREMENT,
//     role_id INT NOT NULL,
//     doctype_id INT NOT NULL,
//     docaction_id INT NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (role_id) REFERENCES wf_roles_master(id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (docaction_id) REFERENCES wf_docactions_master(id),
//     UNIQUE (role_id, doctype_id, docaction_id)
// );
type wf_role_docactions struct {
	Id        int64
	Role      *wf_roles_master      `orm:"rel(fk);unique"`
	Doctype   *wf_doctypes_master   `orm:"rel(fk);unique"`
	Docaction *wf_docactions_master `orm:"rel(fk);unique"`
}

// CREATE TABLE wf_roles_master (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(50) NOT NULL,
//     PRIMARY KEY (id),
//     UNIQUE (name)
// );
type wf_roles_master struct {
	Id   int64
	Name string `orm:"size(50);unique"`
}

// INSERT INTO wf_roles_master(name)
// VALUES('SUPER_ADMIN');

// INSERT INTO wf_roles_master(name)
// VALUES('ADMIN');
// CREATE OR REPLACE VIEW wf_users_master AS
// SELECT id, first_name, last_name, email, active
// FROM users_master;

// CREATE TABLE wf_workflow_nodes (
//     id INT NOT NULL AUTO_INCREMENT,
//     doctype_id INT NOT NULL,
//     docstate_id INT NOT NULL,
//     ac_id INT,
//     workflow_id INT NOT NULL,
//     name VARCHAR(100) NOT NULL,
//     type ENUM('begin', 'end', 'linear', 'branch', 'joinany', 'joinall') NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (docstate_id) REFERENCES wf_docstates_master(id),
//     FOREIGN KEY (ac_id) REFERENCES wf_access_contexts(id),
//     FOREIGN KEY (workflow_id) REFERENCES wf_workflows(id),
//     UNIQUE (doctype_id, docstate_id),
//     UNIQUE (workflow_id, name)
// );
type wf_workflow_nodes struct {
	Id       int64
	Doctype  *wf_doctypes_master  `orm:"rel(fk);unique"`
	Docstate *wf_docstates_master `orm:"rel(fk);unique"`
	Ac       *wf_access_contexts  `orm:"rel(fk)"`
	Workflow *wf_workflows        `orm:"rel(fk);unique"`
	Name     string               `orm:"size(100);unique"`
	NodeType string               `orm:"size(7);column(type)"`
}

// CREATE TABLE wf_workflows (
//     id INT NOT NULL AUTO_INCREMENT,
//     name VARCHAR(100) NOT NULL,
//     doctype_id INT NOT NULL,
//     docstate_id INT NOT NULL,
//     active TINYINT(1) NOT NULL,
//     PRIMARY KEY (id),
//     FOREIGN KEY (doctype_id) REFERENCES wf_doctypes_master(id),
//     FOREIGN KEY (docstate_id) REFERENCES wf_docstates_master(id),
//     UNIQUE (name),
//     UNIQUE (doctype_id)
// );
type wf_workflows struct {
	Id       int64
	Name     string               `orm:"size(100);unique"`
	Doctype  *wf_doctypes_master  `orm:"rel(fk);unique"`
	Docstate *wf_docstates_master `orm:"rel(fk)"`
	Active   bool
}

func init() {
	orm.RegisterModel(new(users_master), new(wf_ac_group_hierarchy), new(wf_ac_group_roles))
	orm.RegisterModel(new(wf_access_contexts), new(wf_docactions_master), new(wf_docevent_application))
	orm.RegisterModel(new(wf_docevents), new(wf_docstate_transitions), new(wf_docstates_master))
	orm.RegisterModel(new(wf_doctypes_master), new(wf_document_blobs), new(wf_document_tags))
	orm.RegisterModel(new(wf_group_users), new(wf_groups_master), new(wf_mailboxes))
	orm.RegisterModel(new(wf_messages), new(wf_role_docactions), new(wf_roles_master))
	orm.RegisterModel(new(wf_workflow_nodes), new(wf_workflows))
}

func InitFlow() {
	sql := fmt.Sprintf("CREATE VIEW wf_ac_perms_v AS " +
		"SELECT wf_ac_group_roles.ac_id, wf_ac_group_roles.group_id, wf_group_users.user_id, wf_ac_group_roles.role_id, wf_role_docactions.doctype_id, wf_role_docactions.docaction_id " +
		"FROM wf_ac_group_roles " +
		"JOIN wf_group_users ON wf_ac_group_roles.group_id = wf_group_users.group_id " +
		"JOIN wf_role_docactions ON wf_ac_group_roles.role_id = wf_role_docactions.role_id;")

	sql2 := fmt.Sprintf("INSERT INTO wf_docstates_master(name) VALUES('__RESERVED_CHILD_STATE__');")

	sql3 := fmt.Sprintf("INSERT INTO wf_roles_master(name) VALUES('SUPER_ADMIN');")

	sql4 := fmt.Sprintf("INSERT INTO wf_roles_master(name) VALUES('ADMIN');")

	sql5 := fmt.Sprintf("CREATE VIEW wf_users_master AS SELECT id, first_name, last_name, email, active FROM users_master;")

	o := orm.NewOrm()
	res, err := o.Raw(sql).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		o.Rollback() // beego.Info("插入t_studentInfo表出错,事务回滚")
	}
	res, err = o.Raw(sql2).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		o.Rollback() // beego.Info("插入t_studentInfo表出错,事务回滚")
	}
	res, err = o.Raw(sql3).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		o.Rollback() // beego.Info("插入t_studentInfo表出错,事务回滚")
	}
	res, err = o.Raw(sql4).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		o.Rollback() // beego.Info("插入t_studentInfo表出错,事务回滚")
	}
	res, err = o.Raw(sql5).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	} else {
		o.Rollback() // beego.Info("插入t_studentInfo表出错,事务回滚")
	}
}
