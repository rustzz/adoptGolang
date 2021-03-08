package controller

import (
	"adoptGolang/internal/engine/system/customAPI"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

type Client struct {
	GroupID	int
	API		*api.VK
}

type Controller struct {
	*Client
	*customAPI.CustomAPI
}

func NewController(
	client *api.VK, group object.GroupsGroup,
) *Controller {
	return &Controller{
		Client: &Client{
			GroupID: group.ID,
			API: client,
		},
		CustomAPI: &customAPI.CustomAPI{
			GroupID: group.ID,
			API: client,
		},
	}
}
