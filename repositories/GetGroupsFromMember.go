package repositories

import "log"

func GetGroupsFromMembers(userId int) ([]int, error) {
	var GroupID []int
	err := db.Table("groupmembers").Where("user_id = ?", userId).Pluck("groep_id", &GroupID)
	log.Println(err, GroupID)
	return GroupID, nil
}
