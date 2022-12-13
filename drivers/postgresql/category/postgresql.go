package category

import (
	"PPOB_BACKEND/businesses/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	conn *gorm.DB
}

func NewCategoryRepository(conn *gorm.DB) category.Repository {
	return &categoryRepository{
		conn: conn,
	}
}

func (cr *categoryRepository) GetAll(CategoryName string) []category.Domain {
	var cat []Category

	if CategoryName != "" {
		cr.conn.Preload("ProductType").Find(&cat, "name = ?", CategoryName)
	} else {
		cr.conn.Preload("ProductType").Find(&cat)
	}

	categoryDomain := []category.Domain{}
	for _, category := range cat {
		categoryDomain = append(categoryDomain, category.ToDomain())
	}

	return categoryDomain
}
func (cr *categoryRepository) Create(categoryDomain *category.Domain) category.Domain {
	cat := FromDomain(categoryDomain)

	cr.conn.Create(&cat)

	return cat.ToDomain()
}
func (cr *categoryRepository) GetDetail(category_id int) category.Domain {
	var cat Category

	cr.conn.First(&cat, category_id)

	return cat.ToDomain()
}
func (cr *categoryRepository) Update(categoryDomain *category.Domain, category_id int) category.Domain {
	cat := FromDomain(categoryDomain)

	cr.conn.Model(&cat).Where("id = ?", category_id).Updates(
		Category{
			Name: categoryDomain.Name,
		},
	)

	return cat.ToDomain()
}
func (cr *categoryRepository) Delete(category_id int) (category.Domain, error) {
	var cat Category

	err := cr.conn.First(&cat, category_id).Error

	if err != nil {
		return cat.ToDomain(), err
	}

	cr.conn.Unscoped().Where("id = ?", category_id).Delete(&cat)

	return cat.ToDomain(), nil
}
