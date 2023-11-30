/*
 * TencentBlueKing is pleased to support the open source community by making 蓝鲸智云-权限中心(BlueKing-IAM) available.
 * Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

package dao

import (
	"testing"

	"iam/pkg/database"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func Test_subjectTemplateGroupManager_BulkCreateWithTx(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`^INSERT INTO subject_template_group`).WithArgs(
			int64(2), int64(1), int64(1), int64(3),
		).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		relations := []SubjectTemplateGroup{{
			SubjectPK:  2,
			TemplateID: 1,
			GroupPK:    1,
			ExpiredAt:  3,
		}}

		tx, err := db.Beginx()
		assert.NoError(t, err)

		manager := &subjectTemplateGroupManager{DB: db}
		err = manager.BulkCreateWithTx(tx, relations)

		tx.Commit()

		assert.NoError(t, err)
	})
}

func Test_subjectTemplateGroupManager_BulkUpdateExpiredAtByRelationWithTx(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^UPDATE subject_template_group`
		mock.ExpectBegin()
		mock.ExpectPrepare(mockQuery)
		mock.ExpectExec(mockQuery).WithArgs(int64(3), int64(2), int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		tx, err := db.Beginx()
		assert.NoError(t, err)
		manager := &subjectTemplateGroupManager{DB: db}
		err = manager.BulkUpdateExpiredAtByRelationWithTx(tx, []SubjectRelation{{
			SubjectPK: int64(2),
			GroupPK:   int64(1),
			ExpiredAt: int64(3),
		}})

		assert.NoError(t, err, "query from db fail.")
	})
}

func Test_subjectTemplateGroupManager_BulkDeleteWithTx(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^DELETE FROM subject_template_group WHERE subject_pk = (.*) AND group_pk = (.*) AND template_id = (.*)`
		mock.ExpectBegin()
		mock.ExpectPrepare(mockQuery)
		mock.ExpectExec(mockQuery).WithArgs(int64(2), int64(1), int64(1)).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		tx, err := db.Beginx()
		assert.NoError(t, err)
		manager := &subjectTemplateGroupManager{DB: db}
		err = manager.BulkDeleteWithTx(tx, []SubjectTemplateGroup{{
			SubjectPK:  int64(2),
			TemplateID: int64(1),
			GroupPK:    int64(1),
			ExpiredAt:  int64(3),
		}})

		assert.NoError(t, err, "query from db fail.")
	})
}

func Test_subjectTemplateGroupManager_GetExpiredAtBySubjectGroup(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^SELECT (.*) FROM subject_template_group WHERE subject_pk =`
		mockRows := sqlmock.NewRows([]string{"policy_expired_at"}).AddRow(int64(1))
		mock.ExpectQuery(mockQuery).WithArgs(int64(1), int64(10), int64(0)).WillReturnRows(mockRows)

		manager := &subjectTemplateGroupManager{DB: db}
		expiredAt, err := manager.GetMaxExpiredAtBySubjectGroup(int64(1), int64(10), int64(0))

		assert.NoError(t, err, "query from db fail.")
		assert.Equal(t, expiredAt, int64(1))
	})
}

func Test_subjectTemplateGroupManager_ListPagingTemplateGroupMember(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^SELECT pk, subject_pk, template_id, group_pk, expired_at, created_at FROM subject_template_group
		 WHERE group_pk = (.*) AND template_id = (.*) ORDER BY pk DESC LIMIT (.*) OFFSET (.*)`
		mockRows := sqlmock.NewRows(
			[]string{"pk", "subject_pk", "template_id", "group_pk", "expired_at"},
		).AddRow(int64(1), int64(2), int64(3), int64(4), int64(0))
		mock.ExpectQuery(mockQuery).WithArgs(int64(1), int64(2), 0, 10).WillReturnRows(mockRows)

		manager := &subjectTemplateGroupManager{DB: db}
		relations, err := manager.ListPagingTemplateGroupMember(int64(1), int64(2), 0, 10)

		assert.NoError(t, err, "query from db fail.")
		assert.Len(t, relations, 1)
	})
}

func Test_subjectTemplateGroupManager_GetCount(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^SELECT COUNT(.*) FROM subject_template_group WHERE group_pk = `
		mockRows := sqlmock.NewRows([]string{"count(*)"}).AddRow(int64(1))
		mock.ExpectQuery(mockQuery).WithArgs(int64(1), int64(2)).WillReturnRows(mockRows)

		manager := &subjectTemplateGroupManager{DB: db}
		cnt, err := manager.GetTemplateGroupMemberCount(int64(1), int64(2))

		assert.NoError(t, err, "query from db fail.")
		assert.Equal(t, cnt, int64(1))
	})
}

func Test_subjectTemplateGroupManager_ListRelationBySubjectPKGroupPKs(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^SELECT
		pk,
		subject_pk,
		template_id,
		group_pk,
		expired_at,
		created_at
		FROM subject_template_group
		WHERE subject_pk = (.*)
		AND group_pk in (.*)`
		mockRows := sqlmock.NewRows(
			[]string{
				"subject_pk",
				"group_pk",
				"expired_at",
			},
		).AddRow(int64(1), int64(2), int64(3))
		mock.ExpectQuery(mockQuery).WithArgs(int64(123), int64(1)).WillReturnRows(mockRows)

		manager := &subjectTemplateGroupManager{DB: db}
		relations, err := manager.ListRelationBySubjectPKGroupPKs(
			123,
			[]int64{1},
		)

		assert.NoError(t, err, "query from db fail.")
		assert.Len(t, relations, 1)
	})
}

func Test_subjectTemplateGroupManager_BulkUpdateExpiredAtWithTx(t *testing.T) {
	database.RunWithMock(t, func(db *sqlx.DB, mock sqlmock.Sqlmock, t *testing.T) {
		mockQuery := `^UPDATE subject_template_group`
		mock.ExpectBegin()
		mock.ExpectPrepare(mockQuery)
		mock.ExpectExec(mockQuery).
			WithArgs(int64(3), int64(2), int64(1), int64(1)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		tx, err := db.Beginx()
		assert.NoError(t, err)
		manager := &subjectTemplateGroupManager{DB: db}
		err = manager.BulkUpdateExpiredAtWithTx(tx, []SubjectTemplateGroup{{
			SubjectPK:  int64(2),
			GroupPK:    int64(1),
			TemplateID: int64(1),
			ExpiredAt:  int64(3),
		}})

		assert.NoError(t, err, "query from db fail.")
	})
}
