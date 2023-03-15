package repositories

import "log"

func GetSepNames(Groepname string) ([]int, error) {
	log.Println(Groepname)
	var GroupID []int
	err := db.Table("groups").Where("groepname IN (?)", Groepname).Pluck("id", &GroupID)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	log.Println("GetSepNames", GroupID)
	return GroupID, nil
}
