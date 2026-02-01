package ioc

import (
	"github.com/RafaelSKaturabara/Flickly/internal/domain/core/mediator"
	simpleRuleIRepositories "github.com/RafaelSKaturabara/Flickly/internal/domain/simplerule/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/domain/identity/repositories"
	"github.com/RafaelSKaturabara/Flickly/internal/infra/crosscutting/utilities"
	dataUsers "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity"
	dataSimplerule "github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule"
	simpleRuleRepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/simplerule/repositories"
	identityRepositories "github.com/RafaelSKaturabara/Flickly/internal/infra/data/identity/repositories"
)

func InjectRepositories(serviceCollection utilities.IServiceCollection) {
	mediatR := mediator.NewMediatR()
	// teste
	utilities.AddService[mediator.Mediator](serviceCollection, mediatR)

	// users
	usersDb := dataUsers.GetIdentityLocalDBConnection()
	dataUsers.MigrateDatabase(usersDb)
	utilities.AddService[repositories.IAccessGrantRepository](serviceCollection, identityRepositories.NewAccessGrantRepository(usersDb))
	utilities.AddService[repositories.IClientRepository](serviceCollection, identityRepositories.NewClientRepository(usersDb))
	utilities.AddService[repositories.IRedirectURIRepository](serviceCollection, identityRepositories.NewRedirectURIRepository(usersDb))
	utilities.AddService[repositories.IRoleRepository](serviceCollection, identityRepositories.NewRoleRespository(usersDb))
	utilities.AddService[repositories.IUserAccessRepository](serviceCollection, identityRepositories.NewUserAccessRepository(usersDb))
	utilities.AddService[repositories.IUserRepository](serviceCollection, identityRepositories.NewUserRepository(usersDb))

	// simple rule
	simpleRuleDb := dataSimplerule.GetLocalSimpleRuleDBConnection()
	dataSimplerule.MigrateSimpleRuleDatabase(simpleRuleDb)

	utilities.AddService[simpleRuleIRepositories.ICategoryRepository](serviceCollection, simpleRuleRepositories.NewCategoryRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.ICommitmentRepository](serviceCollection, simpleRuleRepositories.NewCommitmentRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.IExpenseRepository](serviceCollection, simpleRuleRepositories.NewExpenseRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.IIncomeRepository](serviceCollection, simpleRuleRepositories.NewIncomeRepository(simpleRuleDb))
	utilities.AddService[simpleRuleIRepositories.ISubcategoryRepository](serviceCollection, simpleRuleRepositories.NewSubcategoryRepository(simpleRuleDb))
}
