package ioc

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	simpleRuleIRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/users/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	dbinit "github.com/RafaelSKaturabara/Flickly/internal/infra/data/init"
	simpleRuleRepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/repositories"
	infrarepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/users/repositories"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)

	// users
	usersDb := dbinit.GetUsersLocalDBConnection()
	dbinit.MigrateDatabase(usersDb)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infrarepositories.NewUserRepository(usersDb))

	// simple rule
	simpleRuleDb := dbinit.GetLocalSimpleRuleDBConnection()
	dbinit.MigrateSimpleRuleDatabase(simpleRuleDb)

	utilities.AddService[simpleRuleIRepositories.ICategoryRepository](serviceCollection, simpleRuleRepositories.NewCategoryRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.ICommitmentRepository](serviceCollection, simpleRuleRepositories.NewCommitmentRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.IExpenseRepository](serviceCollection, simpleRuleRepositories.NewExpenseRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.IIncomeRepository](serviceCollection, simpleRuleRepositories.NewIncomeRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.ISubcategoryRepository](serviceCollection, simpleRuleRepositories.NewSubcategoryRepository(simpleRuleDb))
}
