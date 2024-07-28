package database

import (
	"database/sql"
	"github.com/withoutsecondd/ToDo/models"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"reflect"
	"testing"
)

func TestMySqlDB_GetUserById(t *testing.T) {
	// ARRANGE

	conn, mock, err := sqlmock.Newx()
	if err != nil {
		t.Fatalf("an error %v occurred when opening a stub connection", err)
	}
	defer conn.Close()

	db := NewMySqlDB(conn)

	testTable := []struct {
		name    string
		mock    func()
		arg     int64
		want    *models.UserResponse
		wantErr bool
	}{
		{
			name: "testing call to an existing user",
			mock: func() {
				columns := []string{"id", "age", "first_name", "last_name", "city", "email"}
				rows := sqlmock.NewRows(columns).
					AddRow("1", "18", "testFirstName", "testLastName", "testCity", "test@example.com")

				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnRows(rows)
			},
			arg: 1,
			want: &models.UserResponse{
				ID:        1,
				Age:       18,
				FirstName: "testFirstName",
				LastName:  "testLastName",
				City:      "testCity",
				Email:     "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "testing call to user that isn't present in db",
			mock: func() {
				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			},
			arg:     1,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tC := range testTable {
		t.Run(tC.name, func(t *testing.T) {
			// ACT

			tC.mock()

			got, err := db.GetUserById(tC.arg)

			// ASSERT

			if (err != nil) != tC.wantErr {
				t.Fatalf("expected error? %v. got: %v", tC.wantErr, err)
				return
			}

			if !reflect.DeepEqual(got, tC.want) {
				t.Fatalf("expected: %v. got: %v", got, tC.want)
				return
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatalf("unfullfilled expectations: %v", err)
			}
		})
	}
}

func TestMySqlDB_GetUserByEmail(t *testing.T) {

}

func TestMySqlDB_GetUserPasswordByEmail(t *testing.T) {

}

func TestMySqlDB_CreateUser(t *testing.T) {

}

func TestMySqlDB_GetListById(t *testing.T) {

}

func TestMySqlDB_GetListsByUserId(t *testing.T) {

}

func TestMySqlDB_GetTasksByUserId(t *testing.T) {

}

func TestMySqlDB_GetTasksByListId(t *testing.T) {

}

func TestMySqlDB_GetTaskById(t *testing.T) {

}
