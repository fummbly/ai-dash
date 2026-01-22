package http

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"www.github.com/fummbly/ai-dash/internal/domain"
	"www.github.com/fummbly/ai-dash/internal/parser"
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

	question := c.QueryParam("question")

	resChan := make(chan domain.Response)

	go func() {
		err := h.responeService.Generate(resChan, question)
		if err != nil {
			log.Printf("Failed to generate response: %v", err)

			return
		}
	}()

	message := strings.Builder{}

	for {
		select {
		case <-c.Request().Context().Done():
			log.Printf("SSE client disconnected, ip: %v", c.RealIP())

			return nil
		case response, ok := <-resChan:
			if !ok {
				// log.Printf("Generated data finished")
				// TODO create SSE data formater
				if _, err := fmt.Fprintf(w, "event: GenerateFinished\ndata: <div>Generated Data finished</div>\n\n"); err != nil {
					return err
				}

				w.Flush()

				return nil
			}

			message.WriteString(response.Response)
			parsedMessage := parser.ConvertMarkdown(message.String())
			// log.Printf("Generated response: %s", message)

			if _, err := fmt.Fprintf(w, "data: <div>%s</div>\n\n", parsedMessage); err != nil {
				return err
			}

			w.Flush()
		}
	}
}
