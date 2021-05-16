package abstractfactory

import "fmt"

//订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

//订单详情记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

//抽象工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//关系型数据库的ordermaindao实现
type RDBMainDAO struct {
}

func (r *RDBMainDAO) SaveOrderMain() {
	panic("implement me")
}

func (r *RDBMainDAO) OrderMainDAO() {
	fmt.Print("rdb detail save\n")
}

//RDB抽象工厂的实现
type RDBDAOFactory struct {
}

func (r *RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

//
//func (r *RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
//	return &{}
//}
