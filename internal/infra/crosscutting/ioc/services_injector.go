package ioc

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	simpleRuleIRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/data/core"
	simpleRuleRepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/repositories"
	infrarepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/users/repositories"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infrarepositories.NewUserRepository())

	// simple rule

	db := core.GetLocalDBConnection()
	core.MigrateDatabase(db)

	utilities.AddService[simpleRuleIRepositories.ICategoryRepository](serviceCollection, simpleRuleRepositories.NewCategoryRepository(db))
	utilities.AddService[simpleRuleIRepositories.ICommitmentRepository](serviceCollection, simpleRuleRepositories.NewCommitmentRepository(db))
	utilities.AddService[simpleRuleIRepositories.IExpenseRepository](serviceCollection, simpleRuleRepositories.NewExpenseRepository(db))
	utilities.AddService[simpleRuleIRepositories.IIncomeRepository](serviceCollection, simpleRuleRepositories.NewIncomeRepository(db))
	utilities.AddService[simpleRuleIRepositories.ISubcategoryRepository](serviceCollection, simpleRuleRepositories.NewSubcategoryRepository(db))
}
