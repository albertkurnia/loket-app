package usecase

import (
	"errors"
	"loket-app/modules/event/model"
	locModel "loket-app/modules/location/model"
	"testing"
	"time"

	queryMock "loket-app/modules/event/query/mock"
	locUCMock "loket-app/modules/location/usecase/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	errMock = errors.New("error mock")
)

func TestEventUsecaseCreateTicket(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventQuery := queryMock.NewMockEventQuery(ctrl)

	t.Run("should return error: invalid request", func(t *testing.T) {
		var data *model.CreateTicketReq

		impl := eventUseCaseImpl{}

		ticket, err := impl.CreateTicket(data)
		assert.Error(t, err)
		assert.Nil(t, ticket)
	})

	t.Run("should return error since there is error_load_ticket_by_type", func(t *testing.T) {
		var data model.CreateTicketReq
		data.Type = "sample"

		eventQuery.
			EXPECT().
			LoadTicketByType("SAMPLE").
			Return(nil, errMock)

		impl := eventUseCaseImpl{
			EventQuery: eventQuery,
		}

		ticket, err := impl.CreateTicket(&data)
		assert.Error(t, err)
		assert.Nil(t, ticket)
		assert.Equal(t, "error mock", err.Error())
	})

	t.Run("should return error: ticket with that type already exist", func(t *testing.T) {
		var data model.CreateTicketReq
		data.Type = "sample"

		eventQuery.
			EXPECT().
			LoadTicketByType("SAMPLE").
			Return(&model.Ticket{}, nil)

		impl := eventUseCaseImpl{
			EventQuery: eventQuery,
		}

		ticket, err := impl.CreateTicket(&data)
		assert.Error(t, err)
		assert.Nil(t, ticket)
		assert.Equal(t, "ticket with that type already exist", err.Error())
	})

	t.Run("should return error since error_insert_ticket", func(t *testing.T) {
		var data model.CreateTicketReq
		data.Type = "sample"

		eventQuery.
			EXPECT().
			LoadTicketByType("SAMPLE").
			Return(nil, nil)

		eventQuery.
			EXPECT().
			InsertTicket(gomock.Any()).
			Return(nil, errMock)

		impl := eventUseCaseImpl{
			EventQuery: eventQuery,
		}

		ticket, err := impl.CreateTicket(&data)
		assert.Error(t, err)
		assert.Nil(t, ticket)
		assert.Equal(t, "error mock", err.Error())
	})

	t.Run("should success to return ticket", func(t *testing.T) {
		var data model.CreateTicketReq
		data.Type = "sample"

		eventQuery.
			EXPECT().
			LoadTicketByType("SAMPLE").
			Return(nil, nil)

		eventQuery.
			EXPECT().
			InsertTicket(gomock.Any()).
			Return(&model.Ticket{}, nil)

		impl := eventUseCaseImpl{
			EventQuery: eventQuery,
		}

		ticket, err := impl.CreateTicket(&data)
		assert.NoError(t, err)
		assert.NotNil(t, ticket)
	})
}

func TestEventUsecaseCreateEvent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventQuery := queryMock.NewMockEventQuery(ctrl)

	t.Run("should return error invalid request", func(t *testing.T) {
		var data *model.CreateEventReq

		impl := eventUseCaseImpl{}

		event, err := impl.CreateEvent(data)
		assert.Error(t, err)
		assert.Nil(t, event)
		assert.Equal(t, "invalid request", err.Error())
	})

	t.Run("should return error since error_insert_event", func(t *testing.T) {
		var data model.CreateEventReq

		eventQuery.
			EXPECT().
			InsertEvent(gomock.Any()).
			Return(nil, errMock)

		impl := eventUseCaseImpl{
			EventQuery: eventQuery,
		}

		event, err := impl.CreateEvent(&data)
		assert.Error(t, err)
		assert.Nil(t, event)
	})

	t.Run("should succes to return event", func(t *testing.T) {
		var data model.CreateEventReq

		eventQuery.
			EXPECT().
			InsertEvent(gomock.Any()).
			Return(&model.Event{}, nil)

		impl := eventUseCaseImpl{
			EventQuery: eventQuery,
		}

		event, err := impl.CreateEvent(&data)
		assert.NoError(t, err)
		assert.NotNil(t, event)
	})
}

func TestEventUsecaseGetEventInformation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventQuery := queryMock.NewMockEventQuery(ctrl)
	locUsecase := locUCMock.NewMockLocationUseCase(ctrl)

	var id uint64 = 1

	t.Run("should return error since error_load_event_by_id", func(t *testing.T) {
		eventQuery.
			EXPECT().
			LoadEventByID(gomock.Any()).
			Return(nil, errMock)

		impl := NewEventUseCase(eventQuery, nil)

		eventInfo, err := impl.GetEventInformation(id)
		assert.Error(t, err)
		assert.Nil(t, eventInfo)
	})

	t.Run("should return error since error_load_ticket_by_ids", func(t *testing.T) {
		var event model.Event
		event.ID = 1
		event.Title = "aaa"
		event.StartDate = time.Now()
		event.EndDate = time.Now()
		event.TicketID = []uint64{1, 2}

		eventQuery.
			EXPECT().
			LoadEventByID(gomock.Any()).
			Return(&event, nil)

		eventQuery.
			EXPECT().
			LoadTicketByIDs(event.TicketID).
			Return(nil, errMock)

		impl := NewEventUseCase(eventQuery, nil)

		eventInfo, err := impl.GetEventInformation(id)
		assert.Error(t, err)
		assert.Nil(t, eventInfo)
	})

	t.Run("should return error since error_load_location_by_id", func(t *testing.T) {
		var event model.Event
		event.ID = 1
		event.Title = "aaa"
		event.StartDate = time.Now()
		event.EndDate = time.Now()
		event.TicketID = []uint64{1, 2}
		event.LocationID = 1

		tickets := make([]*model.Ticket, 0)
		ticket1 := model.Ticket{
			ID: 1,
		}

		tickets = append(tickets, &ticket1)

		eventQuery.
			EXPECT().
			LoadEventByID(gomock.Any()).
			Return(&event, nil)

		eventQuery.
			EXPECT().
			LoadTicketByIDs(event.TicketID).
			Return(tickets, nil)

		locUsecase.
			EXPECT().
			LoadLocationByID(event.LocationID).
			Return(nil, errMock)

		impl := NewEventUseCase(eventQuery, locUsecase)

		eventInfo, err := impl.GetEventInformation(id)
		assert.Error(t, err)
		assert.Nil(t, eventInfo)
	})

	t.Run("should return event information", func(t *testing.T) {
		var event model.Event
		event.ID = 1
		event.Title = "aaa"
		event.StartDate = time.Now()
		event.EndDate = time.Now()
		event.TicketID = []uint64{1, 2}
		event.LocationID = 1

		tickets := make([]*model.Ticket, 0)
		ticket1 := model.Ticket{
			ID: 1,
		}

		tickets = append(tickets, &ticket1)

		eventQuery.
			EXPECT().
			LoadEventByID(gomock.Any()).
			Return(&event, nil)

		eventQuery.
			EXPECT().
			LoadTicketByIDs(event.TicketID).
			Return(tickets, nil)

		locUsecase.
			EXPECT().
			LoadLocationByID(event.LocationID).
			Return(&locModel.Location{}, nil)

		impl := NewEventUseCase(eventQuery, locUsecase)

		eventInfo, err := impl.GetEventInformation(id)
		assert.NoError(t, err)
		assert.NotNil(t, eventInfo)
	})
}

func TestEventUsecaseGetTicket(t *testing.T) {
	var id uint64

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventQuery := queryMock.NewMockEventQuery(ctrl)

	eventQuery.
		EXPECT().
		GetTicket(id).
		Return(&model.Ticket{}, nil)

	impl := eventUseCaseImpl{
		EventQuery: eventQuery,
	}
	resp, err := impl.GetTicket(id)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
