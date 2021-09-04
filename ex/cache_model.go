package ex

//cache example

type Stu struct {
	ID int64 `json:"id" db:"id"`
}

//func main() {
//	ids := []int64{1, 2, 3, 4, 5, 200}
//	ret, err := GetInfo(ids)
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	for _, v := range ret {
//		b, _ := json.Marshal(v)
//		fmt.Println("ret=", string(b))
//	}
//
//}

//时间复杂度n~2n
//func GetInfo(ids []int64) ([]*Stu, error) {
//
//	ret := make([]*Stu, len(ids))
//
//	rvals := ctl.GetRedis(ids)
//	mpIDIndex := map[int64]int{}
//	sqlids := []int64{}
//	for i, v := range rvals {
//		mpIDIndex[ids[i]] = i
//		if v.IsNil() {
//			sqlids = append(sqlids, ids[i])
//		} else {
//			tmp := &Stu{}
//			json.Unmarshal([]byte(v.String()), tmp)
//			ret[i] = tmp
//		}
//	}
//	if len(sqlids) == 0 {
//		return ret, nil
//	}
//	dbData, err := ctl.SelectSql(sqlids)
//	if err != nil {
//		//有错误了
//		for _, id := range sqlids {
//			if ret[mpIDIndex[id]] == nil {
//				ret[mpIDIndex[id]] = &Stu{}
//			}
//		}
//		return ret, err
//	}
//
//	redisCache := map[string]*Stu{}
//	for i, r := range dbData {
//		ret[mpIDIndex[r.ID]] = &dbData[i]
//		redisCache["test_"+fmt.Sprintf("%v", r.ID)] = &dbData[i]
//	}
//
//	if len(sqlids) != len(dbData) {
//		tmpCache := map[string]*Stu{}
//		for _, id := range sqlids {
//			if ret[mpIDIndex[id]] == nil {
//				tmpCache["test_"+fmt.Sprintf("%v", id)] = nil
//				ret[mpIDIndex[id]] = &Stu{}
//			}
//		}
//		if len(tmpCache) > 0 {
//			//ctl.SetRedis(redisCache)
//		}
//	}
//
//	if len(redisCache) > 0 {
//		//ctl.SetRedis(redisCache)
//	}
//	return ret, nil
//
//}