/**
 * @Author: yzy
 * @Description:
 * @Version: 1.0.0
 * @Date: 2021/8/19 15:32
 * @Copyright: MIN-Group；国家重大科技基础设施——未来网络北大实验室；深圳市信息论与未来网络重点实验室
 */
package kernel

type DB interface {
	// 根据传入的hash 从数据库中取出rlp编码的stateObject 并进行解码
	OpenAccount(addr Address) *stateObject
	// 传入stateObject 对其进行rlp编码 然后插入数据库中去
	SaveToDB(*stateObject) error
	// 数据库是否存在账户
	ExistAccount(Address) bool
	// 更新账户数据
	UpdateAccount(Address, *stateObject) error
}

type MockDB struct{}

func (*MockDB) OpenAccount(addr Address) *stateObject {
	return nil
}

func (*MockDB) SaveToDB(object *stateObject) error {
	return nil
}

func (*MockDB) ExistAccount(address Address) bool {
	return true
}

func (MockDB) UpdateAccount(address Address, object *stateObject) error {
	return nil
}