package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"river0825/cleanarchitecture/core/arch/example/application"
	"river0825/cleanarchitecture/core/arch/example/port/repository/inmem"
	gin2 "river0825/cleanarchitecture/core/arch/port/gin"
	"river0825/cleanarchitecture/core/dependency/log"
)

func main() {
	// using generic repository
	repo := inmem.NewRoomRepository()
	facade := application.NewAFacade(repo)

	g := gin.Default()
	g.GET("/join", gin2.HandlerFunc[application.JoinCommand](facade))
	g.GET("/leave", gin2.HandlerFunc[application.LeaveCommand](facade))
	g.GET("/error", gin2.HandlerFunc[application.ErrorCommand](facade))

	s := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	/**
	zerolog provides a default stack trace marshaling function.
	but the file path is not included in the stack trace.
	So we need to provide our own stack trace marshaling function.
	You could set the Env=local to enable the file path in the stack trace.
	And `curl localhost:8080/error` to see the stack trace.
	*/
	if os.Getenv("Env") == "local" {
		zerolog.ErrorStackMarshaler = log.DebugMarshalStack
	} else {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
