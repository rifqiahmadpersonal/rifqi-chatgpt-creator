package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/verssache/chatgpt-creator/internal/api/handlers"
	"github.com/verssache/chatgpt-creator/internal/config"
)

func SetupRoutes(router *gin.RouterGroup, db *sqlx.DB, logger *logrus.Logger, cfg *config.EnvConfig) {
	h := handlers.NewHandlers(db, logger, cfg)

	accounts := router.Group("/accounts")
	{
		accounts.GET("", h.ListAccounts)
		accounts.GET("/:id", h.GetAccount)
		accounts.POST("", h.CreateAccount)
		accounts.DELETE("/:id", h.DeleteAccount)
		accounts.GET("/export", h.ExportAccounts)
	}

	emailDomains := router.Group("/email-domains")
	{
		emailDomains.GET("", h.ListEmailDomains)
		emailDomains.GET("/:id", h.GetEmailDomain)
		emailDomains.POST("", h.CreateEmailDomain)
		emailDomains.PUT("/:id", h.UpdateEmailDomain)
		emailDomains.DELETE("/:id", h.DeleteEmailDomain)
		emailDomains.POST("/:id/check", h.CheckEmailDomainHealth)
	}

	batchJobs := router.Group("/batch-jobs")
	{
		batchJobs.GET("", h.ListBatchJobs)
		batchJobs.GET("/:id", h.GetBatchJob)
		batchJobs.POST("", h.CreateBatchJob)
		batchJobs.POST("/:id/start", h.StartBatchJob)
		batchJobs.POST("/:id/stop", h.StopBatchJob)
		batchJobs.GET("/:id/attempts", h.GetBatchJobAttempts)
	}

	configurations := router.Group("/configurations")
	{
		configurations.GET("", h.ListConfigurations)
		configurations.GET("/:key", h.GetConfiguration)
		configurations.PUT("/:key", h.UpdateConfiguration)
	}

	blacklistedDomains := router.Group("/blacklisted-domains")
	{
		blacklistedDomains.GET("", h.ListBlacklistedDomains)
		blacklistedDomains.POST("", h.CreateBlacklistedDomain)
		blacklistedDomains.DELETE("/:id", h.DeleteBlacklistedDomain)
	}

	stats := router.Group("/stats")
	{
		stats.GET("/dashboard", h.GetDashboardStats)
	}
}
