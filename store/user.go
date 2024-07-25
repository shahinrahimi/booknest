package store

import "github.com/shahinrahimi/booknest/pkg/user"

func (s *SqliteStore) GetUsers() ([]*user.User, error) {
	rows, err := s.db.Query(user.SelectAll)
	if err != nil {
		return nil, err
	}
	var users []*user.User
	for rows.Next() {
		var user user.User
		if err := rows.Scan(user.ToFeilds()...); err != nil {
			s.logger.Printf("Error scranning rows for a user: %v", err)
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *SqliteStore) GetUser(id string) (*user.User, error) {
	row := s.db.QueryRow(user.Select, id)
	var user user.User
	if err := row.Scan(user.ToFeilds()...); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqliteStore) CreateUser(u user.User) error {
	if _, err := s.db.Exec(user.Insert, u.ToArgs()...); err != nil {
		return err
	}
	return nil
}

func (s *SqliteStore) UpdateUser(id string, u user.User) error {
	if _, err := s.db.Exec(user.Update, u.ToUpdatedArgs()...); err != nil {
		return err
	}
	return nil
}

func (s *SqliteStore) DeleteUser(id string) error {
	if _, err := s.db.Exec(user.Delete, id); err != nil {
		return err
	}
	return nil
}
