package repositories

import "log"

func GetSepNames(Groupname string) []int {
	var GroupID []int
	err := db.Table("groups").Where("groepname = ?", Groupname).Pluck("id", &GroupID)
	log.Println(err, GroupID)
	return GroupID
}
