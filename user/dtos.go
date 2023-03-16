package user

type TransferMoneyDto struct {
	From   uint `json:"from" binding:"required,numeric"`
	To     uint `json:"to" binding:"required,numeric"`
	Amount int  `json:"amount" binding:"required,numeric,gte=0,lte=999999999"`
}

type SaveDto struct {
	ID    uint   `json:"id" binding:"numeric"`
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Money int    `json:"money" binding:"required,gte=0,lte=999999999"`
	Age   int    `json:"age" binding:"required,gte=16,lte=100"`
}
