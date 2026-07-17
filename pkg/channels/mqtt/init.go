package mqtt

import (
	"github.com/BlackOceanX/picoclawquant/pkg/bus"
	"github.com/BlackOceanX/picoclawquant/pkg/channels"
	"github.com/BlackOceanX/picoclawquant/pkg/config"
)

func init() {
	channels.RegisterSafeFactory(
		config.ChannelMQTT,
		func(bc *config.Channel, cfg *config.MQTTSettings, b *bus.MessageBus) (channels.Channel, error) {
			return NewMQTTChannel(bc, cfg, b)
		},
	)
}
