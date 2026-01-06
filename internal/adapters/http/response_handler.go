package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"www.github.com/fummbly/ai-dash/internal/domain"
	"www.github.com/fummbly/ai-dash/internal/service"
)

type ResponseHandler struct {
	responeService service.ResponseService
}

func NewResponseHandler(svc service.ResponseService) *ResponseHandler {
	return &ResponseHandler{
		responeService: svc,
	}
}

func (h *ResponseHandler) StreamResponse(c echo.Context) error {
	w := c.Response()
	w.Header().Set("Content-Type", "text/event-stream")
	w.WriteHeader(http.StatusOK)

	resChan := make(chan domain.Response)

	go func() {
		err := h.responeService.Generate(resChan)
		if err != nil {
			panic(err)
		}
	}()

	message := ""

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected, ip: %v", c.RealIP())

			return nil
		case response, ok := <-resChan:
			if !ok {
				log.Printf("Generated data finished")
				// TODO create SSE data formater
				if _, err := fmt.Fprintf(w, "event: GenerateFinished\ndata: <div>Generated Data finished</div>\n\n"); err != nil {
					return err
				}

				w.Flush()

				return nil
			}

			message += response.Response
			log.Printf("Generated response: %s", message)

			if _, err := fmt.Fprintf(w, "data: <div>%s</div>\n\n", message); err != nil {
				return err
			}

			w.Flush()
		}
	}

	/*
		for {
			select {
			case <-c.Request().Context().Done():
				log.Printf("SSE client disconnected, ip: %v", c.RealIP())
				return nil
			case data, ok := <-dataChan:
				if !ok {
					log.Printf("Generated data finished")
					if _, err := fmt.Fprintf(w, "event: GenerateFinished\ndata: <div>Generated Data finished</div>\n\n"); err != nil {
						return err
					}
					w.Flush()

					return nil
				}
				log.Printf("Generated response: %s", data)
				message += data
				if _, err := fmt.Fprintf(w, "data: <div>%s</div>\n\n", message); err != nil {
					return err
				}
				w.Flush()
			}
		}
	*/
}
