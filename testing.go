package sooslib

import "os/user"

func osGetwdMock() (dir string, err error) {
	return "/root", nil
}

func userCurrentMock() (*user.User, error) {

	usr := user.User{
		HomeDir:  "/root",
		Username: "root",
		Name:     "root",
		Uid:      "0",
		Gid:      "0",
	}

	return &usr, nil
}
