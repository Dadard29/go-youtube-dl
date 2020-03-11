package managers

import "github.com/pkg/errors"

func VideoManagerGet(token string) (error) {
	// check sub
	c, err := sc.CheckToken(token)

	if err != nil {
		return err
	}

	if !c {
		return errors.New(errorNotSubscribed)
	}


	// todo
	return nil
}

func VideoManagerCreate(token string) (error) {
	// check sub
	c, err := sc.CheckToken(token)

	if err != nil {
		return err
	}

	if !c {
		return errors.New(errorNotSubscribed)
	}

	// todo
}

func VideoManagerUpdate(token string) (error) {
	// check sub
	c, err := sc.CheckToken(token)

	if err != nil {
		return err
	}

	if !c {
		return errors.New(errorNotSubscribed)
	}
	// todo
}

func VideoManagerRemove(token string) (error) {

	// check sub
	c, err := sc.CheckToken(token)

	if err != nil {
		return err
	}

	if !c {
		return errors.New(errorNotSubscribed)
	}
	// todo
}
