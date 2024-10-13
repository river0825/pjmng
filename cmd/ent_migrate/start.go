package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"river0825/cleanarchitecture/ent"
	"river0825/cleanarchitecture/ent/group"
	"river0825/cleanarchitecture/ent/user"
)

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Query().Where(user.Name("a8m")).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user was queried: ", u)

	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.
		Create().SetModel("Tesla").
		SetRegisteredAt(time.Now()).Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	am8, err := client.User.Create().
		AddCars(tesla, ford).
		SetAge(30).
		SetName("a8m").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	log.Println("user was created: ", am8)
	return am8, nil
}

func QueryCarUsers(ctx context.Context, user *ent.User) error {
	cars, err := user.QueryCars().All(ctx)

	if err != nil {
		return fmt.Errorf("failed querying cars: %w", err)
	}

	for _, c := range cars {
		owner, err := c.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
		}
		log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}

	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users created above.
	err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		// Attach this car to Ariel.
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()).
		// Attach this car to Ariel.
		SetOwner(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		// Attach this car to Neta.
		SetOwner(neta).
		Exec(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Exec(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}

func QueryGithub(ctx context.Context, client *ent.Client) error {
	cars, err := client.Group.Query().Where(group.Name("GitHub")).QueryUsers().QueryCars().All(ctx)

	if err != nil {
		return fmt.Errorf("failed getting cars: %w", err)
	}

	log.Println("cars returned:", cars)

	return nil

}
