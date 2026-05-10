package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

	"github.com/verssache/chatgpt-creator/internal/config"
	"github.com/verssache/chatgpt-creator/internal/models"
	"github.com/verssache/chatgpt-creator/internal/repository"
)

type Handlers struct {
	db       *sqlx.DB
	logger   *logrus.Logger
	cfg      *config.EnvConfig
	repos    *Repositories
}

type Repositories struct {
	accounts            repository.AccountRepository
	emailDomains        repository.EmailDomainRepository
	batchJobs           repository.BatchJobRepository
	configurations      repository.ConfigurationRepository
	blacklistedDomains  repository.BlacklistedDomainRepository
	registrationAttempts repository.RegistrationAttemptRepository
}

func NewHandlers(db *sqlx.DB, logger *logrus.Logger, cfg *config.EnvConfig) *Handlers {
	return &Handlers{
		db:     db,
		logger: logger,
		cfg:    cfg,
		repos: &Repositories{
			accounts:            repository.NewAccountRepository(db),
			emailDomains:        repository.NewEmailDomainRepository(db),
			batchJobs:           repository.NewBatchJobRepository(db),
			configurations:      repository.NewConfigurationRepository(db),
			blacklistedDomains:  repository.NewBlacklistedDomainRepository(db),
			registrationAttempts: repository.NewRegistrationAttemptRepository(db),
		},
	}
}

func Init() {}

// ==================== ACCOUNTS ====================

func (h *Handlers) ListAccounts(c *gin.Context) {
	accounts, err := h.repos.accounts.List(c.Request.Context(), false)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list accounts")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list accounts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (h *Handlers) GetAccount(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	account, err := h.repos.accounts.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account": account})
}

func (h *Handlers) CreateAccount(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := &models.Account{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  req.Password,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.repos.accounts.Create(c.Request.Context(), account); err != nil {
		h.logger.WithError(err).Error("Failed to create account")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account": account})
}

func (h *Handlers) DeleteAccount(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	if err := h.repos.accounts.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (h *Handlers) ExportAccounts(c *gin.Context) {
	accounts, err := h.repos.accounts.List(c.Request.Context(), false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to export accounts"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

// ==================== EMAIL DOMAINS ====================

func (h *Handlers) ListEmailDomains(c *gin.Context) {
	domains, err := h.repos.emailDomains.List(c.Request.Context(), false)
	if err != nil {
		h.logger.WithError(err).Error("Failed to list email domains")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list email domains"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"domains": domains})
}

func (h *Handlers) GetEmailDomain(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid domain id"})
		return
	}

	domain, err := h.repos.emailDomains.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "domain not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"domain": domain})
}

func (h *Handlers) CreateEmailDomain(c *gin.Context) {
	var req struct {
		Domain   string `json:"domain" binding:"required"`
		Priority int    `json:"priority"`
		Source   string `json:"source"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Source == "" {
		req.Source = models.DomainSourceGenerator
	}
	if req.Priority == 0 {
		req.Priority = 50
	}

	domain := &models.EmailDomain{
		ID:           uuid.New().String(),
		Domain:       req.Domain,
		Priority:     req.Priority,
		IsActive:     true,
		Source:       req.Source,
		HealthStatus: models.DomainHealthStatusUnknown,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := h.repos.emailDomains.Create(c.Request.Context(), domain); err != nil {
		h.logger.WithError(err).Error("Failed to create email domain")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create email domain"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"domain": domain})
}

func (h *Handlers) UpdateEmailDomain(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid domain id"})
		return
	}

	var req struct {
		Domain   string `json:"domain"`
		Priority int    `json:"priority"`
		IsActive *bool  `json:"is_active"`
		Source   string `json:"source"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain, err := h.repos.emailDomains.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "domain not found"})
		return
	}

	if req.Domain != "" {
		domain.Domain = req.Domain
	}
	if req.Priority > 0 {
		domain.Priority = req.Priority
	}
	if req.IsActive != nil {
		domain.IsActive = *req.IsActive
	}
	if req.Source != "" {
		domain.Source = req.Source
	}
	domain.UpdatedAt = time.Now()

	if err := h.repos.emailDomains.Update(c.Request.Context(), domain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update domain"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"domain": domain})
}

func (h *Handlers) DeleteEmailDomain(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid domain id"})
		return
	}

	if err := h.repos.emailDomains.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "domain not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

func (h *Handlers) CheckEmailDomainHealth(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid domain id"})
		return
	}

	domain, err := h.repos.emailDomains.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "domain not found"})
		return
	}

	now := time.Now()
	domain.LastChecked = &now
	domain.HealthStatus = models.DomainHealthStatusHealthy
	domain.UpdatedAt = now

	if err := h.repos.emailDomains.Update(c.Request.Context(), domain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update health status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"health_status": domain.HealthStatus})
}

// ==================== BATCH JOBS ====================

func (h *Handlers) ListBatchJobs(c *gin.Context) {
	jobs, err := h.repos.batchJobs.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list batch jobs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

func (h *Handlers) GetBatchJob(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	job, err := h.repos.batchJobs.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"job": job})
}

func (h *Handlers) CreateBatchJob(c *gin.Context) {
	var req struct {
		TargetCount int `json:"target_count" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job := &models.BatchJob{
		ID:          uuid.New().String(),
		TargetCount: req.TargetCount,
		Status:      models.BatchJobStatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := h.repos.batchJobs.Create(c.Request.Context(), job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create batch job"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"job": job})
}

func (h *Handlers) StartBatchJob(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	job, err := h.repos.batchJobs.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
		return
	}

	job.Status = models.BatchJobStatusRunning
	job.UpdatedAt = time.Now()

	if err := h.repos.batchJobs.Update(c.Request.Context(), job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "running"})
}

func (h *Handlers) StopBatchJob(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	job, err := h.repos.batchJobs.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "job not found"})
		return
	}

	job.Status = models.BatchJobStatusStopped
	job.UpdatedAt = time.Now()

	if err := h.repos.batchJobs.Update(c.Request.Context(), job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to stop job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "stopped"})
}

func (h *Handlers) GetBatchJobAttempts(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid job id"})
		return
	}

	attempts, err := h.repos.registrationAttempts.GetByBatchJobID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get attempts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attempts": attempts})
}

// ==================== CONFIGURATIONS ====================

func (h *Handlers) ListConfigurations(c *gin.Context) {
	configs, err := h.repos.configurations.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list configurations"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"configurations": configs})
}

func (h *Handlers) GetConfiguration(c *gin.Context) {
	key := c.Param("key")

	config, err := h.repos.configurations.GetByKey(c.Request.Context(), key)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"key": key, "value": ""})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": config.Key, "value": config.Value})
}

func (h *Handlers) UpdateConfiguration(c *gin.Context) {
	key := c.Param("key")

	var req struct {
		Value string `json:"value"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config := &models.Configuration{
		Key:       key,
		Value:     req.Value,
		UpdatedAt: time.Now(),
	}

	if err := h.repos.configurations.Upsert(c.Request.Context(), config); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update configuration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": req.Value})
}

// ==================== BLACKLISTED DOMAINS ====================

func (h *Handlers) ListBlacklistedDomains(c *gin.Context) {
	domains, err := h.repos.blacklistedDomains.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list blacklisted domains"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"domains": domains})
}

func (h *Handlers) CreateBlacklistedDomain(c *gin.Context) {
	var req struct {
		Domain string `json:"domain" binding:"required"`
		Reason string `json:"reason"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	domain := &models.BlacklistedDomain{
		ID:        uuid.New().String(),
		Domain:    req.Domain,
		Reason:    req.Reason,
		CreatedAt: time.Now(),
	}

	if err := h.repos.blacklistedDomains.Create(c.Request.Context(), domain); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create blacklisted domain"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"domain": domain})
}

func (h *Handlers) DeleteBlacklistedDomain(c *gin.Context) {
	id := c.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid domain id"})
		return
	}

	if err := h.repos.blacklistedDomains.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "domain not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": true})
}

// ==================== STATS ====================

func (h *Handlers) GetDashboardStats(c *gin.Context) {
	ctx := c.Request.Context()

	accounts, _ := h.repos.accounts.List(ctx, false)
	activeAccounts := 0
	for _, a := range accounts {
		if a.IsActive {
			activeAccounts++
		}
	}

	domains, _ := h.repos.emailDomains.List(ctx, true)
	activeDomains := len(domains)

	jobs, _ := h.repos.batchJobs.List(ctx)
	runningJobs := 0
	for _, j := range jobs {
		if j.Status == models.BatchJobStatusRunning {
			runningJobs++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"total_accounts":       len(accounts),
		"active_accounts":      activeAccounts,
		"total_batch_jobs":     len(jobs),
		"running_batch_jobs":   runningJobs,
		"active_email_domains": activeDomains,
	})
}