package config

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// LoadHTTP inicia e gerencia o ciclo de vida de um servidor HTTP.
func LoadHTTP(ctx context.Context, container *Container, router http.Handler) {
	server := &http.Server{
		Addr:    ":" + container.Environments.PORT,
		Handler: router,
	}

	// Goroutine para iniciar o servidor
	go func() {
		container.Logger.Infof("Server started on port %s\n", container.Environments.PORT)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			container.Logger.Errorf("Error starting server: %s\n", err)
		}
	}()

	// Canal para capturar sinais do sistema operacional
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// Aguarda um sinal de interrupção
	<-sig
	container.Logger.Warn("Shutting down server...")

	// Contexto para shutdown com timeout
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Desligando o servidor
	if err := server.Shutdown(shutdownCtx); err != nil {
		container.Logger.Errorf("Server forced to shutdown: %v\n", err)
	}

	container.Logger.Warn("Server exited gracefully")
}

// Carrega adpatadores de rotas
func LoadRouter(ctx context.Context) *gin.Engine {
	// Configuração do modo de execução do Gin
	envMode := ctx.Value(ginModeKey).(ginMode)

	gin.SetMode(string(envMode))

	// Inicializa o roteador
	router := gin.New()

	// Configuração de CORS
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	// Rota de verificação de status
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Grupo de rotas
	//apiV1Router := router.Group("/v1")

	// Carrega as rotas via adpatadores internos

	return router
}
