package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransactionFormatter := CampaignTransactionFormatter{}
	campaignTransactionFormatter.ID = transaction.ID
	campaignTransactionFormatter.Name = transaction.User.Name
	campaignTransactionFormatter.Amount = transaction.Amount
	campaignTransactionFormatter.CreatedAt = transaction.CreatedAt

	return campaignTransactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	campaignTransactionFormatters := []CampaignTransactionFormatter{}
	for _, transaction := range transactions {
		campaignTransactionFormatter := FormatCampaignTransaction(transaction)
		campaignTransactionFormatters = append(campaignTransactionFormatters, campaignTransactionFormatter)
	}

	return campaignTransactionFormatters
}

type UserTransactionFormatter struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	Campaign  CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{}
	userTransactionFormatter.ID = transaction.ID
	userTransactionFormatter.Amount = transaction.Amount
	userTransactionFormatter.Status = transaction.Status
	userTransactionFormatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	} else {
		campaignFormatter.ImageURL = ""
	}

	userTransactionFormatter.Campaign = campaignFormatter

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	userTransactionFormatters := []UserTransactionFormatter{}
	for _, transaction := range transactions {
		userTransactionFormatter := FormatUserTransaction(transaction)
		userTransactionFormatters = append(userTransactionFormatters, userTransactionFormatter)
	}

	return userTransactionFormatters
}