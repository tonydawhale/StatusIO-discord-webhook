package constants

import (
	"StatusIO-discord-webhook/types"
)

var EmbedColorGreen = 4437377
var EmbedColorYellow = 16426522
var EmbedColorOrange = 15885602
var EmbedColorRed = 15746887
var EmbedColorBlack = 2303786
var EmbedColorBlurple = 5793266

var TestStatusPageIncident = types.StatusPageIncident{
	Id: "l59btkwbxm3t",
	Name: "Hypixel Forums Unavailable",
	Status: "resolved",
	CreatedAt: "2023-04-12T16:01:07.925-04:00",
	UpdatedAt: "2023-04-12T16:08:40.511-04:00",
	ResolvedAt: "2023-04-12T16:08:40.496-04:00",
	Impact: "critical",
	Shortlink: "https://stspg.io/6hrm3cxn6qsr",
	StartedAt: "2023-04-12T16:01:07.918-04:00",
	PageId: "qhddg0t7xyb0",
	IncidentUpdates: []types.StatusPageIncidentUpdate{
		{
			Id: "dj63tgjq5yk3",
			Status: "resolved",
			CreatedAt: "2023-04-12T16:08:40.496-04:00",
			UpdatedAt: "2023-04-12T16:08:40.496-04:00",
			Body: "This incident has been resolved.",
			DisplayAt: "2023-04-12T16:08:40.496-04:00",
			IncidentId: "l59btkwbxm3t",
			DeliverNotifications: true,
			AffectedComponents: []types.StatusPageComponentUpdate{
				{
					Code: "200q0q2bh1xk",
					Name: "Forums",
					OldStatus: "major_outage",
					NewStatus: "operational",
				},
			},
		},
		{
			Id: "fpdknz94kscp",
			Status: "investigating",
			Body: "We are currently investigating an issue with the Hypixel Forums being offline and unavailable",
			IncidentId: "l59btkwbxm3t",
			CreatedAt: "2023-04-12T16:01:08.002-04:00",
			UpdatedAt: "2023-04-12T16:01:08.002-04:00",
			DisplayAt: "2023-04-12T16:01:08.002-04:00",
			DeliverNotifications: true,
			AffectedComponents: []types.StatusPageComponentUpdate{
				{
					Code: "200q0q2bh1xk",
					Name: "Forums",
					OldStatus: "operational",
					NewStatus: "major_outage",
				},
			},
		},
	},
	Components: []types.StatusPageComponent{
		{
			Id: "200q0q2bh1xk",
			Name: "Forums",
			Status: "operational",
			CreatedAt: "2020-10-07T15:24:28.086-04:00",
			UpdatedAt: "2023-04-13T09:05:11.868-04:00",
			Position: 6,
			Description: "https://hypixel.net",
			Showcase: false,
			StartDate: "2019-12-31",
			PageId: "qhddg0t7xyb0",
			Group: false,
			OnlyShowIfDegraded: false,
		},
	},
}