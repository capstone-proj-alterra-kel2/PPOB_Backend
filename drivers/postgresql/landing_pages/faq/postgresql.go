package faq

import (
	"PPOB_BACKEND/businesses/landing_pages/faq"

	"gorm.io/gorm"
)

type faqRepository struct {
	conn *gorm.DB
}

func NewFAQRepository(conn *gorm.DB) faq.Repository {
	return &faqRepository{
		conn: conn,
	}
}

func (fr *faqRepository) GetAll() []faq.Domain {
	var faqData []FAQ

	fr.conn.Find(&faqData)

	faqFromDomain := []faq.Domain{}
	for _, result := range faqData {
		faqFromDomain = append(faqFromDomain, result.ToDomain())
	}

	return faqFromDomain
}
func (fr *faqRepository) Create(faqDomain *faq.Domain) faq.Domain {
	faq := FromDomain(faqDomain)

	fr.conn.Create(&faq)

	return faq.ToDomain()
}
func (fr *faqRepository) Update(faqDomain *faq.Domain, faq_id int) (faq.Domain, error) {
	faqReq := FromDomain(faqDomain)

	err := fr.conn.First(&faqReq, faq_id).Error

	if err != nil {
		return faqReq.ToDomain(), err
	}

	fr.conn.Model(&faqReq).Where("id = ?", faq_id).Updates(FAQ{
		Question: faqDomain.Question,
		Answer:   faqDomain.Answer,
	})

	return faqReq.ToDomain(), nil
}
func (fr *faqRepository) Delete(faq_id int) (faq.Domain, error) {
	var faqData FAQ

	err := fr.conn.First(&faqData, faq_id).Error

	if err != nil {
		return faqData.ToDomain(), err
	}

	fr.conn.Unscoped().Where("id = ?", faq_id).Delete(&faqData)

	return faqData.ToDomain(), nil
}
