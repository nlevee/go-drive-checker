package drivechecker

// StoreScrapper interface to scrap a drive state and update config
type StoreScrapper interface {
	GetStoreById(storeID string) (Store, error)
	GetStoreByPostalCode(postalCode string) ([]Store, error)
	GetStoreIDByPostalCode(postalCode string) ([]string, error)
}

// Retail struct to wrapp function of a drive
type Retail struct {
	Name string
	StoreScrapper
}

// GetStoreIDByPostalCode fetch storeIDs by postal code
func (d Retail) GetStoreIDByPostalCode(postalCode string) ([]string, error) {
	storeIds := []string{}

	stores, err := d.GetStoreByPostalCode(postalCode)
	if err != nil {
		return storeIds, err
	}

	for _, v := range stores {
		storeIds = append(storeIds, v.ID)
	}

	return storeIds, nil
}
