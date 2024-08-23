package doses

import (
	"net/http"
	"strconv"
	"time"
)
import "github.com/gin-gonic/gin"

func GetAllDoses(repo DoseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		doses, err := repo.FindAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, doses)
	}
}

func GetAllDosesToday(repo DoseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		doses, err := repo.FindAllToday()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, doses)
	}
}

func AddDose(repo DoseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var dose Dose
		if err := c.BindJSON(&dose); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dose, err := repo.Add(dose)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dose)
	}
}

func AddDoseNow(repo DoseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		dose := Dose{DateTaken: time.Now()}
		dose, err := repo.Add(dose)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dose)
	}
}

func UpdateDose(repo DoseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var dose Dose
		if err := c.BindJSON(&dose); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dose, err := repo.Update(dose)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dose)
	}
}

func DeleteDose(repo DoseRepository) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		doseId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = repo.Delete(doseId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Dose deleted"})
	}
}
