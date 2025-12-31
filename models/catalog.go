package models

import (
	"time"

	"gorm.io/gorm"
)

type Catalog struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"default:''" json:"title"`
	UserID    uint      `gorm:"index" json:"user_id"`
	SortBy    int       `gorm:"default:0;not null" json:"sort_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Links     []Link    `gorm:"foreignKey:CatalogID" json:"links,omitempty"`
}

func GetCatalogsByUserID(userID uint) ([]Catalog, error) {
	var catalogs []Catalog
	result := DB.Where("user_id = ?", userID).Order("sort_by asc").Preload("Links", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_by asc")
	}).Find(&catalogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return catalogs, nil
}

func GetCatalogByIDAndUserID(id, userID uint) (*Catalog, error) {
	var catalog Catalog
	result := DB.Where("id = ? AND user_id = ?", id, userID).First(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &catalog, nil
}

func CreateCatalog(title string, sortBy int, userID uint) (*Catalog, error) {
	catalog := Catalog{
		Title:  title,
		SortBy: sortBy,
		UserID: userID,
	}
	result := DB.Create(&catalog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &catalog, nil
}

func (c *Catalog) Save() error {
	return DB.Save(c).Error
}

func (c *Catalog) Delete() error {
	// Delete all links in this catalog first
	if err := DB.Where("catalog_id = ?", c.ID).Delete(&Link{}).Error; err != nil {
		return err
	}
	return DB.Delete(c).Error
}
