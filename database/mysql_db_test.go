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

	testTable := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		arg     int64
		want    *models.UserResponse
		wantErr bool
	}{
		{
			name: "testing call to an existing user",
			mock: func(mock sqlmock.Sqlmock) {
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
			mock: func(mock sqlmock.Sqlmock) {
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
			conn, mock, err := sqlmock.Newx()
			if err != nil {
				t.Fatalf("an error %v occurred when opening a stub connection", err)
			}
			defer conn.Close()

			db := NewMySqlDB(conn)
			tC.mock(mock)

			// ACT

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
	// ARRANGE

	testTable := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		arg     string
		want    *models.UserResponse
		wantErr bool
	}{
		{
			name: "testing call when user with such email exists",
			mock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "age", "first_name", "last_name", "city", "email"}
				rows := sqlmock.NewRows(columns).
					AddRow("1", "18", "testFirstName", "testLastName", "testCity", "test@example.com")

				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE email = ?").
					WithArgs("test@example.com").
					WillReturnRows(rows)
			},
			arg: "test@example.com",
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
			name: "testing call when user with such email doesn't exist",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE email = ?").
					WithArgs("invalid@example.com").
					WillReturnError(sql.ErrNoRows)
			},
			arg:     "invalid@example.com",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tC := range testTable {
		t.Run(tC.name, func(t *testing.T) {
			conn, mock, err := sqlmock.Newx()
			if err != nil {
				t.Fatalf("an error %v occurred when opening a stub connection", err)
			}
			defer conn.Close()

			db := NewMySqlDB(conn)
			tC.mock(mock)

			// ACT

			got, err := db.GetUserByEmail(tC.arg)

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

func TestMySqlDB_GetUserPasswordByEmail(t *testing.T) {
	// ARRANGE

	testTable := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		arg     string
		want    []byte
		wantErr bool
	}{
		{
			name: "testing call when user with such email exists",
			mock: func(mock sqlmock.Sqlmock) {
				columns := []string{"password"}
				rows := sqlmock.NewRows(columns).
					AddRow("test password")

				mock.ExpectQuery("SELECT password FROM withoutsecondd.user WHERE email = ?").
					WithArgs("test@example.com").
					WillReturnRows(rows)
			},
			arg:     "test@example.com",
			want:    []byte("test password"),
			wantErr: false,
		},
		{
			name: "testing call when user with such email doesn't exist",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT password FROM withoutsecondd.user WHERE email = ?").
					WithArgs("invalid@example.com").
					WillReturnError(sql.ErrNoRows)
			},
			arg:     "invalid@example.com",
			want:    nil,
			wantErr: true,
		},
		{
			name: "testing call with invalid argument",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT password FROM withoutsecondd.user WHERE email = ?").
					WithArgs("invalidArg").
					WillReturnError(sql.ErrNoRows)
			},
			arg:     "invalidArg",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tC := range testTable {
		t.Run(tC.name, func(t *testing.T) {
			conn, mock, err := sqlmock.Newx()
			if err != nil {
				t.Fatalf("an error %v occurred when opening a stub connection", err)
			}
			defer conn.Close()

			db := NewMySqlDB(conn)
			tC.mock(mock)

			// ACT

			got, err := db.GetUserPasswordByEmail(tC.arg)

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

func TestMySqlDB_CreateUser(t *testing.T) {

}

func TestMySqlDB_GetListById(t *testing.T) {
	// ARRANGE

	testTable := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		arg     int64
		want    *models.List
		wantErr bool
	}{
		{
			name: "testing call to existing list",
			mock: func(mock sqlmock.Sqlmock) {
				columns := []string{"id", "user_id", "title"}
				rows := sqlmock.NewRows(columns).
					AddRow("1", "1", "Test list")

				mock.ExpectQuery("SELECT .+ FROM withoutsecondd.list WHERE id = ?").
					WithArgs(1).
					WillReturnRows(rows)
			},
			arg: 1,
			want: &models.List{
				ID:     1,
				UserID: 1,
				Title:  "Test list",
			},
			wantErr: false,
		},
		{
			name: "testing call to non-existing list",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT .+ FROM withoutsecondd.list WHERE id = ?").
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
			arg:     2,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tC := range testTable {
		t.Run(tC.name, func(t *testing.T) {
			conn, mock, err := sqlmock.Newx()
			if err != nil {
				t.Fatalf("an error %v occurred when opening a stub connection", err)
			}
			defer conn.Close()

			db := NewMySqlDB(conn)
			tC.mock(mock)

			// ACT

			got, err := db.GetListById(tC.arg)

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

func TestMySqlDB_GetListsByUserId(t *testing.T) {
	// ARRANGE

	testTable := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		arg     int64
		want    []models.List
		wantErr bool
	}{
		{
			name: "testing call to lists of existing user",
			mock: func(mock sqlmock.Sqlmock) {
				userColumns := []string{"id", "age", "first_name", "last_name", "city", "email"}
				userRows := sqlmock.NewRows(userColumns).
					AddRow("1", "18", "testFirstName", "testLastName", "testCity", "test@example.com")

				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnRows(userRows)

				listColumns := []string{"id", "user_id", "title"}
				listRows := sqlmock.NewRows(listColumns).
					AddRow("1", "1", "Test list #1").
					AddRow("2", "1", "Test list #2")

				mock.ExpectQuery("SELECT .+ FROM withoutsecondd.list WHERE user_id = ?").
					WithArgs(1).
					WillReturnRows(listRows)
			},
			arg: 1,
			want: []models.List{
				{
					ID:     1,
					UserID: 1,
					Title:  "Test list #1",
				},
				{
					ID:     2,
					UserID: 1,
					Title:  "Test list #2",
				},
			},
			wantErr: false,
		},
		{
			name: "testing call to lists of non-existing user",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			},
			arg:     1,
			want:    nil,
			wantErr: true,
		},
		{
			name: "testing call to user with no lists",
			mock: func(mock sqlmock.Sqlmock) {
				userColumns := []string{"id", "age", "first_name", "last_name", "city", "email"}
				userRows := sqlmock.NewRows(userColumns).
					AddRow("1", "18", "testFirstName", "testLastName", "testCity", "test@example.com")

				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnRows(userRows)

				listColumns := []string{"id", "user_id", "title"}
				listRows := sqlmock.NewRows(listColumns)

				mock.ExpectQuery("SELECT .+ FROM withoutsecondd.list WHERE user_id = ?").
					WithArgs(1).
					WillReturnRows(listRows)
			},
			arg:     1,
			want:    []models.List{},
			wantErr: false,
		},
	}

	for _, tC := range testTable {
		t.Run(tC.name, func(t *testing.T) {
			conn, mock, err := sqlmock.Newx()
			if err != nil {
				t.Fatalf("an error %v occurred when opening a stub connection", err)
			}
			defer conn.Close()

			db := NewMySqlDB(conn)
			tC.mock(mock)

			// ACT

			got, err := db.GetListsByUserId(tC.arg)

			// ASSERT

			if (err != nil) != tC.wantErr {
				t.Fatalf("expected error? %v. got: %v", tC.wantErr, err)
			}

			if !reflect.DeepEqual(got, tC.want) {
				t.Fatalf("expected: %v. got: %v", got, tC.want)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatalf("unfullfilled expectations: %v", err)
			}
		})
	}
}

func TestMySqlDB_GetTasksByUserId(t *testing.T) {
	// ARRANGE

	testTable := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		arg     int64
		want    []models.Task
		wantErr bool
	}{
		{
			name: "testing call to tasks of existing user",
			mock: func(mock sqlmock.Sqlmock) {
				userColumns := []string{"id", "age", "first_name", "last_name", "city", "email"}
				userRows := sqlmock.NewRows(userColumns).
					AddRow("1", "18", "testFirstName", "testLastName", "testCity", "test@example.com")

				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnRows(userRows)

				taskColumns := []string{"id", "list_id", "title", "description", "status", "deadline"}
				taskRows := sqlmock.NewRows(taskColumns).
					AddRow("1", "1", "Test task #1", "This is a test task", "Completed", "Test date").
					AddRow("2", "1", "Test task #2", "This is a test task", "Completed", "Test date").
					AddRow("3", "2", "Test task #3", "This is a test task", "Completed", "Test date").
					AddRow("4", "2", "Test task #4", "This is a test task", "Completed", "Test date")

				query := `
					SELECT t.id, t.list_id, t.title, t.description, t.status, t.deadline FROM
    					\(SELECT id FROM withoutsecondd.list WHERE user_id = \?\) AS l
        			INNER JOIN withoutsecondd.task AS t ON t.list_id = l.id
					ORDER BY t.list_id;
				`
				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(taskRows)
			},
			arg: 1,
			want: []models.Task{
				{
					ID:          1,
					ListID:      1,
					Title:       "Test task #1",
					Description: "This is a test task",
					Status:      "Completed",
					Deadline:    "Test date",
					Tags:        nil,
				},
				{
					ID:          2,
					ListID:      1,
					Title:       "Test task #2",
					Description: "This is a test task",
					Status:      "Completed",
					Deadline:    "Test date",
					Tags:        nil,
				},
				{
					ID:          3,
					ListID:      2,
					Title:       "Test task #3",
					Description: "This is a test task",
					Status:      "Completed",
					Deadline:    "Test date",
					Tags:        nil,
				},
				{
					ID:          4,
					ListID:      2,
					Title:       "Test task #4",
					Description: "This is a test task",
					Status:      "Completed",
					Deadline:    "Test date",
					Tags:        nil,
				},
			},
			wantErr: false,
		},
		{
			name: "testing call to tasks of non-existing user",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnError(sql.ErrNoRows)
			},
			arg:     1,
			want:    nil,
			wantErr: true,
		},
		{
			name: "testing call to user with no tasks",
			mock: func(mock sqlmock.Sqlmock) {
				userColumns := []string{"id", "age", "first_name", "last_name", "city", "email"}
				userRows := sqlmock.NewRows(userColumns).
					AddRow("1", "18", "testFirstName", "testLastName", "testCity", "test@example.com")

				mock.ExpectQuery("SELECT id, age, first_name, last_name, city, email FROM withoutsecondd.user WHERE id = ?").
					WithArgs(1).
					WillReturnRows(userRows)

				taskColumns := []string{"id", "list_id", "title", "description", "status", "deadline"}
				taskRows := sqlmock.NewRows(taskColumns)

				query := `
					SELECT t.id, t.list_id, t.title, t.description, t.status, t.deadline FROM
    					\(SELECT id FROM withoutsecondd.list WHERE user_id = \?\) AS l
        			INNER JOIN withoutsecondd.task AS t ON t.list_id = l.id
					ORDER BY t.list_id;
				`
				mock.ExpectQuery(query).
					WithArgs(1).
					WillReturnRows(taskRows)
			},
			arg:     1,
			want:    []models.Task{},
			wantErr: false,
		},
	}

	for _, tC := range testTable {
		t.Run(tC.name, func(t *testing.T) {
			conn, mock, err := sqlmock.Newx()
			if err != nil {
				t.Fatalf("an error %v occurred when opening a stub connection", err)
			}
			defer conn.Close()

			db := NewMySqlDB(conn)
			tC.mock(mock)

			// ACT

			got, err := db.GetTasksByUserId(tC.arg)

			// ASSERT

			if (err != nil) != tC.wantErr {
				t.Fatalf("expected error? %v. got: %v", tC.wantErr, err)
			}

			if !reflect.DeepEqual(got, tC.want) {
				t.Fatalf("expected: %v. got: %v", got, tC.want)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Fatalf("unfullfilled expectations: %v", err)
			}
		})
	}
}

func TestMySqlDB_GetTasksByListId(t *testing.T) {

}

func TestMySqlDB_GetTaskById(t *testing.T) {

}
