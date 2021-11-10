package RedEnv

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func updateSnatch(env Env, usrcnt UserCount) {
	fmt.Println("usrcnt:", usrcnt)
	db.Save(&usrcnt)
	db.Create(&env)

	conn := redisPool.Get()
	defer conn.Close()

	conn.Do("rpush", usrcnt.ID, env.ID)
	conn.Do("HMSet", env.ID, "UID", env.UID, "money", env.Money, "opened", env.Opened, "getTime", env.GetTime)
	conn.Do("Set", fmt.Sprintf("%s%d", "cnt", usrcnt.ID), usrcnt.Count)
}

func getSnatch(uid int) (int, error) {
	conn := redisPool.Get()
	defer conn.Close()
	cnt, err := redis.Int(conn.Do("Get", fmt.Sprintf("%s%d", "cnt", uid)))
	if err != nil {
		var usrcnt UserCount
		result := db.First(&usrcnt, uid)
		if result.RowsAffected != 0 {
			err = nil
			cnt = usrcnt.Count
		}
	}
	return cnt, err
}

func getOpen(eid, uid int) (int, int, error) {
	rc := redisPool.Get()
	defer rc.Close()
	envVal, err := redis.Int(rc.Do("HGet", eid, "money"))
	money, err2 := redis.Int(rc.Do("Get", uid))

	if err != nil || err2 != nil{
		var env Env
		var usrMon UserMoney
		fmt.Println("eid, uid:", eid, uid)
		result := db.First(&env, eid)
		result2 := db.First(&usrMon, uid)
		fmt.Println("r1, r2:", result.RowsAffected, result2.RowsAffected)
		if result.RowsAffected != 0 && result2.RowsAffected != 0{
			err = nil
			err2 = nil
			envVal = env.Money
			money = usrMon.Money
		}
	}
	if err2 != nil {
		err = err2
	}
	return envVal, money, err
}

func updateOpen(eid, uid, newMoney int) {
	fmt.Println("eid, uid, newMoney", eid, uid, newMoney)
	usrMon := UserMoney{uid, newMoney}
	db.Save(&usrMon)

	env := Env{ID: eid, Opened: 1}
	db.Model(&env).Update("opened", 1)

	conn := redisPool.Get()
	defer conn.Close()
	conn.Do("Set", uid, newMoney)
	conn.Do("HSet", eid, "Opened", 1)

}

func getGwl(uid int) ([]Env, int, int, error) {
	rc := redisPool.Get()
	defer rc.Close()
	money, err := redis.Int(rc.Do("Get", uid))
	if err != nil {
		var usrmon UserMoney
		result := db.First(&usrmon, uid)
		if result.RowsAffected == 0 {
			return nil, 0, 0, err
		} else {
			money = usrmon.Money
		}
	}

	envs := make([]Env, 5)
	result := db.Find(&envs, "uid = ?", uid)
	if result.RowsAffected == 0 {
		return nil, 0, 0, err
	} else {
		return envs, int(result.RowsAffected), money, nil
	}
}