package sql

import (
	"context"

	"github.com/spaceuptech/space-api-go/api/config"
	"github.com/spaceuptech/space-api-go/api/utils"
	"github.com/spaceuptech/space-api-go/api/model"
	"github.com/spaceuptech/space-api-go/api/proto"
	"github.com/spaceuptech/space-api-go/api/realtime"
)

// SQL is the client responsible to commuicate with the SQL crud module
type SQL struct {
	config *config.Config
	db     string
}

// Init returns a SQL client object
func Init(db string, config *config.Config) *SQL {
	return &SQL{db: db, config: config}
}

// Insert returns a helper to fire a insert request
func (s *SQL) Insert(col string) *Insert {
	return initInsert(context.TODO(), s.db, col, s.config)
}

// Get returns a helper to fire a get all request
func (s *SQL) Get(col string) *Get {
	return initGet(context.TODO(), s.db, col, utils.All, s.config)
}

// GetOne returns a helper to fire a get one request
func (s *SQL) GetOne(col string) *Get {
	return initGet(context.TODO(), s.db, col, utils.One, s.config)
}

// Update returns a helper to fire a update all request
func (s *SQL) Update(col string) *Update {
	return initUpdate(context.TODO(), s.db, col, utils.All, s.config)
}

// UpdateOne returns a helper to fire a update one request
func (s *SQL) UpdateOne(col string) *Update {
	return initUpdate(context.TODO(), s.db, col, utils.One, s.config)
}

// Delete returns a helper to fire a delete all request
func (s *SQL) Delete(col string) *Delete {
	return initDelete(context.TODO(), s.db, col, utils.All, s.config)
}

// DeleteOne returns a helper to fire a delete one request
func (s *SQL) DeleteOne(col string) *Delete {
	return initDelete(context.TODO(), s.db, col, utils.One, s.config)
}

// BeginBatch returns a helper to fire a batch request
func (s *SQL) BeginBatch() *Batch {
	return initBatch(context.TODO(), s.db, s.config)
}

// LiveQuery returns a helper to fire a liveQuery request
func (s *SQL) LiveQuery(col string) *realtime.LiveQuery {
	return realtime.Init(s.config, s.db, col)
}

// Profile fires a profile request
func (s *SQL) Profile(id string) (*model.Response, error) {
	m := &proto.Meta{DbType: s.db, Project: s.config.Project, Token: s.config.Token}
	return s.config.Transport.Profile(context.TODO(), m, id)
}

// Profiles fires a profiles request
func (s *SQL) Profiles() (*model.Response, error) {
	m := &proto.Meta{DbType: s.db, Project: s.config.Project, Token: s.config.Token}
	return s.config.Transport.Profiles(context.TODO(), m)
}

// SignIn fires a signIn request
func (s *SQL) SignIn(email, password string) (*model.Response, error) {
	m := &proto.Meta{DbType: s.db, Project: s.config.Project, Token: s.config.Token}
	return s.config.Transport.SignIn(context.TODO(), m, email, password)
}

// SignUp fires a signUp request
func (s *SQL) SignUp(email, name, password, role string) (*model.Response, error) {
	m := &proto.Meta{DbType: s.db, Project: s.config.Project, Token: s.config.Token}
	return s.config.Transport.SignUp(context.TODO(), m, email, name, password, role)
}

// EditProfile fires a editProfile request
func (s *SQL) EditProfile(id string, values model.ProfileParams) (*model.Response, error) {
	m := &proto.Meta{DbType: s.db, Project: s.config.Project, Token: s.config.Token}
	return s.config.Transport.EditProfile(context.TODO(), m, id, values)
}
