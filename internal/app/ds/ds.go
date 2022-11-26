package ds

type Users struct {
	Id_user  uint   `sql:"type:uuid;primary_key;default:" json:"Id_user" gorm:"primarykey"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Id_group string `json:"id_group"`
}

type Objests struct {
	Id_object       uint   `sql:"type:uuid;primary_key;default:" json:"Id_object" gorm:"primarykey"`
	Number          string `json:"number"`
	Floor           uint   `json:"floor"`
	Description     string `json:"description"`
	Id_coordination uint   `json:"id_coordination"`
	Id_type         uint   `json:"id_type"`
}

type Occupation struct {
	Id_occupation uint   `sql:"type:uuid;primary_key;default:" json:"Id_occupation" gorm:"primarykey"`
	Id_Group      uint   `json:"id_group"`
	Time          string `json:"time"`
}

type Groups struct {
	Id_group uint   `sql:"type:uuid;primary_key;default:" json:"Id_group" gorm:"primarykey"`
	Name     string `json:"name"`
}

type FeedBack struct {
	Id_feedback uint   `sql:"type:uuid;primary_key;default:" json:"Id_feedback" gorm:"primarykey"`
	Id_user     uint   `json:"id_user"`
	Message     string `json:"message"`
	Status      string `json:"status"`
}

type Favorites struct {
	Id_favorite uint `sql:"type:uuid;primary_key;default:" json:"Id_favorite" gorm:"primarykey"`
	Id_user     uint `json:"id_user"`
	Id_object   uint `json:"id_object"`
}

type Coordinates struct {
	Id_coordination uint `sql:"type:uuid;primary_key;default:" json:"Id_coordination" gorm:"primarykey"`
	X_coordination  int  `json:"x_coordination"`
	Y_coordination  int  `json:"y_coordination"`
	Height          uint `json:"height"`
	Width           uint `json:"width"`
}

type Type struct {
	Id_type   uint   `sql:"type:uuid;primary_key;default:" json:"Id_type" gorm:"primarykey"`
	Type_name string `json:"type_name"`
}
