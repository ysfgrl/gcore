package grepo

import (
	"context"

	"github.com/ysfgrl/gcore/gerror"
	"github.com/ysfgrl/gcore/gmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRepository[DType any] interface {
	GetById(ctx context.Context, id primitive.ObjectID, isAgg bool) (*DType, *gerror.Error)
	GetByKey(ctx context.Context, key string, value any, isAgg bool) (*DType, *gerror.Error)
	GetByQuery(ctx context.Context, query bson.M, isAgg bool) (*DType, *gerror.Error)
	GetByAggregate(ctx context.Context, query bson.M, agg []bson.M) (*DType, *gerror.Error)
	List(ctx context.Context, filters gmodel.ListRequest) (*gmodel.ListResponse[DType], *gerror.Error)
	Count(ctx context.Context, filters gmodel.ListRequest) (gmodel.Int, *gerror.Error)
	DeleteById(ctx context.Context, id primitive.ObjectID) *gerror.Error

	Create(ctx context.Context, schema DType) (*DType, *gerror.Error)
	CreateMany(ctx context.Context, schemas []DType) (int, *gerror.Error)
	UpdateMany(ctx context.Context, schemas []DType, updateBy []string, createNotExist bool) (int, *gerror.Error)
	Init()

	UpdateField(ctx context.Context, id primitive.ObjectID, key string, val any) (*DType, *gerror.Error)
	UpdateFields(ctx context.Context, id primitive.ObjectID, fields map[string]any) (*DType, *gerror.Error)
	UpdateOne(ctx context.Context, filter bson.M, query bson.M, upsert bool) (*DType, *gerror.Error)
	UpdateOneModel(ctx context.Context, filter bson.M, schema DType, upsert bool) (*DType, *gerror.Error)
	UpdateById(ctx context.Context, id primitive.ObjectID, schema DType) (*DType, *gerror.Error)
}
