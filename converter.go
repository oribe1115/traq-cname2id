package cname2id

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/samber/lo"
	"github.com/traPtitech/go-traq"
)

type Converter struct {
	apiClient *traq.APIClient
	auth      context.Context
}

func NewConverter(apiClient *traq.APIClient, auth context.Context) *Converter {
	return &Converter{
		apiClient: apiClient,
		auth:      auth,
	}
}

func (c *Converter) GetChannelID(channelName string) (string, error) {
	list, res, err := c.apiClient.ChannelApi.GetChannels(c.auth).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to get channel list from traQ API: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code returned from traQ API: %s", res.Status)
	}

	return nameToID(channelName, list.GetPublic())
}

func nameToID(channelName string, list []traq.Channel) (string, error) {
	channelName = strings.TrimPrefix(channelName, "#")
	split := strings.Split(channelName, "/")
	if len(split) < 1 {
		return "", fmt.Errorf("unexpected channel name format: `%s`", channelName)
	}

	root, ok := lo.Find(list, func(item traq.Channel) bool {
		return !item.ParentId.IsSet() && item.GetName() == split[0]
	})

	if !ok {
		return "", fmt.Errorf("root channel `%s` not found", split[0])
	}

	id2Channel := lo.SliceToMap(list, func(item traq.Channel) (string, *traq.Channel) {
		return item.Id, &item
	})

	current := &root
	for i := 1; i < len(split); i++ {
		parent := current
		brothers := parent.Children

		found := false
		for _, brother := range brothers {
			channel, ok := id2Channel[brother]
			if !ok {
				return "", fmt.Errorf("not found channel id `%s` in channel list", brother)
			}

			if channel.GetName() == split[i] {
				current = channel
				found = true
				break
			}
		}

		if !found {
			return "", fmt.Errorf("`#%s` found but `#%s` not found", strings.Join(split[:i], "/"), channelName)
		}
	}

	return current.Id, nil
}
