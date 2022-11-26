package repository

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hac/internal/app/ds"
)

type Repository struct {
	db *gorm.DB
}

func New(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) CreateObject(object *ds.Objects) error {
	err := r.db.Create(object).Error
	return err
}

func (r *Repository) GetObjectByFloor(floor uint) ([]ds.Objects, error) {
	
	objects := []ds.Objects{}
	result := r.db.Find(&objects)
	err := r.db.Find(objects, "floor = ?", floor).Error
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (r *Repository) GetObjectById(id uint) (*ds.Objects, error) {
	object := &ds.Objects{}
	err := r.db.First(object, "id_object = ?", id).Error
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (r *Repository) GetFavoriteByID(id_user uint) ([]ds.Favorites, error) {
	
	favorites := []ds.Favorites{}
	result := r.db.Find(&favorites)
	err := r.db.Find(favorites, "id_user = ?", id_user).Error
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func (r *Repository) CreateFavorite(favorite *ds.Favorites) error {
	err := r.db.Create(favorite).Error
	return err
}


func (r *Repository) DeleteFavorite(id_favorite uint) error {
	err := r.db.First(&ds.Favorites{}, "id_favorite = ?", id_favorite).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&ds.Favorites{}, "id_favorite = ?", id_favorite).Error
	return err
}

func (r *Repository) CreateFeedBack(feedback *ds.Feedback) error {
	err := r.db.Create(feedback).Error
	return err
}


func (r *Repository) GetFeedbackByID(id_user uint) ([]ds.Feedback, error) {
	
	feedback := []ds.Feedback{}
	result := r.db.Find(&feedback)
	err := r.db.Find(feedback, "id_user = ?", id_user).Error
	if err != nil {
		return nil, err
	}
	return feedback, nil
}

func (r *Repository) GetOccupationByID(id_group uint) ([]ds.Feedback, error) {
	
	occupation := []ds.Occupation{}
	result := r.db.Find(&occupation)
	err := r.db.Find(occupation, "id_group = ?", id_group).Error
	if err != nil {
		return nil, err
	}
	return occupation, nil
}


/*func (r *Repository) ChangeProduct(id uint, new_price uint) error {
	err := r.db.Model(&ds.Goods{}).Where("id_good = ?", id).Update("price", new_price).Error
	return err
}
*/

//Users

func (r *Repository) CreateUser(user *ds.Users) error {
	err := r.db.Create(user).Error
	return err
}

func (r *Repository) LoginCheck(user *ds.Users) error {
	user_db := ds.Users{}
	err := r.db.Model(&ds.Users{}).Where("login = ?", user.Login).Take(&user_db).Error
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user_db.Password), []byte(user.Password))
	if err != nil {
		return err
	}
	user.Id_user = user_db.Id_user
	return nil
}

func (r *Repository) GetUserByID(id uint) (*ds.Users, error) {
	user := &ds.Users{}
	err := r.db.First(user, "id_user = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) GetIdByLogin(login string) (uint, error) {
	user := &ds.Users{}
	err := r.db.First(user, "login = ?", login).Error
	if err != nil {
		return 0, err
	}
	return user.Id_user, nil
}

func (r *Repository) CreateBasketRow(basket_row *ds.Basket) error {
	err := r.db.Create(basket_row).Error
	return err
}

func (r *Repository) GetBasket(id_user uint) ([]ds.Basket, error) {
	var basket []ds.Basket
	result := r.db.Find(&basket, "id_user = ?", id_user)
	if result.Error != nil {
		return nil, result.Error
	}
	return basket, nil
}
