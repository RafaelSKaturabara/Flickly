package ioc

import (
	"github.com/rkaturabara/flickly/internal/domain/core/mediator"
	"github.com/rkaturabara/flickly/internal/domain/users/repositories"
	simpleRuleIRepositories"github.com/rkaturabara/flickly/internal/domain/simplerule/repositories"
	"github.com/rkaturabara/flickly/internal/infra/crosscutting/utilities"
	infrarepositories "github.com/rkaturabara/flickly/internal/infra/data/users/repositories"
	simpleRuleRepositories "github.com/rkaturabara/flickly/internal/infra/data/simplerule/repositories"
	"github.com/rkaturabara/flickly/internal/infra/data/core"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infrarepositories.NewUserRepository())

	// simple rule
	
	db := core.GetLocalDBConnection()

	utilities.AddService[simpleRuleIRepositories.ICategoryRepository](serviceCollection, simpleRuleRepositories.NewCategoryRepository(db))
	utilities.AddService[simpleRuleIRepositories.ICommitmentRepository](serviceCollection, simpleRuleRepositories.NewCommitmentRepository(db))
	utilities.AddService[simpleRuleIRepositories.IExpenseRepository](serviceCollection, simpleRuleRepositories.NewExpenseRepository(db))
	utilities.AddService[simpleRuleIRepositories.IIncomeRepository](serviceCollection, simpleRuleRepositories.NewIncomeRepository(db))
	utilities.AddService[simpleRuleIRepositories.ISubcategoryRepository](serviceCollection, simpleRuleRepositories.NewSubcategoryRepository(db))
}
