package bot

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/snowflake/v2"
)

type ButtonEvent struct {
	*events.ComponentInteractionCreate
	Data ButtonData
}

type Button struct {
	Query   string
	Handler func(context.Context, ButtonEvent, *discord.MessageCreateBuilder) error
}

type ButtonData struct {
	params map[string]string
}

func (d *ButtonData) String(name string) string {
	if s, ok := d.params[name]; ok {
		return s
	}
	return ""
}

func (d *ButtonData) Int(name string) int {
	s, ok := d.params[name]
	if !ok {
		return 0
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}

	return n
}

func (d *ButtonData) Snowflake(name string) snowflake.ID {
	s, ok := d.params[name]
	if !ok {
		return 0
	}

	id, err := snowflake.Parse(s)
	if err != nil {
		return 0
	}

	return id
}

func parseButtonID(id string) (string, map[string]string, error) {
	params := make(map[string]string)

	query, form, ok := strings.Cut(id, "?")
	if ok {
		for _, part := range strings.Split(form, "&") {
			k, v, ok := strings.Cut(part, "=")
			if !ok {
				return "", nil, errors.New("malformed form param")
			}

			params[k] = v
		}
	}

	return query, params, nil
}

func AddButtons(bs ...Button) {
	for _, bt := range bs {
		Buttons[bt.Query] = bt
	}
}
