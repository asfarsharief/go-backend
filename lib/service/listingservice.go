package service

import (
	"fmt"

	"github.com/Karkhana-Studio/Backend-server/lib/payload"
	"github.com/google/uuid"
)

type ListingServiceInterface interface {
	GetDataForWomenListingPage() []payload.ProductMinis
}

type ListingService struct {
}

func NewListingService() ListingServiceInterface {
	return &ListingService{}
}

//ForIntegration only
// ProductCatalogueID = 1
// ProductCategoryID : Kaftan = 1, bottoms: 2, tops: 3, dress: 4, crochets: 5, Accessories: 6
func (ls *ListingService) GetDataForWomenListingPage() []payload.ProductMinis {
	kaftanProductMinis := payload.ProductMinis{
		ProductID:          fmt.Sprint(uuid.New()),
		ProductCatalogueID: 1,
		ProductCategoryID:  1,
		ProductName:        "Pink Kaftan",
		Cost:               300.00,
		ImageUrls: []payload.ImageUrls{
			{
				Prefrence: 1,
				ImageURL:  "https://i.ibb.co/qRtSwFm/kaftan.jpg",
			},
		},
	}

	bottomsProductMinis := payload.ProductMinis{
		ProductID:          fmt.Sprint(uuid.New()),
		ProductCatalogueID: 1,
		ProductCategoryID:  2,
		ProductName:        "Blue Jeans",
		Cost:               400.00,
		ImageUrls: []payload.ImageUrls{
			{
				Prefrence: 1,
				ImageURL:  "https://i.ibb.co/xCrPMSf/bottoms.jpg",
			},
		},
	}

	topsProductMinis := payload.ProductMinis{
		ProductID:          fmt.Sprint(uuid.New()),
		ProductCatalogueID: 1,
		ProductCategoryID:  3,
		ProductName:        "Cool Top",
		Cost:               500.00,
		ImageUrls: []payload.ImageUrls{
			{
				Prefrence: 1,
				ImageURL:  "https://i.ibb.co/RysvVVV/topImage.jpg",
			},
		},
	}

	dressProductMinis := payload.ProductMinis{
		ProductID:          fmt.Sprint(uuid.New()),
		ProductCatalogueID: 1,
		ProductCategoryID:  4,
		ProductName:        "Flower Dress",
		Cost:               900.00,
		ImageUrls: []payload.ImageUrls{
			{
				Prefrence: 1,
				ImageURL:  "https://i.ibb.co/hMh9VNh/dress.webp",
			},
		},
	}

	crochetsProductMinis := payload.ProductMinis{
		ProductID:          fmt.Sprint(uuid.New()),
		ProductCatalogueID: 1,
		ProductCategoryID:  5,
		ProductName:        "Nice Crochets",
		Cost:               200.00,
		ImageUrls: []payload.ImageUrls{
			{
				Prefrence: 1,
				ImageURL:  "https://i.ibb.co/FWk8xSj/crochet.jpg",
			},
		},
	}

	AccProductMinis := payload.ProductMinis{
		ProductID:          fmt.Sprint(uuid.New()),
		ProductCatalogueID: 1,
		ProductCategoryID:  6,
		ProductName:        "Cool Sunglassess",
		Cost:               50.00,
		ImageUrls: []payload.ImageUrls{
			{
				Prefrence: 1,
				ImageURL:  "https://i.ibb.co/3spR9Cc/accessories.jpg",
			},
		},
	}

	productMinis := []payload.ProductMinis{
		kaftanProductMinis,
		kaftanProductMinis,
		kaftanProductMinis,
		kaftanProductMinis,
		bottomsProductMinis,
		bottomsProductMinis,
		bottomsProductMinis,
		bottomsProductMinis,
		topsProductMinis,
		topsProductMinis,
		topsProductMinis,
		topsProductMinis,
		dressProductMinis,
		dressProductMinis,
		dressProductMinis,
		dressProductMinis,
		crochetsProductMinis,
		crochetsProductMinis,
		crochetsProductMinis,
		crochetsProductMinis,
		AccProductMinis,
		AccProductMinis,
		AccProductMinis,
		AccProductMinis,
	}

	return productMinis
}
