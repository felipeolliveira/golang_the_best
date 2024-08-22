package database

func Update(id string, firstname string, lastname string, bio string) error {
	mu.Lock()
	defer mu.Unlock()

	foundedUser, err := getUserModelById(id)
	if err != nil {
		return err
	}

	foundedUser.Update(firstname, lastname, bio)
	db[foundedUser.Id] = foundedUser.UserModel

	return nil
}
