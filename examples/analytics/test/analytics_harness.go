package test

import (
	"github.com/orbs-network/orbs-client-sdk-go/orbs"
	"github.com/stretchr/testify/require"
	"testing"
)

func (h *analyticsContract) getEvents(t *testing.T, sender *orbs.OrbsAccount) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.name, "getEvents")
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}

func (h *analyticsContract) getAggregationByActionOverPeriodOfTime(t *testing.T, sender *orbs.OrbsAccount, eventCategory string, aggregationType string, startTime uint64, endTime uint64) interface{} {
	query, err := h.client.CreateQuery(sender.PublicKey, h.name, "getAggregationByActionOverPeriodOfTime", eventCategory, aggregationType, startTime, endTime)
	require.NoError(t, err)

	queryResponse, err := h.client.SendQuery(query)
	require.NoError(t, err)

	return queryResponse.OutputArguments[0]
}
