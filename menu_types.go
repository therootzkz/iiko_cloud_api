package iiko

import (
	"time"

	"github.com/google/uuid"
)

type GetMenuRequest struct {
	OrganizationId uuid.UUID `json:"organizationId"`
	StartRevision  *int      `json:"startRevision,omitempty"`
}

type GetMenuResponse struct {
	CorrelationId     uuid.UUID             `json:"correlationId"`
	Groups            []ProductsGroupInfo   `json:"groups"`
	ProductCategories []ProductCategoryInfo `json:"productCategories"`
	Products          []ProductInfo         `json:"products"`
	Sizes             []Size                `json:"sizes"`
	Revision          int                   `json:"revision"`
}

type GetMenuErrorDetails struct {
	CorrelationId    uuid.UUID `json:"correlationId"`
	ErrorDescription string    `json:"errorDescription"`
	ErrorCode        *string   `json:"error,omitempty"`
}

type ProductsGroupInfo struct {
	ImageLinks       []string   `json:"imageLinks"`
	ParentGroup      *uuid.UUID `json:"parentGroup,omitempty"`
	Order            int        `json:"order"`
	IsIncludedInMenu bool       `json:"isIncludedInMenu"`
	IsGroupModifier  bool       `json:"isGroupModifier"`
	Id               uuid.UUID  `json:"id"`
	Code             *string    `json:"code,omitempty"`
	Name             string     `json:"name"`
	Description      *string    `json:"description,omitempty"`
	AdditionalInfo   *string    `json:"additionalInfo,omitempty"`
	Tags             *[]string  `json:"tags,omitempty"`
	IsDeleted        bool       `json:"isDeleted"`
	SeoDescription   *string    `json:"seoDescription,omitempty"`
	SeoText          *string    `json:"seoText,omitempty"`
	SeoKeywords      *string    `json:"seoKeywords,omitempty"`
	SeoTitle         *string    `json:"seoTitle,omitempty"`
}

type ProductCategoryInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
}
type ProductInfo struct {
	FatAmount               *float64             `json:"fatAmount,omitempty"`
	ProteinsAmount          *float64             `json:"proteinsAmount,omitempty"`
	CarbohydratesAmount     *float64             `json:"carbohydratesAmount,omitempty"`
	EnergyAmount            *float64             `json:"energyAmount,omitempty"`
	FatFullAmount           *float64             `json:"fatFullAmount,omitempty"`
	ProteinsFullAmount      *float64             `json:"proteinsFullAmount,omitempty"`
	CarbohydratesFullAmount *float64             `json:"carbohydratesFullAmount,omitempty"`
	EnergyFullAmount        *float64             `json:"energyFullAmount,omitempty"`
	Weight                  *float64             `json:"weight,omitempty"`
	GroupId                 *uuid.UUID           `json:"groupId,omitempty"`
	ProductCategoryId       *uuid.UUID           `json:"productCategoryId,omitempty"`
	Type                    *string              `json:"type,omitempty"`
	OrderItemType           string               `json:"orderItemType"`
	ModifierSchemaId        *uuid.UUID           `json:"modifierSchemaId,omitempty"`
	ModifierSchemaName      *string              `json:"modifierSchemaName,omitempty"`
	Splittable              bool                 `json:"splittable"`
	MeasureUnit             string               `json:"measureUnit"`
	SizePrices              []SizePrice          `json:"sizePrices"`
	Modifiers               []SimpleModifierInfo `json:"modifiers"`
	GroupModifiers          []GroupModifierInfo  `json:"groupModifiers"`
	ImageLinks              []string             `json:"imageLinks"`
	DoNotPrintInCheque      bool                 `json:"doNotPrintInCheque"`
	ParentGroup             *uuid.UUID           `json:"parentGroup,omitempty"`
	Order                   int                  `json:"order"`
	FullNameEnglish         *string              `json:"fullNameEnglish,omitempty"`
	UseBalanceForSell       bool                 `json:"useBalanceForSell"`
	CanSetOpenPrice         bool                 `json:"canSetOpenPrice"`
	Id                      uuid.UUID            `json:"id"`
	Code                    *string              `json:"code,omitempty"`
	Name                    string               `json:"name"`
	Description             *string              `json:"description,omitempty"`
	AdditionalInfo          *string              `json:"additionalInfo,omitempty"`
	Tags                    *[]string            `json:"tags,omitempty"`
	IsDeleted               bool                 `json:"isDeleted"`
	SeoDescription          *string              `json:"seoDescription,omitempty"`
	SeoText                 *string              `json:"seoText,omitempty"`
	SeoKeywords             *string              `json:"seoKeywords,omitempty"`
	SeoTitle                *string              `json:"seoTitle,omitempty"`
}

type SizePrice struct {
	SizeId *uuid.UUID `json:"sizeId,omitempty"`
	Price  struct {
		CurrentPrice       float64    `json:"currentPrice"`
		IsIncludedInMenu   bool       `json:"isIncludedInMenu"`
		NextPrice          *float64   `json:"nextPrice,omitempty"`
		NextIncludedInMenu bool       `json:"nextIncludedInMenu"`
		NextDatePrice      *time.Time `json:"nextDatePrice"`
	} `json:"price"`
}

type SimpleModifierInfo struct {
	Id                  uuid.UUID `json:"id"`
	DefaultAmount       *int      `json:"defaultAmount,omitempty"`
	MinAmount           int       `json:"minAmount"`
	MaxAmount           int       `json:"maxAmount"`
	Required            *bool     `json:"required,omitempty"`
	HideIfDefaultAmount *bool     `json:"hideIfDefaultAmount,omitempty"`
	Splittable          *bool     `json:"splittable,omitempty"`
	FreeOfChargeAmount  *int      `json:"freeOfChargeAmount,omitempty"`
}

type GroupModifierInfo struct {
	Id                                   uuid.UUID           `json:"id"`
	MinAmount                            int                 `json:"minAmount"`
	MaxAmount                            int                 `json:"maxAmount"`
	Required                             *bool               `json:"required,omitempty"`
	ChildModifiersHaveMinMaxRestrictions *bool               `json:"childModifiersHaveMinMaxRestrictions,omitempty"`
	ChildModifiers                       []ChildModifierInfo `json:"childModifiers"`
	HideIfDefaultAmount                  *bool               `json:"hideIfDefaultAmount,omitempty"`
	DefaultAmount                        *int                `json:"defaultAmount,omitempty"`
	Splittable                           *bool               `json:"splittable,omitempty"`
	FreeOfChargeAmount                   *int                `json:"freeOfChargeAmount,omitempty"`
}

type ChildModifierInfo struct {
	Id                  uuid.UUID `json:"id"`
	DefaultAmount       *int      `json:"defaultAmount,omitempty"`
	MinAmount           int       `json:"minAmount"`
	MaxAmount           int       `json:"maxAmount"`
	Required            *bool     `json:"required,omitempty"`
	HideIfDefaultAmount *bool     `json:"hideIfDefaultAmount,omitempty"`
	Splittable          *bool     `json:"splittable,omitempty"`
	FreeOfChargeAmount  *int      `json:"freeOfChargeAmount,omitempty"`
}

type Size struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Priority  *int      `json:"priority,omitempty"`
	IsDefault *bool     `json:"isDefault,omitempty"`
}
