package models

type User struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"unique"`
	Password      []byte `json:"-"`
	WalletAddress string `json:"walletAddr" gorm:"unique"`
}

//type UserWallet struct {
//	wID           uint
// 	uId           uint
//	walletAddress uint
//}
