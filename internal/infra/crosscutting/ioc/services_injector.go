package ioc

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	simpleRuleIRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	dataUsers "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity"
	dataSimplerule "github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule"
	simpleRuleRepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/repositories"
	infrarepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/repositories"
)

func InjectServices(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)

	// users
	usersDb := dataUsers.GetUsersLocalDBConnection()
	dataUsers.MigrateDatabase(usersDb)
	utilities.AddService[repositories.IUserRepository](serviceCollection, infrarepositories.NewUserRepository(usersDb))

	// simple rule
	simpleRuleDb := dataSimplerule.GetLocalSimpleRuleDBConnection()
	dataSimplerule.MigrateSimpleRuleDatabase(simpleRuleDb)

	utilities.AddService[simpleRuleIRepositories.ICategoryRepository](serviceCollection, simpleRuleRepositories.NewCategoryRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.ICommitmentRepository](serviceCollection, simpleRuleRepositories.NewCommitmentRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.IExpenseRepository](serviceCollection, simpleRuleRepositories.NewExpenseRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.IIncomeRepository](serviceCollection, simpleRuleRepositories.NewIncomeRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.ISubcategoryRepository](serviceCollection, simpleRuleRepositories.NewSubcategoryRepository(simpleRuleDb))
}
