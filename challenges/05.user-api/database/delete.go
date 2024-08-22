package database

func Delete(id string) error {
	mu.Lock()
	defer mu.Unlock()

	foundedUser, err := getUserModelById(id)
	if err != nil {
		return err
	}

	foundedUser.Delete()
	db[foundedUser.Id] = foundedUser.UserModel

	return nil
}
