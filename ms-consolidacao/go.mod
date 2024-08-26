module github.com/sialka/cartola_fc

go 1.22.4

require (
	github.com/go-sql-driver/mysql v1.8.1
	github.com/google/uuid v1.6.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

replace github.com/sialka/cartola_fc/pkg/uow => /home/sidnei/developer/cartola-fc/ms-consolidacao/pkg/uow
