package models

import (
	"time"
)

type Link struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Title     string    `gorm:"default:''" json:"title"`
	URL       string    `gorm:"default:''" json:"url"`
	Icon      string    `gorm:"default:''" json:"icon"`
	UserID    uint      `gorm:"index" json:"user_id"`
	CatalogID uint      `gorm:"index" json:"catalog_id"`
	SortBy    int       `gorm:"default:0;not null" json:"sort_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetLinksByUserID(userID uint) ([]Link, error) {
	var links []Link
	result := DB.Where("user_id = ?", userID).Order("sort_by asc").Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}

func GetLinkByIDAndUserID(id, userID uint) (*Link, error) {
	var link Link
	result := DB.Where("id = ? AND user_id = ?", id, userID).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func CreateLink(title, url, icon string, sortBy int, userID, catalogID uint) (*Link, error) {
	link := Link{
		Title:     title,
		URL:       url,
		Icon:      icon,
		SortBy:    sortBy,
		UserID:    userID,
		CatalogID: catalogID,
	}
	result := DB.Create(&link)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (l *Link) Save() error {
	return DB.Save(l).Error
}

func (l *Link) Delete() error {
	return DB.Delete(l).Error
}

func DeleteLinksByCatalogID(catalogID uint) error {
	return DB.Where("catalog_id = ?", catalogID).Delete(&Link{}).Error
}
