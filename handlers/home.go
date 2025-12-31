package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liubaicai/navi-home/models"
)

func Index(c *gin.Context) {
	user := getCurrentUser(c)

	var catalogsJSON string
	if user != nil {
		catalogs, err := models.GetCatalogsByUserID(user.ID)
		if err != nil {
			catalogs = []models.Catalog{}
		}
		catalogsBytes, _ := json.Marshal(catalogs)
		catalogsJSON = string(catalogsBytes)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"User":     user,
		"Catalogs": catalogsJSON,
	})
}

func Search(c *gin.Context) {
	q := c.Query("q")
	if q == "" {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")
	} else {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com/search?q="+url.QueryEscape(q))
	}
}

type CatalogData struct {
	ID    int        `json:"id"`
	Title string     `json:"title"`
	Links []LinkData `json:"links"`
}

type LinkData struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Icon      string `json:"icon"`
	CatalogID int    `json:"catalog_id"`
}

func SetCatalogs(c *gin.Context) {
	user := getCurrentUser(c)
	if user == nil {
		c.String(http.StatusOK, "not ok")
		return
	}

	var requestData struct {
		Data string `json:"data"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.String(http.StatusOK, "not ok")
		return
	}

	var catalogs []CatalogData
	if err := json.Unmarshal([]byte(requestData.Data), &catalogs); err != nil {
		c.String(http.StatusOK, "not ok")
		return
	}

	// Get existing catalogs
	existingCatalogs, _ := models.GetCatalogsByUserID(user.ID)
	existingCatalogIDs := make(map[int]bool)
	for _, cat := range existingCatalogs {
		existingCatalogIDs[int(cat.ID)] = true
	}

	// Track which catalog IDs are in the new data
	newCatalogIDs := make(map[int]bool)
	for _, item := range catalogs {
		if item.ID > 0 {
			newCatalogIDs[item.ID] = true
		}
	}

	// Delete catalogs that are not in the new data
	for _, cat := range existingCatalogs {
		if !newCatalogIDs[int(cat.ID)] {
			cat.Delete()
		}
	}

	// Process catalogs
	for index, item := range catalogs {
		var catalog *models.Catalog
		var err error

		if item.ID > 0 {
			catalog, err = models.GetCatalogByIDAndUserID(uint(item.ID), user.ID)
			if err != nil {
				// Catalog doesn't exist, create new one
				catalog, err = models.CreateCatalog(item.Title, index+1, user.ID)
				if err != nil {
					continue
				}
			} else {
				catalog.Title = item.Title
				catalog.SortBy = index + 1
				catalog.Save()
			}
		} else {
			// New catalog (zero or negative ID)
			catalog, err = models.CreateCatalog(item.Title, index+1, user.ID)
			if err != nil {
				continue
			}
		}

		// Get existing links for this catalog
		existingLinks := make(map[int]bool)
		for _, link := range catalog.Links {
			existingLinks[int(link.ID)] = true
		}

		// Track which link IDs are in the new data
		newLinkIDs := make(map[int]bool)
		for _, link := range item.Links {
			if link.ID > 0 {
				newLinkIDs[link.ID] = true
			}
		}

		// Delete links that are not in the new data
		for _, link := range catalog.Links {
			if !newLinkIDs[int(link.ID)] {
				link.Delete()
			}
		}

		// Process links
		for linkIndex, linkData := range item.Links {
			if linkData.ID > 0 {
				link, err := models.GetLinkByIDAndUserID(uint(linkData.ID), user.ID)
				if err != nil {
					// Link doesn't exist, create new one
					models.CreateLink(linkData.Title, linkData.URL, linkData.Icon, linkIndex+1, user.ID, catalog.ID)
				} else {
					link.Title = linkData.Title
					link.URL = linkData.URL
					link.Icon = linkData.Icon
					link.SortBy = linkIndex + 1
					link.CatalogID = catalog.ID
					link.Save()
				}
			} else {
				// New link (zero or negative ID)
				models.CreateLink(linkData.Title, linkData.URL, linkData.Icon, linkIndex+1, user.ID, catalog.ID)
			}
		}
	}

	c.String(http.StatusOK, "ok")
}

func TestIconURL(c *gin.Context) {
	iconURL := c.Query("url")
	if iconURL == "" {
		c.String(http.StatusOK, "/ie.png")
		return
	}

	parsedURL, err := url.Parse(iconURL)
	if err != nil {
		c.String(http.StatusOK, "/ie.png")
		return
	}

	iconsPath := filepath.Join("public", "icons")
	if err := os.MkdirAll(iconsPath, 0755); err != nil {
		c.String(http.StatusOK, "/ie.png")
		return
	}

	iconFileName := strings.ReplaceAll(parsedURL.Host, ".", "_") + ".ico"
	iconPath := filepath.Join(iconsPath, iconFileName)

	// Check if icon already exists
	if _, err := os.Stat(iconPath); err == nil {
		c.String(http.StatusOK, "/icons/"+iconFileName)
		return
	}

	// Download the icon
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", iconURL, nil)
	if err != nil {
		c.String(http.StatusOK, "/ie.png")
		return
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		c.String(http.StatusOK, "/ie.png")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 || strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		c.String(http.StatusOK, "/ie.png")
		return
	}

	// Save the icon
	file, err := os.Create(iconPath)
	if err != nil {
		c.String(http.StatusOK, "/ie.png")
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		os.Remove(iconPath)
		c.String(http.StatusOK, "/ie.png")
		return
	}

	c.String(http.StatusOK, "/icons/"+iconFileName)
}
