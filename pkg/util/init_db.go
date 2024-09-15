package util

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/ent"
	"sastoj/ent/group"
	"sastoj/ent/user"
)

func InsertDefaultGroup(ctx context.Context, client *ent.Client) (*ent.Group, error) {
	rootGroup := client.Group.Query().Where(group.IsRootEQ(true)).FirstX(ctx)
	if rootGroup != nil {
		return rootGroup, nil
	}
	rootGroup, err := client.Group.Create().
		SetGroupName("root").SetIsRoot(true).Save(ctx)
	if err != nil {
		return rootGroup, err
	}
	log.Info("Inserted default root group.")
	return rootGroup, err
}

func InsertDefaultUser(ctx context.Context, client *ent.Client, rootGroup *ent.Group, rootName string, rootPassword string) error {
	rootUser := client.User.Query().Where(user.HasGroupsWith(group.IsRootEQ(true))).FirstX(ctx)
	if rootUser != nil {
		return nil
	}
	salt := GenerateRandomString(5, "")
	_, err := client.User.Create().
		SetUsername(rootName).
		SetPassword(GenerateMD5Password(salt, rootPassword)).
		SetSalt(salt).
		AddGroups(rootGroup).
		Save(ctx)
	if err != nil {
		log.Debug("err: ", err)
		return err
	}
	return nil
}
