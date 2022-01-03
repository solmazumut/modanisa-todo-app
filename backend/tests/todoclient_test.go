package tests

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"net/http"
	"testing"
)

func TestAddTodo(t *testing.T) {
	// initialize PACT DSL
	pact := dsl.Pact{
		Consumer: "frontend",
		Provider: "backend",
	}

	// setup a PACT Mock Server
	pact.Setup(true)

	t.Run("add test-todo", func(t *testing.T) {
		task := "test-todo"


		pact.
			AddInteraction(). // specify PACT interaction
			Given("The initial todo list is empty "). // specify Provider state
			UponReceiving(fmt.Sprintf("A POST request for %s",task)). // specify test case name
			WithRequest(
				dsl.Request{
					Method: http.MethodPost,
					Path:   dsl.String("/todo"),
					Headers: dsl.MapMatcher{
						"Accept": dsl.String("application/json"),
					},
					Body: dsl.MapMatcher{
						"task": dsl.String(fmt.Sprintf("%s",task)),
					},
				},
			).
			WillRespondWith(
				dsl.Response{
					Status: http.StatusOK,
					Headers: dsl.MapMatcher{
						"Content-Type": dsl.String("application/json"),
					},
					Body: dsl.Match(fmt.Sprintf("%s was added",task)),
				},
			)

		// verify interaction on client side
		err := pact.Verify(func() error {
			// specify host anf post of PACT Mock Server as actual server
			c := NewClient("http://localhost:8081")
			err := c.AddTodo(fmt.Sprintf("%s",task))

			return err
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	// write Contract into file
	if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}

	// stop PACT mock server
	pact.Teardown()
}