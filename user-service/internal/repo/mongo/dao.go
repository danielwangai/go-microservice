package repo

//go:generate mockgen -destination=mocks/mock_dao.go -package=mocks . DAO
type DAO interface {
}
