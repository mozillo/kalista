package tenants

import (
	"kalista/models"
	"kalista/serializers"

	"github.com/gin-gonic/gin"
)

func Show(ctx *gin.Context) {
	var tenant = &models.Tenant{}
	models.DB().Model(tenant).First(tenant)

	ctx.JSON(200, gin.H{
		"data": serializers.SingleSerializer(tenant, &models.AdminTenantShowSerializer{}),
	})
}